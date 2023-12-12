// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: sharePrice/sharePricePb/sharePricePb.proto

package sharePrice

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

// SharePriceGrpcServiceClient is the client API for SharePriceGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SharePriceGrpcServiceClient interface {
	SharePriceSearch(ctx context.Context, in *SharePriceReq, opts ...grpc.CallOption) (*SharePriceRes, error)
}

type sharePriceGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSharePriceGrpcServiceClient(cc grpc.ClientConnInterface) SharePriceGrpcServiceClient {
	return &sharePriceGrpcServiceClient{cc}
}

func (c *sharePriceGrpcServiceClient) SharePriceSearch(ctx context.Context, in *SharePriceReq, opts ...grpc.CallOption) (*SharePriceRes, error) {
	out := new(SharePriceRes)
	err := c.cc.Invoke(ctx, "/SharePriceGrpcService/SharePriceSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharePriceGrpcServiceServer is the server API for SharePriceGrpcService service.
// All implementations must embed UnimplementedSharePriceGrpcServiceServer
// for forward compatibility
type SharePriceGrpcServiceServer interface {
	SharePriceSearch(context.Context, *SharePriceReq) (*SharePriceRes, error)
	mustEmbedUnimplementedSharePriceGrpcServiceServer()
}

// UnimplementedSharePriceGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSharePriceGrpcServiceServer struct {
}

func (UnimplementedSharePriceGrpcServiceServer) SharePriceSearch(context.Context, *SharePriceReq) (*SharePriceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SharePriceSearch not implemented")
}
func (UnimplementedSharePriceGrpcServiceServer) mustEmbedUnimplementedSharePriceGrpcServiceServer() {}

// UnsafeSharePriceGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SharePriceGrpcServiceServer will
// result in compilation errors.
type UnsafeSharePriceGrpcServiceServer interface {
	mustEmbedUnimplementedSharePriceGrpcServiceServer()
}

func RegisterSharePriceGrpcServiceServer(s grpc.ServiceRegistrar, srv SharePriceGrpcServiceServer) {
	s.RegisterService(&SharePriceGrpcService_ServiceDesc, srv)
}

func _SharePriceGrpcService_SharePriceSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SharePriceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharePriceGrpcServiceServer).SharePriceSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SharePriceGrpcService/SharePriceSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharePriceGrpcServiceServer).SharePriceSearch(ctx, req.(*SharePriceReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SharePriceGrpcService_ServiceDesc is the grpc.ServiceDesc for SharePriceGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SharePriceGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SharePriceGrpcService",
	HandlerType: (*SharePriceGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SharePriceSearch",
			Handler:    _SharePriceGrpcService_SharePriceSearch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sharePrice/sharePricePb/sharePricePb.proto",
}