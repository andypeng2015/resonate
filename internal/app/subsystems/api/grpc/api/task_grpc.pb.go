// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: internal/app/subsystems/api/grpc/api/task.proto

package api

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

// TasksClient is the client API for Tasks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TasksClient interface {
	ClaimTask(ctx context.Context, in *ClaimTaskRequest, opts ...grpc.CallOption) (*ClaimTaskResponse, error)
	CompleteTask(ctx context.Context, in *CompleteTaskRequest, opts ...grpc.CallOption) (*CompleteTaskResponse, error)
}

type tasksClient struct {
	cc grpc.ClientConnInterface
}

func NewTasksClient(cc grpc.ClientConnInterface) TasksClient {
	return &tasksClient{cc}
}

func (c *tasksClient) ClaimTask(ctx context.Context, in *ClaimTaskRequest, opts ...grpc.CallOption) (*ClaimTaskResponse, error) {
	out := new(ClaimTaskResponse)
	err := c.cc.Invoke(ctx, "/task.Tasks/ClaimTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) CompleteTask(ctx context.Context, in *CompleteTaskRequest, opts ...grpc.CallOption) (*CompleteTaskResponse, error) {
	out := new(CompleteTaskResponse)
	err := c.cc.Invoke(ctx, "/task.Tasks/CompleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksServer is the server API for Tasks service.
// All implementations must embed UnimplementedTasksServer
// for forward compatibility
type TasksServer interface {
	ClaimTask(context.Context, *ClaimTaskRequest) (*ClaimTaskResponse, error)
	CompleteTask(context.Context, *CompleteTaskRequest) (*CompleteTaskResponse, error)
	mustEmbedUnimplementedTasksServer()
}

// UnimplementedTasksServer must be embedded to have forward compatible implementations.
type UnimplementedTasksServer struct {
}

func (UnimplementedTasksServer) ClaimTask(context.Context, *ClaimTaskRequest) (*ClaimTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimTask not implemented")
}
func (UnimplementedTasksServer) CompleteTask(context.Context, *CompleteTaskRequest) (*CompleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTask not implemented")
}
func (UnimplementedTasksServer) mustEmbedUnimplementedTasksServer() {}

// UnsafeTasksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TasksServer will
// result in compilation errors.
type UnsafeTasksServer interface {
	mustEmbedUnimplementedTasksServer()
}

func RegisterTasksServer(s grpc.ServiceRegistrar, srv TasksServer) {
	s.RegisterService(&Tasks_ServiceDesc, srv)
}

func _Tasks_ClaimTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).ClaimTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Tasks/ClaimTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).ClaimTask(ctx, req.(*ClaimTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_CompleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).CompleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.Tasks/CompleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).CompleteTask(ctx, req.(*CompleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tasks_ServiceDesc is the grpc.ServiceDesc for Tasks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tasks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.Tasks",
	HandlerType: (*TasksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClaimTask",
			Handler:    _Tasks_ClaimTask_Handler,
		},
		{
			MethodName: "CompleteTask",
			Handler:    _Tasks_CompleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/app/subsystems/api/grpc/api/task.proto",
}