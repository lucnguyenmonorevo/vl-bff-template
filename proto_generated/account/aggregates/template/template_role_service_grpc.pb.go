// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/account/aggregates/template/template_role_service.proto

package prototemplate

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
	TemplateRoleService_List_FullMethodName = "/account.aggregates.template.TemplateRoleService/List"
)

// TemplateRoleServiceClient is the client API for TemplateRoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateRoleServiceClient interface {
	List(ctx context.Context, in *TemplateRoleListRequest, opts ...grpc.CallOption) (*TemplateRoleListResponse, error)
}

type templateRoleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateRoleServiceClient(cc grpc.ClientConnInterface) TemplateRoleServiceClient {
	return &templateRoleServiceClient{cc}
}

func (c *templateRoleServiceClient) List(ctx context.Context, in *TemplateRoleListRequest, opts ...grpc.CallOption) (*TemplateRoleListResponse, error) {
	out := new(TemplateRoleListResponse)
	err := c.cc.Invoke(ctx, TemplateRoleService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateRoleServiceServer is the server API for TemplateRoleService service.
// All implementations must embed UnimplementedTemplateRoleServiceServer
// for forward compatibility
type TemplateRoleServiceServer interface {
	List(context.Context, *TemplateRoleListRequest) (*TemplateRoleListResponse, error)
	mustEmbedUnimplementedTemplateRoleServiceServer()
}

// UnimplementedTemplateRoleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTemplateRoleServiceServer struct {
}

func (UnimplementedTemplateRoleServiceServer) List(context.Context, *TemplateRoleListRequest) (*TemplateRoleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTemplateRoleServiceServer) mustEmbedUnimplementedTemplateRoleServiceServer() {}

// UnsafeTemplateRoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateRoleServiceServer will
// result in compilation errors.
type UnsafeTemplateRoleServiceServer interface {
	mustEmbedUnimplementedTemplateRoleServiceServer()
}

func RegisterTemplateRoleServiceServer(s grpc.ServiceRegistrar, srv TemplateRoleServiceServer) {
	s.RegisterService(&TemplateRoleService_ServiceDesc, srv)
}

func _TemplateRoleService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemplateRoleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateRoleServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TemplateRoleService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateRoleServiceServer).List(ctx, req.(*TemplateRoleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TemplateRoleService_ServiceDesc is the grpc.ServiceDesc for TemplateRoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TemplateRoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.aggregates.template.TemplateRoleService",
	HandlerType: (*TemplateRoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _TemplateRoleService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/account/aggregates/template/template_role_service.proto",
}