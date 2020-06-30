// Code generated by protoc-gen-go. DO NOT EDIT.
// source: register.proto

package auth

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

type RegisterRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterResponse struct {
	Message              string            `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	UserData             *UserIdentityData `protobuf:"bytes,2,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{1}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RegisterResponse) GetUserData() *UserIdentityData {
	if m != nil {
		return m.UserData
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "auth.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "auth.RegisterResponse")
}

func init() { proto.RegisterFile("register.proto", fileDescriptor_1303fe8288f4efb6) }

var fileDescriptor_1303fe8288f4efb6 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x29, 0x9f, 0xed, 0x81, 0x00, 0x59, 0x50, 0x59, 0x99, 0x50, 0x26, 0xa6, 0x0c, 0xed,
	0x0f, 0x60, 0x61, 0x61, 0x35, 0x62, 0x60, 0xaa, 0x0e, 0x72, 0x0a, 0x96, 0x5a, 0x3b, 0xdc, 0x5d,
	0x40, 0xfc, 0x7b, 0x64, 0x3b, 0x50, 0x29, 0x9b, 0xdf, 0xe7, 0xec, 0xd7, 0x8f, 0x0d, 0x97, 0x4c,
	0x9d, 0x17, 0x25, 0x6e, 0x7a, 0x8e, 0x1a, 0xcd, 0x31, 0x0e, 0xfa, 0x51, 0xd9, 0x41, 0x88, 0x37,
	0xbe, 0xa5, 0xa0, 0x5e, 0x7f, 0x36, 0x2d, 0x2a, 0x96, 0x79, 0xfd, 0x0a, 0x57, 0x6e, 0x3c, 0xe1,
	0xe8, 0x73, 0x20, 0x51, 0x73, 0x03, 0x27, 0xdb, 0xd8, 0xf9, 0x60, 0x67, 0x77, 0xb3, 0xfb, 0x85,
	0x2b, 0x21, 0x51, 0xda, 0xa1, 0xdf, 0xda, 0xc3, 0x42, 0x73, 0x30, 0x15, 0xcc, 0x7b, 0x14, 0xf9,
	0x8e, 0xdc, 0xda, 0xa3, 0x3c, 0xf8, 0xcf, 0x35, 0xc2, 0xf5, 0xbe, 0x5a, 0xfa, 0x18, 0x84, 0x8c,
	0x85, 0xb3, 0x1d, 0x89, 0x60, 0x47, 0x63, 0xfb, 0x5f, 0x34, 0x6b, 0x58, 0x64, 0xc9, 0xe4, 0x96,
	0xef, 0x38, 0x5f, 0x2d, 0x9b, 0x24, 0xdf, 0xbc, 0x08, 0xf1, 0xd3, 0xa8, 0xfe, 0x88, 0x8a, 0x6e,
	0x9e, 0x36, 0xa6, 0xd5, 0xca, 0xed, 0xed, 0x9f, 0x89, 0xbf, 0xfc, 0x3b, 0x99, 0x07, 0xb8, 0x28,
	0x88, 0x51, 0x7d, 0x0c, 0xe6, 0xb6, 0x94, 0x4c, 0x1e, 0x59, 0x2d, 0xa7, 0xb8, 0x08, 0xd6, 0x07,
	0x6f, 0xa7, 0xf9, 0x63, 0xd6, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xbc, 0xff, 0x8e, 0x4a,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RegisterServiceClient is the client API for RegisterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegisterServiceClient interface {
	Registration(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type registerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegisterServiceClient(cc grpc.ClientConnInterface) RegisterServiceClient {
	return &registerServiceClient{cc}
}

func (c *registerServiceClient) Registration(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/auth.RegisterService/Registration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterServiceServer is the server API for RegisterService service.
type RegisterServiceServer interface {
	Registration(context.Context, *RegisterRequest) (*RegisterResponse, error)
}

// UnimplementedRegisterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRegisterServiceServer struct {
}

func (*UnimplementedRegisterServiceServer) Registration(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registration not implemented")
}

func RegisterRegisterServiceServer(s *grpc.Server, srv RegisterServiceServer) {
	s.RegisterService(&_RegisterService_serviceDesc, srv)
}

func _RegisterService_Registration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServiceServer).Registration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.RegisterService/Registration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServiceServer).Registration(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegisterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.RegisterService",
	HandlerType: (*RegisterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registration",
			Handler:    _RegisterService_Registration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register.proto",
}
