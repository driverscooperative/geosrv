// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: omeapi/location/v1beta1/location_update.proto

package v1beta1

import (
	v1beta1 "github.com/openmarketplaceengine/openmarketplaceengine/internal/omeapi/type/v1beta1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type LocationUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId   string                 `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	Location   *v1beta1.Location      `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *LocationUpdate) Reset() {
	*x = LocationUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_location_v1beta1_location_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationUpdate) ProtoMessage() {}

func (x *LocationUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_location_v1beta1_location_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationUpdate.ProtoReflect.Descriptor instead.
func (*LocationUpdate) Descriptor() ([]byte, []int) {
	return file_omeapi_location_v1beta1_location_update_proto_rawDescGZIP(), []int{0}
}

func (x *LocationUpdate) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *LocationUpdate) GetLocation() *v1beta1.Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *LocationUpdate) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_omeapi_location_v1beta1_location_update_proto protoreflect.FileDescriptor

var file_omeapi_location_v1beta1_location_update_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6f, 0x6d, 0x65, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01,
	0x0a, 0x0e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x59, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_omeapi_location_v1beta1_location_update_proto_rawDescOnce sync.Once
	file_omeapi_location_v1beta1_location_update_proto_rawDescData = file_omeapi_location_v1beta1_location_update_proto_rawDesc
)

func file_omeapi_location_v1beta1_location_update_proto_rawDescGZIP() []byte {
	file_omeapi_location_v1beta1_location_update_proto_rawDescOnce.Do(func() {
		file_omeapi_location_v1beta1_location_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_omeapi_location_v1beta1_location_update_proto_rawDescData)
	})
	return file_omeapi_location_v1beta1_location_update_proto_rawDescData
}

var file_omeapi_location_v1beta1_location_update_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_omeapi_location_v1beta1_location_update_proto_goTypes = []interface{}{
	(*LocationUpdate)(nil),        // 0: omeapi.location.v1beta1.LocationUpdate
	(*v1beta1.Location)(nil),      // 1: omeapi.type.v1beta1.Location
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_omeapi_location_v1beta1_location_update_proto_depIdxs = []int32{
	1, // 0: omeapi.location.v1beta1.LocationUpdate.location:type_name -> omeapi.type.v1beta1.Location
	2, // 1: omeapi.location.v1beta1.LocationUpdate.update_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_omeapi_location_v1beta1_location_update_proto_init() }
func file_omeapi_location_v1beta1_location_update_proto_init() {
	if File_omeapi_location_v1beta1_location_update_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_omeapi_location_v1beta1_location_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationUpdate); i {
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
			RawDescriptor: file_omeapi_location_v1beta1_location_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_omeapi_location_v1beta1_location_update_proto_goTypes,
		DependencyIndexes: file_omeapi_location_v1beta1_location_update_proto_depIdxs,
		MessageInfos:      file_omeapi_location_v1beta1_location_update_proto_msgTypes,
	}.Build()
	File_omeapi_location_v1beta1_location_update_proto = out.File
	file_omeapi_location_v1beta1_location_update_proto_rawDesc = nil
	file_omeapi_location_v1beta1_location_update_proto_goTypes = nil
	file_omeapi_location_v1beta1_location_update_proto_depIdxs = nil
}
