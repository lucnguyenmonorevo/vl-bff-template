// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/general_info/aggregates/personal_setting/bookmark_service.proto

package protopersonalsetting

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
	BookmarkService_GetByUserID_FullMethodName = "/general_info.aggregates.personal_setting.BookmarkService/GetByUserID"
)

// BookmarkServiceClient is the client API for BookmarkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookmarkServiceClient interface {
	GetByUserID(ctx context.Context, in *BookMarkGetByUserIDRequest, opts ...grpc.CallOption) (*BookMarkGetByUserIDResponse, error)
}

type bookmarkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookmarkServiceClient(cc grpc.ClientConnInterface) BookmarkServiceClient {
	return &bookmarkServiceClient{cc}
}

func (c *bookmarkServiceClient) GetByUserID(ctx context.Context, in *BookMarkGetByUserIDRequest, opts ...grpc.CallOption) (*BookMarkGetByUserIDResponse, error) {
	out := new(BookMarkGetByUserIDResponse)
	err := c.cc.Invoke(ctx, BookmarkService_GetByUserID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookmarkServiceServer is the server API for BookmarkService service.
// All implementations must embed UnimplementedBookmarkServiceServer
// for forward compatibility
type BookmarkServiceServer interface {
	GetByUserID(context.Context, *BookMarkGetByUserIDRequest) (*BookMarkGetByUserIDResponse, error)
	mustEmbedUnimplementedBookmarkServiceServer()
}

// UnimplementedBookmarkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookmarkServiceServer struct {
}

func (UnimplementedBookmarkServiceServer) GetByUserID(context.Context, *BookMarkGetByUserIDRequest) (*BookMarkGetByUserIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUserID not implemented")
}
func (UnimplementedBookmarkServiceServer) mustEmbedUnimplementedBookmarkServiceServer() {}

// UnsafeBookmarkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookmarkServiceServer will
// result in compilation errors.
type UnsafeBookmarkServiceServer interface {
	mustEmbedUnimplementedBookmarkServiceServer()
}

func RegisterBookmarkServiceServer(s grpc.ServiceRegistrar, srv BookmarkServiceServer) {
	s.RegisterService(&BookmarkService_ServiceDesc, srv)
}

func _BookmarkService_GetByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookMarkGetByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookmarkServiceServer).GetByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookmarkService_GetByUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookmarkServiceServer).GetByUserID(ctx, req.(*BookMarkGetByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookmarkService_ServiceDesc is the grpc.ServiceDesc for BookmarkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookmarkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "general_info.aggregates.personal_setting.BookmarkService",
	HandlerType: (*BookmarkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByUserID",
			Handler:    _BookmarkService_GetByUserID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/general_info/aggregates/personal_setting/bookmark_service.proto",
}