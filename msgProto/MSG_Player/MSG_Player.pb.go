// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MSG_Player.proto

package MSG_Player // import "msgProto/MSG_Player"

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
	SUBMSG_Begin          SUBMSG = 0
	SUBMSG_CS_EnterServer SUBMSG = 1
	SUBMSG_SC_EnterServer SUBMSG = 2
	SUBMSG_CS_LeaveServer SUBMSG = 3
	SUBMSG_SC_LeaveServer SUBMSG = 4
	SUBMSG_CS_PlayerInfo  SUBMSG = 5
	SUBMSG_SC_PlayerInfo  SUBMSG = 6
)

var SUBMSG_name = map[int32]string{
	0: "Begin",
	1: "CS_EnterServer",
	2: "SC_EnterServer",
	3: "CS_LeaveServer",
	4: "SC_LeaveServer",
	5: "CS_PlayerInfo",
	6: "SC_PlayerInfo",
}
var SUBMSG_value = map[string]int32{
	"Begin":          0,
	"CS_EnterServer": 1,
	"SC_EnterServer": 2,
	"CS_LeaveServer": 3,
	"SC_LeaveServer": 4,
	"CS_PlayerInfo":  5,
	"SC_PlayerInfo":  6,
}

func (x SUBMSG) String() string {
	return proto.EnumName(SUBMSG_name, int32(x))
}
func (SUBMSG) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_MSG_Player_cfc91f0f1ed3bc12, []int{0}
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
	return fileDescriptor_MSG_Player_cfc91f0f1ed3bc12, []int{1}
}

// 玩家基础信息枚举
type EmBaseInfo int32

const (
	EmBaseInfo_baseInvalid EmBaseInfo = 0
	EmBaseInfo_Name        EmBaseInfo = 1
	EmBaseInfo_Level       EmBaseInfo = 2
	EmBaseInfo_HeadIcon    EmBaseInfo = 3
	EmBaseInfo_DBID        EmBaseInfo = 4
)

var EmBaseInfo_name = map[int32]string{
	0: "baseInvalid",
	1: "Name",
	2: "Level",
	3: "HeadIcon",
	4: "DBID",
}
var EmBaseInfo_value = map[string]int32{
	"baseInvalid": 0,
	"Name":        1,
	"Level":       2,
	"HeadIcon":    3,
	"DBID":        4,
}

func (x EmBaseInfo) String() string {
	return proto.EnumName(EmBaseInfo_name, int32(x))
}
func (EmBaseInfo) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_MSG_Player_cfc91f0f1ed3bc12, []int{2}
}

// 玩家基础金币类枚举
type EmBaseMoney int32

const (
	EmBaseMoney_moneyInvalid EmBaseMoney = 0
	EmBaseMoney_Coin         EmBaseMoney = 1
)

var EmBaseMoney_name = map[int32]string{
	0: "moneyInvalid",
	1: "Coin",
}
var EmBaseMoney_value = map[string]int32{
	"moneyInvalid": 0,
	"Coin":         1,
}

func (x EmBaseMoney) String() string {
	return proto.EnumName(EmBaseMoney_name, int32(x))
}
func (EmBaseMoney) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_MSG_Player_cfc91f0f1ed3bc12, []int{3}
}

// 玩家基础信息结构
type BasePlayerInfo_Repeat struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Level                int32    `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
	HeadIcon             string   `protobuf:"bytes,3,opt,name=HeadIcon,proto3" json:"HeadIcon,omitempty"`
	DBID                 int64    `protobuf:"varint,4,opt,name=DBID,proto3" json:"DBID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BasePlayerInfo_Repeat) Reset()         { *m = BasePlayerInfo_Repeat{} }
func (m *BasePlayerInfo_Repeat) String() string { return proto.CompactTextString(m) }
func (*BasePlayerInfo_Repeat) ProtoMessage()    {}
func (*BasePlayerInfo_Repeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_Player_cfc91f0f1ed3bc12, []int{0}
}
func (m *BasePlayerInfo_Repeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BasePlayerInfo_Repeat.Unmarshal(m, b)
}
func (m *BasePlayerInfo_Repeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BasePlayerInfo_Repeat.Marshal(b, m, deterministic)
}
func (dst *BasePlayerInfo_Repeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasePlayerInfo_Repeat.Merge(dst, src)
}
func (m *BasePlayerInfo_Repeat) XXX_Size() int {
	return xxx_messageInfo_BasePlayerInfo_Repeat.Size(m)
}
func (m *BasePlayerInfo_Repeat) XXX_DiscardUnknown() {
	xxx_messageInfo_BasePlayerInfo_Repeat.DiscardUnknown(m)
}

