// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: omeapi/location/v1beta1/location.proto

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

	WorkerId   string                 `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	Lon        float64                `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
	Lat        float64                `protobuf:"fixed64,3,opt,name=lat,proto3" json:"lat,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *UpdateLocationRequest) Reset() {
	*x = UpdateLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_location_v1beta1_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLocationRequest) ProtoMessage() {}

func (x *UpdateLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_location_v1beta1_location_proto_msgTypes[0]
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
	return file_omeapi_location_v1beta1_location_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateLocationRequest) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *UpdateLocationRequest) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *UpdateLocationRequest) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *UpdateLocationRequest) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type Movement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromLon float64 `protobuf:"fixed64,1,opt,name=from_lon,json=fromLon,proto3" json:"from_lon,omitempty"`
	FromLat float64 `protobuf:"fixed64,2,opt,name=from_lat,json=fromLat,proto3" json:"from_lat,omitempty"`
	ToLon   float64 `protobuf:"fixed64,3,opt,name=to_lon,json=toLon,proto3" json:"to_lon,omitempty"`
	ToLat   float64 `protobuf:"fixed64,4,opt,name=to_lat,json=toLat,proto3" json:"to_lat,omitempty"`
}

func (x *Movement) Reset() {
	*x = Movement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_location_v1beta1_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movement) ProtoMessage() {}

func (x *Movement) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_location_v1beta1_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movement.ProtoReflect.Descriptor instead.
func (*Movement) Descriptor() ([]byte, []int) {
	return file_omeapi_location_v1beta1_location_proto_rawDescGZIP(), []int{1}
}

func (x *Movement) GetFromLon() float64 {
	if x != nil {
		return x.FromLon
	}
	return 0
}

func (x *Movement) GetFromLat() float64 {
	if x != nil {
		return x.FromLat
	}
	return 0
}

func (x *Movement) GetToLon() float64 {
	if x != nil {
		return x.ToLon
	}
	return 0
}

func (x *Movement) GetToLat() float64 {
	if x != nil {
		return x.ToLat
	}
	return 0
}

type TollgateCrossing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TollgateId string    `protobuf:"bytes,1,opt,name=tollgate_id,json=tollgateId,proto3" json:"tollgate_id,omitempty"`
	Movement   *Movement `protobuf:"bytes,2,opt,name=movement,proto3" json:"movement,omitempty"`
	Direction  string    `protobuf:"bytes,3,opt,name=direction,proto3" json:"direction,omitempty"`
}

func (x *TollgateCrossing) Reset() {
	*x = TollgateCrossing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_location_v1beta1_location_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TollgateCrossing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TollgateCrossing) ProtoMessage() {}

func (x *TollgateCrossing) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_location_v1beta1_location_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TollgateCrossing.ProtoReflect.Descriptor instead.
func (*TollgateCrossing) Descriptor() ([]byte, []int) {
	return file_omeapi_location_v1beta1_location_proto_rawDescGZIP(), []int{2}
}

func (x *TollgateCrossing) GetTollgateId() string {
	if x != nil {
		return x.TollgateId
	}
	return ""
}

func (x *TollgateCrossing) GetMovement() *Movement {
	if x != nil {
		return x.Movement
	}
	return nil
}

func (x *TollgateCrossing) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

type UpdateLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId         string                 `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	TollgateCrossing *TollgateCrossing      `protobuf:"bytes,2,opt,name=tollgate_crossing,json=tollgateCrossing,proto3" json:"tollgate_crossing,omitempty"`
	UpdateTime       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *UpdateLocationResponse) Reset() {
	*x = UpdateLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_location_v1beta1_location_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLocationResponse) ProtoMessage() {}

func (x *UpdateLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_location_v1beta1_location_proto_msgTypes[3]
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
	return file_omeapi_location_v1beta1_location_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateLocationResponse) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *UpdateLocationResponse) GetTollgateCrossing() *TollgateCrossing {
	if x != nil {
		return x.TollgateCrossing
	}
	return nil
}

func (x *UpdateLocationResponse) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
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
		mi := &file_omeapi_location_v1beta1_location_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryLocationRequest) ProtoMessage() {}

func (x *QueryLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_location_v1beta1_location_proto_msgTypes[4]
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
	return file_omeapi_location_v1beta1_location_proto_rawDescGZIP(), []int{4}
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
	Lon          float64                `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
	Lat          float64                `protobuf:"fixed64,3,opt,name=lat,proto3" json:"lat,omitempty"`
	LastSeenTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_seen_time,json=lastSeenTime,proto3" json:"last_seen_time,omitempty"`
}

func (x *QueryLocationResponse) Reset() {
	*x = QueryLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_location_v1beta1_location_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryLocationResponse) ProtoMessage() {}

func (x *QueryLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_location_v1beta1_location_proto_msgTypes[5]
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
	return file_omeapi_location_v1beta1_location_proto_rawDescGZIP(), []int{5}
}

