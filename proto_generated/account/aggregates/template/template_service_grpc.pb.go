// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/account/aggregates/template/template_service.proto

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
	TemplateService_List_FullMethodName = "/account.aggregates.template.TemplateService/List"
)

// TemplateServiceClient is the client API for TemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateServiceClient interface {
	List(ctx context.Context, in *TemplateListRequest, opts ...grpc.CallOption) (*TemplateListResponse, error)
}

type templateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateServiceClient(cc grpc.ClientConnInterface) TemplateServiceClient {
	return &templateServiceClient{cc}
}

func (c *templateServiceClient) List(ctx context.Context, in *TemplateListRequest, opts ...grpc.CallOption) (*TemplateListResponse, error) {
	out := new(TemplateListResponse)
	err := c.cc.Invoke(ctx, TemplateService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateServiceServer is the server API for TemplateService service.
// All implementations must embed UnimplementedTemplateServiceServer
// for forward compatibility
type TemplateServiceServer interface {
	List(context.Context, *TemplateListRequest) (*TemplateListResponse, error)
	mustEmbedUnimplementedTemplateServiceServer()
}

// UnimplementedTemplateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTemplateServiceServer struct {
}

func (UnimplementedTemplateServiceServer) List(context.Context, *TemplateListRequest) (*TemplateListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTemplateServiceServer) mustEmbedUnimplementedTemplateServiceServer() {}

// UnsafeTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateServiceServer will
// result in compilation errors.
type UnsafeTemplateServiceServer interface {
	mustEmbedUnimplementedTemplateServiceServer()
}

func RegisterTemplateServiceServer(s grpc.ServiceRegistrar, srv TemplateServiceServer) {
	s.RegisterService(&TemplateService_ServiceDesc, srv)
}

func _TemplateService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemplateListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TemplateService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServiceServer).List(ctx, req.(*TemplateListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TemplateService_ServiceDesc is the grpc.ServiceDesc for TemplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TemplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.aggregates.template.TemplateService",
	HandlerType: (*TemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _TemplateService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/account/aggregates/template/template_service.proto",
}
