// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: transfer/transfer.proto

package transfer

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Body) Reset() {
	*x = Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Body) ProtoMessage() {}

func (x *Body) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Body.ProtoReflect.Descriptor instead.
func (*Body) Descriptor() ([]byte, []int) {
	return file_transfer_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *Body) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Body) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_transfer_transfer_proto protoreflect.FileDescriptor

var file_transfer_transfer_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x65, 0x6c, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x3f, 0x0a, 0x03, 0x41, 0x50,
	0x49, 0x12, 0x38, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x16, 0x2e, 0x65, 0x6c, 0x61, 0x73,
	0x74, 0x69, 0x63, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x64,
	0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x6c, 0x61, 0x73,
	0x74, 0x69, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transfer_transfer_proto_rawDescOnce sync.Once
	file_transfer_transfer_proto_rawDescData = file_transfer_transfer_proto_rawDesc
)

func file_transfer_transfer_proto_rawDescGZIP() []byte {
	file_transfer_transfer_proto_rawDescOnce.Do(func() {
		file_transfer_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_transfer_transfer_proto_rawDescData)
	})
	return file_transfer_transfer_proto_rawDescData
}

var file_transfer_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transfer_transfer_proto_goTypes = []interface{}{
	(*Body)(nil),          // 0: elastic.transfer.Body
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_transfer_transfer_proto_depIdxs = []int32{
	0, // 0: elastic.transfer.API.Push:input_type -> elastic.transfer.Body
	1, // 1: elastic.transfer.API.Push:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transfer_transfer_proto_init() }
func file_transfer_transfer_proto_init() {
	if File_transfer_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transfer_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Body); i {
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
			RawDescriptor: file_transfer_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transfer_transfer_proto_goTypes,
		DependencyIndexes: file_transfer_transfer_proto_depIdxs,
		MessageInfos:      file_transfer_transfer_proto_msgTypes,
	}.Build()
	File_transfer_transfer_proto = out.File
	file_transfer_transfer_proto_rawDesc = nil
	file_transfer_transfer_proto_goTypes = nil
	file_transfer_transfer_proto_depIdxs = nil
}
