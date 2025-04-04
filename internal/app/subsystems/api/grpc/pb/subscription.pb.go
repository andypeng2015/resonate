// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.3
// source: internal/app/subsystems/api/grpc/pb/subscription.proto

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

type CreateSubscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PromiseId string `protobuf:"bytes,2,opt,name=promiseId,proto3" json:"promiseId,omitempty"`
	Timeout   int64  `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Recv      *Recv  `protobuf:"bytes,5,opt,name=recv,proto3" json:"recv,omitempty"`
	RequestId string `protobuf:"bytes,6,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

func (x *CreateSubscriptionRequest) Reset() {
	*x = CreateSubscriptionRequest{}
	mi := &file_internal_app_subsystems_api_grpc_pb_subscription_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubscriptionRequest) ProtoMessage() {}

func (x *CreateSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_subsystems_api_grpc_pb_subscription_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*CreateSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_subsystems_api_grpc_pb_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSubscriptionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateSubscriptionRequest) GetPromiseId() string {
	if x != nil {
		return x.PromiseId
	}
	return ""
}

func (x *CreateSubscriptionRequest) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *CreateSubscriptionRequest) GetRecv() *Recv {
	if x != nil {
		return x.Recv
	}
	return nil
}

func (x *CreateSubscriptionRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type CreateSubscriptionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Noop     bool      `protobuf:"varint,1,opt,name=noop,proto3" json:"noop,omitempty"`
	Callback *Callback `protobuf:"bytes,2,opt,name=callback,proto3" json:"callback,omitempty"`
	Promise  *Promise  `protobuf:"bytes,3,opt,name=promise,proto3" json:"promise,omitempty"`
}

func (x *CreateSubscriptionResponse) Reset() {
	*x = CreateSubscriptionResponse{}
	mi := &file_internal_app_subsystems_api_grpc_pb_subscription_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSubscriptionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubscriptionResponse) ProtoMessage() {}

func (x *CreateSubscriptionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_subsystems_api_grpc_pb_subscription_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubscriptionResponse.ProtoReflect.Descriptor instead.
func (*CreateSubscriptionResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_subsystems_api_grpc_pb_subscription_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSubscriptionResponse) GetNoop() bool {
	if x != nil {
		return x.Noop
	}
	return false
}

func (x *CreateSubscriptionResponse) GetCallback() *Callback {
	if x != nil {
		return x.Callback
	}
	return nil
}

func (x *CreateSubscriptionResponse) GetPromise() *Promise {
	if x != nil {
		return x.Promise
	}
	return nil
}

var File_internal_app_subsystems_api_grpc_pb_subscription_proto protoreflect.FileDescriptor

var file_internal_app_subsystems_api_grpc_pb_subscription_proto_rawDesc = []byte{
	0x0a, 0x36, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73,
	0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x34, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x5f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x72, 0x65, 0x63, 0x76, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x5f, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x76, 0x52, 0x04, 0x72, 0x65, 0x63, 0x76, 0x12, 0x1c, 0x0a,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x1a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f,
	0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6e, 0x6f, 0x6f, 0x70, 0x12, 0x30,
	0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x2e, 0x43, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x12, 0x2c, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x5f, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x6d, 0x69, 0x73, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x32, 0x7a,
	0x0a, 0x0d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x69, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x74,
	0x65, 0x68, 0x71, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_app_subsystems_api_grpc_pb_subscription_proto_rawDescOnce sync.Once
	file_internal_app_subsystems_api_grpc_pb_subscription_proto_rawDescData = file_internal_app_subsystems_api_grpc_pb_subscription_proto_rawDesc
)

func file_internal_app_subsystems_api_grpc_pb_subscription_proto_rawDescGZIP() []byte {
	file_internal_app_subsystems_api_grpc_pb_subscription_proto_rawDescOnce.Do(func() {
		file_internal_app_subsystems_api_grpc_pb_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_app_subsystems_api_grpc_pb_subscription_proto_rawDescData)
	})
	return file_internal_app_subsystems_api_grpc_pb_subscription_proto_rawDescData
}

var file_internal_app_subsystems_api_grpc_pb_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_app_subsystems_api_grpc_pb_subscription_proto_goTypes = []any{
	(*CreateSubscriptionRequest)(nil),  // 0: subscription.CreateSubscriptionRequest
	(*CreateSubscriptionResponse)(nil), // 1: subscription.CreateSubscriptionResponse
	(*Recv)(nil),                       // 2: callback_t.Recv
	(*Callback)(nil),                   // 3: callback_t.Callback
	(*Promise)(nil),                    // 4: promise_t.Promise
}
var file_internal_app_subsystems_api_grpc_pb_subscription_proto_depIdxs = []int32{
	2, // 0: subscription.CreateSubscriptionRequest.recv:type_name -> callback_t.Recv
	3, // 1: subscription.CreateSubscriptionResponse.callback:type_name -> callback_t.Callback
	4, // 2: subscription.CreateSubscriptionResponse.promise:type_name -> promise_t.Promise
	0, // 3: subscription.Subscriptions.CreateSubscription:input_type -> subscription.CreateSubscriptionRequest
	1, // 4: subscription.Subscriptions.CreateSubscription:output_type -> subscription.CreateSubscriptionResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_app_subsystems_api_grpc_pb_subscription_proto_init() }
func file_internal_app_subsystems_api_grpc_pb_subscription_proto_init() {
	if File_internal_app_subsystems_api_grpc_pb_subscription_proto != nil {
		return
	}
	file_internal_app_subsystems_api_grpc_pb_callback_t_proto_init()
	file_internal_app_subsystems_api_grpc_pb_promise_t_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_app_subsystems_api_grpc_pb_subscription_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_app_subsystems_api_grpc_pb_subscription_proto_goTypes,
		DependencyIndexes: file_internal_app_subsystems_api_grpc_pb_subscription_proto_depIdxs,
		MessageInfos:      file_internal_app_subsystems_api_grpc_pb_subscription_proto_msgTypes,
	}.Build()
	File_internal_app_subsystems_api_grpc_pb_subscription_proto = out.File
	file_internal_app_subsystems_api_grpc_pb_subscription_proto_rawDesc = nil
	file_internal_app_subsystems_api_grpc_pb_subscription_proto_goTypes = nil
	file_internal_app_subsystems_api_grpc_pb_subscription_proto_depIdxs = nil
}
