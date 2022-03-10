// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: idl/ome/location/v1beta1/location.proto

package v1beta1

import (
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

type UpdateLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId  string  `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude  float64 `protobuf:"fixed64,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
}

func (x *UpdateLocationRequest) Reset() {
	*x = UpdateLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_ome_location_v1beta1_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLocationRequest) ProtoMessage() {}

func (x *UpdateLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_ome_location_v1beta1_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLocationRequest.ProtoReflect.Descriptor instead.
func (*UpdateLocationRequest) Descriptor() ([]byte, []int) {
	return file_idl_ome_location_v1beta1_location_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateLocationRequest) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *UpdateLocationRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *UpdateLocationRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

type UpdateLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId string `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
}

func (x *UpdateLocationResponse) Reset() {
	*x = UpdateLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_ome_location_v1beta1_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLocationResponse) ProtoMessage() {}

func (x *UpdateLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_ome_location_v1beta1_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLocationResponse.ProtoReflect.Descriptor instead.
func (*UpdateLocationResponse) Descriptor() ([]byte, []int) {
	return file_idl_ome_location_v1beta1_location_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateLocationResponse) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

type QueryLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId string `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
}

func (x *QueryLocationRequest) Reset() {
	*x = QueryLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_ome_location_v1beta1_location_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryLocationRequest) ProtoMessage() {}

func (x *QueryLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_ome_location_v1beta1_location_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryLocationRequest.ProtoReflect.Descriptor instead.
func (*QueryLocationRequest) Descriptor() ([]byte, []int) {
	return file_idl_ome_location_v1beta1_location_proto_rawDescGZIP(), []int{2}
}

func (x *QueryLocationRequest) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

type QueryLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId     string                 `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	Longitude    float64                `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude     float64                `protobuf:"fixed64,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	LastSeenTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_seen_time,json=lastSeenTime,proto3" json:"last_seen_time,omitempty"`
}

func (x *QueryLocationResponse) Reset() {
	*x = QueryLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_ome_location_v1beta1_location_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryLocationResponse) ProtoMessage() {}

func (x *QueryLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_ome_location_v1beta1_location_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryLocationResponse.ProtoReflect.Descriptor instead.
func (*QueryLocationResponse) Descriptor() ([]byte, []int) {
	return file_idl_ome_location_v1beta1_location_proto_rawDescGZIP(), []int{3}
}

func (x *QueryLocationResponse) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *QueryLocationResponse) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *QueryLocationResponse) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *QueryLocationResponse) GetLastSeenTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSeenTime
	}
	return nil
}

var File_idl_ome_location_v1beta1_location_proto protoreflect.FileDescriptor

var file_idl_ome_location_v1beta1_location_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x64, 0x6c, 0x2f, 0x6f, 0x6d, 0x65, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6f, 0x6d, 0x65, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6e, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x22, 0x35, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x22, 0xb0, 0x01, 0x0a,
	0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x40, 0x0a,
	0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x32,
	0xec, 0x01, 0x0a, 0x0f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x6f, 0x6d, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6f, 0x6d, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x6a, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x6f, 0x6d, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x6f, 0x6d, 0x65, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x16,
	0x5a, 0x14, 0x6f, 0x6d, 0x65, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_ome_location_v1beta1_location_proto_rawDescOnce sync.Once
	file_idl_ome_location_v1beta1_location_proto_rawDescData = file_idl_ome_location_v1beta1_location_proto_rawDesc
)

func file_idl_ome_location_v1beta1_location_proto_rawDescGZIP() []byte {
	file_idl_ome_location_v1beta1_location_proto_rawDescOnce.Do(func() {
		file_idl_ome_location_v1beta1_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_ome_location_v1beta1_location_proto_rawDescData)
	})
	return file_idl_ome_location_v1beta1_location_proto_rawDescData
}

var file_idl_ome_location_v1beta1_location_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_idl_ome_location_v1beta1_location_proto_goTypes = []interface{}{
	(*UpdateLocationRequest)(nil),  // 0: ome.location.v1beta1.UpdateLocationRequest
	(*UpdateLocationResponse)(nil), // 1: ome.location.v1beta1.UpdateLocationResponse
	(*QueryLocationRequest)(nil),   // 2: ome.location.v1beta1.QueryLocationRequest
	(*QueryLocationResponse)(nil),  // 3: ome.location.v1beta1.QueryLocationResponse
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
}
var file_idl_ome_location_v1beta1_location_proto_depIdxs = []int32{
	4, // 0: ome.location.v1beta1.QueryLocationResponse.last_seen_time:type_name -> google.protobuf.Timestamp
	0, // 1: ome.location.v1beta1.LocationService.UpdateLocation:input_type -> ome.location.v1beta1.UpdateLocationRequest
	2, // 2: ome.location.v1beta1.LocationService.QueryLocation:input_type -> ome.location.v1beta1.QueryLocationRequest
	1, // 3: ome.location.v1beta1.LocationService.UpdateLocation:output_type -> ome.location.v1beta1.UpdateLocationResponse
	3, // 4: ome.location.v1beta1.LocationService.QueryLocation:output_type -> ome.location.v1beta1.QueryLocationResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_idl_ome_location_v1beta1_location_proto_init() }
func file_idl_ome_location_v1beta1_location_proto_init() {
	if File_idl_ome_location_v1beta1_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_ome_location_v1beta1_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLocationRequest); i {
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
		file_idl_ome_location_v1beta1_location_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLocationResponse); i {
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
		file_idl_ome_location_v1beta1_location_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryLocationRequest); i {
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
		file_idl_ome_location_v1beta1_location_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryLocationResponse); i {
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
			RawDescriptor: file_idl_ome_location_v1beta1_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_ome_location_v1beta1_location_proto_goTypes,
		DependencyIndexes: file_idl_ome_location_v1beta1_location_proto_depIdxs,
		MessageInfos:      file_idl_ome_location_v1beta1_location_proto_msgTypes,
	}.Build()
	File_idl_ome_location_v1beta1_location_proto = out.File
	file_idl_ome_location_v1beta1_location_proto_rawDesc = nil
	file_idl_ome_location_v1beta1_location_proto_goTypes = nil
	file_idl_ome_location_v1beta1_location_proto_depIdxs = nil
}
