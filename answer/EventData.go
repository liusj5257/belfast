package answer

import (
	"github.com/bettercallmolly/belfast/connection"

	"github.com/bettercallmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

func EventData(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_26120
	response.WeeklyFree = proto.Uint32(0)
	response.MonthlyTicket = proto.Uint32(0)
	response.PayCoinCount = proto.Uint32(0)
	response.FirstEnter = proto.Uint32(0)
	return client.SendMessage(26120, &response)
}

func init() {
	data := []byte{0x08, 0x00, 0x10, 0x00, 0x20, 0x00, 0x28, 0x01}
	proto.Unmarshal(data, &validSC26120)
}
