// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api/worker/v1beta1/worker.proto

package v1beta1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type WorkerStatus int32

const (
	WorkerStatus_WORKER_STATUS_UNSPECIFIED WorkerStatus = 0
	WorkerStatus_WORKER_STATUS_OFFLINE     WorkerStatus = 1
	WorkerStatus_WORKER_STATUS_READY       WorkerStatus = 2
	WorkerStatus_WORKER_STATUS_ON_JOB      WorkerStatus = 3
	WorkerStatus_WORKER_STATUS_PAUSED      WorkerStatus = 4
	WorkerStatus_WORKER_STATUS_DISABLED    WorkerStatus = 5
)

// Enum value maps for WorkerStatus.
var (
	WorkerStatus_name = map[int32]string{
		0: "WORKER_STATUS_UNSPECIFIED",
		1: "WORKER_STATUS_OFFLINE",
		2: "WORKER_STATUS_READY",
		3: "WORKER_STATUS_ON_JOB",
		4: "WORKER_STATUS_PAUSED",
		5: "WORKER_STATUS_DISABLED",
	}
	WorkerStatus_value = map[string]int32{
		"WORKER_STATUS_UNSPECIFIED": 0,
		"WORKER_STATUS_OFFLINE":     1,
		"WORKER_STATUS_READY":       2,
		"WORKER_STATUS_ON_JOB":      3,
		"WORKER_STATUS_PAUSED":      4,
		"WORKER_STATUS_DISABLED":    5,
	}
)

func (x WorkerStatus) Enum() *WorkerStatus {
	p := new(WorkerStatus)
	*p = x
	return p
}

func (x WorkerStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkerStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_worker_v1beta1_worker_proto_enumTypes[0].Descriptor()
}

func (WorkerStatus) Type() protoreflect.EnumType {
	return &file_api_worker_v1beta1_worker_proto_enumTypes[0]
}

func (x WorkerStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkerStatus.Descriptor instead.
func (WorkerStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_worker_v1beta1_worker_proto_rawDescGZIP(), []int{0}
}

type GetWorkerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId string `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
}

func (x *GetWorkerRequest) Reset() {
	*x = GetWorkerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_worker_v1beta1_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkerRequest) ProtoMessage() {}

func (x *GetWorkerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_worker_v1beta1_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkerRequest.ProtoReflect.Descriptor instead.
func (*GetWorkerRequest) Descriptor() ([]byte, []int) {
	return file_api_worker_v1beta1_worker_proto_rawDescGZIP(), []int{0}
}

func (x *GetWorkerRequest) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

type GetWorkerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Worker *Worker `protobuf:"bytes,1,opt,name=worker,proto3" json:"worker,omitempty"`
}

func (x *GetWorkerResponse) Reset() {
	*x = GetWorkerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_worker_v1beta1_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkerResponse) ProtoMessage() {}

func (x *GetWorkerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_worker_v1beta1_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkerResponse.ProtoReflect.Descriptor instead.
func (*GetWorkerResponse) Descriptor() ([]byte, []int) {
	return file_api_worker_v1beta1_worker_proto_rawDescGZIP(), []int{1}
}

func (x *GetWorkerResponse) GetWorker() *Worker {
	if x != nil {
		return x.Worker
	}
	return nil
}

type UpdateWorkerStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId string       `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	Status   WorkerStatus `protobuf:"varint,2,opt,name=status,proto3,enum=api.worker.v1beta1.WorkerStatus" json:"status,omitempty"`
}

func (x *UpdateWorkerStatusRequest) Reset() {
	*x = UpdateWorkerStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_worker_v1beta1_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWorkerStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkerStatusRequest) ProtoMessage() {}

func (x *UpdateWorkerStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_worker_v1beta1_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkerStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateWorkerStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_worker_v1beta1_worker_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateWorkerStatusRequest) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *UpdateWorkerStatusRequest) GetStatus() WorkerStatus {
	if x != nil {
		return x.Status
	}
	return WorkerStatus_WORKER_STATUS_UNSPECIFIED
}

type UpdateWorkerStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId string       `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	Status   WorkerStatus `protobuf:"varint,2,opt,name=status,proto3,enum=api.worker.v1beta1.WorkerStatus" json:"status,omitempty"`
}

func (x *UpdateWorkerStatusResponse) Reset() {
	*x = UpdateWorkerStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_worker_v1beta1_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWorkerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkerStatusResponse) ProtoMessage() {}

func (x *UpdateWorkerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_worker_v1beta1_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkerStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateWorkerStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_worker_v1beta1_worker_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateWorkerStatusResponse) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *UpdateWorkerStatusResponse) GetStatus() WorkerStatus {
	if x != nil {
		return x.Status
	}
	return WorkerStatus_WORKER_STATUS_UNSPECIFIED
}

type ListWorkersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    WorkerStatus `protobuf:"varint,1,opt,name=status,proto3,enum=api.worker.v1beta1.WorkerStatus" json:"status,omitempty"`
	PageSize  int32        `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string       `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListWorkersRequest) Reset() {
	*x = ListWorkersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_worker_v1beta1_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkersRequest) ProtoMessage() {}

func (x *ListWorkersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_worker_v1beta1_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkersRequest.ProtoReflect.Descriptor instead.
func (*ListWorkersRequest) Descriptor() ([]byte, []int) {
	return file_api_worker_v1beta1_worker_proto_rawDescGZIP(), []int{4}
}

func (x *ListWorkersRequest) GetStatus() WorkerStatus {
	if x != nil {
		return x.Status
	}
	return WorkerStatus_WORKER_STATUS_UNSPECIFIED
}

func (x *ListWorkersRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListWorkersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type Worker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId string       `protobuf:"bytes,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	Status   WorkerStatus `protobuf:"varint,2,opt,name=status,proto3,enum=api.worker.v1beta1.WorkerStatus" json:"status,omitempty"`
}

func (x *Worker) Reset() {
	*x = Worker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_worker_v1beta1_worker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Worker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Worker) ProtoMessage() {}

