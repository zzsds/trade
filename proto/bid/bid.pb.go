// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: proto/bid/bid.proto

package bid

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Type int32

const (
	Type_Buy  Type = 0
	Type_Sell Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "Buy",
		1: "Sell",
	}
	Type_value = map[string]int32{
		"Buy":  0,
		"Sell": 1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_bid_bid_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_proto_bid_bid_proto_enumTypes[0]
}

func (x Type) Amount() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_bid_bid_proto_rawDescGZIP(), []int{0}
}

var File_proto_bid_bid_proto protoreflect.FileDescriptor

var file_proto_bid_bid_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x64, 0x2f, 0x62, 0x69, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x62, 0x69, 0x64, 0x2a, 0x19, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x75, 0x79, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53,
	0x65, 0x6c, 0x6c, 0x10, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62,
	0x69, 0x64, 0x3b, 0x62, 0x69, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_bid_bid_proto_rawDescOnce sync.Once
	file_proto_bid_bid_proto_rawDescData = file_proto_bid_bid_proto_rawDesc
)

func file_proto_bid_bid_proto_rawDescGZIP() []byte {
	file_proto_bid_bid_proto_rawDescOnce.Do(func() {
		file_proto_bid_bid_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bid_bid_proto_rawDescData)
	})
	return file_proto_bid_bid_proto_rawDescData
}

var file_proto_bid_bid_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_bid_bid_proto_goTypes = []interface{}{
	(Type)(0), // 0: bid.Type
}
var file_proto_bid_bid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_bid_bid_proto_init() }
func file_proto_bid_bid_proto_init() {
	if File_proto_bid_bid_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_bid_bid_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_bid_bid_proto_goTypes,
		DependencyIndexes: file_proto_bid_bid_proto_depIdxs,
		EnumInfos:         file_proto_bid_bid_proto_enumTypes,
	}.Build()
	File_proto_bid_bid_proto = out.File
	file_proto_bid_bid_proto_rawDesc = nil
	file_proto_bid_bid_proto_goTypes = nil
	file_proto_bid_bid_proto_depIdxs = nil
}
