// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bistream_service.proto

package pb

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

type Request struct {
	Tx                   []byte   `protobuf:"bytes,1,opt,name=Tx,proto3" json:"Tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_bistream_service_985f1314e491115e, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

type Response struct {
	Result               string   `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=Method,proto3" json:"Method,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	Error                string   `protobuf:"bytes,4,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_bistream_service_985f1314e491115e, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *Response) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Response) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_bistream_service_985f1314e491115e, []int{2}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Request)(nil), "pb.Request")
	proto.RegisterType((*Response)(nil), "pb.Response")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BistreamServiceClient is the client API for BistreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BistreamServiceClient interface {
	RunICode(ctx context.Context, opts ...grpc.CallOption) (BistreamService_RunICodeClient, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type bistreamServiceClient struct {
	cc *grpc.ClientConn
}

func NewBistreamServiceClient(cc *grpc.ClientConn) BistreamServiceClient {
	return &bistreamServiceClient{cc}
}

func (c *bistreamServiceClient) RunICode(ctx context.Context, opts ...grpc.CallOption) (BistreamService_RunICodeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BistreamService_serviceDesc.Streams[0], "/pb.BistreamService/RunICode", opts...)
	if err != nil {
		return nil, err
	}
	x := &bistreamServiceRunICodeClient{stream}
	return x, nil
}

type BistreamService_RunICodeClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type bistreamServiceRunICodeClient struct {
	grpc.ClientStream
}

func (x *bistreamServiceRunICodeClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bistreamServiceRunICodeClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bistreamServiceClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.BistreamService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BistreamServiceServer is the server API for BistreamService service.
type BistreamServiceServer interface {
	RunICode(BistreamService_RunICodeServer) error
	Ping(context.Context, *Empty) (*Empty, error)
}

func RegisterBistreamServiceServer(s *grpc.Server, srv BistreamServiceServer) {
	s.RegisterService(&_BistreamService_serviceDesc, srv)
}

func _BistreamService_RunICode_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BistreamServiceServer).RunICode(&bistreamServiceRunICodeServer{stream})
}

type BistreamService_RunICodeServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type bistreamServiceRunICodeServer struct {
	grpc.ServerStream
}

func (x *bistreamServiceRunICodeServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bistreamServiceRunICodeServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BistreamService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BistreamServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BistreamService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BistreamServiceServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _BistreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BistreamService",
	HandlerType: (*BistreamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _BistreamService_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunICode",
			Handler:       _BistreamService_RunICode_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bistream_service.proto",
}

func init() {
	proto.RegisterFile("bistream_service.proto", fileDescriptor_bistream_service_985f1314e491115e)
}

var fileDescriptor_bistream_service_985f1314e491115e = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x27, 0xb1, 0xf3, 0x77, 0x1d, 0x14, 0x2e, 0x32, 0xc4, 0x59, 0x88, 0x64, 0x55, 0x10,
	0x8a, 0xe8, 0x1b, 0xa8, 0x5d, 0xb8, 0x10, 0x24, 0x76, 0xad, 0x34, 0xf6, 0xa2, 0x05, 0xdb, 0xc4,
	0x24, 0x95, 0xfa, 0xf6, 0xd2, 0xa4, 0xe0, 0xee, 0x7e, 0xe7, 0x84, 0xc3, 0x17, 0xd8, 0xeb, 0xd6,
	0x07, 0x47, 0x75, 0xf7, 0xe6, 0xc9, 0xfd, 0xb4, 0xef, 0x54, 0x58, 0x67, 0x82, 0x41, 0x6e, 0xb5,
	0x3c, 0x87, 0xb5, 0xa2, 0xef, 0x81, 0x7c, 0xc0, 0x13, 0xe0, 0xd5, 0x28, 0xd8, 0x25, 0xcb, 0x77,
	0x8a, 0x57, 0xa3, 0x6c, 0x60, 0xa3, 0xc8, 0x5b, 0xd3, 0x7b, 0xc2, 0x3d, 0xac, 0x14, 0xf9, 0xe1,
	0x2b, 0xc4, 0x7e, 0xab, 0x66, 0x9a, 0xf2, 0x27, 0x0a, 0x9f, 0xa6, 0x11, 0x3c, 0xe5, 0x89, 0x10,
	0x21, 0x7b, 0xa8, 0x43, 0x2d, 0x8e, 0xe2, 0x5a, 0xbc, 0xf1, 0x0c, 0x96, 0xa5, 0x73, 0xc6, 0x89,
	0x2c, 0x3e, 0x4d, 0x20, 0xd7, 0xb0, 0x2c, 0x3b, 0x1b, 0x7e, 0x6f, 0x5e, 0xe1, 0xf4, 0x6e, 0xf6,
	0x7c, 0x49, 0x9a, 0x78, 0x05, 0x1b, 0x35, 0xf4, 0x8f, 0xf7, 0xa6, 0x21, 0x3c, 0x2e, 0xac, 0x2e,
	0x66, 0xd5, 0xc3, 0x2e, 0x41, 0x92, 0x93, 0x8b, 0x9c, 0x5d, 0x33, 0xbc, 0x80, 0xec, 0xb9, 0xed,
	0x3f, 0x70, 0x3b, 0x75, 0x71, 0xf2, 0xf0, 0x7f, 0xca, 0x85, 0x5e, 0xc5, 0x4f, 0xdf, 0xfe, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x32, 0x1e, 0x55, 0xcd, 0x0e, 0x01, 0x00, 0x00,
}
