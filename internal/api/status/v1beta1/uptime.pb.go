// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: api/status/v1beta1/uptime.proto

package v1beta1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetUptimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUptimeRequest) Reset() {
	*x = GetUptimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_status_v1beta1_uptime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUptimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUptimeRequest) ProtoMessage() {}

func (x *GetUptimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_status_v1beta1_uptime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUptimeRequest.ProtoReflect.Descriptor instead.
func (*GetUptimeRequest) Descriptor() ([]byte, []int) {
	return file_api_status_v1beta1_uptime_proto_rawDescGZIP(), []int{0}
}

type GetUptimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uptime  int64                  `protobuf:"varint,1,opt,name=uptime,proto3" json:"uptime,omitempty"`
	Started *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=started,proto3" json:"started,omitempty"`
}

func (x *GetUptimeResponse) Reset() {
	*x = GetUptimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_status_v1beta1_uptime_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUptimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUptimeResponse) ProtoMessage() {}

func (x *GetUptimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_status_v1beta1_uptime_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUptimeResponse.ProtoReflect.Descriptor instead.
func (*GetUptimeResponse) Descriptor() ([]byte, []int) {
	return file_api_status_v1beta1_uptime_proto_rawDescGZIP(), []int{1}
}

func (x *GetUptimeResponse) GetUptime() int64 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

func (x *GetUptimeResponse) GetStarted() *timestamppb.Timestamp {
	if x != nil {
		return x.Started
	}
	return nil
}

var File_api_status_v1beta1_uptime_proto protoreflect.FileDescriptor

var file_api_status_v1beta1_uptime_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x61, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x32, 0x6b, 0x0a, 0x0d, 0x55, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_status_v1beta1_uptime_proto_rawDescOnce sync.Once
	file_api_status_v1beta1_uptime_proto_rawDescData = file_api_status_v1beta1_uptime_proto_rawDesc
)

func file_api_status_v1beta1_uptime_proto_rawDescGZIP() []byte {
	file_api_status_v1beta1_uptime_proto_rawDescOnce.Do(func() {
		file_api_status_v1beta1_uptime_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_status_v1beta1_uptime_proto_rawDescData)
	})
	return file_api_status_v1beta1_uptime_proto_rawDescData
}

var file_api_status_v1beta1_uptime_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_status_v1beta1_uptime_proto_goTypes = []interface{}{
	(*GetUptimeRequest)(nil),      // 0: api.status.v1beta1.GetUptimeRequest
	(*GetUptimeResponse)(nil),     // 1: api.status.v1beta1.GetUptimeResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_api_status_v1beta1_uptime_proto_depIdxs = []int32{
	2, // 0: api.status.v1beta1.GetUptimeResponse.started:type_name -> google.protobuf.Timestamp
	0, // 1: api.status.v1beta1.UptimeService.GetUptime:input_type -> api.status.v1beta1.GetUptimeRequest
	1, // 2: api.status.v1beta1.UptimeService.GetUptime:output_type -> api.status.v1beta1.GetUptimeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_status_v1beta1_uptime_proto_init() }
func file_api_status_v1beta1_uptime_proto_init() {
	if File_api_status_v1beta1_uptime_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_status_v1beta1_uptime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUptimeRequest); i {
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
		file_api_status_v1beta1_uptime_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUptimeResponse); i {
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
			RawDescriptor: file_api_status_v1beta1_uptime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_status_v1beta1_uptime_proto_goTypes,
		DependencyIndexes: file_api_status_v1beta1_uptime_proto_depIdxs,
		MessageInfos:      file_api_status_v1beta1_uptime_proto_msgTypes,
	}.Build()
	File_api_status_v1beta1_uptime_proto = out.File
	file_api_status_v1beta1_uptime_proto_rawDesc = nil
	file_api_status_v1beta1_uptime_proto_goTypes = nil
	file_api_status_v1beta1_uptime_proto_depIdxs = nil
}
