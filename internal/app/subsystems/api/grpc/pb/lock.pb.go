// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: internal/app/subsystems/api/grpc/pb/lock.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AcquireLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId  string `protobuf:"bytes,1,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
	ExecutionId string `protobuf:"bytes,2,opt,name=executionId,proto3" json:"executionId,omitempty"`
	ProcessId   string `protobuf:"bytes,3,opt,name=processId,proto3" json:"processId,omitempty"`
	Ttl         int64  `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	RequestId   string `protobuf:"bytes,5,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

func (x *AcquireLockRequest) Reset() {
	*x = AcquireLockRequest{}
	mi := &file_internal_app_subsystems_api_grpc_pb_lock_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AcquireLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcquireLockRequest) ProtoMessage() {}

func (x *AcquireLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_subsystems_api_grpc_pb_lock_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcquireLockRequest.ProtoReflect.Descriptor instead.
func (*AcquireLockRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDescGZIP(), []int{0}
}

func (x *AcquireLockRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *AcquireLockRequest) GetExecutionId() string {
	if x != nil {
		return x.ExecutionId
	}
	return ""
}

func (x *AcquireLockRequest) GetProcessId() string {
	if x != nil {
		return x.ProcessId
	}
	return ""
}

func (x *AcquireLockRequest) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *AcquireLockRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type AcquireLockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Acquired bool `protobuf:"varint,1,opt,name=acquired,proto3" json:"acquired,omitempty"`
}

func (x *AcquireLockResponse) Reset() {
	*x = AcquireLockResponse{}
	mi := &file_internal_app_subsystems_api_grpc_pb_lock_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AcquireLockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcquireLockResponse) ProtoMessage() {}

