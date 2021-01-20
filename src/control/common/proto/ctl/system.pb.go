// Code generated by protoc-gen-go. DO NOT EDIT.
// source: system.proto

package ctl

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

// SystemMember refers to a data-plane instance that is a member of DAOS
// system running on host with the control-plane listening at "Addr".
type SystemMember struct {
	Addr  string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Uuid  string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Rank  uint32 `protobuf:"varint,3,opt,name=rank,proto3" json:"rank,omitempty"`
	State uint32 `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`
	// ancillary info e.g. error msg or reason for state change
	Info                 string   `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	FaultDomain          string   `protobuf:"bytes,6,opt,name=faultDomain,proto3" json:"faultDomain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemMember) Reset()         { *m = SystemMember{} }
func (m *SystemMember) String() string { return proto.CompactTextString(m) }
func (*SystemMember) ProtoMessage()    {}
func (*SystemMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_86a7260ebdc12f47, []int{0}
}

func (m *SystemMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemMember.Unmarshal(m, b)
}
func (m *SystemMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemMember.Marshal(b, m, deterministic)
}
func (m *SystemMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemMember.Merge(m, src)
}
func (m *SystemMember) XXX_Size() int {
	return xxx_messageInfo_SystemMember.Size(m)
}
func (m *SystemMember) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemMember.DiscardUnknown(m)
}

var xxx_messageInfo_SystemMember proto.InternalMessageInfo

func (m *SystemMember) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *SystemMember) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *SystemMember) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *SystemMember) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *SystemMember) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *SystemMember) GetFaultDomain() string {
	if m != nil {
		return m.FaultDomain
	}
	return ""
}

// RankResult is a generic result for a system operation on a rank.
// Identical to mgmt.proto RanksResp_RankResult.
type RankResult struct {
	Rank                 uint32   `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Action               string   `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Errored              bool     `protobuf:"varint,3,opt,name=errored,proto3" json:"errored,omitempty"`
	Msg                  string   `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	State                uint32   `protobuf:"varint,5,opt,name=state,proto3" json:"state,omitempty"`
	Addr                 string   `protobuf:"bytes,6,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankResult) Reset()         { *m = RankResult{} }
func (m *RankResult) String() string { return proto.CompactTextString(m) }
func (*RankResult) ProtoMessage()    {}
func (*RankResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_86a7260ebdc12f47, []int{1}
}

func (m *RankResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankResult.Unmarshal(m, b)
}
func (m *RankResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankResult.Marshal(b, m, deterministic)
}
func (m *RankResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankResult.Merge(m, src)
}
func (m *RankResult) XXX_Size() int {
	return xxx_messageInfo_RankResult.Size(m)
}
func (m *RankResult) XXX_DiscardUnknown() {
	xxx_messageInfo_RankResult.DiscardUnknown(m)
}

var xxx_messageInfo_RankResult proto.InternalMessageInfo

func (m *RankResult) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *RankResult) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *RankResult) GetErrored() bool {
	if m != nil {
		return m.Errored
	}
	return false
}

func (m *RankResult) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RankResult) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *RankResult) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

