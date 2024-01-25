// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/account/aggregates/permission/permission_policy_service.proto

package protopermission

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	grpc1 "vl-account/proto_generated/account/aggregates/permission/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PermissionPolicyService_List_FullMethodName   = "/account.aggregates.permission.PermissionPolicyService/List"
	PermissionPolicyService_Get_FullMethodName    = "/account.aggregates.permission.PermissionPolicyService/Get"
	PermissionPolicyService_Verify_FullMethodName = "/account.aggregates.permission.PermissionPolicyService/Verify"
	PermissionPolicyService_Delete_FullMethodName = "/account.aggregates.permission.PermissionPolicyService/Delete"
	PermissionPolicyService_Create_FullMethodName = "/account.aggregates.permission.PermissionPolicyService/Create"
	PermissionPolicyService_Update_FullMethodName = "/account.aggregates.permission.PermissionPolicyService/Update"
)

// PermissionPolicyServiceClient is the client API for PermissionPolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionPolicyServiceClient interface {
	List(ctx context.Context, in *grpc1.PermissionPolicyListRequest, opts ...grpc.CallOption) (*grpc1.PermissionPolicyListResponse, error)
	Get(ctx context.Context, in *grpc1.PermissionPolicyGetRequest, opts ...grpc.CallOption) (*grpc1.PermissionPolicyGetResponse, error)
	Verify(ctx context.Context, in *grpc1.PermissionPolicyVerifyRequest, opts ...grpc.CallOption) (*grpc1.PermissionPolicyVerifyResponse, error)
	Delete(ctx context.Context, in *grpc1.PermissionPolicyDeleteRequest, opts ...grpc.CallOption) (*grpc1.PermissionPolicyDeleteResponse, error)
	Create(ctx context.Context, in *grpc1.PermissionPolicyCreateRequest, opts ...grpc.CallOption) (*grpc1.PermissionPolicyCreateResponse, error)
	Update(ctx context.Context, in *grpc1.PermissionPolicyUpdateRequest, opts ...grpc.CallOption) (*grpc1.PermissionPolicyUpdateResponse, error)
}

type permissionPolicyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionPolicyServiceClient(cc grpc.ClientConnInterface) PermissionPolicyServiceClient {
	return &permissionPolicyServiceClient{cc}
}