func (x *Worker) ProtoReflect() protoreflect.Message {
	mi := &file_api_worker_v1beta1_worker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Worker.ProtoReflect.Descriptor instead.
func (*Worker) Descriptor() ([]byte, []int) {
	return file_api_worker_v1beta1_worker_proto_rawDescGZIP(), []int{5}
}

func (x *Worker) GetWorkerId() string {
	if x != nil {
		return x.WorkerId
	}
	return ""
}

func (x *Worker) GetStatus() WorkerStatus {
	if x != nil {
		return x.Status
	}
	return WorkerStatus_WORKER_STATUS_UNSPECIFIED
}

type ListWorkersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workers       []*Worker `protobuf:"bytes,1,rep,name=workers,proto3" json:"workers,omitempty"`
	NextPageToken string    `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListWorkersResponse) Reset() {
	*x = ListWorkersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_worker_v1beta1_worker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWorkersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkersResponse) ProtoMessage() {}

func (x *ListWorkersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_worker_v1beta1_worker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkersResponse.ProtoReflect.Descriptor instead.
func (*ListWorkersResponse) Descriptor() ([]byte, []int) {
	return file_api_worker_v1beta1_worker_proto_rawDescGZIP(), []int{6}
}

func (x *ListWorkersResponse) GetWorkers() []*Worker {
	if x != nil {
		return x.Workers
	}
	return nil
}

func (x *ListWorkersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_api_worker_v1beta1_worker_proto protoreflect.FileDescriptor

var file_api_worker_v1beta1_worker_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x24, 0x52, 0x08,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x06, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x22, 0x87, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x24, 0x52, 0x08, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x82, 0x01, 0x04, 0x10,
	0x01, 0x20, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x73, 0x0a, 0x1a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x8a, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5f, 0x0a,
	0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x73,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x2a, 0xb1, 0x01, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x57, 0x4f, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x57, 0x4f, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x17,
	0x0a, 0x13, 0x57, 0x4f, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x57, 0x4f, 0x52, 0x4b, 0x45,
	0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4e, 0x5f, 0x4a, 0x4f, 0x42, 0x10,
	0x03, 0x12, 0x18, 0x0a, 0x14, 0x57, 0x4f, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x57,
	0x4f, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x49, 0x53,
	0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x32, 0xc4, 0x02, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x75, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x26, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x54,
	0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_worker_v1beta1_worker_proto_rawDescOnce sync.Once
	file_api_worker_v1beta1_worker_proto_rawDescData = file_api_worker_v1beta1_worker_proto_rawDesc
)

func file_api_worker_v1beta1_worker_proto_rawDescGZIP() []byte {
	file_api_worker_v1beta1_worker_proto_rawDescOnce.Do(func() {
		file_api_worker_v1beta1_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_worker_v1beta1_worker_proto_rawDescData)
	})
	return file_api_worker_v1beta1_worker_proto_rawDescData
}

var file_api_worker_v1beta1_worker_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_worker_v1beta1_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_worker_v1beta1_worker_proto_goTypes = []interface{}{
	(WorkerStatus)(0),                  // 0: api.worker.v1beta1.WorkerStatus
	(*GetWorkerRequest)(nil),           // 1: api.worker.v1beta1.GetWorkerRequest
	(*GetWorkerResponse)(nil),          // 2: api.worker.v1beta1.GetWorkerResponse
	(*UpdateWorkerStatusRequest)(nil),  // 3: api.worker.v1beta1.UpdateWorkerStatusRequest
	(*UpdateWorkerStatusResponse)(nil), // 4: api.worker.v1beta1.UpdateWorkerStatusResponse
	(*ListWorkersRequest)(nil),         // 5: api.worker.v1beta1.ListWorkersRequest
	(*Worker)(nil),                     // 6: api.worker.v1beta1.Worker
	(*ListWorkersResponse)(nil),        // 7: api.worker.v1beta1.ListWorkersResponse
}
var file_api_worker_v1beta1_worker_proto_depIdxs = []int32{
	6, // 0: api.worker.v1beta1.GetWorkerResponse.worker:type_name -> api.worker.v1beta1.Worker
	0, // 1: api.worker.v1beta1.UpdateWorkerStatusRequest.status:type_name -> api.worker.v1beta1.WorkerStatus
	0, // 2: api.worker.v1beta1.UpdateWorkerStatusResponse.status:type_name -> api.worker.v1beta1.WorkerStatus
	0, // 3: api.worker.v1beta1.ListWorkersRequest.status:type_name -> api.worker.v1beta1.WorkerStatus
	0, // 4: api.worker.v1beta1.Worker.status:type_name -> api.worker.v1beta1.WorkerStatus
	6, // 5: api.worker.v1beta1.ListWorkersResponse.workers:type_name -> api.worker.v1beta1.Worker
	1, // 6: api.worker.v1beta1.WorkerService.GetWorker:input_type -> api.worker.v1beta1.GetWorkerRequest
	3, // 7: api.worker.v1beta1.WorkerService.UpdateWorkerStatus:input_type -> api.worker.v1beta1.UpdateWorkerStatusRequest
	5, // 8: api.worker.v1beta1.WorkerService.ListWorkers:input_type -> api.worker.v1beta1.ListWorkersRequest
	2, // 9: api.worker.v1beta1.WorkerService.GetWorker:output_type -> api.worker.v1beta1.GetWorkerResponse
	4, // 10: api.worker.v1beta1.WorkerService.UpdateWorkerStatus:output_type -> api.worker.v1beta1.UpdateWorkerStatusResponse
	7, // 11: api.worker.v1beta1.WorkerService.ListWorkers:output_type -> api.worker.v1beta1.ListWorkersResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_worker_v1beta1_worker_proto_init() }
func file_api_worker_v1beta1_worker_proto_init() {
	if File_api_worker_v1beta1_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_worker_v1beta1_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkerRequest); i {
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
		file_api_worker_v1beta1_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkerResponse); i {
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
		file_api_worker_v1beta1_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWorkerStatusRequest); i {
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
		file_api_worker_v1beta1_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWorkerStatusResponse); i {
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
		file_api_worker_v1beta1_worker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkersRequest); i {
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
		file_api_worker_v1beta1_worker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Worker); i {
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
		file_api_worker_v1beta1_worker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWorkersResponse); i {
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
			RawDescriptor: file_api_worker_v1beta1_worker_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_worker_v1beta1_worker_proto_goTypes,
		DependencyIndexes: file_api_worker_v1beta1_worker_proto_depIdxs,
		EnumInfos:         file_api_worker_v1beta1_worker_proto_enumTypes,
		MessageInfos:      file_api_worker_v1beta1_worker_proto_msgTypes,
	}.Build()
	File_api_worker_v1beta1_worker_proto = out.File
	file_api_worker_v1beta1_worker_proto_rawDesc = nil
	file_api_worker_v1beta1_worker_proto_goTypes = nil
	file_api_worker_v1beta1_worker_proto_depIdxs = nil
}
