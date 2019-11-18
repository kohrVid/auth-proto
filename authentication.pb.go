// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication.proto

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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN         HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING         HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING     HealthCheckResponse_ServingStatus = 2
	HealthCheckResponse_SERVICE_UNKNOWN HealthCheckResponse_ServingStatus = 3
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
	3: "SERVICE_UNKNOWN",
}

var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":         0,
	"SERVING":         1,
	"NOT_SERVING":     2,
	"SERVICE_UNKNOWN": 3,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}

func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{1, 0}
}

type HealthCheckRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{0}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status               HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=proto.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{1}
}

func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(m, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

type AuthenticationRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationRequest) Reset()         { *m = AuthenticationRequest{} }
func (m *AuthenticationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticationRequest) ProtoMessage()    {}
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{2}
}

func (m *AuthenticationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationRequest.Unmarshal(m, b)
}
func (m *AuthenticationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationRequest.Marshal(b, m, deterministic)
}
func (m *AuthenticationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationRequest.Merge(m, src)
}
func (m *AuthenticationRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticationRequest.Size(m)
}
func (m *AuthenticationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationRequest proto.InternalMessageInfo

func (m *AuthenticationRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthenticationRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthenticationResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationResponse) Reset()         { *m = AuthenticationResponse{} }
func (m *AuthenticationResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticationResponse) ProtoMessage()    {}
func (*AuthenticationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{3}
}

func (m *AuthenticationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationResponse.Unmarshal(m, b)
}
func (m *AuthenticationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationResponse.Merge(m, src)
}
func (m *AuthenticationResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticationResponse.Size(m)
}
func (m *AuthenticationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationResponse proto.InternalMessageInfo

func (m *AuthenticationResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
	proto.RegisterType((*HealthCheckRequest)(nil), "proto.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "proto.HealthCheckResponse")
	proto.RegisterType((*AuthenticationRequest)(nil), "proto.AuthenticationRequest")
	proto.RegisterType((*AuthenticationResponse)(nil), "proto.AuthenticationResponse")
}

func init() { proto.RegisterFile("authentication.proto", fileDescriptor_d0dbc99083440df2) }

var fileDescriptor_d0dbc99083440df2 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xdf, 0x4b, 0xc3, 0x30,
	0x18, 0xb4, 0x93, 0x75, 0xfa, 0x0d, 0xed, 0xc8, 0x74, 0xd4, 0xa2, 0x20, 0x7d, 0xda, 0x53, 0x91,
	0xfa, 0x2e, 0x4a, 0x19, 0x4e, 0x84, 0x14, 0x5a, 0x7f, 0x3c, 0x8e, 0xd8, 0x85, 0xb5, 0x58, 0xdb,
	0x9a, 0xa4, 0xfa, 0x2f, 0xf9, 0x67, 0x4a, 0x93, 0x66, 0xd0, 0xd1, 0x3d, 0x85, 0xcb, 0x7d, 0x77,
	0x5f, 0xee, 0x02, 0x67, 0xa4, 0x16, 0x29, 0x2d, 0x44, 0x96, 0x10, 0x91, 0x95, 0x85, 0x57, 0xb1,
	0x52, 0x94, 0x68, 0x28, 0x0f, 0xd7, 0x03, 0xb4, 0xa4, 0x24, 0x17, 0x69, 0x90, 0xd2, 0xe4, 0x33,
	0xa2, 0xdf, 0x35, 0xe5, 0x02, 0xd9, 0x30, 0xe2, 0x94, 0xfd, 0x64, 0x09, 0xb5, 0x8d, 0x6b, 0x63,
	0x7e, 0x1c, 0x69, 0xe8, 0xfe, 0x19, 0x30, 0xed, 0x08, 0x78, 0x55, 0x16, 0x9c, 0xa2, 0x7b, 0x30,
	0xb9, 0x20, 0xa2, 0xe6, 0x52, 0x70, 0xea, 0xcf, 0xd5, 0x1a, 0xaf, 0x67, 0xd6, 0x8b, 0x1b, 0xaf,
	0x62, 0x13, 0xcb, 0xf9, 0xa8, 0xd5, 0xb9, 0x21, 0x9c, 0x74, 0x08, 0x34, 0x86, 0xd1, 0x2b, 0x7e,
	0xc6, 0xe1, 0x3b, 0x9e, 0x1c, 0x34, 0x20, 0x5e, 0x44, 0x6f, 0x4f, 0xf8, 0x71, 0x62, 0x20, 0x0b,
	0xc6, 0x38, 0x7c, 0x59, 0xe9, 0x8b, 0x01, 0x9a, 0x82, 0x25, 0x41, 0xb0, 0x58, 0x69, 0xc9, 0xa1,
	0x1b, 0xc2, 0xf9, 0x43, 0x27, 0xb9, 0x4e, 0xe7, 0xc0, 0x51, 0xcd, 0x29, 0x2b, 0xc8, 0x97, 0x8e,
	0xb7, 0xc5, 0x0d, 0x57, 0x11, 0xce, 0x7f, 0x4b, 0xb6, 0xb6, 0x07, 0x8a, 0xd3, 0xd8, 0xbd, 0x81,
	0xd9, 0xae, 0x61, 0x9b, 0x7e, 0x06, 0x26, 0xa3, 0xbc, 0xce, 0x45, 0xeb, 0xd7, 0x22, 0x7f, 0x09,
	0xa6, 0x2a, 0x00, 0xdd, 0xc1, 0x50, 0x96, 0x80, 0x2e, 0xfa, 0x8a, 0x91, 0xef, 0x72, 0x9c, 0xfd,
	0x9d, 0xf9, 0x9b, 0xdd, 0x30, 0xb1, 0xfa, 0x10, 0x84, 0xc1, 0x0a, 0x18, 0x5d, 0x37, 0x04, 0xc9,
	0xd5, 0x8a, 0xcb, 0xd6, 0xa7, 0x37, 0xbd, 0x73, 0xb5, 0x87, 0x55, 0x8b, 0x3e, 0x4c, 0xc9, 0xde,
	0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xda, 0x99, 0xd4, 0xf6, 0x36, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthClient interface {
	// If the requested service is unknown, the call will fail with status
	// NOT_FOUND.
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/proto.Health/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
type HealthServer interface {
	// If the requested service is unknown, the call will fail with status
	// NOT_FOUND.
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

// UnimplementedHealthServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (*UnimplementedHealthServer) Check(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	CredentialCheck(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
}

type authenticationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationServiceClient(cc *grpc.ClientConn) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) CredentialCheck(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	out := new(AuthenticationResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthenticationService/CredentialCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
type AuthenticationServiceServer interface {
	CredentialCheck(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error)
}

// UnimplementedAuthenticationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (*UnimplementedAuthenticationServiceServer) CredentialCheck(ctx context.Context, req *AuthenticationRequest) (*AuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CredentialCheck not implemented")
}

func RegisterAuthenticationServiceServer(s *grpc.Server, srv AuthenticationServiceServer) {
	s.RegisterService(&_AuthenticationService_serviceDesc, srv)
}

func _AuthenticationService_CredentialCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).CredentialCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthenticationService/CredentialCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).CredentialCheck(ctx, req.(*AuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthenticationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CredentialCheck",
			Handler:    _AuthenticationService_CredentialCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}