// SystemStopReq supplies system shutdown parameters.
type SystemStopReq struct {
	Prep                 bool     `protobuf:"varint,1,opt,name=prep,proto3" json:"prep,omitempty"`
	Kill                 bool     `protobuf:"varint,2,opt,name=kill,proto3" json:"kill,omitempty"`
	Force                bool     `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
	Ranks                string   `protobuf:"bytes,5,opt,name=ranks,proto3" json:"ranks,omitempty"`
	Hosts                string   `protobuf:"bytes,6,opt,name=hosts,proto3" json:"hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemStopReq) Reset()         { *m = SystemStopReq{} }
func (m *SystemStopReq) String() string { return proto.CompactTextString(m) }
func (*SystemStopReq) ProtoMessage()    {}
func (*SystemStopReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_86a7260ebdc12f47, []int{2}
}

func (m *SystemStopReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemStopReq.Unmarshal(m, b)
}
func (m *SystemStopReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemStopReq.Marshal(b, m, deterministic)
}
func (m *SystemStopReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemStopReq.Merge(m, src)
}
func (m *SystemStopReq) XXX_Size() int {
	return xxx_messageInfo_SystemStopReq.Size(m)
}
func (m *SystemStopReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemStopReq.DiscardUnknown(m)
}

var xxx_messageInfo_SystemStopReq proto.InternalMessageInfo

func (m *SystemStopReq) GetPrep() bool {
	if m != nil {
		return m.Prep
	}
	return false
}

func (m *SystemStopReq) GetKill() bool {
	if m != nil {
		return m.Kill
	}
	return false
}

func (m *SystemStopReq) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

func (m *SystemStopReq) GetRanks() string {
	if m != nil {
		return m.Ranks
	}
	return ""
}

func (m *SystemStopReq) GetHosts() string {
	if m != nil {
		return m.Hosts
	}
	return ""
}

// SystemStopResp returns status of shutdown attempt and results
// of attempts to stop system members.
type SystemStopResp struct {
	Results              []*RankResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Absentranks          string        `protobuf:"bytes,2,opt,name=absentranks,proto3" json:"absentranks,omitempty"`
	Absenthosts          string        `protobuf:"bytes,3,opt,name=absenthosts,proto3" json:"absenthosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SystemStopResp) Reset()         { *m = SystemStopResp{} }
func (m *SystemStopResp) String() string { return proto.CompactTextString(m) }
func (*SystemStopResp) ProtoMessage()    {}
func (*SystemStopResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_86a7260ebdc12f47, []int{3}
}

func (m *SystemStopResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemStopResp.Unmarshal(m, b)
}
func (m *SystemStopResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemStopResp.Marshal(b, m, deterministic)
}
func (m *SystemStopResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemStopResp.Merge(m, src)
}
func (m *SystemStopResp) XXX_Size() int {
	return xxx_messageInfo_SystemStopResp.Size(m)
}
func (m *SystemStopResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemStopResp.DiscardUnknown(m)
}

var xxx_messageInfo_SystemStopResp proto.InternalMessageInfo

func (m *SystemStopResp) GetResults() []*RankResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *SystemStopResp) GetAbsentranks() string {
	if m != nil {
		return m.Absentranks
	}
	return ""
}

func (m *SystemStopResp) GetAbsenthosts() string {
	if m != nil {
		return m.Absenthosts
	}
	return ""
}

// SystemResetFormatReq supplies system reset format parameters.
type SystemResetFormatReq struct {
	Ranks                string   `protobuf:"bytes,2,opt,name=ranks,proto3" json:"ranks,omitempty"`
	Hosts                string   `protobuf:"bytes,3,opt,name=hosts,proto3" json:"hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemResetFormatReq) Reset()         { *m = SystemResetFormatReq{} }
func (m *SystemResetFormatReq) String() string { return proto.CompactTextString(m) }
func (*SystemResetFormatReq) ProtoMessage()    {}
func (*SystemResetFormatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_86a7260ebdc12f47, []int{4}
}

func (m *SystemResetFormatReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemResetFormatReq.Unmarshal(m, b)
}
func (m *SystemResetFormatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemResetFormatReq.Marshal(b, m, deterministic)
}
func (m *SystemResetFormatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemResetFormatReq.Merge(m, src)
}
func (m *SystemResetFormatReq) XXX_Size() int {
	return xxx_messageInfo_SystemResetFormatReq.Size(m)
}
func (m *SystemResetFormatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemResetFormatReq.DiscardUnknown(m)
}

var xxx_messageInfo_SystemResetFormatReq proto.InternalMessageInfo

func (m *SystemResetFormatReq) GetRanks() string {
	if m != nil {
		return m.Ranks
	}
	return ""
}

func (m *SystemResetFormatReq) GetHosts() string {
	if m != nil {
		return m.Hosts
	}
	return ""
}

// SystemResetFormatResp returns status of reset format attempt and results
// of attempts to reset format of system members.
type SystemResetFormatResp struct {
	Results              []*RankResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Absentranks          string        `protobuf:"bytes,2,opt,name=absentranks,proto3" json:"absentranks,omitempty"`
	Absenthosts          string        `protobuf:"bytes,3,opt,name=absenthosts,proto3" json:"absenthosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SystemResetFormatResp) Reset()         { *m = SystemResetFormatResp{} }
func (m *SystemResetFormatResp) String() string { return proto.CompactTextString(m) }
func (*SystemResetFormatResp) ProtoMessage()    {}
func (*SystemResetFormatResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_86a7260ebdc12f47, []int{5}
}

func (m *SystemResetFormatResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemResetFormatResp.Unmarshal(m, b)
}
func (m *SystemResetFormatResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemResetFormatResp.Marshal(b, m, deterministic)
}
func (m *SystemResetFormatResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemResetFormatResp.Merge(m, src)
}
func (m *SystemResetFormatResp) XXX_Size() int {
	return xxx_messageInfo_SystemResetFormatResp.Size(m)
}
func (m *SystemResetFormatResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemResetFormatResp.DiscardUnknown(m)
}

var xxx_messageInfo_SystemResetFormatResp proto.InternalMessageInfo

func (m *SystemResetFormatResp) GetResults() []*RankResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *SystemResetFormatResp) GetAbsentranks() string {
	if m != nil {
		return m.Absentranks
	}
	return ""
}

func (m *SystemResetFormatResp) GetAbsenthosts() string {
	if m != nil {
		return m.Absenthosts
	}
	return ""
}

// SystemStartReq supplies system restart parameters.
type SystemStartReq struct {
	Ranks                string   `protobuf:"bytes,2,opt,name=ranks,proto3" json:"ranks,omitempty"`
	Hosts                string   `protobuf:"bytes,3,opt,name=hosts,proto3" json:"hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemStartReq) Reset()         { *m = SystemStartReq{} }
func (m *SystemStartReq) String() string { return proto.CompactTextString(m) }
func (*SystemStartReq) ProtoMessage()    {}
func (*SystemStartReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_86a7260ebdc12f47, []int{6}
}

func (m *SystemStartReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemStartReq.Unmarshal(m, b)
}
func (m *SystemStartReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemStartReq.Marshal(b, m, deterministic)
}
func (m *SystemStartReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemStartReq.Merge(m, src)
}
func (m *SystemStartReq) XXX_Size() int {
	return xxx_messageInfo_SystemStartReq.Size(m)
}
func (m *SystemStartReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemStartReq.DiscardUnknown(m)
}

var xxx_messageInfo_SystemStartReq proto.InternalMessageInfo

func (m *SystemStartReq) GetRanks() string {
	if m != nil {
		return m.Ranks
	}
	return ""
}

func (m *SystemStartReq) GetHosts() string {
	if m != nil {
		return m.Hosts
	}
	return ""
}

// SystemStartResp returns status of restart attempt and results
// of attempts to start system members.
type SystemStartResp struct {
	Results              []*RankResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Absentranks          string        `protobuf:"bytes,2,opt,name=absentranks,proto3" json:"absentranks,omitempty"`
	Absenthosts          string        `protobuf:"bytes,3,opt,name=absenthosts,proto3" json:"absenthosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SystemStartResp) Reset()         { *m = SystemStartResp{} }
func (m *SystemStartResp) String() string { return proto.CompactTextString(m) }
func (*SystemStartResp) ProtoMessage()    {}
func (*SystemStartResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_86a7260ebdc12f47, []int{7}
}

func (m *SystemStartResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemStartResp.Unmarshal(m, b)
}
func (m *SystemStartResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemStartResp.Marshal(b, m, deterministic)
}
func (m *SystemStartResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemStartResp.Merge(m, src)
}
func (m *SystemStartResp) XXX_Size() int {
	return xxx_messageInfo_SystemStartResp.Size(m)
}
func (m *SystemStartResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemStartResp.DiscardUnknown(m)
}

var xxx_messageInfo_SystemStartResp proto.InternalMessageInfo

func (m *SystemStartResp) GetResults() []*RankResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *SystemStartResp) GetAbsentranks() string {
	if m != nil {
		return m.Absentranks
	}
	return ""
}

func (m *SystemStartResp) GetAbsenthosts() string {
	if m != nil {
		return m.Absenthosts
	}
	return ""
}

// SystemQueryReq supplies system query parameters.
type SystemQueryReq struct {
	Ranks                string   `protobuf:"bytes,2,opt,name=ranks,proto3" json:"ranks,omitempty"`
	Hosts                string   `protobuf:"bytes,3,opt,name=hosts,proto3" json:"hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemQueryReq) Reset()         { *m = SystemQueryReq{} }
func (m *SystemQueryReq) String() string { return proto.CompactTextString(m) }
func (*SystemQueryReq) ProtoMessage()    {}
func (*SystemQueryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_86a7260ebdc12f47, []int{8}
}

func (m *SystemQueryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemQueryReq.Unmarshal(m, b)
}
func (m *SystemQueryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemQueryReq.Marshal(b, m, deterministic)
}
func (m *SystemQueryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemQueryReq.Merge(m, src)
}
func (m *SystemQueryReq) XXX_Size() int {
	return xxx_messageInfo_SystemQueryReq.Size(m)
}
func (m *SystemQueryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemQueryReq.DiscardUnknown(m)
}

var xxx_messageInfo_SystemQueryReq proto.InternalMessageInfo

func (m *SystemQueryReq) GetRanks() string {
	if m != nil {
		return m.Ranks
	}
	return ""
}

func (m *SystemQueryReq) GetHosts() string {
	if m != nil {
		return m.Hosts
	}
	return ""
}

// SystemQueryResp returns active system members.
type SystemQueryResp struct {
	Members              []*SystemMember `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	Absentranks          string          `protobuf:"bytes,2,opt,name=absentranks,proto3" json:"absentranks,omitempty"`
	Absenthosts          string          `protobuf:"bytes,3,opt,name=absenthosts,proto3" json:"absenthosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SystemQueryResp) Reset()         { *m = SystemQueryResp{} }
func (m *SystemQueryResp) String() string { return proto.CompactTextString(m) }
func (*SystemQueryResp) ProtoMessage()    {}
func (*SystemQueryResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_86a7260ebdc12f47, []int{9}
}

func (m *SystemQueryResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemQueryResp.Unmarshal(m, b)
}
func (m *SystemQueryResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemQueryResp.Marshal(b, m, deterministic)
}
func (m *SystemQueryResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemQueryResp.Merge(m, src)
}
func (m *SystemQueryResp) XXX_Size() int {
	return xxx_messageInfo_SystemQueryResp.Size(m)
}
func (m *SystemQueryResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemQueryResp.DiscardUnknown(m)
}

var xxx_messageInfo_SystemQueryResp proto.InternalMessageInfo

func (m *SystemQueryResp) GetMembers() []*SystemMember {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *SystemQueryResp) GetAbsentranks() string {
	if m != nil {
		return m.Absentranks
	}
	return ""
}

func (m *SystemQueryResp) GetAbsenthosts() string {
	if m != nil {
		return m.Absenthosts
	}
	return ""
}

func init() {
	proto.RegisterType((*SystemMember)(nil), "ctl.SystemMember")
	proto.RegisterType((*RankResult)(nil), "ctl.RankResult")
	proto.RegisterType((*SystemStopReq)(nil), "ctl.SystemStopReq")
	proto.RegisterType((*SystemStopResp)(nil), "ctl.SystemStopResp")
	proto.RegisterType((*SystemResetFormatReq)(nil), "ctl.SystemResetFormatReq")
	proto.RegisterType((*SystemResetFormatResp)(nil), "ctl.SystemResetFormatResp")
	proto.RegisterType((*SystemStartReq)(nil), "ctl.SystemStartReq")
	proto.RegisterType((*SystemStartResp)(nil), "ctl.SystemStartResp")
	proto.RegisterType((*SystemQueryReq)(nil), "ctl.SystemQueryReq")
	proto.RegisterType((*SystemQueryResp)(nil), "ctl.SystemQueryResp")
}

func init() {
	proto.RegisterFile("system.proto", fileDescriptor_86a7260ebdc12f47)
}

var fileDescriptor_86a7260ebdc12f47 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0x3d, 0x6f, 0xdb, 0x30,
	0x10, 0x05, 0x2d, 0x5b, 0xb6, 0xcf, 0x76, 0xdd, 0x12, 0x6e, 0xc1, 0x51, 0xd0, 0xe4, 0xa2, 0x80,
	0x87, 0x76, 0xed, 0x54, 0x14, 0xdd, 0x3a, 0x94, 0xfe, 0x05, 0xb4, 0x44, 0x37, 0x82, 0x25, 0x51,
	0x21, 0xa9, 0xc1, 0x53, 0x82, 0x0c, 0x19, 0x33, 0xe6, 0xf7, 0x06, 0x3c, 0x4a, 0x96, 0x12, 0x78,
	0x31, 0x10, 0x24, 0xdb, 0xbb, 0xc7, 0xa3, 0xde, 0x07, 0x20, 0xc2, 0xdc, 0x1c, 0x8d, 0x95, 0xc5,
	0xa6, 0xd2, 0xca, 0x2a, 0x1a, 0x24, 0x36, 0x8f, 0x1f, 0x09, 0xcc, 0xb7, 0xc8, 0xfe, 0x95, 0xc5,
	0x4e, 0x6a, 0x4a, 0x61, 0x28, 0xd2, 0x54, 0x33, 0x12, 0x91, 0xf5, 0x94, 0x23, 0x76, 0x5c, 0x5d,
	0x67, 0x29, 0x1b, 0x78, 0xce, 0x61, 0xc7, 0x69, 0x51, 0x1e, 0x58, 0x10, 0x91, 0xf5, 0x82, 0x23,
	0xa6, 0x2b, 0x18, 0x19, 0x2b, 0xac, 0x64, 0x43, 0x24, 0xfd, 0xe0, 0x36, 0xb3, 0x72, 0xaf, 0xd8,
	0xc8, 0xdf, 0x76, 0x98, 0x46, 0x30, 0xdb, 0x8b, 0x3a, 0xb7, 0xbf, 0x55, 0x21, 0xb2, 0x92, 0x85,
	0x78, 0xd4, 0xa7, 0xe2, 0x07, 0x02, 0xc0, 0x45, 0x79, 0xe0, 0xd2, 0xd4, 0xb9, 0x3d, 0xc9, 0x91,
	0x9e, 0xdc, 0x17, 0x08, 0x45, 0x62, 0x33, 0x55, 0x36, 0xc6, 0x9a, 0x89, 0x32, 0x18, 0x4b, 0xad,
	0x95, 0x96, 0x29, 0xba, 0x9b, 0xf0, 0x76, 0xa4, 0x1f, 0x21, 0x28, 0xcc, 0x7f, 0xb4, 0x37, 0xe5,
	0x0e, 0x76, 0x96, 0x47, 0x2f, 0x2c, 0x63, 0x09, 0x61, 0x57, 0x42, 0x7c, 0x84, 0x85, 0x2f, 0x6a,
	0x6b, 0x55, 0xc5, 0xe5, 0xb5, 0x5b, 0xaa, 0xb4, 0xac, 0xd0, 0xd2, 0x84, 0x23, 0x76, 0xdc, 0x21,
	0xcb, 0x73, 0x34, 0x34, 0xe1, 0x88, 0x9d, 0xc4, 0x5e, 0xe9, 0x44, 0x36, 0x66, 0xfc, 0xe0, 0x58,
	0x17, 0xc2, 0x34, 0xb5, 0xf8, 0xc1, 0xb1, 0x57, 0xca, 0x58, 0xd3, 0x28, 0xfb, 0x21, 0xbe, 0x81,
	0x0f, 0x7d, 0x69, 0x53, 0xd1, 0xaf, 0x30, 0xd6, 0x58, 0x8c, 0x61, 0x24, 0x0a, 0xd6, 0xb3, 0xef,
	0xcb, 0x4d, 0x62, 0xf3, 0x4d, 0x57, 0x18, 0x6f, 0xcf, 0x5d, 0xd5, 0x62, 0x67, 0x64, 0x69, 0xbd,
	0x9c, 0xaf, 0xaa, 0x4f, 0x75, 0x1b, 0x5e, 0x3a, 0xe8, 0x6f, 0x78, 0x03, 0xbf, 0x60, 0xe5, 0x0d,
	0x70, 0x69, 0xa4, 0xfd, 0xa3, 0x74, 0x21, 0xac, 0xab, 0xe0, 0x14, 0x62, 0x70, 0x36, 0x44, 0xd0,
	0x0f, 0x71, 0x4f, 0xe0, 0xf3, 0x99, 0x8f, 0xbc, 0x7d, 0x98, 0x9f, 0x5d, 0x9b, 0x42, 0x5f, 0x1c,
	0xe3, 0x96, 0xc0, 0xf2, 0xd9, 0xf5, 0x77, 0x0c, 0xf0, 0xaf, 0x96, 0xfa, 0x78, 0x69, 0x80, 0xbb,
	0x53, 0x80, 0xe6, 0xba, 0xa9, 0xe8, 0x37, 0x18, 0x17, 0xf8, 0xfb, 0xb7, 0x01, 0x3e, 0x61, 0x80,
	0xfe, 0xc3, 0xc0, 0xdb, 0x8d, 0xd7, 0x88, 0xb0, 0x0b, 0xf1, 0x09, 0xfa, 0xf1, 0x14, 0x00, 0x00,
	0xff, 0xff, 0x15, 0xf5, 0xf5, 0x90, 0x92, 0x04, 0x00, 0x00,
}