// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: v1/ratewatcher/rw.proto

package ratewatcher

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

// RateWatcherServiceClient is the client API for RateWatcherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RateWatcherServiceClient interface {
	GetRate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RateResponse, error)
}

type rateWatcherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRateWatcherServiceClient(cc grpc.ClientConnInterface) RateWatcherServiceClient {
	return &rateWatcherServiceClient{cc}
}

func (c *rateWatcherServiceClient) GetRate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, "/ratewatcher.v1.RateWatcherService/GetRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateWatcherServiceServer is the server API for RateWatcherService service.
// All implementations must embed UnimplementedRateWatcherServiceServer
// for forward compatibility
type RateWatcherServiceServer interface {
	GetRate(context.Context, *emptypb.Empty) (*RateResponse, error)
	mustEmbedUnimplementedRateWatcherServiceServer()
}

// UnimplementedRateWatcherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRateWatcherServiceServer struct {
}

func (UnimplementedRateWatcherServiceServer) GetRate(context.Context, *emptypb.Empty) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRate not implemented")
}
func (UnimplementedRateWatcherServiceServer) mustEmbedUnimplementedRateWatcherServiceServer() {}

// UnsafeRateWatcherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RateWatcherServiceServer will
// result in compilation errors.
type UnsafeRateWatcherServiceServer interface {
	mustEmbedUnimplementedRateWatcherServiceServer()
}

func RegisterRateWatcherServiceServer(s grpc.ServiceRegistrar, srv RateWatcherServiceServer) {
	s.RegisterService(&RateWatcherService_ServiceDesc, srv)
}

func _RateWatcherService_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateWatcherServiceServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ratewatcher.v1.RateWatcherService/GetRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateWatcherServiceServer).GetRate(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RateWatcherService_ServiceDesc is the grpc.ServiceDesc for RateWatcherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RateWatcherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ratewatcher.v1.RateWatcherService",
	HandlerType: (*RateWatcherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _RateWatcherService_GetRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/ratewatcher/rw.proto",
}