func (c *permissionPolicyServiceClient) List(ctx context.Context, in *grpc1.PermissionPolicyListRequest, opts ...grpc.CallOption) (*grpc1.PermissionPolicyListResponse, error) {
	out := new(grpc1.PermissionPolicyListResponse)
	err := c.cc.Invoke(ctx, PermissionPolicyService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionPolicyServiceClient) Get(ctx context.Context, in *grpc1.PermissionPolicyGetRequest, opts ...grpc.CallOption) (*grpc1.PermissionPolicyGetResponse, error) {
	out := new(grpc1.PermissionPolicyGetResponse)
	err := c.cc.Invoke(ctx, PermissionPolicyService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionPolicyServiceClient) Verify(ctx context.Context, in *grpc1.PermissionPolicyVerifyRequest, opts ...grpc.CallOption) (*grpc1.PermissionPolicyVerifyResponse, error) {
	out := new(grpc1.PermissionPolicyVerifyResponse)
	err := c.cc.Invoke(ctx, PermissionPolicyService_Verify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionPolicyServiceClient) Delete(ctx context.Context, in *grpc1.PermissionPolicyDeleteRequest, opts ...grpc.CallOption) (*grpc1.PermissionPolicyDeleteResponse, error) {
	out := new(grpc1.PermissionPolicyDeleteResponse)
	err := c.cc.Invoke(ctx, PermissionPolicyService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionPolicyServiceClient) Create(ctx context.Context, in *grpc1.PermissionPolicyCreateRequest, opts ...grpc.CallOption) (*grpc1.PermissionPolicyCreateResponse, error) {
	out := new(grpc1.PermissionPolicyCreateResponse)
	err := c.cc.Invoke(ctx, PermissionPolicyService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionPolicyServiceClient) Update(ctx context.Context, in *grpc1.PermissionPolicyUpdateRequest, opts ...grpc.CallOption) (*grpc1.PermissionPolicyUpdateResponse, error) {
	out := new(grpc1.PermissionPolicyUpdateResponse)
	err := c.cc.Invoke(ctx, PermissionPolicyService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionPolicyServiceServer is the server API for PermissionPolicyService service.
// All implementations must embed UnimplementedPermissionPolicyServiceServer
// for forward compatibility
type PermissionPolicyServiceServer interface {
	List(context.Context, *grpc1.PermissionPolicyListRequest) (*grpc1.PermissionPolicyListResponse, error)
	Get(context.Context, *grpc1.PermissionPolicyGetRequest) (*grpc1.PermissionPolicyGetResponse, error)
	Verify(context.Context, *grpc1.PermissionPolicyVerifyRequest) (*grpc1.PermissionPolicyVerifyResponse, error)
	Delete(context.Context, *grpc1.PermissionPolicyDeleteRequest) (*grpc1.PermissionPolicyDeleteResponse, error)
	Create(context.Context, *grpc1.PermissionPolicyCreateRequest) (*grpc1.PermissionPolicyCreateResponse, error)
	Update(context.Context, *grpc1.PermissionPolicyUpdateRequest) (*grpc1.PermissionPolicyUpdateResponse, error)
	mustEmbedUnimplementedPermissionPolicyServiceServer()
}

// UnimplementedPermissionPolicyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionPolicyServiceServer struct {
}

func (UnimplementedPermissionPolicyServiceServer) List(context.Context, *grpc1.PermissionPolicyListRequest) (*grpc1.PermissionPolicyListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPermissionPolicyServiceServer) Get(context.Context, *grpc1.PermissionPolicyGetRequest) (*grpc1.PermissionPolicyGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPermissionPolicyServiceServer) Verify(context.Context, *grpc1.PermissionPolicyVerifyRequest) (*grpc1.PermissionPolicyVerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedPermissionPolicyServiceServer) Delete(context.Context, *grpc1.PermissionPolicyDeleteRequest) (*grpc1.PermissionPolicyDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPermissionPolicyServiceServer) Create(context.Context, *grpc1.PermissionPolicyCreateRequest) (*grpc1.PermissionPolicyCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPermissionPolicyServiceServer) Update(context.Context, *grpc1.PermissionPolicyUpdateRequest) (*grpc1.PermissionPolicyUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPermissionPolicyServiceServer) mustEmbedUnimplementedPermissionPolicyServiceServer() {
}

// UnsafePermissionPolicyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionPolicyServiceServer will
// result in compilation errors.
type UnsafePermissionPolicyServiceServer interface {
	mustEmbedUnimplementedPermissionPolicyServiceServer()
}

func RegisterPermissionPolicyServiceServer(s grpc.ServiceRegistrar, srv PermissionPolicyServiceServer) {
	s.RegisterService(&PermissionPolicyService_ServiceDesc, srv)
}

func _PermissionPolicyService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc1.PermissionPolicyListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionPolicyServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionPolicyService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionPolicyServiceServer).List(ctx, req.(*grpc1.PermissionPolicyListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionPolicyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc1.PermissionPolicyGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionPolicyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionPolicyService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionPolicyServiceServer).Get(ctx, req.(*grpc1.PermissionPolicyGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionPolicyService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc1.PermissionPolicyVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionPolicyServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionPolicyService_Verify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionPolicyServiceServer).Verify(ctx, req.(*grpc1.PermissionPolicyVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionPolicyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc1.PermissionPolicyDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionPolicyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionPolicyService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionPolicyServiceServer).Delete(ctx, req.(*grpc1.PermissionPolicyDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionPolicyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc1.PermissionPolicyCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionPolicyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionPolicyService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionPolicyServiceServer).Create(ctx, req.(*grpc1.PermissionPolicyCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionPolicyService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc1.PermissionPolicyUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionPolicyServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionPolicyService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionPolicyServiceServer).Update(ctx, req.(*grpc1.PermissionPolicyUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionPolicyService_ServiceDesc is the grpc.ServiceDesc for PermissionPolicyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionPolicyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.aggregates.permission.PermissionPolicyService",
	HandlerType: (*PermissionPolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _PermissionPolicyService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PermissionPolicyService_Get_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _PermissionPolicyService_Verify_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PermissionPolicyService_Delete_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PermissionPolicyService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PermissionPolicyService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/account/aggregates/permission/permission_policy_service.proto",
}
