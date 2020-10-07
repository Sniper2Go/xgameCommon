// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MSG_HeartBeat.proto

package MSG_HeartBeat // import "msgProto/MSG_HeartBeat"

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
// server
type SUBMSG int32

const (
	SUBMSG_Begin        SUBMSG = 0
	SUBMSG_CS_HeartBeat SUBMSG = 1
	SUBMSG_SC_HeartBeat SUBMSG = 2
)

var SUBMSG_name = map[int32]string{
	0: "Begin",
	1: "CS_HeartBeat",
	2: "SC_HeartBeat",
}
var SUBMSG_value = map[string]int32{
	"Begin":        0,
	"CS_HeartBeat": 1,
	"SC_HeartBeat": 2,
}

func (x SUBMSG) String() string {
	return proto.EnumName(SUBMSG_name, int32(x))
}
func (SUBMSG) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_MSG_HeartBeat_5647d3692269f30f, []int{0}
}

type ErrorCode int32

const (
	ErrorCode_Invalid ErrorCode = 0
	ErrorCode_Success ErrorCode = 1
	ErrorCode_Fail    ErrorCode = 2
)

var ErrorCode_name = map[int32]string{
	0: "Invalid",
	1: "Success",
	2: "Fail",
}
var ErrorCode_value = map[string]int32{
	"Invalid": 0,
	"Success": 1,
	"Fail":    2,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_MSG_HeartBeat_5647d3692269f30f, []int{1}
}

// CS_HeartBeat
type CS_HeartBeat_Req struct {
	SvrPoint             uint32   `protobuf:"varint,1,opt,name=SvrPoint,proto3" json:"SvrPoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_HeartBeat_Req) Reset()         { *m = CS_HeartBeat_Req{} }
func (m *CS_HeartBeat_Req) String() string { return proto.CompactTextString(m) }
func (*CS_HeartBeat_Req) ProtoMessage()    {}
func (*CS_HeartBeat_Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_HeartBeat_5647d3692269f30f, []int{0}
}
func (m *CS_HeartBeat_Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_HeartBeat_Req.Unmarshal(m, b)
}
func (m *CS_HeartBeat_Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_HeartBeat_Req.Marshal(b, m, deterministic)
}
func (dst *CS_HeartBeat_Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_HeartBeat_Req.Merge(dst, src)
}
func (m *CS_HeartBeat_Req) XXX_Size() int {
	return xxx_messageInfo_CS_HeartBeat_Req.Size(m)
}
func (m *CS_HeartBeat_Req) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_HeartBeat_Req.DiscardUnknown(m)
}

var xxx_messageInfo_CS_HeartBeat_Req proto.InternalMessageInfo

func (m *CS_HeartBeat_Req) GetSvrPoint() uint32 {
	if m != nil {
		return m.SvrPoint
	}
	return 0
}

// SC_HeartBeat
type SC_HeartBeat_Rsp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_HeartBeat_Rsp) Reset()         { *m = SC_HeartBeat_Rsp{} }
func (m *SC_HeartBeat_Rsp) String() string { return proto.CompactTextString(m) }
func (*SC_HeartBeat_Rsp) ProtoMessage()    {}
func (*SC_HeartBeat_Rsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_HeartBeat_5647d3692269f30f, []int{1}
}
func (m *SC_HeartBeat_Rsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_HeartBeat_Rsp.Unmarshal(m, b)
}
func (m *SC_HeartBeat_Rsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_HeartBeat_Rsp.Marshal(b, m, deterministic)
}
func (dst *SC_HeartBeat_Rsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_HeartBeat_Rsp.Merge(dst, src)
}
func (m *SC_HeartBeat_Rsp) XXX_Size() int {
	return xxx_messageInfo_SC_HeartBeat_Rsp.Size(m)
}
func (m *SC_HeartBeat_Rsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_HeartBeat_Rsp.DiscardUnknown(m)
}

var xxx_messageInfo_SC_HeartBeat_Rsp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CS_HeartBeat_Req)(nil), "MSG_HeartBeat.CS_HeartBeat_Req")
	proto.RegisterType((*SC_HeartBeat_Rsp)(nil), "MSG_HeartBeat.SC_HeartBeat_Rsp")
	proto.RegisterEnum("MSG_HeartBeat.SUBMSG", SUBMSG_name, SUBMSG_value)
	proto.RegisterEnum("MSG_HeartBeat.ErrorCode", ErrorCode_name, ErrorCode_value)
}

func init() { proto.RegisterFile("MSG_HeartBeat.proto", fileDescriptor_MSG_HeartBeat_5647d3692269f30f) }

var fileDescriptor_MSG_HeartBeat_5647d3692269f30f = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xf6, 0x0d, 0x76, 0x8f,
	0xf7, 0x48, 0x4d, 0x2c, 0x2a, 0x71, 0x4a, 0x4d, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x45, 0x11, 0x54, 0xd2, 0xe3, 0x12, 0x70, 0x0e, 0x46, 0xf0, 0xe3, 0x83, 0x52, 0x0b, 0x85,
	0xa4, 0xb8, 0x38, 0x82, 0xcb, 0x8a, 0x02, 0xf2, 0x33, 0xf3, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35,
	0x78, 0x83, 0xe0, 0x7c, 0x25, 0x21, 0x2e, 0x81, 0x60, 0x67, 0x64, 0xf5, 0xc5, 0x05, 0x5a, 0xe6,
	0x5c, 0x6c, 0xc1, 0xa1, 0x4e, 0xbe, 0xc1, 0xee, 0x42, 0x9c, 0x5c, 0xac, 0x4e, 0xa9, 0xe9, 0x99,
	0x79, 0x02, 0x0c, 0x42, 0x02, 0x5c, 0x3c, 0xc8, 0x06, 0x0b, 0x30, 0x82, 0x44, 0x90, 0xb5, 0x0a,
	0x30, 0x69, 0xe9, 0x73, 0x71, 0xba, 0x16, 0x15, 0xe5, 0x17, 0x39, 0xe7, 0xa7, 0xa4, 0x0a, 0x71,
	0x73, 0xb1, 0x7b, 0xe6, 0x95, 0x25, 0xe6, 0x64, 0xa6, 0x08, 0x30, 0x80, 0x38, 0xc1, 0xa5, 0xc9,
	0xc9, 0xa9, 0xc5, 0xc5, 0x02, 0x8c, 0x42, 0x1c, 0x5c, 0x2c, 0x6e, 0x89, 0x99, 0x39, 0x02, 0x4c,
	0x4e, 0x12, 0x51, 0x62, 0xb9, 0xc5, 0xe9, 0x01, 0x20, 0x8f, 0xe8, 0xa3, 0xf8, 0x23, 0x89, 0x0d,
	0xec, 0x3b, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0x60, 0xe4, 0x4b, 0xf4, 0x00, 0x00,
	0x00,
}
