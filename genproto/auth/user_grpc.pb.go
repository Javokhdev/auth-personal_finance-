// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: user.proto

package auth

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetProfile(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*UserRes, error)
	EditProfile(ctx context.Context, in *UserRes, opts ...grpc.CallOption) (*UserRes, error)
	ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*ChangePasswordRes, error)
	GetSetting(ctx context.Context, in *GetSettingReq, opts ...grpc.CallOption) (*Setting, error)
	EditSetting(ctx context.Context, in *SettingReq, opts ...grpc.CallOption) (*SettingRes, error)
	DeleteUser(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRes, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetProfile(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/auth.UserService/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EditProfile(ctx context.Context, in *UserRes, opts ...grpc.CallOption) (*UserRes, error) {
	out := new(UserRes)
	err := c.cc.Invoke(ctx, "/auth.UserService/EditProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*ChangePasswordRes, error) {
	out := new(ChangePasswordRes)
	err := c.cc.Invoke(ctx, "/auth.UserService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetSetting(ctx context.Context, in *GetSettingReq, opts ...grpc.CallOption) (*Setting, error) {
	out := new(Setting)
	err := c.cc.Invoke(ctx, "/auth.UserService/GetSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EditSetting(ctx context.Context, in *SettingReq, opts ...grpc.CallOption) (*SettingRes, error) {
	out := new(SettingRes)
	err := c.cc.Invoke(ctx, "/auth.UserService/EditSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRes, error) {
	out := new(DeleteRes)
	err := c.cc.Invoke(ctx, "/auth.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetProfile(context.Context, *GetByIdReq) (*UserRes, error)
	EditProfile(context.Context, *UserRes) (*UserRes, error)
	ChangePassword(context.Context, *ChangePasswordReq) (*ChangePasswordRes, error)
	GetSetting(context.Context, *GetSettingReq) (*Setting, error)
	EditSetting(context.Context, *SettingReq) (*SettingRes, error)
	DeleteUser(context.Context, *DeleteReq) (*DeleteRes, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetProfile(context.Context, *GetByIdReq) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedUserServiceServer) EditProfile(context.Context, *UserRes) (*UserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProfile not implemented")
}
func (UnimplementedUserServiceServer) ChangePassword(context.Context, *ChangePasswordReq) (*ChangePasswordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedUserServiceServer) GetSetting(context.Context, *GetSettingReq) (*Setting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSetting not implemented")
}
func (UnimplementedUserServiceServer) EditSetting(context.Context, *SettingReq) (*SettingRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditSetting not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteReq) (*DeleteRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.UserService/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetProfile(ctx, req.(*GetByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EditProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EditProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.UserService/EditProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EditProfile(ctx, req.(*UserRes))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.UserService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePassword(ctx, req.(*ChangePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSettingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.UserService/GetSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetSetting(ctx, req.(*GetSettingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EditSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SettingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EditSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.UserService/EditSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EditSetting(ctx, req.(*SettingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfile",
			Handler:    _UserService_GetProfile_Handler,
		},
		{
			MethodName: "EditProfile",
			Handler:    _UserService_EditProfile_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserService_ChangePassword_Handler,
		},
		{
			MethodName: "GetSetting",
			Handler:    _UserService_GetSetting_Handler,
		},
		{
			MethodName: "EditSetting",
			Handler:    _UserService_EditSetting_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
