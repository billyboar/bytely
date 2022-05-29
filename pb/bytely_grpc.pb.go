// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BytelyServiceClient is the client API for BytelyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BytelyServiceClient interface {
	AddURL(ctx context.Context, in *AddURLRequest, opts ...grpc.CallOption) (*AddURLResponse, error)
	GetOriginalURL(ctx context.Context, in *GetOriginalURLRequest, opts ...grpc.CallOption) (*GetOriginalURLResponse, error)
	GetURLStats(ctx context.Context, in *GetURLStatsRequest, opts ...grpc.CallOption) (*GetURLStatsResponse, error)
	DeleteURL(ctx context.Context, in *DeleteURLRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type bytelyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBytelyServiceClient(cc grpc.ClientConnInterface) BytelyServiceClient {
	return &bytelyServiceClient{cc}
}

func (c *bytelyServiceClient) AddURL(ctx context.Context, in *AddURLRequest, opts ...grpc.CallOption) (*AddURLResponse, error) {
	out := new(AddURLResponse)
	err := c.cc.Invoke(ctx, "/bytely.BytelyService/AddURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bytelyServiceClient) GetOriginalURL(ctx context.Context, in *GetOriginalURLRequest, opts ...grpc.CallOption) (*GetOriginalURLResponse, error) {
	out := new(GetOriginalURLResponse)
	err := c.cc.Invoke(ctx, "/bytely.BytelyService/GetOriginalURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bytelyServiceClient) GetURLStats(ctx context.Context, in *GetURLStatsRequest, opts ...grpc.CallOption) (*GetURLStatsResponse, error) {
	out := new(GetURLStatsResponse)
	err := c.cc.Invoke(ctx, "/bytely.BytelyService/GetURLStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bytelyServiceClient) DeleteURL(ctx context.Context, in *DeleteURLRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/bytely.BytelyService/DeleteURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BytelyServiceServer is the server API for BytelyService service.
// All implementations must embed UnimplementedBytelyServiceServer
// for forward compatibility
type BytelyServiceServer interface {
	AddURL(context.Context, *AddURLRequest) (*AddURLResponse, error)
	GetOriginalURL(context.Context, *GetOriginalURLRequest) (*GetOriginalURLResponse, error)
	GetURLStats(context.Context, *GetURLStatsRequest) (*GetURLStatsResponse, error)
	DeleteURL(context.Context, *DeleteURLRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedBytelyServiceServer()
}

// UnimplementedBytelyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBytelyServiceServer struct {
}

func (UnimplementedBytelyServiceServer) AddURL(context.Context, *AddURLRequest) (*AddURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddURL not implemented")
}
func (UnimplementedBytelyServiceServer) GetOriginalURL(context.Context, *GetOriginalURLRequest) (*GetOriginalURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOriginalURL not implemented")
}
func (UnimplementedBytelyServiceServer) GetURLStats(context.Context, *GetURLStatsRequest) (*GetURLStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetURLStats not implemented")
}
func (UnimplementedBytelyServiceServer) DeleteURL(context.Context, *DeleteURLRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteURL not implemented")
}
func (UnimplementedBytelyServiceServer) mustEmbedUnimplementedBytelyServiceServer() {}

// UnsafeBytelyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BytelyServiceServer will
// result in compilation errors.
type UnsafeBytelyServiceServer interface {
	mustEmbedUnimplementedBytelyServiceServer()
}

func RegisterBytelyServiceServer(s grpc.ServiceRegistrar, srv BytelyServiceServer) {
	s.RegisterService(&BytelyService_ServiceDesc, srv)
}

func _BytelyService_AddURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BytelyServiceServer).AddURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bytely.BytelyService/AddURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BytelyServiceServer).AddURL(ctx, req.(*AddURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BytelyService_GetOriginalURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOriginalURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BytelyServiceServer).GetOriginalURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bytely.BytelyService/GetOriginalURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BytelyServiceServer).GetOriginalURL(ctx, req.(*GetOriginalURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BytelyService_GetURLStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetURLStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BytelyServiceServer).GetURLStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bytely.BytelyService/GetURLStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BytelyServiceServer).GetURLStats(ctx, req.(*GetURLStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BytelyService_DeleteURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BytelyServiceServer).DeleteURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bytely.BytelyService/DeleteURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BytelyServiceServer).DeleteURL(ctx, req.(*DeleteURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BytelyService_ServiceDesc is the grpc.ServiceDesc for BytelyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BytelyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bytely.BytelyService",
	HandlerType: (*BytelyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddURL",
			Handler:    _BytelyService_AddURL_Handler,
		},
		{
			MethodName: "GetOriginalURL",
			Handler:    _BytelyService_GetOriginalURL_Handler,
		},
		{
			MethodName: "GetURLStats",
			Handler:    _BytelyService_GetURLStats_Handler,
		},
		{
			MethodName: "DeleteURL",
			Handler:    _BytelyService_DeleteURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bytely.proto",
}