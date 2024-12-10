// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: internal/app/subsystems/api/grpc/pb/task.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Tasks_ClaimTask_FullMethodName      = "/task.Tasks/ClaimTask"
	Tasks_CompleteTask_FullMethodName   = "/task.Tasks/CompleteTask"
	Tasks_HeartbeatTasks_FullMethodName = "/task.Tasks/HeartbeatTasks"
)

// TasksClient is the client API for Tasks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TasksClient interface {
	ClaimTask(ctx context.Context, in *ClaimTaskRequest, opts ...grpc.CallOption) (*ClaimTaskResponse, error)
	CompleteTask(ctx context.Context, in *CompleteTaskRequest, opts ...grpc.CallOption) (*CompleteTaskResponse, error)
	HeartbeatTasks(ctx context.Context, in *HeartbeatTasksRequest, opts ...grpc.CallOption) (*HeartbeatTasksResponse, error)
}

type tasksClient struct {
	cc grpc.ClientConnInterface
}

func NewTasksClient(cc grpc.ClientConnInterface) TasksClient {
	return &tasksClient{cc}
}

func (c *tasksClient) ClaimTask(ctx context.Context, in *ClaimTaskRequest, opts ...grpc.CallOption) (*ClaimTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClaimTaskResponse)
	err := c.cc.Invoke(ctx, Tasks_ClaimTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) CompleteTask(ctx context.Context, in *CompleteTaskRequest, opts ...grpc.CallOption) (*CompleteTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CompleteTaskResponse)
	err := c.cc.Invoke(ctx, Tasks_CompleteTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) HeartbeatTasks(ctx context.Context, in *HeartbeatTasksRequest, opts ...grpc.CallOption) (*HeartbeatTasksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HeartbeatTasksResponse)
	err := c.cc.Invoke(ctx, Tasks_HeartbeatTasks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksServer is the server API for Tasks service.
// All implementations must embed UnimplementedTasksServer
// for forward compatibility.
type TasksServer interface {
	ClaimTask(context.Context, *ClaimTaskRequest) (*ClaimTaskResponse, error)
	CompleteTask(context.Context, *CompleteTaskRequest) (*CompleteTaskResponse, error)
	HeartbeatTasks(context.Context, *HeartbeatTasksRequest) (*HeartbeatTasksResponse, error)
	mustEmbedUnimplementedTasksServer()
}

// UnimplementedTasksServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTasksServer struct{}

func (UnimplementedTasksServer) ClaimTask(context.Context, *ClaimTaskRequest) (*ClaimTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimTask not implemented")
}
func (UnimplementedTasksServer) CompleteTask(context.Context, *CompleteTaskRequest) (*CompleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTask not implemented")
}
func (UnimplementedTasksServer) HeartbeatTasks(context.Context, *HeartbeatTasksRequest) (*HeartbeatTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartbeatTasks not implemented")
}
func (UnimplementedTasksServer) mustEmbedUnimplementedTasksServer() {}
func (UnimplementedTasksServer) testEmbeddedByValue()               {}

// UnsafeTasksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TasksServer will
// result in compilation errors.
type UnsafeTasksServer interface {
	mustEmbedUnimplementedTasksServer()
}

func RegisterTasksServer(s grpc.ServiceRegistrar, srv TasksServer) {
	// If the following call pancis, it indicates UnimplementedTasksServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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
		FullMethod: Tasks_ClaimTask_FullMethodName,
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
		FullMethod: Tasks_CompleteTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).CompleteTask(ctx, req.(*CompleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_HeartbeatTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).HeartbeatTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tasks_HeartbeatTasks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).HeartbeatTasks(ctx, req.(*HeartbeatTasksRequest))
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
		{
			MethodName: "HeartbeatTasks",
			Handler:    _Tasks_HeartbeatTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/app/subsystems/api/grpc/pb/task.proto",
}
