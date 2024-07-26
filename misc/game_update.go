package misc

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/consts"
	"github.com/ggmolly/belfast/logger"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

const (
	versionURL = "https://gitee.com/aqeqqq/belfast-data/raw/main/versions.json"
)

var (
	errRegionMismatch  = errors.New("Region mismatch")
	errVersionMismatch = errors.New("Version mismatch")
)

type VersionMap map[string]Version
type HashMap []GameChecksum

type Version struct {
	Region  string
	Version string
}

type GameChecksum struct {
	Category string
	Hash     string
}

type hashCache struct {
	Region  string
	Version string
	Hashes  HashMap
}

var azurLaneHashes HashMap
var azurLaneVersions VersionMap

func GetLatestVersions() VersionMap {
	return azurLaneVersions
}

func hashFromCache() (HashMap, error) {
	// use gob to read the cache file
	file, err := os.Open(".cached_hashes")
	if err != nil { // no cache
		return nil, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	var cache hashCache
	err = decoder.Decode(&cache)
	if err != nil {
		logger.LogEvent("GameUpdate", "GetHashes", err.Error(), logger.LOG_LEVEL_ERROR)
		return nil, err
	}
	if cache.Region != os.Getenv("AL_REGION") {
		return nil, errRegionMismatch
	}
	if cache.Version != azurLaneVersions[os.Getenv("AL_REGION")].Version {
		return nil, errVersionMismatch
	}
	return cache.Hashes, nil

}

func GetGameHashes() HashMap {
	version := azurLaneVersions[os.Getenv("AL_REGION")].Version

	if azurLaneHashes != nil && azurLaneVersions[os.Getenv("AL_REGION")].Version == version {
		return azurLaneHashes
	}

	// check if we have a cached version
	hashes, err := hashFromCache()
	if err == nil && azurLaneVersions[os.Getenv("AL_REGION")].Version == version {
		azurLaneHashes = hashes
		return hashes
	}

	// no cache, get the hashes from the server
	logger.LogEvent("GameUpdate", "GetHashes", "No cached hashes, fetching from server", logger.LOG_LEVEL_INFO)
	sock, err := net.Dial("tcp", fmt.Sprintf("%s:80", consts.RegionGateways[os.Getenv("AL_REGION")]))
	if err != nil {
		logger.LogEvent("GameUpdate", "GetHashes", err.Error(), logger.LOG_LEVEL_ERROR)
		return nil
	}
	defer sock.Close()

	// Forge an update packet, CS_10800
	promptUpdate := protobuf.CS_10800{
		State:    proto.Uint32(59),  // 59 is something, might need to update this later?
		Platform: proto.String("1"), // iOS
	}
	if os.Getenv("AL_REGION") == "CN" {
		// fix version invalid?
		promptUpdate.State = proto.Uint32(56)
	}
	packet, err := proto.Marshal(&promptUpdate)
	if err != nil {
		logger.LogEvent("GameUpdate", "GetHashes", err.Error(), logger.LOG_LEVEL_ERROR)
		return nil
	}
	connection.InjectPacketHeader(10800, &packet, 0)
	// Send the packet
	logger.LogEvent("GameUpdate", "GetHashes", "Sending update prompt", logger.LOG_LEVEL_INFO)
	if _, err := sock.Write(packet); err != nil {
		logger.LogEvent("GameUpdate", "GetHashes", err.Error(), logger.LOG_LEVEL_ERROR)
		return nil
	}
	// Read the response
	logger.LogEvent("GameUpdate", "GetHashes", "Reading update response", logger.LOG_LEVEL_INFO)
	var responseData protobuf.SC_10801
	response := make([]byte, 1024)
	n, err := sock.Read(response)
	if err != nil || n < 8 {
		logger.LogEvent("GameUpdate", "GetHashes", "Failed to receive response, or invalid response.", logger.LOG_LEVEL_ERROR)
		return nil
	}
	response = response[7:n]
	if err := proto.Unmarshal(response, &responseData); err != nil {
		logger.LogEvent("GameUpdate", "GetHashes", err.Error(), logger.LOG_LEVEL_ERROR)
		return nil
	}
	// Parse the response
	logger.LogEvent("GameUpdate", "GetHashes", "Parsing update response", logger.LOG_LEVEL_INFO)
	for _, hash := range responseData.GetVersion() {
		if !strings.Contains(hash, "hash$") {
			continue
		}
		fields := strings.Split(hash, "$")
		azurLaneHashes = append(azurLaneHashes, GameChecksum{
			Category: fields[1],
			Hash:     hash,
		})
	}
	// Cache the hashes
	cache := hashCache{
		Region:  os.Getenv("AL_REGION"),
		Version: azurLaneVersions[os.Getenv("AL_REGION")].Version,
		Hashes:  azurLaneHashes,
	}
	file, err := os.OpenFile(".cached_hashes", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		logger.LogEvent("GameUpdate", "GetHashes", err.Error(), logger.LOG_LEVEL_ERROR)
		return nil
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(cache)
	if err != nil {
		logger.LogEvent("GameUpdate", "GetHashes", err.Error(), logger.LOG_LEVEL_ERROR)
		return nil
	}
	go UpdateAllData(os.Getenv("AL_REGION"))
	return azurLaneHashes
}

func LastCacheUpdate() time.Time {
	file, err := os.Stat(".cached_hashes")
	if err != nil {
		return time.Time{}
	}
	return file.ModTime()
}

func LastCacheUpdateVersion() string {
	if azurLaneHashes == nil {
		return ""
	}
	return azurLaneVersions[os.Getenv("AL_REGION")].Version
}

func init() {
	// Download the latest versions and parse them to the map
	resp, err := http.Get(versionURL)
	if err != nil {
		logger.LogEvent("GameUpdate", "init", fmt.Sprintf("failed to fetch versions: %s", err.Error()), logger.LOG_LEVEL_ERROR)
		return
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	deserializedMap := make(map[string]string)
	if err := decoder.Decode(&deserializedMap); err != nil {
		logger.LogEvent("GameUpdate", "init", fmt.Sprintf("failed to parse versions: %s", err.Error()), logger.LOG_LEVEL_ERROR)
		return
	}
	azurLaneVersions = make(VersionMap)
	for region, version := range deserializedMap {
		azurLaneVersions[region] = Version{
			Region:  region,
			Version: version,
		}
	}
}
