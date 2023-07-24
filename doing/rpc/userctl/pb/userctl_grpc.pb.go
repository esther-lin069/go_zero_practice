// 聲明 proto 語法版本，固定值

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/userctl.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserCtl_Login_FullMethodName    = "/pb.UserCtl/Login"
	UserCtl_Register_FullMethodName = "/pb.UserCtl/Register"
)

// UserCtlClient is the client API for UserCtl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserCtlClient interface {
	// Login rpc Logic
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	// Register rpc Logic
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
}

type userCtlClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCtlClient(cc grpc.ClientConnInterface) UserCtlClient {
	return &userCtlClient{cc}
}

func (c *userCtlClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, UserCtl_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCtlClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, UserCtl_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCtlServer is the server API for UserCtl service.
// All implementations must embed UnimplementedUserCtlServer
// for forward compatibility
type UserCtlServer interface {
	// Login rpc Logic
	Login(context.Context, *LoginReq) (*LoginResp, error)
	// Register rpc Logic
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	mustEmbedUnimplementedUserCtlServer()
}

// UnimplementedUserCtlServer must be embedded to have forward compatible implementations.
type UnimplementedUserCtlServer struct {
}

func (UnimplementedUserCtlServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserCtlServer) Register(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserCtlServer) mustEmbedUnimplementedUserCtlServer() {}

// UnsafeUserCtlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserCtlServer will
// result in compilation errors.
type UnsafeUserCtlServer interface {
	mustEmbedUnimplementedUserCtlServer()
}

func RegisterUserCtlServer(s grpc.ServiceRegistrar, srv UserCtlServer) {
	s.RegisterService(&UserCtl_ServiceDesc, srv)
}

func _UserCtl_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCtlServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCtl_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCtlServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCtl_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCtlServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCtl_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCtlServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserCtl_ServiceDesc is the grpc.ServiceDesc for UserCtl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserCtl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserCtl",
	HandlerType: (*UserCtlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserCtl_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserCtl_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/userctl.proto",
}