var xxx_messageInfo_BasePlayerInfo_Repeat proto.InternalMessageInfo

func (m *BasePlayerInfo_Repeat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BasePlayerInfo_Repeat) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *BasePlayerInfo_Repeat) GetHeadIcon() string {
	if m != nil {
		return m.HeadIcon
	}
	return ""
}

func (m *BasePlayerInfo_Repeat) GetDBID() int64 {
	if m != nil {
		return m.DBID
	}
	return 0
}

// 玩家基础金币类结构
type BaseMoney_Repeat struct {
	Coin                 int64    `protobuf:"varint,1,opt,name=Coin,proto3" json:"Coin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseMoney_Repeat) Reset()         { *m = BaseMoney_Repeat{} }
func (m *BaseMoney_Repeat) String() string { return proto.CompactTextString(m) }
func (*BaseMoney_Repeat) ProtoMessage()    {}
func (*BaseMoney_Repeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_Player_cfc91f0f1ed3bc12, []int{1}
}
func (m *BaseMoney_Repeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseMoney_Repeat.Unmarshal(m, b)
}
func (m *BaseMoney_Repeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseMoney_Repeat.Marshal(b, m, deterministic)
}
func (dst *BaseMoney_Repeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseMoney_Repeat.Merge(dst, src)
}
func (m *BaseMoney_Repeat) XXX_Size() int {
	return xxx_messageInfo_BaseMoney_Repeat.Size(m)
}
func (m *BaseMoney_Repeat) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseMoney_Repeat.DiscardUnknown(m)
}

var xxx_messageInfo_BaseMoney_Repeat proto.InternalMessageInfo

func (m *BaseMoney_Repeat) GetCoin() int64 {
	if m != nil {
		return m.Coin
	}
	return 0
}

// CS_EnterServer
type CS_EnterServer_Req struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_EnterServer_Req) Reset()         { *m = CS_EnterServer_Req{} }
func (m *CS_EnterServer_Req) String() string { return proto.CompactTextString(m) }
func (*CS_EnterServer_Req) ProtoMessage()    {}
func (*CS_EnterServer_Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_Player_cfc91f0f1ed3bc12, []int{2}
}
func (m *CS_EnterServer_Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_EnterServer_Req.Unmarshal(m, b)
}
func (m *CS_EnterServer_Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_EnterServer_Req.Marshal(b, m, deterministic)
}
func (dst *CS_EnterServer_Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_EnterServer_Req.Merge(dst, src)
}
func (m *CS_EnterServer_Req) XXX_Size() int {
	return xxx_messageInfo_CS_EnterServer_Req.Size(m)
}
func (m *CS_EnterServer_Req) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_EnterServer_Req.DiscardUnknown(m)
}

var xxx_messageInfo_CS_EnterServer_Req proto.InternalMessageInfo

// SC_EnterServer
type SC_EnterServer_Rsp struct {
	Ret                  ErrorCode `protobuf:"varint,1,opt,name=Ret,proto3,enum=MSG_Player.ErrorCode" json:"Ret,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SC_EnterServer_Rsp) Reset()         { *m = SC_EnterServer_Rsp{} }
func (m *SC_EnterServer_Rsp) String() string { return proto.CompactTextString(m) }
func (*SC_EnterServer_Rsp) ProtoMessage()    {}
func (*SC_EnterServer_Rsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_Player_cfc91f0f1ed3bc12, []int{3}
}
func (m *SC_EnterServer_Rsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_EnterServer_Rsp.Unmarshal(m, b)
}
func (m *SC_EnterServer_Rsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_EnterServer_Rsp.Marshal(b, m, deterministic)
}
func (dst *SC_EnterServer_Rsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_EnterServer_Rsp.Merge(dst, src)
}
func (m *SC_EnterServer_Rsp) XXX_Size() int {
	return xxx_messageInfo_SC_EnterServer_Rsp.Size(m)
}
func (m *SC_EnterServer_Rsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_EnterServer_Rsp.DiscardUnknown(m)
}

var xxx_messageInfo_SC_EnterServer_Rsp proto.InternalMessageInfo

func (m *SC_EnterServer_Rsp) GetRet() ErrorCode {
	if m != nil {
		return m.Ret
	}
	return ErrorCode_Invalid
}

// CS_LeaveServer
type CS_LeaveServer_Req struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_LeaveServer_Req) Reset()         { *m = CS_LeaveServer_Req{} }
func (m *CS_LeaveServer_Req) String() string { return proto.CompactTextString(m) }
func (*CS_LeaveServer_Req) ProtoMessage()    {}
func (*CS_LeaveServer_Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_Player_cfc91f0f1ed3bc12, []int{4}
}
func (m *CS_LeaveServer_Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_LeaveServer_Req.Unmarshal(m, b)
}
func (m *CS_LeaveServer_Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_LeaveServer_Req.Marshal(b, m, deterministic)
}
func (dst *CS_LeaveServer_Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_LeaveServer_Req.Merge(dst, src)
}
func (m *CS_LeaveServer_Req) XXX_Size() int {
	return xxx_messageInfo_CS_LeaveServer_Req.Size(m)
}
func (m *CS_LeaveServer_Req) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_LeaveServer_Req.DiscardUnknown(m)
}