func (x *QueryLocationResponse) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *QueryLocationResponse) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *QueryLocationResponse) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *QueryLocationResponse) GetLastSeenTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSeenTime
	}
	return nil
}

var File_omeapi_location_v1beta1_location_proto protoreflect.FileDescriptor

var file_omeapi_location_v1beta1_location_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x6e, 0x0a, 0x08, 0x4d, 0x6f,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6c,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x4c, 0x6f,
	0x6e, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x4c, 0x61, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x74, 0x6f,
	0x4c, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x6f, 0x5f, 0x6c, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x4c, 0x61, 0x74, 0x22, 0x90, 0x01, 0x0a, 0x10, 0x54,
	0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x3d, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xca, 0x01,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x56, 0x0a, 0x11, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74,
	0x65, 0x5f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x6c, 0x6c, 0x67,
	0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x74, 0x6f, 0x6c,
	0x6c, 0x67, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x14, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x9a, 0x01, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xf8, 0x01, 0x0a,
	0x0f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x73, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2e, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x6f, 0x6d, 0x65, 0x61, 0x70,
	0x69, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_omeapi_location_v1beta1_location_proto_rawDescOnce sync.Once
	file_omeapi_location_v1beta1_location_proto_rawDescData = file_omeapi_location_v1beta1_location_proto_rawDesc
)

func file_omeapi_location_v1beta1_location_proto_rawDescGZIP() []byte {
	file_omeapi_location_v1beta1_location_proto_rawDescOnce.Do(func() {
		file_omeapi_location_v1beta1_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_omeapi_location_v1beta1_location_proto_rawDescData)
	})
	return file_omeapi_location_v1beta1_location_proto_rawDescData
}

var file_omeapi_location_v1beta1_location_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_omeapi_location_v1beta1_location_proto_goTypes = []interface{}{
	(*UpdateLocationRequest)(nil),  // 0: omeapi.location.v1beta1.UpdateLocationRequest
	(*Movement)(nil),               // 1: omeapi.location.v1beta1.Movement
	(*TollgateCrossing)(nil),       // 2: omeapi.location.v1beta1.TollgateCrossing
	(*UpdateLocationResponse)(nil), // 3: omeapi.location.v1beta1.UpdateLocationResponse
	(*QueryLocationRequest)(nil),   // 4: omeapi.location.v1beta1.QueryLocationRequest
	(*QueryLocationResponse)(nil),  // 5: omeapi.location.v1beta1.QueryLocationResponse
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
}
var file_omeapi_location_v1beta1_location_proto_depIdxs = []int32{
	6, // 0: omeapi.location.v1beta1.UpdateLocationRequest.update_time:type_name -> google.protobuf.Timestamp
	1, // 1: omeapi.location.v1beta1.TollgateCrossing.movement:type_name -> omeapi.location.v1beta1.Movement
	2, // 2: omeapi.location.v1beta1.UpdateLocationResponse.tollgate_crossing:type_name -> omeapi.location.v1beta1.TollgateCrossing
	6, // 3: omeapi.location.v1beta1.UpdateLocationResponse.update_time:type_name -> google.protobuf.Timestamp
	6, // 4: omeapi.location.v1beta1.QueryLocationResponse.last_seen_time:type_name -> google.protobuf.Timestamp
	0, // 5: omeapi.location.v1beta1.LocationService.UpdateLocation:input_type -> omeapi.location.v1beta1.UpdateLocationRequest
	4, // 6: omeapi.location.v1beta1.LocationService.QueryLocation:input_type -> omeapi.location.v1beta1.QueryLocationRequest
	3, // 7: omeapi.location.v1beta1.LocationService.UpdateLocation:output_type -> omeapi.location.v1beta1.UpdateLocationResponse
	5, // 8: omeapi.location.v1beta1.LocationService.QueryLocation:output_type -> omeapi.location.v1beta1.QueryLocationResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_omeapi_location_v1beta1_location_proto_init() }
func file_omeapi_location_v1beta1_location_proto_init() {
	if File_omeapi_location_v1beta1_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_omeapi_location_v1beta1_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_omeapi_location_v1beta1_location_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movement); i {
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
		file_omeapi_location_v1beta1_location_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TollgateCrossing); i {
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
		file_omeapi_location_v1beta1_location_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_omeapi_location_v1beta1_location_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_omeapi_location_v1beta1_location_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_omeapi_location_v1beta1_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_omeapi_location_v1beta1_location_proto_goTypes,
		DependencyIndexes: file_omeapi_location_v1beta1_location_proto_depIdxs,
		MessageInfos:      file_omeapi_location_v1beta1_location_proto_msgTypes,
	}.Build()
	File_omeapi_location_v1beta1_location_proto = out.File
	file_omeapi_location_v1beta1_location_proto_rawDesc = nil
	file_omeapi_location_v1beta1_location_proto_goTypes = nil
	file_omeapi_location_v1beta1_location_proto_depIdxs = nil
}
