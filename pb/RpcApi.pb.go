// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.2
// source: RpcApi.proto

package pb

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

type RpcReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RpcReq) Reset() {
	*x = RpcReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RpcApi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcReq) ProtoMessage() {}

func (x *RpcReq) ProtoReflect() protoreflect.Message {
	mi := &file_RpcApi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcReq.ProtoReflect.Descriptor instead.
func (*RpcReq) Descriptor() ([]byte, []int) {
	return file_RpcApi_proto_rawDescGZIP(), []int{0}
}

func (x *RpcReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type RpcRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *RpcRes) Reset() {
	*x = RpcRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RpcApi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcRes) ProtoMessage() {}

func (x *RpcRes) ProtoReflect() protoreflect.Message {
	mi := &file_RpcApi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcRes.ProtoReflect.Descriptor instead.
func (*RpcRes) Descriptor() ([]byte, []int) {
	return file_RpcApi_proto_rawDescGZIP(), []int{1}
}

func (x *RpcRes) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *RpcRes) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

var File_RpcApi_proto protoreflect.FileDescriptor

var file_RpcApi_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x52, 0x70, 0x63, 0x41, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a,
	0x0a, 0x06, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x30, 0x0a, 0x06, 0x52, 0x70,
	0x63, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x32, 0x3d, 0x0a, 0x03,
	0x52, 0x70, 0x63, 0x12, 0x19, 0x0a, 0x03, 0x4d, 0x64, 0x35, 0x12, 0x07, 0x2e, 0x52, 0x70, 0x63,
	0x52, 0x65, 0x71, 0x1a, 0x07, 0x2e, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x1b,
	0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x07, 0x2e, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71,
	0x1a, 0x07, 0x2e, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RpcApi_proto_rawDescOnce sync.Once
	file_RpcApi_proto_rawDescData = file_RpcApi_proto_rawDesc
)

func file_RpcApi_proto_rawDescGZIP() []byte {
	file_RpcApi_proto_rawDescOnce.Do(func() {
		file_RpcApi_proto_rawDescData = protoimpl.X.CompressGZIP(file_RpcApi_proto_rawDescData)
	})
	return file_RpcApi_proto_rawDescData
}

var file_RpcApi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_RpcApi_proto_goTypes = []interface{}{
	(*RpcReq)(nil), // 0: RpcReq
	(*RpcRes)(nil), // 1: RpcRes
}
var file_RpcApi_proto_depIdxs = []int32{
	0, // 0: Rpc.Md5:input_type -> RpcReq
	0, // 1: Rpc.Hello:input_type -> RpcReq
	1, // 2: Rpc.Md5:output_type -> RpcRes
	1, // 3: Rpc.Hello:output_type -> RpcRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RpcApi_proto_init() }
func file_RpcApi_proto_init() {
	if File_RpcApi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RpcApi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcReq); i {
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
		file_RpcApi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcRes); i {
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
			RawDescriptor: file_RpcApi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_RpcApi_proto_goTypes,
		DependencyIndexes: file_RpcApi_proto_depIdxs,
		MessageInfos:      file_RpcApi_proto_msgTypes,
	}.Build()
	File_RpcApi_proto = out.File
	file_RpcApi_proto_rawDesc = nil
	file_RpcApi_proto_goTypes = nil
	file_RpcApi_proto_depIdxs = nil
}