var xxx_messageInfo_CS_LeaveServer_Req proto.InternalMessageInfo

// SC_LeaveServer
type SC_LeaveServer_Rsp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_LeaveServer_Rsp) Reset()         { *m = SC_LeaveServer_Rsp{} }
func (m *SC_LeaveServer_Rsp) String() string { return proto.CompactTextString(m) }
func (*SC_LeaveServer_Rsp) ProtoMessage()    {}
func (*SC_LeaveServer_Rsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_Player_cfc91f0f1ed3bc12, []int{5}
}
func (m *SC_LeaveServer_Rsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_LeaveServer_Rsp.Unmarshal(m, b)
}
func (m *SC_LeaveServer_Rsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_LeaveServer_Rsp.Marshal(b, m, deterministic)
}
func (dst *SC_LeaveServer_Rsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_LeaveServer_Rsp.Merge(dst, src)
}
func (m *SC_LeaveServer_Rsp) XXX_Size() int {
	return xxx_messageInfo_SC_LeaveServer_Rsp.Size(m)
}
func (m *SC_LeaveServer_Rsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_LeaveServer_Rsp.DiscardUnknown(m)
}

var xxx_messageInfo_SC_LeaveServer_Rsp proto.InternalMessageInfo

// CS_PlayerInfo
type CS_PlayerInfo_Req struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_PlayerInfo_Req) Reset()         { *m = CS_PlayerInfo_Req{} }
func (m *CS_PlayerInfo_Req) String() string { return proto.CompactTextString(m) }
func (*CS_PlayerInfo_Req) ProtoMessage()    {}
func (*CS_PlayerInfo_Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_Player_cfc91f0f1ed3bc12, []int{6}
}
func (m *CS_PlayerInfo_Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_PlayerInfo_Req.Unmarshal(m, b)
}
func (m *CS_PlayerInfo_Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_PlayerInfo_Req.Marshal(b, m, deterministic)
}
func (dst *CS_PlayerInfo_Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_PlayerInfo_Req.Merge(dst, src)
}
func (m *CS_PlayerInfo_Req) XXX_Size() int {
	return xxx_messageInfo_CS_PlayerInfo_Req.Size(m)
}
func (m *CS_PlayerInfo_Req) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_PlayerInfo_Req.DiscardUnknown(m)
}

var xxx_messageInfo_CS_PlayerInfo_Req proto.InternalMessageInfo

