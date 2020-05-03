// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo.proto

package amassa

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

// frontend and inputhandler
type Request struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Extension            string   `protobuf:"bytes,3,opt,name=extension,proto3" json:"extension,omitempty"`
	Bytes                string   `protobuf:"bytes,4,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Command              string   `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetExtension() string {
	if m != nil {
		return m.Extension
	}
	return ""
}

func (m *Request) GetBytes() string {
	if m != nil {
		return m.Bytes
	}
	return ""
}

func (m *Request) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

type Response struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "amassa.Request")
	proto.RegisterType((*Response)(nil), "amassa.Response")
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor_ca53982754088a9d) }

var fileDescriptor_ca53982754088a9d = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xc1, 0x4a, 0xc5, 0x30,
	0x10, 0x45, 0xad, 0xef, 0xbd, 0x3e, 0x3b, 0x20, 0xca, 0x28, 0x12, 0xc4, 0x85, 0x64, 0xe5, 0xaa,
	0x0b, 0xf5, 0x07, 0xdc, 0x29, 0xee, 0xe2, 0x17, 0x4c, 0xdb, 0x59, 0x14, 0x9a, 0xa4, 0x76, 0xa6,
	0xa2, 0x4b, 0xff, 0x5c, 0x9a, 0x54, 0xba, 0xcb, 0x39, 0x81, 0x7b, 0x6f, 0x02, 0xd0, 0xb1, 0x8f,
	0xf5, 0x38, 0x45, 0x8d, 0x58, 0x92, 0x27, 0x11, 0xb2, 0xbf, 0x05, 0x1c, 0x1d, 0x7f, 0xce, 0x2c,
	0x8a, 0xd7, 0x70, 0x60, 0x4f, 0xfd, 0x60, 0x8a, 0xfb, 0xe2, 0xa1, 0x72, 0x19, 0x10, 0x61, 0x1f,
	0xc8, 0xb3, 0x39, 0x4d, 0x32, 0x9d, 0xf1, 0x0e, 0x2a, 0xfe, 0x56, 0x0e, 0xd2, 0xc7, 0x60, 0x76,
	0xe9, 0x62, 0x13, 0x4b, 0x4e, 0xf3, 0xa3, 0x2c, 0x66, 0x9f, 0x73, 0x12, 0xa0, 0x81, 0x63, 0x1b,
	0xbd, 0xa7, 0xd0, 0x99, 0x43, 0xf2, 0xff, 0x68, 0x2d, 0x9c, 0x39, 0x96, 0x31, 0x06, 0x61, 0xbc,
	0x81, 0x52, 0x94, 0x74, 0x96, 0x34, 0x62, 0xe7, 0x56, 0x7a, 0x7c, 0x87, 0xab, 0xb7, 0x30, 0xce,
	0xfa, 0x4a, 0xa1, 0x1b, 0x78, 0xfa, 0xe0, 0xe9, 0xab, 0x6f, 0x19, 0x9f, 0xe1, 0x7c, 0x5d, 0xff,
	0xd2, 0xea, 0xd2, 0x7d, 0x51, 0xe7, 0x87, 0xd5, 0xab, 0xbe, 0xbd, 0xdc, 0x44, 0xae, 0xb0, 0x27,
	0x4d, 0x99, 0xfe, 0xe0, 0xe9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x37, 0xb3, 0xca, 0x7b, 0x11, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InputHandlerServiceClient is the client API for InputHandlerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InputHandlerServiceClient interface {
	RequestAction(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type inputHandlerServiceClient struct {
	cc *grpc.ClientConn
}

func NewInputHandlerServiceClient(cc *grpc.ClientConn) InputHandlerServiceClient {
	return &inputHandlerServiceClient{cc}
}

func (c *inputHandlerServiceClient) RequestAction(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/amassa.InputHandlerService/RequestAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InputHandlerServiceServer is the server API for InputHandlerService service.
type InputHandlerServiceServer interface {
	RequestAction(context.Context, *Request) (*Response, error)
}

// UnimplementedInputHandlerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInputHandlerServiceServer struct {
}

func (*UnimplementedInputHandlerServiceServer) RequestAction(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAction not implemented")
}

func RegisterInputHandlerServiceServer(s *grpc.Server, srv InputHandlerServiceServer) {
	s.RegisterService(&_InputHandlerService_serviceDesc, srv)
}

func _InputHandlerService_RequestAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InputHandlerServiceServer).RequestAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/amassa.InputHandlerService/RequestAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InputHandlerServiceServer).RequestAction(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _InputHandlerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "amassa.InputHandlerService",
	HandlerType: (*InputHandlerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestAction",
			Handler:    _InputHandlerService_RequestAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}
