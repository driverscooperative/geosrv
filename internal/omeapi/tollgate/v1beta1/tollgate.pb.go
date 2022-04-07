// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: omeapi/tollgate/v1beta1/tollgate.proto

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

type QueryOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TollgateId string `protobuf:"bytes,1,opt,name=tollgate_id,json=tollgateId,proto3" json:"tollgate_id,omitempty"`
}

func (x *QueryOneRequest) Reset() {
	*x = QueryOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOneRequest) ProtoMessage() {}

func (x *QueryOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOneRequest.ProtoReflect.Descriptor instead.
func (*QueryOneRequest) Descriptor() ([]byte, []int) {
	return file_omeapi_tollgate_v1beta1_tollgate_proto_rawDescGZIP(), []int{0}
}

func (x *QueryOneRequest) GetTollgateId() string {
	if x != nil {
		return x.TollgateId
	}
	return ""
}

type QueryOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tollgate *Tollgate `protobuf:"bytes,1,opt,name=tollgate,proto3" json:"tollgate,omitempty"`
}

func (x *QueryOneResponse) Reset() {
	*x = QueryOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOneResponse) ProtoMessage() {}

func (x *QueryOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOneResponse.ProtoReflect.Descriptor instead.
func (*QueryOneResponse) Descriptor() ([]byte, []int) {
	return file_omeapi_tollgate_v1beta1_tollgate_proto_rawDescGZIP(), []int{1}
}

func (x *QueryOneResponse) GetTollgate() *Tollgate {
	if x != nil {
		return x.Tollgate
	}
	return nil
}

type QueryAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *QueryAllRequest) Reset() {
	*x = QueryAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllRequest) ProtoMessage() {}

func (x *QueryAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAllRequest.ProtoReflect.Descriptor instead.
func (*QueryAllRequest) Descriptor() ([]byte, []int) {
	return file_omeapi_tollgate_v1beta1_tollgate_proto_rawDescGZIP(), []int{2}
}

func (x *QueryAllRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type QueryAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tollgates []*Tollgate `protobuf:"bytes,1,rep,name=tollgates,proto3" json:"tollgates,omitempty"`
}

func (x *QueryAllResponse) Reset() {
	*x = QueryAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAllResponse) ProtoMessage() {}

func (x *QueryAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAllResponse.ProtoReflect.Descriptor instead.
func (*QueryAllResponse) Descriptor() ([]byte, []int) {
	return file_omeapi_tollgate_v1beta1_tollgate_proto_rawDescGZIP(), []int{3}
}

func (x *QueryAllResponse) GetTollgates() []*Tollgate {
	if x != nil {
		return x.Tollgates
	}
	return nil
}

type Tollgate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BBoxes   *BBoxes                `protobuf:"bytes,3,opt,name=b_boxes,json=bBoxes,proto3" json:"b_boxes,omitempty"`
	GateLine *GateLine              `protobuf:"bytes,4,opt,name=gate_line,json=gateLine,proto3" json:"gate_line,omitempty"`
	Created  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Updated  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *Tollgate) Reset() {
	*x = Tollgate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tollgate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tollgate) ProtoMessage() {}

func (x *Tollgate) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tollgate.ProtoReflect.Descriptor instead.
func (*Tollgate) Descriptor() ([]byte, []int) {
	return file_omeapi_tollgate_v1beta1_tollgate_proto_rawDescGZIP(), []int{4}
}

func (x *Tollgate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tollgate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tollgate) GetBBoxes() *BBoxes {
	if x != nil {
		return x.BBoxes
	}
	return nil
}

func (x *Tollgate) GetGateLine() *GateLine {
	if x != nil {
		return x.GateLine
	}
	return nil
}

func (x *Tollgate) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Tollgate) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

