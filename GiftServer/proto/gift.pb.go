// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gift.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Manager struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Manager) Reset()         { *m = Manager{} }
func (m *Manager) String() string { return proto.CompactTextString(m) }
func (*Manager) ProtoMessage()    {}
func (*Manager) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a02b901f1be8e63, []int{0}
}

func (m *Manager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manager.Unmarshal(m, b)
}
func (m *Manager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manager.Marshal(b, m, deterministic)
}
func (m *Manager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manager.Merge(m, src)
}
func (m *Manager) XXX_Size() int {
	return xxx_messageInfo_Manager.Size(m)
}
func (m *Manager) XXX_DiscardUnknown() {
	xxx_messageInfo_Manager.DiscardUnknown(m)
}

var xxx_messageInfo_Manager proto.InternalMessageInfo

type Manager_Nil struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Manager_Nil) Reset()         { *m = Manager_Nil{} }
func (m *Manager_Nil) String() string { return proto.CompactTextString(m) }
func (*Manager_Nil) ProtoMessage()    {}
func (*Manager_Nil) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a02b901f1be8e63, []int{0, 0}
}

func (m *Manager_Nil) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manager_Nil.Unmarshal(m, b)
}
func (m *Manager_Nil) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manager_Nil.Marshal(b, m, deterministic)
}
func (m *Manager_Nil) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manager_Nil.Merge(m, src)
}
func (m *Manager_Nil) XXX_Size() int {
	return xxx_messageInfo_Manager_Nil.Size(m)
}
func (m *Manager_Nil) XXX_DiscardUnknown() {
	xxx_messageInfo_Manager_Nil.DiscardUnknown(m)
}

var xxx_messageInfo_Manager_Nil proto.InternalMessageInfo

