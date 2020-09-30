// Code generated by protoc-gen-go. DO NOT EDIT.
// source: geoip.proto

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

type Geo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Geo) Reset()         { *m = Geo{} }
func (m *Geo) String() string { return proto.CompactTextString(m) }
func (*Geo) ProtoMessage()    {}
func (*Geo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a5b7ee991490f54, []int{0}
}

func (m *Geo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Geo.Unmarshal(m, b)
}
func (m *Geo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Geo.Marshal(b, m, deterministic)
}
func (m *Geo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Geo.Merge(m, src)
}
func (m *Geo) XXX_Size() int {
	return xxx_messageInfo_Geo.Size(m)
}
func (m *Geo) XXX_DiscardUnknown() {
	xxx_messageInfo_Geo.DiscardUnknown(m)
}

var xxx_messageInfo_Geo proto.InternalMessageInfo

type Geo_IP struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Geo_IP) Reset()         { *m = Geo_IP{} }
func (m *Geo_IP) String() string { return proto.CompactTextString(m) }
func (*Geo_IP) ProtoMessage()    {}
func (*Geo_IP) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a5b7ee991490f54, []int{0, 0}
}

func (m *Geo_IP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Geo_IP.Unmarshal(m, b)
}
func (m *Geo_IP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Geo_IP.Marshal(b, m, deterministic)
}
func (m *Geo_IP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Geo_IP.Merge(m, src)
}
func (m *Geo_IP) XXX_Size() int {
	return xxx_messageInfo_Geo_IP.Size(m)
}
func (m *Geo_IP) XXX_DiscardUnknown() {
	xxx_messageInfo_Geo_IP.DiscardUnknown(m)
}

var xxx_messageInfo_Geo_IP proto.InternalMessageInfo

func (m *Geo_IP) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type String struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a5b7ee991490f54, []int{1}
}

func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (m *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(m, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Geo)(nil), "proto.Geo")
	proto.RegisterType((*Geo_IP)(nil), "proto.Geo.IP")
	proto.RegisterType((*String)(nil), "proto.String")
}

func init() { proto.RegisterFile("geoip.proto", fileDescriptor_9a5b7ee991490f54) }

var fileDescriptor_9a5b7ee991490f54 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4f, 0xcd, 0xcf,
	0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xd2, 0x5c, 0xcc, 0xee,
	0xa9, 0xf9, 0x52, 0x22, 0x5c, 0x4c, 0x9e, 0x01, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x05, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x4c, 0x99, 0x05, 0x4a, 0x52, 0x5c, 0x6c, 0xc1, 0x25, 0x45, 0x99,
	0x79, 0xe9, 0x42, 0x02, 0x5c, 0xcc, 0xb9, 0xc5, 0xe9, 0x50, 0x29, 0x10, 0xd3, 0x48, 0x8f, 0x8b,
	0xd5, 0x3d, 0x35, 0xdf, 0xb3, 0x40, 0x48, 0x95, 0x8b, 0x35, 0xb0, 0x34, 0xb5, 0xa8, 0x52, 0x88,
	0x17, 0x62, 0xb2, 0x9e, 0x7b, 0x6a, 0xbe, 0x9e, 0x67, 0x80, 0x14, 0x8c, 0x0b, 0x31, 0x21, 0x89,
	0x0d, 0xcc, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x22, 0xb4, 0x16, 0x32, 0x85, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GeoIpClient is the client API for GeoIp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeoIpClient interface {
	Query(ctx context.Context, in *Geo_IP, opts ...grpc.CallOption) (*String, error)
}

type geoIpClient struct {
	cc *grpc.ClientConn
}

func NewGeoIpClient(cc *grpc.ClientConn) GeoIpClient {
	return &geoIpClient{cc}
}

func (c *geoIpClient) Query(ctx context.Context, in *Geo_IP, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/proto.GeoIp/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeoIpServer is the server API for GeoIp service.
type GeoIpServer interface {
	Query(context.Context, *Geo_IP) (*String, error)
}

// UnimplementedGeoIpServer can be embedded to have forward compatible implementations.
type UnimplementedGeoIpServer struct {
}

func (*UnimplementedGeoIpServer) Query(ctx context.Context, req *Geo_IP) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

func RegisterGeoIpServer(s *grpc.Server, srv GeoIpServer) {
	s.RegisterService(&_GeoIp_serviceDesc, srv)
}

func _GeoIp_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Geo_IP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoIpServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GeoIp/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoIpServer).Query(ctx, req.(*Geo_IP))
	}
	return interceptor(ctx, in, info, handler)
}

var _GeoIp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GeoIp",
	HandlerType: (*GeoIpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _GeoIp_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "geoip.proto",
}
