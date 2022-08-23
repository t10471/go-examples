// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: protos/other.proto

package proto

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

type OtherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *OtherRequest) Reset() {
	*x = OtherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_other_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherRequest) ProtoMessage() {}

func (x *OtherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_other_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherRequest.ProtoReflect.Descriptor instead.
func (*OtherRequest) Descriptor() ([]byte, []int) {
	return file_protos_other_proto_rawDescGZIP(), []int{0}
}

func (x *OtherRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type OtherReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OtherReply) Reset() {
	*x = OtherReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_other_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtherReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherReply) ProtoMessage() {}

func (x *OtherReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_other_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherReply.ProtoReflect.Descriptor instead.
func (*OtherReply) Descriptor() ([]byte, []int) {
	return file_protos_other_proto_rawDescGZIP(), []int{1}
}

func (x *OtherReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_other_proto protoreflect.FileDescriptor

var file_protos_other_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x0c, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x26, 0x0a, 0x0a, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x77, 0x0a, 0x05, 0x4f, 0x74, 0x68, 0x65, 0x72,
	0x12, 0x35, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x6c, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x13, 0x2e,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x56, 0x32, 0x12, 0x13, 0x2e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x31, 0x30, 0x34, 0x37, 0x31, 0x2f, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_other_proto_rawDescOnce sync.Once
	file_protos_other_proto_rawDescData = file_protos_other_proto_rawDesc
)

func file_protos_other_proto_rawDescGZIP() []byte {
	file_protos_other_proto_rawDescOnce.Do(func() {
		file_protos_other_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_other_proto_rawDescData)
	})
	return file_protos_other_proto_rawDescData
}

var file_protos_other_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_other_proto_goTypes = []interface{}{
	(*OtherRequest)(nil), // 0: other.OtherRequest
	(*OtherReply)(nil),   // 1: other.OtherReply
}
var file_protos_other_proto_depIdxs = []int32{
	0, // 0: other.Other.CallOther:input_type -> other.OtherRequest
	0, // 1: other.Other.CallOtherV2:input_type -> other.OtherRequest
	1, // 2: other.Other.CallOther:output_type -> other.OtherReply
	1, // 3: other.Other.CallOtherV2:output_type -> other.OtherReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_other_proto_init() }
func file_protos_other_proto_init() {
	if File_protos_other_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_other_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtherRequest); i {
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
		file_protos_other_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtherReply); i {
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
			RawDescriptor: file_protos_other_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_other_proto_goTypes,
		DependencyIndexes: file_protos_other_proto_depIdxs,
		MessageInfos:      file_protos_other_proto_msgTypes,
	}.Build()
	File_protos_other_proto = out.File
	file_protos_other_proto_rawDesc = nil
	file_protos_other_proto_goTypes = nil
	file_protos_other_proto_depIdxs = nil
}
