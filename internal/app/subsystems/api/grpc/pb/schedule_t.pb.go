// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: internal/app/subsystems/api/grpc/pb/schedule_t.proto

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

type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description    string            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Cron           string            `protobuf:"bytes,3,opt,name=cron,proto3" json:"cron,omitempty"`
	Tags           map[string]string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PromiseId      string            `protobuf:"bytes,5,opt,name=promiseId,proto3" json:"promiseId,omitempty"`
	PromiseTimeout int64             `protobuf:"varint,6,opt,name=promiseTimeout,proto3" json:"promiseTimeout,omitempty"`
	PromiseParam   *Value            `protobuf:"bytes,7,opt,name=promiseParam,proto3" json:"promiseParam,omitempty"`
	PromiseTags    map[string]string `protobuf:"bytes,8,rep,name=promiseTags,proto3" json:"promiseTags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IdempotencyKey string            `protobuf:"bytes,9,opt,name=idempotencyKey,proto3" json:"idempotencyKey,omitempty"`
	LastRunTime    int64             `protobuf:"varint,10,opt,name=lastRunTime,proto3" json:"lastRunTime,omitempty"`
	NextRunTime    int64             `protobuf:"varint,11,opt,name=nextRunTime,proto3" json:"nextRunTime,omitempty"`
	CreatedOn      int64             `protobuf:"varint,12,opt,name=createdOn,proto3" json:"createdOn,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_rawDescGZIP(), []int{0}
}

func (x *Schedule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Schedule) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Schedule) GetCron() string {
	if x != nil {
		return x.Cron
	}
	return ""
}

func (x *Schedule) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Schedule) GetPromiseId() string {
	if x != nil {
		return x.PromiseId
	}
	return ""
}

func (x *Schedule) GetPromiseTimeout() int64 {
	if x != nil {
		return x.PromiseTimeout
	}
	return 0
}

func (x *Schedule) GetPromiseParam() *Value {
	if x != nil {
		return x.PromiseParam
	}
	return nil
}

func (x *Schedule) GetPromiseTags() map[string]string {
	if x != nil {
		return x.PromiseTags
	}
	return nil
}

func (x *Schedule) GetIdempotencyKey() string {
	if x != nil {
		return x.IdempotencyKey
	}
	return ""
}

func (x *Schedule) GetLastRunTime() int64 {
	if x != nil {
		return x.LastRunTime
	}
	return 0
}

func (x *Schedule) GetNextRunTime() int64 {
	if x != nil {
		return x.NextRunTime
	}
	return 0
}

func (x *Schedule) GetCreatedOn() int64 {
	if x != nil {
		return x.CreatedOn
	}
	return 0
}

var File_internal_app_subsystems_api_grpc_pb_schedule_t_proto protoreflect.FileDescriptor

var file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_rawDesc = []byte{
	0x0a, 0x34, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73,
	0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x5f, 0x74, 0x1a, 0x33, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x5f,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x04, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x5f, 0x74, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x54,
	0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e,
	0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x34, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x6d, 0x69, 0x73, 0x65, 0x5f, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x6d, 0x69, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x47, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x6d, 0x69, 0x73, 0x65, 0x54, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x74, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x54, 0x61, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x54,
	0x61, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e,
	0x63, 0x79, 0x4b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x65,
	0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x1a, 0x37, 0x0a,
	0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73,
	0x65, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x68, 0x71, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_rawDescOnce sync.Once
	file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_rawDescData = file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_rawDesc
)

func file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_rawDescGZIP() []byte {
	file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_rawDescOnce.Do(func() {
		file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_rawDescData)
	})
	return file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_rawDescData
}

var file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_goTypes = []interface{}{
	(*Schedule)(nil), // 0: schedule_t.Schedule
	nil,              // 1: schedule_t.Schedule.TagsEntry
	nil,              // 2: schedule_t.Schedule.PromiseTagsEntry
	(*Value)(nil),    // 3: promise_t.Value
}
var file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_depIdxs = []int32{
	1, // 0: schedule_t.Schedule.tags:type_name -> schedule_t.Schedule.TagsEntry
	3, // 1: schedule_t.Schedule.promiseParam:type_name -> promise_t.Value
	2, // 2: schedule_t.Schedule.promiseTags:type_name -> schedule_t.Schedule.PromiseTagsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_init() }
func file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_init() {
	if File_internal_app_subsystems_api_grpc_pb_schedule_t_proto != nil {
		return
	}
	file_internal_app_subsystems_api_grpc_pb_promise_t_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_goTypes,
		DependencyIndexes: file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_depIdxs,
		MessageInfos:      file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_msgTypes,
	}.Build()
	File_internal_app_subsystems_api_grpc_pb_schedule_t_proto = out.File
	file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_rawDesc = nil
	file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_goTypes = nil
	file_internal_app_subsystems_api_grpc_pb_schedule_t_proto_depIdxs = nil
}