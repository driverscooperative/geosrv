// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: omeapi/location/v1beta1/location.proto

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

type UpdateLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value        *LocationUpdate `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	ValidateOnly bool            `protobuf:"varint,2,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
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

func (x *UpdateLocationRequest) GetValue() *LocationUpdate {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *UpdateLocationRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

type UpdateLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId         string                    `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	TollgateCrossing *v1beta1.TollgateCrossing `protobuf:"bytes,2,opt,name=tollgate_crossing,json=tollgateCrossing,proto3" json:"tollgate_crossing,omitempty"`
	UpdateTime       *timestamppb.Timestamp    `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *UpdateLocationResponse) Reset() {
	*x = UpdateLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_location_v1beta1_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLocationResponse) ProtoMessage() {}

func (x *UpdateLocationResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpdateLocationResponse.ProtoReflect.Descriptor instead.
func (*UpdateLocationResponse) Descriptor() ([]byte, []int) {
	return file_omeapi_location_v1beta1_location_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateLocationResponse) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *UpdateLocationResponse) GetTollgateCrossing() *v1beta1.TollgateCrossing {
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

type GetLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId string `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
}

func (x *GetLocationRequest) Reset() {
	*x = GetLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_location_v1beta1_location_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationRequest) ProtoMessage() {}

func (x *GetLocationRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetLocationRequest.ProtoReflect.Descriptor instead.
func (*GetLocationRequest) Descriptor() ([]byte, []int) {
	return file_omeapi_location_v1beta1_location_proto_rawDescGZIP(), []int{2}
}

func (x *GetLocationRequest) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

type GetLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId     string                 `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	Location     *v1beta1.Location      `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	LastSeenTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_seen_time,json=lastSeenTime,proto3" json:"last_seen_time,omitempty"`
}

func (x *GetLocationResponse) Reset() {
	*x = GetLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omeapi_location_v1beta1_location_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationResponse) ProtoMessage() {}

func (x *GetLocationResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetLocationResponse.ProtoReflect.Descriptor instead.
func (*GetLocationResponse) Descriptor() ([]byte, []int) {
	return file_omeapi_location_v1beta1_location_proto_rawDescGZIP(), []int{3}
}

func (x *GetLocationResponse) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *GetLocationResponse) GetLocation() *v1beta1.Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *GetLocationResponse) GetLastSeenTime() *timestamppb.Timestamp {
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
	0x74, 0x6f, 0x1a, 0x22, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x6f, 0x6c, 0x6c,
	0x67, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6f, 0x6d, 0x65,
	0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22,
	0xc6, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x52, 0x0a, 0x11, 0x74, 0x6f, 0x6c, 0x6c, 0x67,
	0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x6f, 0x6c, 0x6c, 0x67, 0x61, 0x74,
	0x65, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x74, 0x6f, 0x6c, 0x6c, 0x67,
	0x61, 0x74, 0x65, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x22, 0xaf, 0x01, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x39, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0e, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xf2, 0x01,
	0x0a, 0x0f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x73, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x59, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6f, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_omeapi_location_v1beta1_location_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_omeapi_location_v1beta1_location_proto_goTypes = []interface{}{
	(*UpdateLocationRequest)(nil),    // 0: omeapi.location.v1beta1.UpdateLocationRequest
	(*UpdateLocationResponse)(nil),   // 1: omeapi.location.v1beta1.UpdateLocationResponse
	(*GetLocationRequest)(nil),       // 2: omeapi.location.v1beta1.GetLocationRequest
	(*GetLocationResponse)(nil),      // 3: omeapi.location.v1beta1.GetLocationResponse
	(*LocationUpdate)(nil),           // 4: omeapi.location.v1beta1.LocationUpdate
	(*v1beta1.TollgateCrossing)(nil), // 5: omeapi.type.v1beta1.TollgateCrossing
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
	(*v1beta1.Location)(nil),         // 7: omeapi.type.v1beta1.Location
}
var file_omeapi_location_v1beta1_location_proto_depIdxs = []int32{
	4, // 0: omeapi.location.v1beta1.UpdateLocationRequest.value:type_name -> omeapi.location.v1beta1.LocationUpdate
	5, // 1: omeapi.location.v1beta1.UpdateLocationResponse.tollgate_crossing:type_name -> omeapi.type.v1beta1.TollgateCrossing
	6, // 2: omeapi.location.v1beta1.UpdateLocationResponse.update_time:type_name -> google.protobuf.Timestamp
	7, // 3: omeapi.location.v1beta1.GetLocationResponse.location:type_name -> omeapi.type.v1beta1.Location
	6, // 4: omeapi.location.v1beta1.GetLocationResponse.last_seen_time:type_name -> google.protobuf.Timestamp
	0, // 5: omeapi.location.v1beta1.LocationService.UpdateLocation:input_type -> omeapi.location.v1beta1.UpdateLocationRequest
	2, // 6: omeapi.location.v1beta1.LocationService.GetLocation:input_type -> omeapi.location.v1beta1.GetLocationRequest
	1, // 7: omeapi.location.v1beta1.LocationService.UpdateLocation:output_type -> omeapi.location.v1beta1.UpdateLocationResponse
	3, // 8: omeapi.location.v1beta1.LocationService.GetLocation:output_type -> omeapi.location.v1beta1.GetLocationResponse
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
	file_omeapi_location_v1beta1_location_update_proto_init()
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
		file_omeapi_location_v1beta1_location_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocationRequest); i {
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
			switch v := v.(*GetLocationResponse); i {
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
			NumMessages:   4,
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
