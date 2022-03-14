// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: omeapi/worker/v1beta1/worker.proto

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

// WorkerServiceClient is the client API for WorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerServiceClient interface {
	GetWorker(ctx context.Context, in *GetWorkerRequest, opts ...grpc.CallOption) (*GetWorkerResponse, error)
	SetState(ctx context.Context, in *SetStateRequest, opts ...grpc.CallOption) (*SetStateResponse, error)
	QueryByState(ctx context.Context, in *QueryByStateRequest, opts ...grpc.CallOption) (*QueryByStateResponse, error)
}

type workerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerServiceClient(cc grpc.ClientConnInterface) WorkerServiceClient {
	return &workerServiceClient{cc}
}

func (c *workerServiceClient) GetWorker(ctx context.Context, in *GetWorkerRequest, opts ...grpc.CallOption) (*GetWorkerResponse, error) {
	out := new(GetWorkerResponse)
	err := c.cc.Invoke(ctx, "/omeapi.worker.v1beta1.WorkerService/GetWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) SetState(ctx context.Context, in *SetStateRequest, opts ...grpc.CallOption) (*SetStateResponse, error) {
	out := new(SetStateResponse)
	err := c.cc.Invoke(ctx, "/omeapi.worker.v1beta1.WorkerService/SetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) QueryByState(ctx context.Context, in *QueryByStateRequest, opts ...grpc.CallOption) (*QueryByStateResponse, error) {
	out := new(QueryByStateResponse)
	err := c.cc.Invoke(ctx, "/omeapi.worker.v1beta1.WorkerService/QueryByState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServiceServer is the server API for WorkerService service.
// All implementations should embed UnimplementedWorkerServiceServer
// for forward compatibility
type WorkerServiceServer interface {
	GetWorker(context.Context, *GetWorkerRequest) (*GetWorkerResponse, error)
	SetState(context.Context, *SetStateRequest) (*SetStateResponse, error)
	QueryByState(context.Context, *QueryByStateRequest) (*QueryByStateResponse, error)
}

// UnimplementedWorkerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWorkerServiceServer struct {
}

func (UnimplementedWorkerServiceServer) GetWorker(context.Context, *GetWorkerRequest) (*GetWorkerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorker not implemented")
}
func (UnimplementedWorkerServiceServer) SetState(context.Context, *SetStateRequest) (*SetStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetState not implemented")
}
func (UnimplementedWorkerServiceServer) QueryByState(context.Context, *QueryByStateRequest) (*QueryByStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryByState not implemented")
}

// UnsafeWorkerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServiceServer will
// result in compilation errors.
type UnsafeWorkerServiceServer interface {
	mustEmbedUnimplementedWorkerServiceServer()
}

func RegisterWorkerServiceServer(s grpc.ServiceRegistrar, srv WorkerServiceServer) {
	s.RegisterService(&WorkerService_ServiceDesc, srv)
}

func _WorkerService_GetWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).GetWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omeapi.worker.v1beta1.WorkerService/GetWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).GetWorker(ctx, req.(*GetWorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_SetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).SetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omeapi.worker.v1beta1.WorkerService/SetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).SetState(ctx, req.(*SetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_QueryByState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryByStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).QueryByState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omeapi.worker.v1beta1.WorkerService/QueryByState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).QueryByState(ctx, req.(*QueryByStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkerService_ServiceDesc is the grpc.ServiceDesc for WorkerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "omeapi.worker.v1beta1.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWorker",
			Handler:    _WorkerService_GetWorker_Handler,
		},
		{
			MethodName: "SetState",
			Handler:    _WorkerService_SetState_Handler,
		},
		{
			MethodName: "QueryByState",
			Handler:    _WorkerService_QueryByState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "omeapi/worker/v1beta1/worker.proto",
}
