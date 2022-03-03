// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: core/model/location/protos/v1/location.proto

package v1

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

// LocationServiceClient is the client API for LocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationServiceClient interface {
	UpdateLocation(ctx context.Context, in *UpdateLocationRequest, opts ...grpc.CallOption) (*UpdateLocationResponse, error)
	QueryLocation(ctx context.Context, in *QueryLocationRequest, opts ...grpc.CallOption) (*QueryLocationResponse, error)
	UpdateLocationStreaming(ctx context.Context, opts ...grpc.CallOption) (LocationService_UpdateLocationStreamingClient, error)
	QueryLocationStreaming(ctx context.Context, in *QueryLocationRequest, opts ...grpc.CallOption) (LocationService_QueryLocationStreamingClient, error)
}

type locationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationServiceClient(cc grpc.ClientConnInterface) LocationServiceClient {
	return &locationServiceClient{cc}
}

func (c *locationServiceClient) UpdateLocation(ctx context.Context, in *UpdateLocationRequest, opts ...grpc.CallOption) (*UpdateLocationResponse, error) {
	out := new(UpdateLocationResponse)
	err := c.cc.Invoke(ctx, "/protos.location.v1.LocationService/UpdateLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) QueryLocation(ctx context.Context, in *QueryLocationRequest, opts ...grpc.CallOption) (*QueryLocationResponse, error) {
	out := new(QueryLocationResponse)
	err := c.cc.Invoke(ctx, "/protos.location.v1.LocationService/QueryLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) UpdateLocationStreaming(ctx context.Context, opts ...grpc.CallOption) (LocationService_UpdateLocationStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &LocationService_ServiceDesc.Streams[0], "/protos.location.v1.LocationService/UpdateLocationStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &locationServiceUpdateLocationStreamingClient{stream}
	return x, nil
}

type LocationService_UpdateLocationStreamingClient interface {
	Send(*UpdateLocationRequest) error
	CloseAndRecv() (*UpdateLocationResponse, error)
	grpc.ClientStream
}

type locationServiceUpdateLocationStreamingClient struct {
	grpc.ClientStream
}

func (x *locationServiceUpdateLocationStreamingClient) Send(m *UpdateLocationRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *locationServiceUpdateLocationStreamingClient) CloseAndRecv() (*UpdateLocationResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UpdateLocationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *locationServiceClient) QueryLocationStreaming(ctx context.Context, in *QueryLocationRequest, opts ...grpc.CallOption) (LocationService_QueryLocationStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &LocationService_ServiceDesc.Streams[1], "/protos.location.v1.LocationService/QueryLocationStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &locationServiceQueryLocationStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LocationService_QueryLocationStreamingClient interface {
	Recv() (*QueryLocationResponse, error)
	grpc.ClientStream
}

type locationServiceQueryLocationStreamingClient struct {
	grpc.ClientStream
}

func (x *locationServiceQueryLocationStreamingClient) Recv() (*QueryLocationResponse, error) {
	m := new(QueryLocationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LocationServiceServer is the server API for LocationService service.
// All implementations must embed UnimplementedLocationServiceServer
// for forward compatibility
type LocationServiceServer interface {
	UpdateLocation(context.Context, *UpdateLocationRequest) (*UpdateLocationResponse, error)
	QueryLocation(context.Context, *QueryLocationRequest) (*QueryLocationResponse, error)
	UpdateLocationStreaming(LocationService_UpdateLocationStreamingServer) error
	QueryLocationStreaming(*QueryLocationRequest, LocationService_QueryLocationStreamingServer) error
	mustEmbedUnimplementedLocationServiceServer()
}

// UnimplementedLocationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLocationServiceServer struct {
}

func (UnimplementedLocationServiceServer) UpdateLocation(context.Context, *UpdateLocationRequest) (*UpdateLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLocation not implemented")
}
func (UnimplementedLocationServiceServer) QueryLocation(context.Context, *QueryLocationRequest) (*QueryLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryLocation not implemented")
}
func (UnimplementedLocationServiceServer) UpdateLocationStreaming(LocationService_UpdateLocationStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateLocationStreaming not implemented")
}
func (UnimplementedLocationServiceServer) QueryLocationStreaming(*QueryLocationRequest, LocationService_QueryLocationStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryLocationStreaming not implemented")
}
func (UnimplementedLocationServiceServer) mustEmbedUnimplementedLocationServiceServer() {}

// UnsafeLocationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationServiceServer will
// result in compilation errors.
type UnsafeLocationServiceServer interface {
	mustEmbedUnimplementedLocationServiceServer()
}

func RegisterLocationServiceServer(s grpc.ServiceRegistrar, srv LocationServiceServer) {
	s.RegisterService(&LocationService_ServiceDesc, srv)
}

func _LocationService_UpdateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).UpdateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.location.v1.LocationService/UpdateLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).UpdateLocation(ctx, req.(*UpdateLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_QueryLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).QueryLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.location.v1.LocationService/QueryLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).QueryLocation(ctx, req.(*QueryLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_UpdateLocationStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LocationServiceServer).UpdateLocationStreaming(&locationServiceUpdateLocationStreamingServer{stream})
}

type LocationService_UpdateLocationStreamingServer interface {
	SendAndClose(*UpdateLocationResponse) error
	Recv() (*UpdateLocationRequest, error)
	grpc.ServerStream
}

type locationServiceUpdateLocationStreamingServer struct {
	grpc.ServerStream
}

func (x *locationServiceUpdateLocationStreamingServer) SendAndClose(m *UpdateLocationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *locationServiceUpdateLocationStreamingServer) Recv() (*UpdateLocationRequest, error) {
	m := new(UpdateLocationRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _LocationService_QueryLocationStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryLocationRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LocationServiceServer).QueryLocationStreaming(m, &locationServiceQueryLocationStreamingServer{stream})
}

type LocationService_QueryLocationStreamingServer interface {
	Send(*QueryLocationResponse) error
	grpc.ServerStream
}

type locationServiceQueryLocationStreamingServer struct {
	grpc.ServerStream
}

func (x *locationServiceQueryLocationStreamingServer) Send(m *QueryLocationResponse) error {
	return x.ServerStream.SendMsg(m)
}

// LocationService_ServiceDesc is the grpc.ServiceDesc for LocationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.location.v1.LocationService",
	HandlerType: (*LocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateLocation",
			Handler:    _LocationService_UpdateLocation_Handler,
		},
		{
			MethodName: "QueryLocation",
			Handler:    _LocationService_QueryLocation_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateLocationStreaming",
			Handler:       _LocationService_UpdateLocationStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "QueryLocationStreaming",
			Handler:       _LocationService_QueryLocationStreaming_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "core/model/location/protos/v1/location.proto",
}