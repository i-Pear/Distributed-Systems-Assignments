// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: protocol/protocol.proto

package protocol

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

type TimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TimeRequest) Reset() {
	*x = TimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRequest) ProtoMessage() {}

func (x *TimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRequest.ProtoReflect.Descriptor instead.
func (*TimeRequest) Descriptor() ([]byte, []int) {
	return file_protocol_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *TimeRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type TimeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerTime uint64 `protobuf:"varint,1,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`
}

func (x *TimeReply) Reset() {
	*x = TimeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeReply) ProtoMessage() {}

func (x *TimeReply) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeReply.ProtoReflect.Descriptor instead.
func (*TimeReply) Descriptor() ([]byte, []int) {
	return file_protocol_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *TimeReply) GetServerTime() uint64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

var File_protocol_protocol_proto protoreflect.FileDescriptor

var file_protocol_protocol_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x30, 0x31, 0x22, 0x23, 0x0a, 0x0b, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2c, 0x0a, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x4e, 0x0a, 0x0b, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x30, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x30, 0x31, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x2d, 0x50, 0x65, 0x61, 0x72, 0x2f,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x2d, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x73, 0x2d, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x30, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_protocol_proto_rawDescOnce sync.Once
	file_protocol_protocol_proto_rawDescData = file_protocol_protocol_proto_rawDesc
)

func file_protocol_protocol_proto_rawDescGZIP() []byte {
	file_protocol_protocol_proto_rawDescOnce.Do(func() {
		file_protocol_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_protocol_proto_rawDescData)
	})
	return file_protocol_protocol_proto_rawDescData
}

var file_protocol_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protocol_protocol_proto_goTypes = []interface{}{
	(*TimeRequest)(nil), // 0: Assignment01.TimeRequest
	(*TimeReply)(nil),   // 1: Assignment01.TimeReply
}
var file_protocol_protocol_proto_depIdxs = []int32{
	0, // 0: Assignment01.TimeService.GetTime:input_type -> Assignment01.TimeRequest
	1, // 1: Assignment01.TimeService.GetTime:output_type -> Assignment01.TimeReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protocol_protocol_proto_init() }
func file_protocol_protocol_proto_init() {
	if File_protocol_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRequest); i {
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
		file_protocol_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeReply); i {
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
			RawDescriptor: file_protocol_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_protocol_proto_goTypes,
		DependencyIndexes: file_protocol_protocol_proto_depIdxs,
		MessageInfos:      file_protocol_protocol_proto_msgTypes,
	}.Build()
	File_protocol_protocol_proto = out.File
	file_protocol_protocol_proto_rawDesc = nil
	file_protocol_protocol_proto_goTypes = nil
	file_protocol_protocol_proto_depIdxs = nil
}
