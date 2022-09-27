// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api/crossing/v1beta1/crossing.proto

package v1beta1

import (
	v1beta1 "github.com/driverscooperative/geosrv/internal/idl/api/type/v1beta1"
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

type ListCrossingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TollgateId string `protobuf:"bytes,1,opt,name=tollgate_id,json=tollgateId,proto3" json:"tollgate_id,omitempty"`
	WorkerId   string `protobuf:"bytes,2,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	PageSize   int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken  string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListCrossingsRequest) Reset() {
	*x = ListCrossingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crossing_v1beta1_crossing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCrossingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCrossingsRequest) ProtoMessage() {}

func (x *ListCrossingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_crossing_v1beta1_crossing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCrossingsRequest.ProtoReflect.Descriptor instead.
func (*ListCrossingsRequest) Descriptor() ([]byte, []int) {
	return file_api_crossing_v1beta1_crossing_proto_rawDescGZIP(), []int{0}
}

func (x *ListCrossingsRequest) GetTollgateId() string {
	if x != nil {
		return x.TollgateId
	}
	return ""
}

func (x *ListCrossingsRequest) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *ListCrossingsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCrossingsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListCrossingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crossings     []*v1beta1.Crossing `protobuf:"bytes,1,rep,name=crossings,proto3" json:"crossings,omitempty"`
	NextPageToken string              `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListCrossingsResponse) Reset() {
	*x = ListCrossingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crossing_v1beta1_crossing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCrossingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCrossingsResponse) ProtoMessage() {}

func (x *ListCrossingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_crossing_v1beta1_crossing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCrossingsResponse.ProtoReflect.Descriptor instead.
func (*ListCrossingsResponse) Descriptor() ([]byte, []int) {
	return file_api_crossing_v1beta1_crossing_proto_rawDescGZIP(), []int{1}
}

func (x *ListCrossingsResponse) GetCrossings() []*v1beta1.Crossing {
	if x != nil {
		return x.Crossings
	}
	return nil
}

func (x *ListCrossingsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_api_crossing_v1beta1_crossing_proto protoreflect.FileDescriptor

var file_api_crossing_v1beta1_crossing_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x72,
	0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x6c, 0x6c,
	0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x79, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x6f, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43,
	0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x7d, 0x0a, 0x0f, 0x43, 0x72,
	0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x5a, 0x5a, 0x58, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x64, 0x6c,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_crossing_v1beta1_crossing_proto_rawDescOnce sync.Once
	file_api_crossing_v1beta1_crossing_proto_rawDescData = file_api_crossing_v1beta1_crossing_proto_rawDesc
)

func file_api_crossing_v1beta1_crossing_proto_rawDescGZIP() []byte {
	file_api_crossing_v1beta1_crossing_proto_rawDescOnce.Do(func() {
		file_api_crossing_v1beta1_crossing_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_crossing_v1beta1_crossing_proto_rawDescData)
	})
	return file_api_crossing_v1beta1_crossing_proto_rawDescData
}

var file_api_crossing_v1beta1_crossing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_crossing_v1beta1_crossing_proto_goTypes = []interface{}{
	(*ListCrossingsRequest)(nil),  // 0: api.crossing.v1beta1.ListCrossingsRequest
	(*ListCrossingsResponse)(nil), // 1: api.crossing.v1beta1.ListCrossingsResponse
	(*v1beta1.Crossing)(nil),      // 2: api.type.v1beta1.Crossing
}
var file_api_crossing_v1beta1_crossing_proto_depIdxs = []int32{
	2, // 0: api.crossing.v1beta1.ListCrossingsResponse.crossings:type_name -> api.type.v1beta1.Crossing
	0, // 1: api.crossing.v1beta1.CrossingService.ListCrossings:input_type -> api.crossing.v1beta1.ListCrossingsRequest
	1, // 2: api.crossing.v1beta1.CrossingService.ListCrossings:output_type -> api.crossing.v1beta1.ListCrossingsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_crossing_v1beta1_crossing_proto_init() }
func file_api_crossing_v1beta1_crossing_proto_init() {
	if File_api_crossing_v1beta1_crossing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_crossing_v1beta1_crossing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCrossingsRequest); i {
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
		file_api_crossing_v1beta1_crossing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCrossingsResponse); i {
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
			RawDescriptor: file_api_crossing_v1beta1_crossing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_crossing_v1beta1_crossing_proto_goTypes,
		DependencyIndexes: file_api_crossing_v1beta1_crossing_proto_depIdxs,
		MessageInfos:      file_api_crossing_v1beta1_crossing_proto_msgTypes,
	}.Build()
	File_api_crossing_v1beta1_crossing_proto = out.File
	file_api_crossing_v1beta1_crossing_proto_rawDesc = nil
	file_api_crossing_v1beta1_crossing_proto_goTypes = nil
	file_api_crossing_v1beta1_crossing_proto_depIdxs = nil
}
