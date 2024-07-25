// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: KEYVALUELIST_CN_EN_JP_TW.proto

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

type KEYVALUELIST_CN_EN_JP_TW struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       *uint32     `protobuf:"varint,1,req,name=key" json:"key,omitempty"`
	ValueList []*KEYVALUE `protobuf:"bytes,2,rep,name=value_list,json=valueList" json:"value_list,omitempty"`
}

func (x *KEYVALUELIST_CN_EN_JP_TW) Reset() {
	*x = KEYVALUELIST_CN_EN_JP_TW{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KEYVALUELIST_CN_EN_JP_TW_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KEYVALUELIST_CN_EN_JP_TW) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KEYVALUELIST_CN_EN_JP_TW) ProtoMessage() {}

func (x *KEYVALUELIST_CN_EN_JP_TW) ProtoReflect() protoreflect.Message {
	mi := &file_KEYVALUELIST_CN_EN_JP_TW_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KEYVALUELIST_CN_EN_JP_TW.ProtoReflect.Descriptor instead.
func (*KEYVALUELIST_CN_EN_JP_TW) Descriptor() ([]byte, []int) {
	return file_KEYVALUELIST_CN_EN_JP_TW_proto_rawDescGZIP(), []int{0}
}

func (x *KEYVALUELIST_CN_EN_JP_TW) GetKey() uint32 {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return 0
}

func (x *KEYVALUELIST_CN_EN_JP_TW) GetValueList() []*KEYVALUE {
	if x != nil {
		return x.ValueList
	}
	return nil
}

var File_KEYVALUELIST_CN_EN_JP_TW_proto protoreflect.FileDescriptor

var file_KEYVALUELIST_CN_EN_JP_TW_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x4b, 0x45, 0x59, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x43,
	0x4e, 0x5f, 0x45, 0x4e, 0x5f, 0x4a, 0x50, 0x5f, 0x54, 0x57, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x1a, 0x0e, 0x4b, 0x45, 0x59, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x18, 0x4b, 0x45, 0x59,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x43, 0x4e, 0x5f, 0x45, 0x4e, 0x5f,
	0x4a, 0x50, 0x5f, 0x54, 0x57, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x65,
	0x6c, 0x66, 0x61, 0x73, 0x74, 0x2e, 0x4b, 0x45, 0x59, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x52, 0x09,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_KEYVALUELIST_CN_EN_JP_TW_proto_rawDescOnce sync.Once
	file_KEYVALUELIST_CN_EN_JP_TW_proto_rawDescData = file_KEYVALUELIST_CN_EN_JP_TW_proto_rawDesc
)

func file_KEYVALUELIST_CN_EN_JP_TW_proto_rawDescGZIP() []byte {
	file_KEYVALUELIST_CN_EN_JP_TW_proto_rawDescOnce.Do(func() {
		file_KEYVALUELIST_CN_EN_JP_TW_proto_rawDescData = protoimpl.X.CompressGZIP(file_KEYVALUELIST_CN_EN_JP_TW_proto_rawDescData)
	})
	return file_KEYVALUELIST_CN_EN_JP_TW_proto_rawDescData
}

var file_KEYVALUELIST_CN_EN_JP_TW_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_KEYVALUELIST_CN_EN_JP_TW_proto_goTypes = []any{
	(*KEYVALUELIST_CN_EN_JP_TW)(nil), // 0: belfast.KEYVALUELIST_CN_EN_JP_TW
	(*KEYVALUE)(nil),                 // 1: belfast.KEYVALUE
}
var file_KEYVALUELIST_CN_EN_JP_TW_proto_depIdxs = []int32{
	1, // 0: belfast.KEYVALUELIST_CN_EN_JP_TW.value_list:type_name -> belfast.KEYVALUE
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_KEYVALUELIST_CN_EN_JP_TW_proto_init() }
func file_KEYVALUELIST_CN_EN_JP_TW_proto_init() {
	if File_KEYVALUELIST_CN_EN_JP_TW_proto != nil {
		return
	}
	file_KEYVALUE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_KEYVALUELIST_CN_EN_JP_TW_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*KEYVALUELIST_CN_EN_JP_TW); i {
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
			RawDescriptor: file_KEYVALUELIST_CN_EN_JP_TW_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_KEYVALUELIST_CN_EN_JP_TW_proto_goTypes,
		DependencyIndexes: file_KEYVALUELIST_CN_EN_JP_TW_proto_depIdxs,
		MessageInfos:      file_KEYVALUELIST_CN_EN_JP_TW_proto_msgTypes,
	}.Build()
	File_KEYVALUELIST_CN_EN_JP_TW_proto = out.File
	file_KEYVALUELIST_CN_EN_JP_TW_proto_rawDesc = nil
	file_KEYVALUELIST_CN_EN_JP_TW_proto_goTypes = nil
	file_KEYVALUELIST_CN_EN_JP_TW_proto_depIdxs = nil
}