type Manager_Item struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Num                  int32    `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Manager_Item) Reset()         { *m = Manager_Item{} }
func (m *Manager_Item) String() string { return proto.CompactTextString(m) }
func (*Manager_Item) ProtoMessage()    {}
func (*Manager_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a02b901f1be8e63, []int{0, 1}
}

func (m *Manager_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manager_Item.Unmarshal(m, b)
}
func (m *Manager_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manager_Item.Marshal(b, m, deterministic)
}
func (m *Manager_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manager_Item.Merge(m, src)
}
func (m *Manager_Item) XXX_Size() int {
	return xxx_messageInfo_Manager_Item.Size(m)
}
func (m *Manager_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Manager_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Manager_Item proto.InternalMessageInfo

func (m *Manager_Item) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Manager_Item) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type Manager_GenReq struct {
	FixCode              string          `protobuf:"bytes,1,opt,name=FixCode,proto3" json:"FixCode,omitempty"`
	Num                  uint32          `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
	StartTime            int64           `protobuf:"varint,3,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime              int64           `protobuf:"varint,4,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	TimesPerCode         uint32          `protobuf:"varint,5,opt,name=TimesPerCode,proto3" json:"TimesPerCode,omitempty"`
	TimesPerUser         uint32          `protobuf:"varint,6,opt,name=TimesPerUser,proto3" json:"TimesPerUser,omitempty"`
	ZoneIds              []uint32        `protobuf:"varint,7,rep,packed,name=ZoneIds,proto3" json:"ZoneIds,omitempty"`
	Items                []*Manager_Item `protobuf:"bytes,8,rep,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Manager_GenReq) Reset()         { *m = Manager_GenReq{} }
func (m *Manager_GenReq) String() string { return proto.CompactTextString(m) }
func (*Manager_GenReq) ProtoMessage()    {}
func (*Manager_GenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a02b901f1be8e63, []int{0, 2}
}

func (m *Manager_GenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manager_GenReq.Unmarshal(m, b)
}
func (m *Manager_GenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manager_GenReq.Marshal(b, m, deterministic)
}
func (m *Manager_GenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manager_GenReq.Merge(m, src)
}
func (m *Manager_GenReq) XXX_Size() int {
	return xxx_messageInfo_Manager_GenReq.Size(m)
}
func (m *Manager_GenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_Manager_GenReq.DiscardUnknown(m)
}

var xxx_messageInfo_Manager_GenReq proto.InternalMessageInfo

func (m *Manager_GenReq) GetFixCode() string {
	if m != nil {
		return m.FixCode
	}
	return ""
}

func (m *Manager_GenReq) GetNum() uint32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *Manager_GenReq) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Manager_GenReq) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *Manager_GenReq) GetTimesPerCode() uint32 {
	if m != nil {
		return m.TimesPerCode
	}
	return 0
}

func (m *Manager_GenReq) GetTimesPerUser() uint32 {
	if m != nil {
		return m.TimesPerUser
	}
	return 0
}

func (m *Manager_GenReq) GetZoneIds() []uint32 {
	if m != nil {
		return m.ZoneIds
	}
	return nil
}

func (m *Manager_GenReq) GetItems() []*Manager_Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type Manager_ExportReq struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Manager_ExportReq) Reset()         { *m = Manager_ExportReq{} }
func (m *Manager_ExportReq) String() string { return proto.CompactTextString(m) }
func (*Manager_ExportReq) ProtoMessage()    {}
func (*Manager_ExportReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a02b901f1be8e63, []int{0, 3}
}

func (m *Manager_ExportReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manager_ExportReq.Unmarshal(m, b)
}
func (m *Manager_ExportReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manager_ExportReq.Marshal(b, m, deterministic)
}
func (m *Manager_ExportReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manager_ExportReq.Merge(m, src)
}
func (m *Manager_ExportReq) XXX_Size() int {
	return xxx_messageInfo_Manager_ExportReq.Size(m)
}
func (m *Manager_ExportReq) XXX_DiscardUnknown() {
	xxx_messageInfo_Manager_ExportReq.DiscardUnknown(m)
}

var xxx_messageInfo_Manager_ExportReq proto.InternalMessageInfo

func (m *Manager_ExportReq) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Manager_CodeStatus struct {
	Code                 string   `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ZoneId               uint32   `protobuf:"varint,3,opt,name=ZoneId,proto3" json:"ZoneId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Manager_CodeStatus) Reset()         { *m = Manager_CodeStatus{} }
func (m *Manager_CodeStatus) String() string { return proto.CompactTextString(m) }
func (*Manager_CodeStatus) ProtoMessage()    {}
func (*Manager_CodeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a02b901f1be8e63, []int{0, 4}
}

func (m *Manager_CodeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manager_CodeStatus.Unmarshal(m, b)
}
func (m *Manager_CodeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manager_CodeStatus.Marshal(b, m, deterministic)
}
func (m *Manager_CodeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manager_CodeStatus.Merge(m, src)
}
func (m *Manager_CodeStatus) XXX_Size() int {
	return xxx_messageInfo_Manager_CodeStatus.Size(m)
}
func (m *Manager_CodeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_Manager_CodeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_Manager_CodeStatus proto.InternalMessageInfo

func (m *Manager_CodeStatus) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Manager_CodeStatus) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Manager_CodeStatus) GetZoneId() uint32 {
	if m != nil {
		return m.ZoneId
	}
	return 0
}

type Manager_ExportResp struct {
	Details              []*Manager_CodeStatus `protobuf:"bytes,1,rep,name=Details,proto3" json:"Details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Manager_ExportResp) Reset()         { *m = Manager_ExportResp{} }
func (m *Manager_ExportResp) String() string { return proto.CompactTextString(m) }
func (*Manager_ExportResp) ProtoMessage()    {}
func (*Manager_ExportResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a02b901f1be8e63, []int{0, 5}
}