// SC_PlayerInfo
type SC_PlayerInfo_Rsp struct {
	BaseInfo             *BasePlayerInfo_Repeat `protobuf:"bytes,1,opt,name=baseInfo,proto3" json:"baseInfo,omitempty"`
	BaseMoney            *BaseMoney_Repeat      `protobuf:"bytes,2,opt,name=baseMoney,proto3" json:"baseMoney,omitempty"`
	Ret                  ErrorCode              `protobuf:"varint,3,opt,name=Ret,proto3,enum=MSG_Player.ErrorCode" json:"Ret,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SC_PlayerInfo_Rsp) Reset()         { *m = SC_PlayerInfo_Rsp{} }
func (m *SC_PlayerInfo_Rsp) String() string { return proto.CompactTextString(m) }
func (*SC_PlayerInfo_Rsp) ProtoMessage()    {}
func (*SC_PlayerInfo_Rsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_Player_cfc91f0f1ed3bc12, []int{7}
}
func (m *SC_PlayerInfo_Rsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_PlayerInfo_Rsp.Unmarshal(m, b)
}
func (m *SC_PlayerInfo_Rsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_PlayerInfo_Rsp.Marshal(b, m, deterministic)
}
func (dst *SC_PlayerInfo_Rsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_PlayerInfo_Rsp.Merge(dst, src)
}
func (m *SC_PlayerInfo_Rsp) XXX_Size() int {
	return xxx_messageInfo_SC_PlayerInfo_Rsp.Size(m)
}
func (m *SC_PlayerInfo_Rsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_PlayerInfo_Rsp.DiscardUnknown(m)
}

var xxx_messageInfo_SC_PlayerInfo_Rsp proto.InternalMessageInfo

func (m *SC_PlayerInfo_Rsp) GetBaseInfo() *BasePlayerInfo_Repeat {
	if m != nil {
		return m.BaseInfo
	}
	return nil
}

func (m *SC_PlayerInfo_Rsp) GetBaseMoney() *BaseMoney_Repeat {
	if m != nil {
		return m.BaseMoney
	}
	return nil
}

func (m *SC_PlayerInfo_Rsp) GetRet() ErrorCode {
	if m != nil {
		return m.Ret
	}
	return ErrorCode_Invalid
}

func init() {
	proto.RegisterType((*BasePlayerInfo_Repeat)(nil), "MSG_Player.BasePlayerInfo_Repeat")
	proto.RegisterType((*BaseMoney_Repeat)(nil), "MSG_Player.BaseMoney_Repeat")
	proto.RegisterType((*CS_EnterServer_Req)(nil), "MSG_Player.CS_EnterServer_Req")
	proto.RegisterType((*SC_EnterServer_Rsp)(nil), "MSG_Player.SC_EnterServer_Rsp")
	proto.RegisterType((*CS_LeaveServer_Req)(nil), "MSG_Player.CS_LeaveServer_Req")
	proto.RegisterType((*SC_LeaveServer_Rsp)(nil), "MSG_Player.SC_LeaveServer_Rsp")
	proto.RegisterType((*CS_PlayerInfo_Req)(nil), "MSG_Player.CS_PlayerInfo_Req")
	proto.RegisterType((*SC_PlayerInfo_Rsp)(nil), "MSG_Player.SC_PlayerInfo_Rsp")
	proto.RegisterEnum("MSG_Player.SUBMSG", SUBMSG_name, SUBMSG_value)
	proto.RegisterEnum("MSG_Player.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("MSG_Player.EmBaseInfo", EmBaseInfo_name, EmBaseInfo_value)
	proto.RegisterEnum("MSG_Player.EmBaseMoney", EmBaseMoney_name, EmBaseMoney_value)
}

func init() { proto.RegisterFile("MSG_Player.proto", fileDescriptor_MSG_Player_cfc91f0f1ed3bc12) }

var fileDescriptor_MSG_Player_cfc91f0f1ed3bc12 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0xb3, 0xd8, 0x50, 0x18, 0xa7, 0xe9, 0x30, 0x01, 0x09, 0x55, 0x3d, 0x50, 0x1f, 0x5a,
	0xca, 0x21, 0x91, 0xe8, 0xad, 0x52, 0x2e, 0x76, 0xd2, 0x94, 0x2a, 0x54, 0xd1, 0x5a, 0xbd, 0xf4,
	0x62, 0x2d, 0x30, 0x8d, 0x90, 0xc0, 0x36, 0x36, 0x45, 0xe2, 0x11, 0xfa, 0x44, 0x7d, 0xbd, 0x6a,
	0xd7, 0xfc, 0xf1, 0x56, 0x95, 0x72, 0x9b, 0xf9, 0x34, 0xf3, 0xed, 0xcc, 0x4f, 0xb3, 0x80, 0x93,
	0xe8, 0x3e, 0x7e, 0x5c, 0xaa, 0x1d, 0xe7, 0x57, 0x59, 0x9e, 0x6e, 0x52, 0x82, 0x93, 0xe2, 0xaf,
	0xa1, 0x1b, 0xa8, 0x82, 0xcb, 0x6c, 0x9c, 0xfc, 0x4c, 0x63, 0xc9, 0x19, 0xab, 0x0d, 0x11, 0xb8,
	0xdf, 0xd4, 0x8a, 0x7b, 0xa2, 0x2f, 0x06, 0x2d, 0x69, 0x62, 0xea, 0x40, 0xfd, 0x81, 0xb7, 0xbc,
	0xec, 0xd5, 0xfa, 0x62, 0x50, 0x97, 0x65, 0x42, 0xaf, 0xa1, 0xf9, 0x85, 0xd5, 0x7c, 0x3c, 0x4b,
	0x93, 0x9e, 0x63, 0xaa, 0x8f, 0xb9, 0x76, 0xb9, 0x0d, 0xc6, 0xb7, 0x3d, 0xb7, 0x2f, 0x06, 0x8e,
	0x34, 0xb1, 0xff, 0x0e, 0x50, 0x3f, 0x39, 0x49, 0x13, 0xde, 0x55, 0x5e, 0x0b, 0xd3, 0x45, 0x62,
	0x5e, 0x73, 0xa4, 0x89, 0xfd, 0x0e, 0x50, 0x18, 0xc5, 0x77, 0xc9, 0x86, 0xf3, 0x88, 0xf3, 0x2d,
	0xe7, 0xb1, 0xe4, 0xb5, 0x7f, 0x03, 0x14, 0x85, 0xb6, 0x5a, 0x64, 0xf4, 0x1e, 0x1c, 0xc9, 0x1b,
	0xd3, 0x7e, 0x31, 0xea, 0x5e, 0x55, 0x56, 0xbe, 0xcb, 0xf3, 0x34, 0x0f, 0xd3, 0x39, 0x4b, 0x5d,
	0xb1, 0x37, 0x7d, 0x60, 0xb5, 0xe5, 0x8a, 0x69, 0xc7, 0x98, 0x5a, 0x6a, 0x91, 0xf9, 0x97, 0xd0,
	0x0e, 0xa3, 0xd8, 0x42, 0xb3, 0xf6, 0xff, 0x08, 0x68, 0x47, 0xa1, 0xa5, 0x16, 0x19, 0xdd, 0x40,
	0x73, 0xaa, 0x0a, 0xd6, 0xb9, 0x19, 0xc2, 0x1b, 0xbd, 0xad, 0x0e, 0xf1, 0x5f, 0xc4, 0xf2, 0xd8,
	0x42, 0x9f, 0xa0, 0x35, 0x3d, 0x20, 0x31, 0x70, 0xbd, 0xd1, 0x9b, 0x7f, 0xfb, 0xab, 0xbc, 0xe4,
	0xa9, 0xfc, 0xb0, 0xba, 0xf3, 0xdc, 0xea, 0xc3, 0xdf, 0x02, 0x1a, 0xd1, 0xf7, 0x60, 0x12, 0xdd,
	0x53, 0x0b, 0xea, 0x01, 0x3f, 0x2d, 0x12, 0x3c, 0x23, 0x82, 0x0b, 0x9b, 0x32, 0x0a, 0xad, 0xd9,
	0x8c, 0xb1, 0xb6, 0xaf, 0xab, 0x20, 0x42, 0x67, 0x5f, 0x57, 0xd5, 0x5c, 0x6a, 0xc3, 0x4b, 0x0b,
	0x1a, 0xd6, 0xb5, 0x64, 0x11, 0xc3, 0xc6, 0xf0, 0x1a, 0x5a, 0xc7, 0xe9, 0xc8, 0x83, 0x17, 0xe3,
	0x64, 0xab, 0x96, 0x8b, 0x39, 0x9e, 0xe9, 0x24, 0xfa, 0x35, 0x9b, 0x71, 0x51, 0xa0, 0xa0, 0x26,
	0xb8, 0x9f, 0xd5, 0x62, 0x89, 0xb5, 0xe1, 0x57, 0x00, 0x5e, 0x05, 0x07, 0x5e, 0xaf, 0xc0, 0x2b,
	0xd9, 0x1d, 0xba, 0x9a, 0xe5, 0xb5, 0xa2, 0xd0, 0xab, 0x99, 0xb3, 0xc4, 0x1a, 0x9d, 0x9f, 0x0e,
	0x13, 0x1d, 0x5d, 0xa2, 0xcf, 0x0f, 0xdd, 0xe1, 0x07, 0xf0, 0x4a, 0xaf, 0x12, 0x20, 0xc2, 0xf9,
	0x4a, 0x07, 0x96, 0x9b, 0xbe, 0x40, 0x14, 0x41, 0xf7, 0xc7, 0xe5, 0xaa, 0x78, 0x7a, 0xd4, 0xdf,
	0xe6, 0xfa, 0x44, 0x76, 0xda, 0x30, 0x1f, 0xe9, 0xe3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd,
	0xd0, 0x25, 0xf7, 0x5c, 0x03, 0x00, 0x00,
}
