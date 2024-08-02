// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.3
// source: yandex/cloud/logging/v1/log_ingestion_service.proto

package logging

import (
	_ "github.com/doublecloud/tross/cloud/bitbucket/public-api/yandex/cloud"
	context "context"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Log entries destination.
	//
	// See [Destination] for details.
	Destination *Destination `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	// Common resource (type, ID) specification for log entries.
	Resource *LogEntryResource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	// List of log entries.
	Entries []*IncomingLogEntry `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty"`
	// Log entries defaults.
	//
	// See [LogEntryDefaults] for details.
	Defaults *LogEntryDefaults `protobuf:"bytes,4,opt,name=defaults,proto3" json:"defaults,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_logging_v1_log_ingestion_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_logging_v1_log_ingestion_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescGZIP(), []int{0}
}

func (x *WriteRequest) GetDestination() *Destination {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *WriteRequest) GetResource() *LogEntryResource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *WriteRequest) GetEntries() []*IncomingLogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *WriteRequest) GetDefaults() *LogEntryDefaults {
	if x != nil {
		return x.Defaults
	}
	return nil
}

type WriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Map<idx, status> of ingest failures.
	//
	// If entry with idx N is absent, it was ingested successfully.
	Errors map[int64]*status.Status `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_logging_v1_log_ingestion_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_logging_v1_log_ingestion_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescGZIP(), []int{1}
}

func (x *WriteResponse) GetErrors() map[int64]*status.Status {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_yandex_cloud_logging_v1_log_ingestion_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDesc = []byte{
	0x0a, 0x33, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x6e,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x02, 0x0a, 0x0c,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x4e, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63,
	0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x09, 0x82,
	0xc8, 0x31, 0x05, 0x31, 0x2d, 0x31, 0x30, 0x30, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x45, 0x0a, 0x08, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x08,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x22, 0xaa, 0x01, 0x0a, 0x0d, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x4d, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x6d, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x05,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6a, 0x0a, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x5a, 0x4b, 0x61, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74, 0x65,
	0x61, 0x6d, 0x2e, 0x72, 0x75, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x74, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescData = file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDesc
)

func file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescData)
	})
	return file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescData
}

var file_yandex_cloud_logging_v1_log_ingestion_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_logging_v1_log_ingestion_service_proto_goTypes = []interface{}{
	(*WriteRequest)(nil),     // 0: yandex.cloud.logging.v1.WriteRequest
	(*WriteResponse)(nil),    // 1: yandex.cloud.logging.v1.WriteResponse
	nil,                      // 2: yandex.cloud.logging.v1.WriteResponse.ErrorsEntry
	(*Destination)(nil),      // 3: yandex.cloud.logging.v1.Destination
	(*LogEntryResource)(nil), // 4: yandex.cloud.logging.v1.LogEntryResource
	(*IncomingLogEntry)(nil), // 5: yandex.cloud.logging.v1.IncomingLogEntry
	(*LogEntryDefaults)(nil), // 6: yandex.cloud.logging.v1.LogEntryDefaults
	(*status.Status)(nil),    // 7: google.rpc.Status
}
var file_yandex_cloud_logging_v1_log_ingestion_service_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.logging.v1.WriteRequest.destination:type_name -> yandex.cloud.logging.v1.Destination
	4, // 1: yandex.cloud.logging.v1.WriteRequest.resource:type_name -> yandex.cloud.logging.v1.LogEntryResource
	5, // 2: yandex.cloud.logging.v1.WriteRequest.entries:type_name -> yandex.cloud.logging.v1.IncomingLogEntry
	6, // 3: yandex.cloud.logging.v1.WriteRequest.defaults:type_name -> yandex.cloud.logging.v1.LogEntryDefaults
	2, // 4: yandex.cloud.logging.v1.WriteResponse.errors:type_name -> yandex.cloud.logging.v1.WriteResponse.ErrorsEntry
	7, // 5: yandex.cloud.logging.v1.WriteResponse.ErrorsEntry.value:type_name -> google.rpc.Status
	0, // 6: yandex.cloud.logging.v1.LogIngestionService.Write:input_type -> yandex.cloud.logging.v1.WriteRequest
	1, // 7: yandex.cloud.logging.v1.LogIngestionService.Write:output_type -> yandex.cloud.logging.v1.WriteResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_yandex_cloud_logging_v1_log_ingestion_service_proto_init() }
func file_yandex_cloud_logging_v1_log_ingestion_service_proto_init() {
	if File_yandex_cloud_logging_v1_log_ingestion_service_proto != nil {
		return
	}
	file_yandex_cloud_logging_v1_log_entry_proto_init()
	file_yandex_cloud_logging_v1_log_resource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_logging_v1_log_ingestion_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_yandex_cloud_logging_v1_log_ingestion_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteResponse); i {
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
			RawDescriptor: file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_logging_v1_log_ingestion_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_logging_v1_log_ingestion_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_logging_v1_log_ingestion_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_logging_v1_log_ingestion_service_proto = out.File
	file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDesc = nil
	file_yandex_cloud_logging_v1_log_ingestion_service_proto_goTypes = nil
	file_yandex_cloud_logging_v1_log_ingestion_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogIngestionServiceClient is the client API for LogIngestionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogIngestionServiceClient interface {
	// Write log entries to specified destination.
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
}

type logIngestionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogIngestionServiceClient(cc grpc.ClientConnInterface) LogIngestionServiceClient {
	return &logIngestionServiceClient{cc}
}

func (c *logIngestionServiceClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.logging.v1.LogIngestionService/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogIngestionServiceServer is the server API for LogIngestionService service.
type LogIngestionServiceServer interface {
	// Write log entries to specified destination.
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
}

// UnimplementedLogIngestionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogIngestionServiceServer struct {
}

func (*UnimplementedLogIngestionServiceServer) Write(context.Context, *WriteRequest) (*WriteResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method Write not implemented")
}

func RegisterLogIngestionServiceServer(s *grpc.Server, srv LogIngestionServiceServer) {
	s.RegisterService(&_LogIngestionService_serviceDesc, srv)
}

func _LogIngestionService_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogIngestionServiceServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.logging.v1.LogIngestionService/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogIngestionServiceServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogIngestionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.logging.v1.LogIngestionService",
	HandlerType: (*LogIngestionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _LogIngestionService_Write_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/logging/v1/log_ingestion_service.proto",
}
