// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: omeapi/tollgate_crossing/v1beta1/tollgate_crossing.proto

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

// TollgateCrossingServiceClient is the client API for TollgateCrossingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TollgateCrossingServiceClient interface {
	ListTollgateCrossings(ctx context.Context, in *ListTollgateCrossingsRequest, opts ...grpc.CallOption) (*ListTollgateCrossingsResponse, error)
}

type tollgateCrossingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTollgateCrossingServiceClient(cc grpc.ClientConnInterface) TollgateCrossingServiceClient {
	return &tollgateCrossingServiceClient{cc}
}

func (c *tollgateCrossingServiceClient) ListTollgateCrossings(ctx context.Context, in *ListTollgateCrossingsRequest, opts ...grpc.CallOption) (*ListTollgateCrossingsResponse, error) {
	out := new(ListTollgateCrossingsResponse)
	err := c.cc.Invoke(ctx, "/omeapi.tollgate_crossing.v1beta1.TollgateCrossingService/ListTollgateCrossings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TollgateCrossingServiceServer is the server API for TollgateCrossingService service.
// All implementations should embed UnimplementedTollgateCrossingServiceServer
// for forward compatibility
type TollgateCrossingServiceServer interface {
	ListTollgateCrossings(context.Context, *ListTollgateCrossingsRequest) (*ListTollgateCrossingsResponse, error)
}

// UnimplementedTollgateCrossingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTollgateCrossingServiceServer struct {
}

func (UnimplementedTollgateCrossingServiceServer) ListTollgateCrossings(context.Context, *ListTollgateCrossingsRequest) (*ListTollgateCrossingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTollgateCrossings not implemented")
}

// UnsafeTollgateCrossingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TollgateCrossingServiceServer will
// result in compilation errors.
type UnsafeTollgateCrossingServiceServer interface {
	mustEmbedUnimplementedTollgateCrossingServiceServer()
}

func RegisterTollgateCrossingServiceServer(s grpc.ServiceRegistrar, srv TollgateCrossingServiceServer) {
	s.RegisterService(&TollgateCrossingService_ServiceDesc, srv)
}

func _TollgateCrossingService_ListTollgateCrossings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTollgateCrossingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TollgateCrossingServiceServer).ListTollgateCrossings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omeapi.tollgate_crossing.v1beta1.TollgateCrossingService/ListTollgateCrossings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TollgateCrossingServiceServer).ListTollgateCrossings(ctx, req.(*ListTollgateCrossingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TollgateCrossingService_ServiceDesc is the grpc.ServiceDesc for TollgateCrossingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TollgateCrossingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "omeapi.tollgate_crossing.v1beta1.TollgateCrossingService",
	HandlerType: (*TollgateCrossingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTollgateCrossings",
			Handler:    _TollgateCrossingService_ListTollgateCrossings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "omeapi/tollgate_crossing/v1beta1/tollgate_crossing.proto",
}
