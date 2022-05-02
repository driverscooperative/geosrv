// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: omeapi/crossing/v1beta1/crossing.proto

package v1beta1

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

// CrossingServiceClient is the client API for CrossingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrossingServiceClient interface {
	ListCrossings(ctx context.Context, in *ListCrossingsRequest, opts ...grpc.CallOption) (*ListCrossingsResponse, error)
}

type crossingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrossingServiceClient(cc grpc.ClientConnInterface) CrossingServiceClient {
	return &crossingServiceClient{cc}
}

func (c *crossingServiceClient) ListCrossings(ctx context.Context, in *ListCrossingsRequest, opts ...grpc.CallOption) (*ListCrossingsResponse, error) {
	out := new(ListCrossingsResponse)
	err := c.cc.Invoke(ctx, "/omeapi.crossing.v1beta1.CrossingService/ListCrossings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrossingServiceServer is the server API for CrossingService service.
// All implementations should embed UnimplementedCrossingServiceServer
// for forward compatibility
type CrossingServiceServer interface {
	ListCrossings(context.Context, *ListCrossingsRequest) (*ListCrossingsResponse, error)
}

// UnimplementedCrossingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCrossingServiceServer struct {
}

func (UnimplementedCrossingServiceServer) ListCrossings(context.Context, *ListCrossingsRequest) (*ListCrossingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCrossings not implemented")
}

// UnsafeCrossingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrossingServiceServer will
// result in compilation errors.
type UnsafeCrossingServiceServer interface {
	mustEmbedUnimplementedCrossingServiceServer()
}

func RegisterCrossingServiceServer(s grpc.ServiceRegistrar, srv CrossingServiceServer) {
	s.RegisterService(&CrossingService_ServiceDesc, srv)
}

func _CrossingService_ListCrossings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCrossingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrossingServiceServer).ListCrossings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omeapi.crossing.v1beta1.CrossingService/ListCrossings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrossingServiceServer).ListCrossings(ctx, req.(*ListCrossingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CrossingService_ServiceDesc is the grpc.ServiceDesc for CrossingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrossingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "omeapi.crossing.v1beta1.CrossingService",
	HandlerType: (*CrossingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCrossings",
			Handler:    _CrossingService_ListCrossings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "omeapi/crossing/v1beta1/crossing.proto",
}