type BBoxes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BBoxes   []*BBox `protobuf:"bytes,1,rep,name=b_boxes,json=bBoxes,proto3" json:"b_boxes,omitempty"`
	Required int32   `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
}

func (x *BBoxes) Reset() {
	*x = BBoxes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BBoxes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BBoxes) ProtoMessage() {}

func (x *BBoxes) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BBoxes.ProtoReflect.Descriptor instead.
func (*BBoxes) Descriptor() ([]byte, []int) {
	return file_omeapi_tollgate_v1beta1_tollgate_proto_rawDescGZIP(), []int{5}
}

func (x *BBoxes) GetBBoxes() []*BBox {
	if x != nil {
		return x.BBoxes
	}
	return nil
}

func (x *BBoxes) GetRequired() int32 {
	if x != nil {
		return x.Required
	}
	return 0
}

type BBox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LonMin float64 `protobuf:"fixed64,1,opt,name=lon_min,json=lonMin,proto3" json:"lon_min,omitempty"`
	LatMin float64 `protobuf:"fixed64,2,opt,name=lat_min,json=latMin,proto3" json:"lat_min,omitempty"`
	LonMax float64 `protobuf:"fixed64,3,opt,name=lon_max,json=lonMax,proto3" json:"lon_max,omitempty"`
	LatMax float64 `protobuf:"fixed64,4,opt,name=lat_max,json=latMax,proto3" json:"lat_max,omitempty"`
}

func (x *BBox) Reset() {
	*x = BBox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BBox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BBox) ProtoMessage() {}

func (x *BBox) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BBox.ProtoReflect.Descriptor instead.
func (*BBox) Descriptor() ([]byte, []int) {
	return file_omeapi_tollgate_v1beta1_tollgate_proto_rawDescGZIP(), []int{6}
}

func (x *BBox) GetLonMin() float64 {
	if x != nil {
		return x.LonMin
	}
	return 0
}

func (x *BBox) GetLatMin() float64 {
	if x != nil {
		return x.LatMin
	}
	return 0
}

func (x *BBox) GetLonMax() float64 {
	if x != nil {
		return x.LonMax
	}
	return 0
}

func (x *BBox) GetLatMax() float64 {
	if x != nil {
		return x.LatMax
	}
	return 0
}

type GateLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lon1 float64 `protobuf:"fixed64,1,opt,name=lon1,proto3" json:"lon1,omitempty"`
	Lat1 float64 `protobuf:"fixed64,2,opt,name=lat1,proto3" json:"lat1,omitempty"`
	Lon2 float64 `protobuf:"fixed64,3,opt,name=lon2,proto3" json:"lon2,omitempty"`
	Lat2 float64 `protobuf:"fixed64,4,opt,name=lat2,proto3" json:"lat2,omitempty"`
}

func (x *GateLine) Reset() {
	*x = GateLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateLine) ProtoMessage() {}

func (x *GateLine) ProtoReflect() protoreflect.Message {
	mi := &file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateLine.ProtoReflect.Descriptor instead.
func (*GateLine) Descriptor() ([]byte, []int) {
	return file_omeapi_tollgate_v1beta1_tollgate_proto_rawDescGZIP(), []int{7}
}

func (x *GateLine) GetLon1() float64 {
	if x != nil {
		return x.Lon1
	}
	return 0
}

func (x *GateLine) GetLat1() float64 {
	if x != nil {
		return x.Lat1
	}
	return 0
}

func (x *GateLine) GetLon2() float64 {
	if x != nil {
		return x.Lon2
	}
	return 0
}

func (x *GateLine) GetLat2() float64 {
	if x != nil {
		return x.Lat2
	}
	return 0
}

var File_omeapi_tollgate_v1beta1_tollgate_proto protoreflect.FileDescriptor

var file_omeapi_tollgate_v1beta1_tollgate_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x32, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x6c, 0x6c,
	0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f,
	0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x74, 0x6f,
	0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f,
	0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x52,
	0x08, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x53, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x6d, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x52, 0x09, 0x74, 0x6f,
	0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x73, 0x22, 0x94, 0x02, 0x0a, 0x08, 0x54, 0x6f, 0x6c, 0x6c,
	0x67, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x62, 0x5f, 0x62, 0x6f,
	0x78, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x6d, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x42, 0x42, 0x6f, 0x78, 0x65, 0x73, 0x52, 0x06, 0x62, 0x42, 0x6f, 0x78,
	0x65, 0x73, 0x12, 0x3e, 0x0a, 0x09, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x74,
	0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x47, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x67, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6e, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x5c,
	0x0a, 0x06, 0x42, 0x42, 0x6f, 0x78, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x07, 0x62, 0x5f, 0x62, 0x6f,
	0x78, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x6d, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x42, 0x42, 0x6f, 0x78, 0x52, 0x06, 0x62, 0x42, 0x6f, 0x78, 0x65, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x6a, 0x0a, 0x04,
	0x42, 0x42, 0x6f, 0x78, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6c, 0x6f, 0x6e, 0x4d, 0x69, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x6c, 0x61, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x6c, 0x61, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x6e, 0x5f, 0x6d, 0x61,
	0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6c, 0x6f, 0x6e, 0x4d, 0x61, 0x78, 0x12,
	0x17, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x6c, 0x61, 0x74, 0x4d, 0x61, 0x78, 0x22, 0x5a, 0x0a, 0x08, 0x47, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x31, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x74, 0x31,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6c, 0x61, 0x74, 0x31, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x6f, 0x6e, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x32,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x74, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x6c, 0x61, 0x74, 0x32, 0x32, 0xd7, 0x01, 0x0a, 0x0f, 0x54, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x4f, 0x6e, 0x65, 0x12, 0x28, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f,
	0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x08, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x12, 0x28, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x6c, 0x6c, 0x67,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19,
	0x5a, 0x17, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_omeapi_tollgate_v1beta1_tollgate_proto_rawDescOnce sync.Once
	file_omeapi_tollgate_v1beta1_tollgate_proto_rawDescData = file_omeapi_tollgate_v1beta1_tollgate_proto_rawDesc
)

func file_omeapi_tollgate_v1beta1_tollgate_proto_rawDescGZIP() []byte {
	file_omeapi_tollgate_v1beta1_tollgate_proto_rawDescOnce.Do(func() {
		file_omeapi_tollgate_v1beta1_tollgate_proto_rawDescData = protoimpl.X.CompressGZIP(file_omeapi_tollgate_v1beta1_tollgate_proto_rawDescData)
	})
	return file_omeapi_tollgate_v1beta1_tollgate_proto_rawDescData
}

var file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_omeapi_tollgate_v1beta1_tollgate_proto_goTypes = []interface{}{
	(*QueryOneRequest)(nil),       // 0: omeapi.tollgate.v1beta1.QueryOneRequest
	(*QueryOneResponse)(nil),      // 1: omeapi.tollgate.v1beta1.QueryOneResponse
	(*QueryAllRequest)(nil),       // 2: omeapi.tollgate.v1beta1.QueryAllRequest
	(*QueryAllResponse)(nil),      // 3: omeapi.tollgate.v1beta1.QueryAllResponse
	(*Tollgate)(nil),              // 4: omeapi.tollgate.v1beta1.Tollgate
	(*BBoxes)(nil),                // 5: omeapi.tollgate.v1beta1.BBoxes
	(*BBox)(nil),                  // 6: omeapi.tollgate.v1beta1.BBox
	(*GateLine)(nil),              // 7: omeapi.tollgate.v1beta1.GateLine
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_omeapi_tollgate_v1beta1_tollgate_proto_depIdxs = []int32{
	4, // 0: omeapi.tollgate.v1beta1.QueryOneResponse.tollgate:type_name -> omeapi.tollgate.v1beta1.Tollgate
	4, // 1: omeapi.tollgate.v1beta1.QueryAllResponse.tollgates:type_name -> omeapi.tollgate.v1beta1.Tollgate
	5, // 2: omeapi.tollgate.v1beta1.Tollgate.b_boxes:type_name -> omeapi.tollgate.v1beta1.BBoxes
	7, // 3: omeapi.tollgate.v1beta1.Tollgate.gate_line:type_name -> omeapi.tollgate.v1beta1.GateLine
	8, // 4: omeapi.tollgate.v1beta1.Tollgate.created:type_name -> google.protobuf.Timestamp
	8, // 5: omeapi.tollgate.v1beta1.Tollgate.updated:type_name -> google.protobuf.Timestamp
	6, // 6: omeapi.tollgate.v1beta1.BBoxes.b_boxes:type_name -> omeapi.tollgate.v1beta1.BBox
	0, // 7: omeapi.tollgate.v1beta1.TollgateService.QueryOne:input_type -> omeapi.tollgate.v1beta1.QueryOneRequest
	2, // 8: omeapi.tollgate.v1beta1.TollgateService.QueryAll:input_type -> omeapi.tollgate.v1beta1.QueryAllRequest
	1, // 9: omeapi.tollgate.v1beta1.TollgateService.QueryOne:output_type -> omeapi.tollgate.v1beta1.QueryOneResponse
	3, // 10: omeapi.tollgate.v1beta1.TollgateService.QueryAll:output_type -> omeapi.tollgate.v1beta1.QueryAllResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_omeapi_tollgate_v1beta1_tollgate_proto_init() }
func file_omeapi_tollgate_v1beta1_tollgate_proto_init() {
	if File_omeapi_tollgate_v1beta1_tollgate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOneRequest); i {
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
		file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOneResponse); i {
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
		file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllRequest); i {
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
		file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAllResponse); i {
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
		file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tollgate); i {
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
		file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BBoxes); i {
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
		file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BBox); i {
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
		file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateLine); i {
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
			RawDescriptor: file_omeapi_tollgate_v1beta1_tollgate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_omeapi_tollgate_v1beta1_tollgate_proto_goTypes,
		DependencyIndexes: file_omeapi_tollgate_v1beta1_tollgate_proto_depIdxs,
		MessageInfos:      file_omeapi_tollgate_v1beta1_tollgate_proto_msgTypes,
	}.Build()
	File_omeapi_tollgate_v1beta1_tollgate_proto = out.File
	file_omeapi_tollgate_v1beta1_tollgate_proto_rawDesc = nil
	file_omeapi_tollgate_v1beta1_tollgate_proto_goTypes = nil
	file_omeapi_tollgate_v1beta1_tollgate_proto_depIdxs = nil
}
