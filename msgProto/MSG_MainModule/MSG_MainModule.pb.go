// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MSG_MainModule.proto

package MSG_MainModule // import "msgProto/MSG_MainModule"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// add by stefanchen
// 主模块消息结构
type MAINMSG int32

const (
	MAINMSG_Begin      MAINMSG = 0
	MAINMSG_SERVER     MAINMSG = 1
	MAINMSG_LOGIN      MAINMSG = 2
	MAINMSG_RPC        MAINMSG = 3
	MAINMSG_HEARTBEAT  MAINMSG = 4
	MAINMSG_CENTERGATE MAINMSG = 5
	// for game server
	MAINMSG_PLAYER MAINMSG = 11
	MAINMSG_CHAT   MAINMSG = 12
)

var MAINMSG_name = map[int32]string{
	0:  "Begin",
	1:  "SERVER",
	2:  "LOGIN",
	3:  "RPC",
	4:  "HEARTBEAT",
	5:  "CENTERGATE",
	11: "PLAYER",
	12: "CHAT",
}
var MAINMSG_value = map[string]int32{
	"Begin":      0,
	"SERVER":     1,
	"LOGIN":      2,
	"RPC":        3,
	"HEARTBEAT":  4,
	"CENTERGATE": 5,
	"PLAYER":     11,
	"CHAT":       12,
}

func (x MAINMSG) String() string {
	return proto.EnumName(MAINMSG_name, int32(x))
}
func (MAINMSG) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_MSG_MainModule_caa6a577b768750f, []int{0}
}

func init() {
	proto.RegisterEnum("MSG_MainModule.MAINMSG", MAINMSG_name, MAINMSG_value)
}

func init() {
	proto.RegisterFile("MSG_MainModule.proto", fileDescriptor_MSG_MainModule_caa6a577b768750f)
}

var fileDescriptor_MSG_MainModule_caa6a577b768750f = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xf1, 0x0d, 0x76, 0x8f,
	0xf7, 0x4d, 0xcc, 0xcc, 0xf3, 0xcd, 0x4f, 0x29, 0xcd, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x43, 0x15, 0xd5, 0xca, 0xe4, 0x62, 0xf7, 0x75, 0xf4, 0xf4, 0xf3, 0x0d, 0x76, 0x17,
	0xe2, 0xe4, 0x62, 0x75, 0x4a, 0x4d, 0xcf, 0xcc, 0x13, 0x60, 0x10, 0xe2, 0xe2, 0x62, 0x0b, 0x76,
	0x0d, 0x0a, 0x73, 0x0d, 0x12, 0x60, 0x04, 0x09, 0xfb, 0xf8, 0xbb, 0x7b, 0xfa, 0x09, 0x30, 0x09,
	0xb1, 0x73, 0x31, 0x07, 0x05, 0x38, 0x0b, 0x30, 0x0b, 0xf1, 0x72, 0x71, 0x7a, 0xb8, 0x3a, 0x06,
	0x85, 0x38, 0xb9, 0x3a, 0x86, 0x08, 0xb0, 0x08, 0xf1, 0x71, 0x71, 0x39, 0xbb, 0xfa, 0x85, 0xb8,
	0x06, 0xb9, 0x3b, 0x86, 0xb8, 0x0a, 0xb0, 0x82, 0xb4, 0x07, 0xf8, 0x38, 0x46, 0xba, 0x06, 0x09,
	0x70, 0x0b, 0x71, 0x70, 0xb1, 0x38, 0x7b, 0x38, 0x86, 0x08, 0xf0, 0x38, 0x49, 0x46, 0x89, 0xe7,
	0x16, 0xa7, 0x07, 0x80, 0x9c, 0xa1, 0x8f, 0xea, 0x8a, 0x24, 0x36, 0xb0, 0xe3, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x5d, 0xc3, 0x1b, 0xd0, 0xb4, 0x00, 0x00, 0x00,
}
