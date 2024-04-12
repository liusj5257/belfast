// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: SC_19114.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SC_19114 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  *int32     `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	Theme   *DORMTHEME `protobuf:"bytes,2,req,name=theme" json:"theme,omitempty"`
	HasFav  *bool      `protobuf:"varint,3,req,name=has_fav,json=hasFav" json:"has_fav,omitempty"`
	HasLike *bool      `protobuf:"varint,4,req,name=has_like,json=hasLike" json:"has_like,omitempty"`
}

func (x *SC_19114) Reset() {
	*x = SC_19114{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SC_19114_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_19114) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_19114) ProtoMessage() {}

func (x *SC_19114) ProtoReflect() protoreflect.Message {
	mi := &file_SC_19114_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_19114.ProtoReflect.Descriptor instead.
func (*SC_19114) Descriptor() ([]byte, []int) {
	return file_SC_19114_proto_rawDescGZIP(), []int{0}
}

func (x *SC_19114) GetResult() int32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *SC_19114) GetTheme() *DORMTHEME {
	if x != nil {
		return x.Theme
	}
	return nil
}

func (x *SC_19114) GetHasFav() bool {
	if x != nil && x.HasFav != nil {
		return *x.HasFav
	}
	return false
}

func (x *SC_19114) GetHasLike() bool {
	if x != nil && x.HasLike != nil {
		return *x.HasLike
	}
	return false
}

var File_SC_19114_proto protoreflect.FileDescriptor

var file_SC_19114_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x43, 0x5f, 0x31, 0x39, 0x31, 0x31, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x0f, 0x44, 0x4f, 0x52, 0x4d, 0x54,
	0x48, 0x45, 0x4d, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x53,
	0x43, 0x5f, 0x31, 0x39, 0x31, 0x31, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x28, 0x0a, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x44, 0x4f, 0x52, 0x4d, 0x54, 0x48, 0x45,
	0x4d, 0x45, 0x52, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x61, 0x73,
	0x5f, 0x66, 0x61, 0x76, 0x18, 0x03, 0x20, 0x02, 0x28, 0x08, 0x52, 0x06, 0x68, 0x61, 0x73, 0x46,
	0x61, 0x76, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x18, 0x04,
	0x20, 0x02, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_SC_19114_proto_rawDescOnce sync.Once
	file_SC_19114_proto_rawDescData = file_SC_19114_proto_rawDesc
)

func file_SC_19114_proto_rawDescGZIP() []byte {
	file_SC_19114_proto_rawDescOnce.Do(func() {
		file_SC_19114_proto_rawDescData = protoimpl.X.CompressGZIP(file_SC_19114_proto_rawDescData)
	})
	return file_SC_19114_proto_rawDescData
}

var file_SC_19114_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SC_19114_proto_goTypes = []interface{}{
	(*SC_19114)(nil),  // 0: belfast.SC_19114
	(*DORMTHEME)(nil), // 1: belfast.DORMTHEME
}
var file_SC_19114_proto_depIdxs = []int32{
	1, // 0: belfast.SC_19114.theme:type_name -> belfast.DORMTHEME
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SC_19114_proto_init() }
func file_SC_19114_proto_init() {
	if File_SC_19114_proto != nil {
		return
	}
	file_DORMTHEME_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SC_19114_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_19114); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SC_19114_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SC_19114_proto_goTypes,
		DependencyIndexes: file_SC_19114_proto_depIdxs,
		MessageInfos:      file_SC_19114_proto_msgTypes,
	}.Build()
	File_SC_19114_proto = out.File
	file_SC_19114_proto_rawDesc = nil
	file_SC_19114_proto_goTypes = nil
	file_SC_19114_proto_depIdxs = nil
}