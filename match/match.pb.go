// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/match/match.proto

package match

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type State int32

const (
	State_Ready    State = 0
	State_Progress State = 1
	State_Suspend  State = 2
	State_Stop     State = 3
)

var State_name = map[int32]string{
	0: "Ready",
	1: "Progress",
	2: "Suspend",
	3: "Stop",
}

var State_value = map[string]int32{
	"Ready":    0,
	"Progress": 1,
	"Suspend":  2,
	"Stop":     3,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1ebb343b10e6eea4, []int{0}
}

func init() {
	proto.RegisterEnum("match.State", State_name, State_value)
}

func init() {
	proto.RegisterFile("proto/match/match.proto", fileDescriptor_1ebb343b10e6eea4)
}

var fileDescriptor_1ebb343b10e6eea4 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x80, 0x90, 0x7a, 0x60, 0x11, 0x21, 0x56, 0x30, 0x47,
	0xcb, 0x9c, 0x8b, 0x35, 0xb8, 0x24, 0xb1, 0x24, 0x55, 0x88, 0x93, 0x8b, 0x35, 0x28, 0x35, 0x31,
	0xa5, 0x52, 0x80, 0x41, 0x88, 0x87, 0x8b, 0x23, 0xa0, 0x28, 0x3f, 0xbd, 0x28, 0xb5, 0xb8, 0x58,
	0x80, 0x51, 0x88, 0x9b, 0x8b, 0x3d, 0xb8, 0xb4, 0xb8, 0x20, 0x35, 0x2f, 0x45, 0x80, 0x49, 0x88,
	0x83, 0x8b, 0x25, 0xb8, 0x24, 0xbf, 0x40, 0x80, 0xd9, 0x89, 0x37, 0x8a, 0x1b, 0x6c, 0x82, 0x35,
	0x98, 0x4c, 0x62, 0x03, 0x9b, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x04, 0x8f, 0xe3, 0x45,
	0x70, 0x00, 0x00, 0x00,
}