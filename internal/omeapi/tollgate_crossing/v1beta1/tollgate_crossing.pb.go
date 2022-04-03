// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: omeapi/tollgate_crossing/v1beta1/tollgate_crossing.proto

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

type QueryTollgateCrossingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TollgateId string `protobuf:"bytes,1,opt,name=tollgate_id,json=tollgateId,proto3" json:"tollgate_id,omitempty"`
	DriverId   string `protobuf:"bytes,2,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
}

func (x *QueryTollgateCrossingsRequest) Reset() {
	*x = QueryTollgateCrossingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTollgateCrossingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTollgateCrossingsRequest) ProtoMessage() {}

func (x *QueryTollgateCrossingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTollgateCrossingsRequest.ProtoReflect.Descriptor instead.
func (*QueryTollgateCrossingsRequest) Descriptor() ([]byte, []int) {
	return file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_rawDescGZIP(), []int{0}
}

func (x *QueryTollgateCrossingsRequest) GetTollgateId() string {
	if x != nil {
		return x.TollgateId
	}
	return ""
}

func (x *QueryTollgateCrossingsRequest) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

type QueryTollgateCrossingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tollgate []*TollgateCrossing `protobuf:"bytes,1,rep,name=tollgate,proto3" json:"tollgate,omitempty"`
}

func (x *QueryTollgateCrossingsResponse) Reset() {
	*x = QueryTollgateCrossingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTollgateCrossingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTollgateCrossingsResponse) ProtoMessage() {}

func (x *QueryTollgateCrossingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTollgateCrossingsResponse.ProtoReflect.Descriptor instead.
func (*QueryTollgateCrossingsResponse) Descriptor() ([]byte, []int) {
	return file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_rawDescGZIP(), []int{1}
}

func (x *QueryTollgateCrossingsResponse) GetTollgate() []*TollgateCrossing {
	if x != nil {
		return x.Tollgate
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
		mi := &file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movement) ProtoMessage() {}

func (x *Movement) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_msgTypes[2]
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
	return file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_rawDescGZIP(), []int{2}
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

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TollgateId  string                 `protobuf:"bytes,2,opt,name=tollgate_id,json=tollgateId,proto3" json:"tollgate_id,omitempty"`
	DriverId    string                 `protobuf:"bytes,3,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	Direction   string                 `protobuf:"bytes,4,opt,name=direction,proto3" json:"direction,omitempty"`
	Alg         string                 `protobuf:"bytes,5,opt,name=alg,proto3" json:"alg,omitempty"`
	Movement    *Movement              `protobuf:"bytes,6,opt,name=movement,proto3" json:"movement,omitempty"`
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
}

func (x *TollgateCrossing) Reset() {
	*x = TollgateCrossing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TollgateCrossing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TollgateCrossing) ProtoMessage() {}

func (x *TollgateCrossing) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_msgTypes[3]
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
	return file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_rawDescGZIP(), []int{3}
}

func (x *TollgateCrossing) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TollgateCrossing) GetTollgateId() string {
	if x != nil {
		return x.TollgateId
	}
	return ""
}

func (x *TollgateCrossing) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *TollgateCrossing) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *TollgateCrossing) GetAlg() string {
	if x != nil {
		return x.Alg
	}
	return ""
}

func (x *TollgateCrossing) GetMovement() *Movement {
	if x != nil {
		return x.Movement
	}
	return nil
}

func (x *TollgateCrossing) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

var File_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto protoreflect.FileDescriptor

var file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_rawDesc = []byte{
	0x0a, 0x38, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74,
	0x65, 0x5f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x6f, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x6f, 0x6d, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x6f, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a,
	0x1d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x43, 0x72,
	0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x70, 0x0a, 0x1e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e,
	0x0a, 0x08, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61,
	0x74, 0x65, 0x5f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x22, 0x6e,
	0x0a, 0x08, 0x4d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x6c, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x66, 0x72,
	0x6f, 0x6d, 0x4c, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6c, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x4c, 0x61, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x74, 0x6f, 0x4c, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x6f, 0x5f, 0x6c, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x4c, 0x61, 0x74, 0x22, 0x97,
	0x02, 0x0a, 0x10, 0x54, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x6c, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6c,
	0x67, 0x12, 0x46, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x6c,
	0x6c, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x08, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xb9, 0x01, 0x0a, 0x17, 0x54, 0x6f, 0x6c,
	0x6c, 0x67, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f,
	0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x3f, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74,
	0x65, 0x5f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65,
	0x43, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x40, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61,
	0x74, 0x65, 0x5f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74,
	0x65, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_rawDescOnce sync.Once
	file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_rawDescData = file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_rawDesc
)

func file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_rawDescGZIP() []byte {
	file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_rawDescOnce.Do(func() {
		file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_rawDescData = protoimpl.X.CompressGZIP(file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_rawDescData)
	})
	return file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_rawDescData
}

var file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_goTypes = []interface{}{
	(*QueryTollgateCrossingsRequest)(nil),  // 0: omeapi.tollgate_crossing.v1beta1.QueryTollgateCrossingsRequest
	(*QueryTollgateCrossingsResponse)(nil), // 1: omeapi.tollgate_crossing.v1beta1.QueryTollgateCrossingsResponse
	(*Movement)(nil),                       // 2: omeapi.tollgate_crossing.v1beta1.Movement
	(*TollgateCrossing)(nil),               // 3: omeapi.tollgate_crossing.v1beta1.TollgateCrossing
	(*timestamppb.Timestamp)(nil),          // 4: google.protobuf.Timestamp
}
var file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_depIdxs = []int32{
	3, // 0: omeapi.tollgate_crossing.v1beta1.QueryTollgateCrossingsResponse.tollgate:type_name -> omeapi.tollgate_crossing.v1beta1.TollgateCrossing
	2, // 1: omeapi.tollgate_crossing.v1beta1.TollgateCrossing.movement:type_name -> omeapi.tollgate_crossing.v1beta1.Movement
	4, // 2: omeapi.tollgate_crossing.v1beta1.TollgateCrossing.created_time:type_name -> google.protobuf.Timestamp
	0, // 3: omeapi.tollgate_crossing.v1beta1.TollgateCrossingService.QueryTollgateCrossings:input_type -> omeapi.tollgate_crossing.v1beta1.QueryTollgateCrossingsRequest
	1, // 4: omeapi.tollgate_crossing.v1beta1.TollgateCrossingService.QueryTollgateCrossings:output_type -> omeapi.tollgate_crossing.v1beta1.QueryTollgateCrossingsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_init() }
func file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_init() {
	if File_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTollgateCrossingsRequest); i {
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
		file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTollgateCrossingsResponse); i {
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
		file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_goTypes,
		DependencyIndexes: file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_depIdxs,
		MessageInfos:      file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_msgTypes,
	}.Build()
	File_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto = out.File
	file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_rawDesc = nil
	file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_goTypes = nil
	file_omeapi_tollgate_crossing_v1beta1_tollgate_crossing_proto_depIdxs = nil
}
