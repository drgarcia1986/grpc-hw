// Code generated by protoc-gen-go.
// source: hw.proto
// DO NOT EDIT!

/*
Package hw is a generated protocol buffer package.

It is generated from these files:
	hw.proto

It has these top-level messages:
	Req
	Res
*/
package hw

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

type Req struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Req) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Res struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Res) Reset()                    { *m = Res{} }
func (m *Res) String() string            { return proto.CompactTextString(m) }
func (*Res) ProtoMessage()               {}
func (*Res) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Res) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Req)(nil), "hw.Req")
	proto.RegisterType((*Res)(nil), "hw.Res")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HelloWorld service

type HelloWorldClient interface {
	Say(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
}

type helloWorldClient struct {
	cc *grpc.ClientConn
}

func NewHelloWorldClient(cc *grpc.ClientConn) HelloWorldClient {
	return &helloWorldClient{cc}
}

func (c *helloWorldClient) Say(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := grpc.Invoke(ctx, "/hw.HelloWorld/Say", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloWorld service

type HelloWorldServer interface {
	Say(context.Context, *Req) (*Res, error)
}

func RegisterHelloWorldServer(s *grpc.Server, srv HelloWorldServer) {
	s.RegisterService(&_HelloWorld_serviceDesc, srv)
}

func _HelloWorld_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hw.HelloWorld/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).Say(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloWorld_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hw.HelloWorld",
	HandlerType: (*HelloWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _HelloWorld_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hw.proto",
}

func init() { proto.RegisterFile("hw.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x28, 0xd7, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xca, 0x28, 0x57, 0x92, 0xe4, 0x62, 0x0e, 0x4a, 0x2d, 0x14,
	0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3,
	0x95, 0xc4, 0x41, 0x52, 0xc5, 0x42, 0x02, 0x5c, 0xcc, 0xb9, 0xc5, 0xe9, 0x50, 0x19, 0x10, 0xd3,
	0x48, 0x95, 0x8b, 0xcb, 0x23, 0x35, 0x27, 0x27, 0x3f, 0x3c, 0xbf, 0x28, 0x27, 0x45, 0x48, 0x9c,
	0x8b, 0x39, 0x38, 0xb1, 0x52, 0x88, 0x5d, 0x2f, 0xa3, 0x5c, 0x2f, 0x28, 0xb5, 0x50, 0x0a, 0xca,
	0x28, 0x4e, 0x62, 0x03, 0xdb, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x28, 0x56, 0x85, 0x53,
	0x71, 0x00, 0x00, 0x00,
}