func (x *AcquireLockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_subsystems_api_grpc_pb_lock_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcquireLockResponse.ProtoReflect.Descriptor instead.
func (*AcquireLockResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDescGZIP(), []int{1}
}

func (x *AcquireLockResponse) GetAcquired() bool {
	if x != nil {
		return x.Acquired
	}
	return false
}

type ReleaseLockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId  string `protobuf:"bytes,1,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
	ExecutionId string `protobuf:"bytes,2,opt,name=executionId,proto3" json:"executionId,omitempty"`
	RequestId   string `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

func (x *ReleaseLockRequest) Reset() {
	*x = ReleaseLockRequest{}
	mi := &file_internal_app_subsystems_api_grpc_pb_lock_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReleaseLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseLockRequest) ProtoMessage() {}

func (x *ReleaseLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_subsystems_api_grpc_pb_lock_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseLockRequest.ProtoReflect.Descriptor instead.
func (*ReleaseLockRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDescGZIP(), []int{2}
}

func (x *ReleaseLockRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *ReleaseLockRequest) GetExecutionId() string {
	if x != nil {
		return x.ExecutionId
	}
	return ""
}

func (x *ReleaseLockRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type ReleaseLockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Released bool `protobuf:"varint,1,opt,name=released,proto3" json:"released,omitempty"`
}

func (x *ReleaseLockResponse) Reset() {
	*x = ReleaseLockResponse{}
	mi := &file_internal_app_subsystems_api_grpc_pb_lock_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReleaseLockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseLockResponse) ProtoMessage() {}

func (x *ReleaseLockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_subsystems_api_grpc_pb_lock_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseLockResponse.ProtoReflect.Descriptor instead.
func (*ReleaseLockResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDescGZIP(), []int{3}
}

func (x *ReleaseLockResponse) GetReleased() bool {
	if x != nil {
		return x.Released
	}
	return false
}

type HeartbeatLocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcessId string `protobuf:"bytes,1,opt,name=processId,proto3" json:"processId,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

func (x *HeartbeatLocksRequest) Reset() {
	*x = HeartbeatLocksRequest{}
	mi := &file_internal_app_subsystems_api_grpc_pb_lock_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatLocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatLocksRequest) ProtoMessage() {}

func (x *HeartbeatLocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_subsystems_api_grpc_pb_lock_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatLocksRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatLocksRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDescGZIP(), []int{4}
}

func (x *HeartbeatLocksRequest) GetProcessId() string {
	if x != nil {
		return x.ProcessId
	}
	return ""
}

func (x *HeartbeatLocksRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type HeartbeatLocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocksAffected int32 `protobuf:"varint,1,opt,name=locksAffected,proto3" json:"locksAffected,omitempty"`
}

func (x *HeartbeatLocksResponse) Reset() {
	*x = HeartbeatLocksResponse{}
	mi := &file_internal_app_subsystems_api_grpc_pb_lock_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatLocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatLocksResponse) ProtoMessage() {}

func (x *HeartbeatLocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_subsystems_api_grpc_pb_lock_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatLocksResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatLocksResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDescGZIP(), []int{5}
}

func (x *HeartbeatLocksResponse) GetLocksAffected() int32 {
	if x != nil {
		return x.LocksAffected
	}
	return 0
}

var File_internal_app_subsystems_api_grpc_pb_lock_proto protoreflect.FileDescriptor

var file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73,
	0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0xa4, 0x01, 0x0a, 0x12, 0x41, 0x63, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x74, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x31, 0x0a,
	0x13, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x22, 0x74, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x22, 0x53, 0x0a, 0x15, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x3e,
	0x0a, 0x16, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x32, 0xe2,
	0x01, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x44, 0x0a, 0x0b, 0x41, 0x63, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x41,
	0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x2e,
	0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x68, 0x71, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDescOnce sync.Once
	file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDescData = file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDesc
)

func file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDescGZIP() []byte {
	file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDescOnce.Do(func() {
		file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDescData)
	})
	return file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDescData
}

var file_internal_app_subsystems_api_grpc_pb_lock_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_app_subsystems_api_grpc_pb_lock_proto_goTypes = []any{
	(*AcquireLockRequest)(nil),     // 0: lock.AcquireLockRequest
	(*AcquireLockResponse)(nil),    // 1: lock.AcquireLockResponse
	(*ReleaseLockRequest)(nil),     // 2: lock.ReleaseLockRequest
	(*ReleaseLockResponse)(nil),    // 3: lock.ReleaseLockResponse
	(*HeartbeatLocksRequest)(nil),  // 4: lock.HeartbeatLocksRequest
	(*HeartbeatLocksResponse)(nil), // 5: lock.HeartbeatLocksResponse
}
var file_internal_app_subsystems_api_grpc_pb_lock_proto_depIdxs = []int32{
	0, // 0: lock.Locks.AcquireLock:input_type -> lock.AcquireLockRequest
	2, // 1: lock.Locks.ReleaseLock:input_type -> lock.ReleaseLockRequest
	4, // 2: lock.Locks.HeartbeatLocks:input_type -> lock.HeartbeatLocksRequest
	1, // 3: lock.Locks.AcquireLock:output_type -> lock.AcquireLockResponse
	3, // 4: lock.Locks.ReleaseLock:output_type -> lock.ReleaseLockResponse
	5, // 5: lock.Locks.HeartbeatLocks:output_type -> lock.HeartbeatLocksResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_app_subsystems_api_grpc_pb_lock_proto_init() }
func file_internal_app_subsystems_api_grpc_pb_lock_proto_init() {
	if File_internal_app_subsystems_api_grpc_pb_lock_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_app_subsystems_api_grpc_pb_lock_proto_goTypes,
		DependencyIndexes: file_internal_app_subsystems_api_grpc_pb_lock_proto_depIdxs,
		MessageInfos:      file_internal_app_subsystems_api_grpc_pb_lock_proto_msgTypes,
	}.Build()
	File_internal_app_subsystems_api_grpc_pb_lock_proto = out.File
	file_internal_app_subsystems_api_grpc_pb_lock_proto_rawDesc = nil
	file_internal_app_subsystems_api_grpc_pb_lock_proto_goTypes = nil
	file_internal_app_subsystems_api_grpc_pb_lock_proto_depIdxs = nil
}
