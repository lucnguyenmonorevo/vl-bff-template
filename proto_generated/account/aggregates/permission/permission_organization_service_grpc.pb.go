// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/account/aggregates/permission/permission_organization_service.proto

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
	PermissionOrganizationService_List_FullMethodName   = "/account.aggregates.permission.PermissionOrganizationService/List"
	PermissionOrganizationService_Get_FullMethodName    = "/account.aggregates.permission.PermissionOrganizationService/Get"
	PermissionOrganizationService_Delete_FullMethodName = "/account.aggregates.permission.PermissionOrganizationService/Delete"
	PermissionOrganizationService_Create_FullMethodName = "/account.aggregates.permission.PermissionOrganizationService/Create"
	PermissionOrganizationService_Update_FullMethodName = "/account.aggregates.permission.PermissionOrganizationService/Update"
)

// PermissionOrganizationServiceClient is the client API for PermissionOrganizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionOrganizationServiceClient interface {
	List(ctx context.Context, in *grpc1.PermissionOrganizationListRequest, opts ...grpc.CallOption) (*grpc1.PermissionOrganizationListResponse, error)
	Get(ctx context.Context, in *grpc1.PermissionOrganizationGetRequest, opts ...grpc.CallOption) (*grpc1.PermissionOrganizationGetResponse, error)
	Delete(ctx context.Context, in *grpc1.PermissionOrganizationDeleteRequest, opts ...grpc.CallOption) (*grpc1.PermissionOrganizationDeleteResponse, error)
	Create(ctx context.Context, in *grpc1.PermissionOrganizationCreateRequest, opts ...grpc.CallOption) (*grpc1.PermissionOrganizationCreateResponse, error)
	Update(ctx context.Context, in *grpc1.PermissionOrganizationUpdateRequest, opts ...grpc.CallOption) (*grpc1.PermissionOrganizationUpdateResponse, error)
}

type permissionOrganizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionOrganizationServiceClient(cc grpc.ClientConnInterface) PermissionOrganizationServiceClient {
	return &permissionOrganizationServiceClient{cc}
}

func (c *permissionOrganizationServiceClient) List(ctx context.Context, in *grpc1.PermissionOrganizationListRequest, opts ...grpc.CallOption) (*grpc1.PermissionOrganizationListResponse, error) {
	out := new(grpc1.PermissionOrganizationListResponse)
	err := c.cc.Invoke(ctx, PermissionOrganizationService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionOrganizationServiceClient) Get(ctx context.Context, in *grpc1.PermissionOrganizationGetRequest, opts ...grpc.CallOption) (*grpc1.PermissionOrganizationGetResponse, error) {
	out := new(grpc1.PermissionOrganizationGetResponse)
	err := c.cc.Invoke(ctx, PermissionOrganizationService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionOrganizationServiceClient) Delete(ctx context.Context, in *grpc1.PermissionOrganizationDeleteRequest, opts ...grpc.CallOption) (*grpc1.PermissionOrganizationDeleteResponse, error) {
	out := new(grpc1.PermissionOrganizationDeleteResponse)
	err := c.cc.Invoke(ctx, PermissionOrganizationService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionOrganizationServiceClient) Create(ctx context.Context, in *grpc1.PermissionOrganizationCreateRequest, opts ...grpc.CallOption) (*grpc1.PermissionOrganizationCreateResponse, error) {
	out := new(grpc1.PermissionOrganizationCreateResponse)
	err := c.cc.Invoke(ctx, PermissionOrganizationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionOrganizationServiceClient) Update(ctx context.Context, in *grpc1.PermissionOrganizationUpdateRequest, opts ...grpc.CallOption) (*grpc1.PermissionOrganizationUpdateResponse, error) {
	out := new(grpc1.PermissionOrganizationUpdateResponse)
	err := c.cc.Invoke(ctx, PermissionOrganizationService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionOrganizationServiceServer is the server API for PermissionOrganizationService service.
// All implementations must embed UnimplementedPermissionOrganizationServiceServer
// for forward compatibility
type PermissionOrganizationServiceServer interface {
	List(context.Context, *grpc1.PermissionOrganizationListRequest) (*grpc1.PermissionOrganizationListResponse, error)
	Get(context.Context, *grpc1.PermissionOrganizationGetRequest) (*grpc1.PermissionOrganizationGetResponse, error)
	Delete(context.Context, *grpc1.PermissionOrganizationDeleteRequest) (*grpc1.PermissionOrganizationDeleteResponse, error)
	Create(context.Context, *grpc1.PermissionOrganizationCreateRequest) (*grpc1.PermissionOrganizationCreateResponse, error)
	Update(context.Context, *grpc1.PermissionOrganizationUpdateRequest) (*grpc1.PermissionOrganizationUpdateResponse, error)
	mustEmbedUnimplementedPermissionOrganizationServiceServer()
}

// UnimplementedPermissionOrganizationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionOrganizationServiceServer struct {
}

func (UnimplementedPermissionOrganizationServiceServer) List(context.Context, *grpc1.PermissionOrganizationListRequest) (*grpc1.PermissionOrganizationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPermissionOrganizationServiceServer) Get(context.Context, *grpc1.PermissionOrganizationGetRequest) (*grpc1.PermissionOrganizationGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPermissionOrganizationServiceServer) Delete(context.Context, *grpc1.PermissionOrganizationDeleteRequest) (*grpc1.PermissionOrganizationDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPermissionOrganizationServiceServer) Create(context.Context, *grpc1.PermissionOrganizationCreateRequest) (*grpc1.PermissionOrganizationCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPermissionOrganizationServiceServer) Update(context.Context, *grpc1.PermissionOrganizationUpdateRequest) (*grpc1.PermissionOrganizationUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPermissionOrganizationServiceServer) mustEmbedUnimplementedPermissionOrganizationServiceServer() {
}

// UnsafePermissionOrganizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionOrganizationServiceServer will
// result in compilation errors.
type UnsafePermissionOrganizationServiceServer interface {
	mustEmbedUnimplementedPermissionOrganizationServiceServer()
}

func RegisterPermissionOrganizationServiceServer(s grpc.ServiceRegistrar, srv PermissionOrganizationServiceServer) {
	s.RegisterService(&PermissionOrganizationService_ServiceDesc, srv)
}

func _PermissionOrganizationService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc1.PermissionOrganizationListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionOrganizationServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionOrganizationService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionOrganizationServiceServer).List(ctx, req.(*grpc1.PermissionOrganizationListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionOrganizationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc1.PermissionOrganizationGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionOrganizationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionOrganizationService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionOrganizationServiceServer).Get(ctx, req.(*grpc1.PermissionOrganizationGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionOrganizationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc1.PermissionOrganizationDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionOrganizationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionOrganizationService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionOrganizationServiceServer).Delete(ctx, req.(*grpc1.PermissionOrganizationDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionOrganizationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc1.PermissionOrganizationCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionOrganizationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionOrganizationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionOrganizationServiceServer).Create(ctx, req.(*grpc1.PermissionOrganizationCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionOrganizationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc1.PermissionOrganizationUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionOrganizationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PermissionOrganizationService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionOrganizationServiceServer).Update(ctx, req.(*grpc1.PermissionOrganizationUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionOrganizationService_ServiceDesc is the grpc.ServiceDesc for PermissionOrganizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionOrganizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.aggregates.permission.PermissionOrganizationService",
	HandlerType: (*PermissionOrganizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _PermissionOrganizationService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PermissionOrganizationService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PermissionOrganizationService_Delete_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PermissionOrganizationService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PermissionOrganizationService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/account/aggregates/permission/permission_organization_service.proto",
}