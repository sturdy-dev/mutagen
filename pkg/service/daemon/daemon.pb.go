// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/daemon/daemon.proto

package daemon // import "github.com/havoc-io/mutagen/pkg/service/daemon"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_daemon_cfcbff340fa2bcf6, []int{0}
}
func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (dst *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(dst, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

type VersionResponse struct {
	// TODO: Should we encapsulate these inside a Version message type, perhaps
	// in the mutagen package?
	Major                uint64   `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor                uint64   `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch                uint64   `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_daemon_cfcbff340fa2bcf6, []int{1}
}
func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (dst *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(dst, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetMajor() uint64 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *VersionResponse) GetMinor() uint64 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *VersionResponse) GetPatch() uint64 {
	if m != nil {
		return m.Patch
	}
	return 0
}

type TerminateRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminateRequest) Reset()         { *m = TerminateRequest{} }
func (m *TerminateRequest) String() string { return proto.CompactTextString(m) }
func (*TerminateRequest) ProtoMessage()    {}
func (*TerminateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_daemon_cfcbff340fa2bcf6, []int{2}
}
func (m *TerminateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminateRequest.Unmarshal(m, b)
}
func (m *TerminateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminateRequest.Marshal(b, m, deterministic)
}
func (dst *TerminateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateRequest.Merge(dst, src)
}
func (m *TerminateRequest) XXX_Size() int {
	return xxx_messageInfo_TerminateRequest.Size(m)
}
func (m *TerminateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateRequest proto.InternalMessageInfo

type TerminateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminateResponse) Reset()         { *m = TerminateResponse{} }
func (m *TerminateResponse) String() string { return proto.CompactTextString(m) }
func (*TerminateResponse) ProtoMessage()    {}
func (*TerminateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_daemon_cfcbff340fa2bcf6, []int{3}
}
func (m *TerminateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminateResponse.Unmarshal(m, b)
}
func (m *TerminateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminateResponse.Marshal(b, m, deterministic)
}
func (dst *TerminateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateResponse.Merge(dst, src)
}
func (m *TerminateResponse) XXX_Size() int {
	return xxx_messageInfo_TerminateResponse.Size(m)
}
func (m *TerminateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VersionRequest)(nil), "daemon.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "daemon.VersionResponse")
	proto.RegisterType((*TerminateRequest)(nil), "daemon.TerminateRequest")
	proto.RegisterType((*TerminateResponse)(nil), "daemon.TerminateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DaemonClient is the client API for Daemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DaemonClient interface {
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error)
}

type daemonClient struct {
	cc *grpc.ClientConn
}

func NewDaemonClient(cc *grpc.ClientConn) DaemonClient {
	return &daemonClient{cc}
}

func (c *daemonClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/daemon.Daemon/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error) {
	out := new(TerminateResponse)
	err := c.cc.Invoke(ctx, "/daemon.Daemon/Terminate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaemonServer is the server API for Daemon service.
type DaemonServer interface {
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error)
}

func RegisterDaemonServer(s *grpc.Server, srv DaemonServer) {
	s.RegisterService(&_Daemon_serviceDesc, srv)
}

func _Daemon_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.Daemon/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Terminate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Terminate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daemon.Daemon/Terminate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Terminate(ctx, req.(*TerminateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Daemon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "daemon.Daemon",
	HandlerType: (*DaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Daemon_Version_Handler,
		},
		{
			MethodName: "Terminate",
			Handler:    _Daemon_Terminate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/daemon/daemon.proto",
}

func init() { proto.RegisterFile("service/daemon/daemon.proto", fileDescriptor_daemon_cfcbff340fa2bcf6) }

var fileDescriptor_daemon_cfcbff340fa2bcf6 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbf, 0x4e, 0x04, 0x21,
	0x10, 0xc6, 0x5d, 0xff, 0xac, 0x71, 0x0a, 0x3d, 0xd1, 0xe8, 0x7a, 0x36, 0x66, 0x2b, 0x1b, 0xc1,
	0x68, 0x6b, 0x75, 0xf1, 0x09, 0x4e, 0x63, 0x61, 0xc7, 0xe1, 0x64, 0x17, 0x0d, 0x0c, 0x02, 0x7b,
	0x0f, 0xe1, 0x53, 0x9b, 0x03, 0xf6, 0xe2, 0xa9, 0x15, 0xf9, 0x7e, 0x4c, 0xbe, 0xf9, 0x65, 0xe0,
	0x32, 0xa0, 0x5f, 0x6a, 0x85, 0xe2, 0x4d, 0xa2, 0x21, 0x5b, 0x1e, 0xee, 0x3c, 0x45, 0x62, 0x75,
	0x4e, 0xed, 0x04, 0x0e, 0x5f, 0xd0, 0x07, 0x4d, 0x76, 0x8e, 0x9f, 0x03, 0x86, 0xd8, 0x3e, 0xc1,
	0xd1, 0x9a, 0x04, 0x47, 0x36, 0x20, 0x3b, 0x85, 0x3d, 0x23, 0xdf, 0xc9, 0x37, 0xd5, 0x55, 0x75,
	0xbd, 0x3b, 0xcf, 0x21, 0x51, 0x6d, 0xc9, 0x37, 0xdb, 0x85, 0xae, 0xc2, 0x8a, 0x3a, 0x19, 0x55,
	0xdf, 0xec, 0x64, 0x9a, 0x42, 0xcb, 0x60, 0xf2, 0x8c, 0xde, 0x68, 0x2b, 0x23, 0x8e, 0x8b, 0x4e,
	0xe0, 0xf8, 0x07, 0xcb, 0xab, 0xee, 0xbe, 0x2a, 0xa8, 0x1f, 0x93, 0x1a, 0x7b, 0x80, 0xfd, 0x22,
	0xc2, 0xce, 0x78, 0x91, 0xdf, 0x74, 0x9d, 0x9e, 0xff, 0xe1, 0xb9, 0xa6, 0xdd, 0x62, 0x33, 0x38,
	0x58, 0xb7, 0xb3, 0x66, 0x9c, 0xfb, 0x2d, 0x31, 0xbd, 0xf8, 0xe7, 0x67, 0xec, 0x98, 0xdd, 0xbe,
	0xf2, 0x4e, 0xc7, 0x7e, 0x58, 0x70, 0x45, 0x46, 0xf4, 0x72, 0x49, 0xea, 0x46, 0x93, 0x30, 0x43,
	0x94, 0x1d, 0x5a, 0xe1, 0x3e, 0x3a, 0xb1, 0x79, 0xe3, 0x45, 0x9d, 0xae, 0x7b, 0xff, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x0a, 0x84, 0xd4, 0xfa, 0x7c, 0x01, 0x00, 0x00,
}