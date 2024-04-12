// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: CS_17607.proto

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

type CS_17607 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shipgroup  *uint32 `protobuf:"varint,1,req,name=shipgroup" json:"shipgroup,omitempty"`
	Shareid    *uint32 `protobuf:"varint,2,req,name=shareid" json:"shareid,omitempty"`
	ReportType *uint32 `protobuf:"varint,3,req,name=report_type,json=reportType" json:"report_type,omitempty"`
}

func (x *CS_17607) Reset() {
	*x = CS_17607{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CS_17607_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_17607) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_17607) ProtoMessage() {}

func (x *CS_17607) ProtoReflect() protoreflect.Message {
	mi := &file_CS_17607_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_17607.ProtoReflect.Descriptor instead.
func (*CS_17607) Descriptor() ([]byte, []int) {
	return file_CS_17607_proto_rawDescGZIP(), []int{0}
}

func (x *CS_17607) GetShipgroup() uint32 {
	if x != nil && x.Shipgroup != nil {
		return *x.Shipgroup
	}
	return 0
}

func (x *CS_17607) GetShareid() uint32 {
	if x != nil && x.Shareid != nil {
		return *x.Shareid
	}
	return 0
}

func (x *CS_17607) GetReportType() uint32 {
	if x != nil && x.ReportType != nil {
		return *x.ReportType
	}
	return 0
}

var File_CS_17607_proto protoreflect.FileDescriptor

var file_CS_17607_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x43, 0x53, 0x5f, 0x31, 0x37, 0x36, 0x30, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x62, 0x65, 0x6c, 0x66, 0x61, 0x73, 0x74, 0x22, 0x63, 0x0a, 0x08, 0x43, 0x53, 0x5f,
	0x31, 0x37, 0x36, 0x30, 0x37, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68, 0x61, 0x72, 0x65, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x65, 0x69, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
}

var (
	file_CS_17607_proto_rawDescOnce sync.Once
	file_CS_17607_proto_rawDescData = file_CS_17607_proto_rawDesc
)

func file_CS_17607_proto_rawDescGZIP() []byte {
	file_CS_17607_proto_rawDescOnce.Do(func() {
		file_CS_17607_proto_rawDescData = protoimpl.X.CompressGZIP(file_CS_17607_proto_rawDescData)
	})
	return file_CS_17607_proto_rawDescData
}

var file_CS_17607_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CS_17607_proto_goTypes = []interface{}{
	(*CS_17607)(nil), // 0: belfast.CS_17607
}
var file_CS_17607_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CS_17607_proto_init() }
func file_CS_17607_proto_init() {
	if File_CS_17607_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CS_17607_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_17607); i {
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
			RawDescriptor: file_CS_17607_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CS_17607_proto_goTypes,
		DependencyIndexes: file_CS_17607_proto_depIdxs,
		MessageInfos:      file_CS_17607_proto_msgTypes,
	}.Build()
	File_CS_17607_proto = out.File
	file_CS_17607_proto_rawDesc = nil
	file_CS_17607_proto_goTypes = nil
	file_CS_17607_proto_depIdxs = nil
}