func (m *Manager_ExportResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manager_ExportResp.Unmarshal(m, b)
}
func (m *Manager_ExportResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manager_ExportResp.Marshal(b, m, deterministic)
}
func (m *Manager_ExportResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manager_ExportResp.Merge(m, src)
}
func (m *Manager_ExportResp) XXX_Size() int {
	return xxx_messageInfo_Manager_ExportResp.Size(m)
}
func (m *Manager_ExportResp) XXX_DiscardUnknown() {
	xxx_messageInfo_Manager_ExportResp.DiscardUnknown(m)
}

var xxx_messageInfo_Manager_ExportResp proto.InternalMessageInfo

func (m *Manager_ExportResp) GetDetails() []*Manager_CodeStatus {
	if m != nil {
		return m.Details
	}
	return nil
}

type Manager_CodeInfo struct {
	Id                   uint32          `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Used                 uint32          `protobuf:"varint,2,opt,name=Used,proto3" json:"Used,omitempty"`
	GenInfo              *Manager_GenReq `protobuf:"bytes,3,opt,name=GenInfo,proto3" json:"GenInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Manager_CodeInfo) Reset()         { *m = Manager_CodeInfo{} }
func (m *Manager_CodeInfo) String() string { return proto.CompactTextString(m) }
func (*Manager_CodeInfo) ProtoMessage()    {}
func (*Manager_CodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a02b901f1be8e63, []int{0, 6}
}

func (m *Manager_CodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manager_CodeInfo.Unmarshal(m, b)
}
func (m *Manager_CodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manager_CodeInfo.Marshal(b, m, deterministic)
}
func (m *Manager_CodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manager_CodeInfo.Merge(m, src)
}
func (m *Manager_CodeInfo) XXX_Size() int {
	return xxx_messageInfo_Manager_CodeInfo.Size(m)
}
func (m *Manager_CodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Manager_CodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Manager_CodeInfo proto.InternalMessageInfo

func (m *Manager_CodeInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Manager_CodeInfo) GetUsed() uint32 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *Manager_CodeInfo) GetGenInfo() *Manager_GenReq {
	if m != nil {
		return m.GenInfo
	}
	return nil
}

type Manager_ListResp struct {
	Details              []*Manager_CodeInfo `protobuf:"bytes,1,rep,name=Details,proto3" json:"Details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Manager_ListResp) Reset()         { *m = Manager_ListResp{} }
func (m *Manager_ListResp) String() string { return proto.CompactTextString(m) }
func (*Manager_ListResp) ProtoMessage()    {}
func (*Manager_ListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a02b901f1be8e63, []int{0, 7}
}

func (m *Manager_ListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manager_ListResp.Unmarshal(m, b)
}
func (m *Manager_ListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manager_ListResp.Marshal(b, m, deterministic)
}
func (m *Manager_ListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manager_ListResp.Merge(m, src)
}
func (m *Manager_ListResp) XXX_Size() int {
	return xxx_messageInfo_Manager_ListResp.Size(m)
}
func (m *Manager_ListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_Manager_ListResp.DiscardUnknown(m)
}

var xxx_messageInfo_Manager_ListResp proto.InternalMessageInfo

func (m *Manager_ListResp) GetDetails() []*Manager_CodeInfo {
	if m != nil {
		return m.Details
	}
	return nil
}

type VerifyReq struct {
	Code                 string   `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Zone                 uint32   `protobuf:"varint,2,opt,name=Zone,proto3" json:"Zone,omitempty"`
	UserId               uint64   `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyReq) Reset()         { *m = VerifyReq{} }
func (m *VerifyReq) String() string { return proto.CompactTextString(m) }
func (*VerifyReq) ProtoMessage()    {}
func (*VerifyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a02b901f1be8e63, []int{1}
}

func (m *VerifyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyReq.Unmarshal(m, b)
}
func (m *VerifyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyReq.Marshal(b, m, deterministic)
}
func (m *VerifyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyReq.Merge(m, src)
}
func (m *VerifyReq) XXX_Size() int {
	return xxx_messageInfo_VerifyReq.Size(m)
}
func (m *VerifyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyReq.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyReq proto.InternalMessageInfo

func (m *VerifyReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *VerifyReq) GetZone() uint32 {
	if m != nil {
		return m.Zone
	}
	return 0
}

func (m *VerifyReq) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type VerifyResp struct {
	Status               uint32   `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyResp) Reset()         { *m = VerifyResp{} }
func (m *VerifyResp) String() string { return proto.CompactTextString(m) }
func (*VerifyResp) ProtoMessage()    {}
func (*VerifyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a02b901f1be8e63, []int{2}
}

func (m *VerifyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyResp.Unmarshal(m, b)
}
func (m *VerifyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyResp.Marshal(b, m, deterministic)
}
func (m *VerifyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyResp.Merge(m, src)
}
func (m *VerifyResp) XXX_Size() int {
	return xxx_messageInfo_VerifyResp.Size(m)
}
func (m *VerifyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyResp.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyResp proto.InternalMessageInfo

func (m *VerifyResp) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type SyncReq struct {
	Zone                 uint32   `protobuf:"varint,1,opt,name=Zone,proto3" json:"Zone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncReq) Reset()         { *m = SyncReq{} }
func (m *SyncReq) String() string { return proto.CompactTextString(m) }
func (*SyncReq) ProtoMessage()    {}
func (*SyncReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a02b901f1be8e63, []int{3}
}

func (m *SyncReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncReq.Unmarshal(m, b)
}
func (m *SyncReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncReq.Marshal(b, m, deterministic)
}
func (m *SyncReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncReq.Merge(m, src)
}
func (m *SyncReq) XXX_Size() int {
	return xxx_messageInfo_SyncReq.Size(m)
}
func (m *SyncReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncReq.DiscardUnknown(m)
}

var xxx_messageInfo_SyncReq proto.InternalMessageInfo

func (m *SyncReq) GetZone() uint32 {
	if m != nil {
		return m.Zone
	}
	return 0
}

func init() {
	proto.RegisterType((*Manager)(nil), "proto.Manager")
	proto.RegisterType((*Manager_Nil)(nil), "proto.Manager.Nil")
	proto.RegisterType((*Manager_Item)(nil), "proto.Manager.Item")
	proto.RegisterType((*Manager_GenReq)(nil), "proto.Manager.GenReq")
	proto.RegisterType((*Manager_ExportReq)(nil), "proto.Manager.ExportReq")
	proto.RegisterType((*Manager_CodeStatus)(nil), "proto.Manager.CodeStatus")
	proto.RegisterType((*Manager_ExportResp)(nil), "proto.Manager.ExportResp")
	proto.RegisterType((*Manager_CodeInfo)(nil), "proto.Manager.CodeInfo")
	proto.RegisterType((*Manager_ListResp)(nil), "proto.Manager.ListResp")
	proto.RegisterType((*VerifyReq)(nil), "proto.VerifyReq")
	proto.RegisterType((*VerifyResp)(nil), "proto.VerifyResp")
	proto.RegisterType((*SyncReq)(nil), "proto.SyncReq")
}

func init() { proto.RegisterFile("gift.proto", fileDescriptor_2a02b901f1be8e63) }

var fileDescriptor_2a02b901f1be8e63 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x63, 0x3b, 0x4e, 0xa6, 0xa4, 0x2a, 0x83, 0x00, 0x77, 0x01, 0x29, 0xb2, 0x38, 0x98,
	0x4b, 0x20, 0x89, 0x38, 0xf6, 0x80, 0xa0, 0x44, 0x16, 0x10, 0x55, 0x1b, 0xca, 0x81, 0x0b, 0x32,
	0xf5, 0xa6, 0x5a, 0x29, 0xb1, 0x8d, 0x77, 0x8b, 0xda, 0x37, 0xe0, 0x29, 0x78, 0x4f, 0x6e, 0x68,
	0xd6, 0xeb, 0x24, 0x75, 0xd3, 0xd3, 0xce, 0xcf, 0x37, 0x33, 0xdf, 0x7c, 0xb3, 0x00, 0x97, 0x72,
	0xa9, 0x47, 0x65, 0x55, 0xe8, 0x02, 0x7d, 0xf3, 0x44, 0x7f, 0x7c, 0x08, 0xbe, 0xa4, 0x79, 0x7a,
	0x29, 0x2a, 0xe6, 0x83, 0x3b, 0x97, 0x2b, 0x16, 0x83, 0x97, 0x68, 0xb1, 0xc6, 0x43, 0xe8, 0x24,
	0x59, 0xe8, 0x0c, 0x9d, 0x78, 0xc0, 0x3b, 0x49, 0x86, 0x47, 0xe0, 0xce, 0xaf, 0xd6, 0x61, 0x67,
	0xe8, 0xc4, 0x3e, 0x27, 0x93, 0xfd, 0x73, 0xa0, 0x3b, 0x13, 0x39, 0x17, 0xbf, 0x30, 0x84, 0xe0,
	0xa3, 0xbc, 0x7e, 0x5f, 0x64, 0xc2, 0x54, 0xf4, 0x79, 0xe3, 0xee, 0x96, 0x0d, 0x4c, 0x19, 0x3e,
	0x87, 0xfe, 0x42, 0xa7, 0x95, 0xfe, 0x2a, 0xd7, 0x22, 0x74, 0x87, 0x4e, 0xec, 0xf2, 0x6d, 0x80,
	0x3a, 0x9d, 0xe6, 0x99, 0xc9, 0x79, 0x26, 0xd7, 0xb8, 0x18, 0xc1, 0x03, 0x7a, 0xd5, 0x99, 0xa8,
	0xcc, 0x20, 0xdf, 0xb4, 0xbc, 0x15, 0xdb, 0xc5, 0x9c, 0x2b, 0x51, 0x85, 0xdd, 0xdb, 0x18, 0x8a,
	0xd1, 0x84, 0xef, 0x45, 0x2e, 0x92, 0x4c, 0x85, 0xc1, 0xd0, 0x8d, 0x07, 0xbc, 0x71, 0xf1, 0x15,
	0xf8, 0xb4, 0xba, 0x0a, 0x7b, 0x43, 0x37, 0x3e, 0x98, 0x3c, 0xaa, 0xb5, 0x1a, 0x59, 0x81, 0x46,
	0x94, 0xe3, 0x35, 0x82, 0x3d, 0x83, 0xfe, 0xe9, 0x75, 0x59, 0x54, 0x9a, 0xb6, 0x6f, 0x49, 0xc5,
	0xce, 0x00, 0x88, 0xcd, 0x42, 0xa7, 0xfa, 0x4a, 0x21, 0x82, 0xb7, 0x23, 0x8c, 0xb1, 0xf1, 0x09,
	0x74, 0x89, 0x4b, 0x92, 0x19, 0x61, 0x5c, 0x6e, 0x3d, 0x8a, 0xd7, 0x64, 0x8c, 0x30, 0x03, 0x6e,
	0x3d, 0xf6, 0x0e, 0xa0, 0x19, 0xa7, 0x4a, 0x9c, 0x42, 0xf0, 0x41, 0xe8, 0x54, 0xae, 0x54, 0xe8,
	0x18, 0xa6, 0xc7, 0x2d, 0xa6, 0xdb, 0xe9, 0xbc, 0x41, 0xb2, 0x1f, 0xd0, 0xa3, 0x70, 0x92, 0x2f,
	0x8b, 0x3b, 0xb7, 0x45, 0xf0, 0xce, 0x95, 0xc8, 0xec, 0x95, 0x8c, 0x8d, 0xaf, 0x21, 0x98, 0x89,
	0x9c, 0xe0, 0x86, 0xcb, 0xc1, 0xe4, 0x71, 0x6b, 0x48, 0x7d, 0x7a, 0xde, 0xa0, 0xd8, 0x09, 0xf4,
	0x3e, 0x4b, 0x55, 0x33, 0x1c, 0xb7, 0x19, 0x3e, 0xdd, 0xc3, 0x90, 0xaa, 0x36, 0xfc, 0xa2, 0x4f,
	0xd0, 0xff, 0x26, 0x2a, 0xb9, 0xbc, 0x21, 0x45, 0xf7, 0x69, 0x86, 0xe0, 0x91, 0x1a, 0x0d, 0x49,
	0xb2, 0x77, 0x74, 0x24, 0x8e, 0x5e, 0xa3, 0x63, 0xf4, 0x12, 0xa0, 0x69, 0xa6, 0x4a, 0x42, 0xd5,
	0x6a, 0xd8, 0x95, 0xad, 0x17, 0xbd, 0x80, 0x60, 0x71, 0x93, 0x5f, 0xd8, 0x81, 0xa6, 0xb9, 0xb3,
	0x6d, 0x3e, 0xf9, 0xdb, 0x81, 0x83, 0x99, 0x5c, 0xea, 0x85, 0xa8, 0x7e, 0xcb, 0x0b, 0x81, 0x63,
	0xf0, 0x08, 0x8e, 0x87, 0x76, 0x17, 0x5b, 0xcb, 0xee, 0xdb, 0xed, 0x8d, 0x83, 0xe3, 0xfa, 0x27,
	0xd4, 0x5c, 0xf0, 0xc8, 0x02, 0x37, 0x7b, 0xb2, 0x87, 0xad, 0x88, 0x2a, 0xf1, 0x2d, 0xf4, 0x66,
	0x22, 0x17, 0x55, 0xaa, 0x05, 0xee, 0x97, 0x9c, 0x61, 0x2b, 0x3c, 0x97, 0x2b, 0x9c, 0x82, 0x47,
	0xea, 0xe3, 0x9e, 0xdc, 0x1d, 0x82, 0x9b, 0x33, 0x9d, 0x40, 0xb7, 0xfe, 0x56, 0x18, 0xb6, 0x20,
	0x9b, 0xcf, 0xcd, 0x8e, 0xef, 0xc9, 0xa8, 0xf2, 0x67, 0xd7, 0x64, 0xa6, 0xff, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xdd, 0x35, 0x9a, 0x17, 0x59, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GiftServiceClient is the client API for GiftService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GiftServiceClient interface {
	Sync(ctx context.Context, in *SyncReq, opts ...grpc.CallOption) (GiftService_SyncClient, error)
	CodeVerify(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyResp, error)
	Generate(ctx context.Context, in *Manager_GenReq, opts ...grpc.CallOption) (*Manager_Nil, error)
	List(ctx context.Context, in *Manager_Nil, opts ...grpc.CallOption) (*Manager_ListResp, error)
	Export(ctx context.Context, in *Manager_ExportReq, opts ...grpc.CallOption) (*Manager_ExportResp, error)
}

type giftServiceClient struct {
	cc *grpc.ClientConn
}

func NewGiftServiceClient(cc *grpc.ClientConn) GiftServiceClient {
	return &giftServiceClient{cc}
}

func (c *giftServiceClient) Sync(ctx context.Context, in *SyncReq, opts ...grpc.CallOption) (GiftService_SyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GiftService_serviceDesc.Streams[0], "/proto.GiftService/Sync", opts...)
	if err != nil {
		return nil, err
	}
	x := &giftServiceSyncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GiftService_SyncClient interface {
	Recv() (*Manager_CodeInfo, error)
	grpc.ClientStream
}

type giftServiceSyncClient struct {
	grpc.ClientStream
}

func (x *giftServiceSyncClient) Recv() (*Manager_CodeInfo, error) {
	m := new(Manager_CodeInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *giftServiceClient) CodeVerify(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyResp, error) {
	out := new(VerifyResp)
	err := c.cc.Invoke(ctx, "/proto.GiftService/CodeVerify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftServiceClient) Generate(ctx context.Context, in *Manager_GenReq, opts ...grpc.CallOption) (*Manager_Nil, error) {
	out := new(Manager_Nil)
	err := c.cc.Invoke(ctx, "/proto.GiftService/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftServiceClient) List(ctx context.Context, in *Manager_Nil, opts ...grpc.CallOption) (*Manager_ListResp, error) {
	out := new(Manager_ListResp)
	err := c.cc.Invoke(ctx, "/proto.GiftService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftServiceClient) Export(ctx context.Context, in *Manager_ExportReq, opts ...grpc.CallOption) (*Manager_ExportResp, error) {
	out := new(Manager_ExportResp)
	err := c.cc.Invoke(ctx, "/proto.GiftService/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GiftServiceServer is the server API for GiftService service.
type GiftServiceServer interface {
	Sync(*SyncReq, GiftService_SyncServer) error
	CodeVerify(context.Context, *VerifyReq) (*VerifyResp, error)
	Generate(context.Context, *Manager_GenReq) (*Manager_Nil, error)
	List(context.Context, *Manager_Nil) (*Manager_ListResp, error)
	Export(context.Context, *Manager_ExportReq) (*Manager_ExportResp, error)
}

// UnimplementedGiftServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGiftServiceServer struct {
}

func (*UnimplementedGiftServiceServer) Sync(req *SyncReq, srv GiftService_SyncServer) error {
	return status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (*UnimplementedGiftServiceServer) CodeVerify(ctx context.Context, req *VerifyReq) (*VerifyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CodeVerify not implemented")
}
func (*UnimplementedGiftServiceServer) Generate(ctx context.Context, req *Manager_GenReq) (*Manager_Nil, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (*UnimplementedGiftServiceServer) List(ctx context.Context, req *Manager_Nil) (*Manager_ListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedGiftServiceServer) Export(ctx context.Context, req *Manager_ExportReq) (*Manager_ExportResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}

func RegisterGiftServiceServer(s *grpc.Server, srv GiftServiceServer) {
	s.RegisterService(&_GiftService_serviceDesc, srv)
}

func _GiftService_Sync_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyncReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GiftServiceServer).Sync(m, &giftServiceSyncServer{stream})
}

type GiftService_SyncServer interface {
	Send(*Manager_CodeInfo) error
	grpc.ServerStream
}

type giftServiceSyncServer struct {
	grpc.ServerStream
}

func (x *giftServiceSyncServer) Send(m *Manager_CodeInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _GiftService_CodeVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftServiceServer).CodeVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GiftService/CodeVerify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftServiceServer).CodeVerify(ctx, req.(*VerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GiftService_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Manager_GenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftServiceServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GiftService/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftServiceServer).Generate(ctx, req.(*Manager_GenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GiftService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Manager_Nil)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GiftService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftServiceServer).List(ctx, req.(*Manager_Nil))
	}
	return interceptor(ctx, in, info, handler)
}

func _GiftService_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Manager_ExportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftServiceServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GiftService/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftServiceServer).Export(ctx, req.(*Manager_ExportReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GiftService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GiftService",
	HandlerType: (*GiftServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CodeVerify",
			Handler:    _GiftService_CodeVerify_Handler,
		},
		{
			MethodName: "Generate",
			Handler:    _GiftService_Generate_Handler,
		},
		{
			MethodName: "List",
			Handler:    _GiftService_List_Handler,
		},
		{
			MethodName: "Export",
			Handler:    _GiftService_Export_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sync",
			Handler:       _GiftService_Sync_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gift.proto",
}