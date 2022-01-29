// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ome/v1beta1/driver.proto

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	ome/v1beta1/driver.proto
	ome/v1beta1/latlng.proto
	ome/v1beta1/api.proto

It has these top-level messages:
	Driver
	LatLng
	UpdateDriverLocationRequest
	UpdateDriverLocationResponse
	UpdateWorkerStateRequest
	UpdateWorkerStateResponse
	GetLocationsRequest
	GetLocationsResponse
	GetWorkersRequest
	GetWorkersResponse
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkerState int32

const (
	WorkerState_WORKER_STATE_INVALID              WorkerState = 0
	WorkerState_WORKER_STATE_AWAITING_START       WorkerState = 1
	WorkerState_WORKER_STATE_GOING_TO_PICKUP      WorkerState = 2
	WorkerState_WORKER_STATE_NEAR_PICKUP          WorkerState = 3
	WorkerState_WORKER_STATE_AWAITING_PICKUP      WorkerState = 4
	WorkerState_WORKER_STATE_ON_TRIP              WorkerState = 5
	WorkerState_WORKER_STATE_NEAR_DESTINATION     WorkerState = 6
	WorkerState_WORKER_STATE_AT_DESTINATION       WorkerState = 7
	WorkerState_WORKER_STATE_COMPLETE_DROPPED_OFF WorkerState = 8
)

var WorkerState_name = map[int32]string{
	0: "WORKER_STATE_INVALID",
	1: "WORKER_STATE_AWAITING_START",
	2: "WORKER_STATE_GOING_TO_PICKUP",
	3: "WORKER_STATE_NEAR_PICKUP",
	4: "WORKER_STATE_AWAITING_PICKUP",
	5: "WORKER_STATE_ON_TRIP",
	6: "WORKER_STATE_NEAR_DESTINATION",
	7: "WORKER_STATE_AT_DESTINATION",
	8: "WORKER_STATE_COMPLETE_DROPPED_OFF",
}
var WorkerState_value = map[string]int32{
	"WORKER_STATE_INVALID":              0,
	"WORKER_STATE_AWAITING_START":       1,
	"WORKER_STATE_GOING_TO_PICKUP":      2,
	"WORKER_STATE_NEAR_PICKUP":          3,
	"WORKER_STATE_AWAITING_PICKUP":      4,
	"WORKER_STATE_ON_TRIP":              5,
	"WORKER_STATE_NEAR_DESTINATION":     6,
	"WORKER_STATE_AT_DESTINATION":       7,
	"WORKER_STATE_COMPLETE_DROPPED_OFF": 8,
}

func (x WorkerState) String() string {
	return proto.EnumName(WorkerState_name, int32(x))
}
func (WorkerState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Driver struct {
}

func (m *Driver) Reset()                    { *m = Driver{} }
func (m *Driver) String() string            { return proto.CompactTextString(m) }
func (*Driver) ProtoMessage()               {}
func (*Driver) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Driver)(nil), "ome.v1beta1.Driver")
	proto.RegisterEnum("ome.v1beta1.WorkerState", WorkerState_name, WorkerState_value)
}

func init() { proto.RegisterFile("ome/v1beta1/driver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x5d, 0xa7, 0x75, 0xa4, 0x97, 0x10, 0x3c, 0x14, 0x9c, 0xe8, 0x04, 0x2f, 0x1e, 0x3a,
	0x8a, 0x4f, 0x90, 0xae, 0xd9, 0x08, 0x5d, 0x93, 0x90, 0x7e, 0xae, 0x20, 0x42, 0xd8, 0x30, 0x07,
	0x91, 0x51, 0x29, 0x65, 0x0f, 0xe4, 0xd1, 0x8b, 0xef, 0xe1, 0x53, 0x49, 0x6b, 0x95, 0x85, 0x1e,
	0xff, 0xf9, 0xfd, 0x92, 0x7c, 0xfc, 0x3f, 0x14, 0x56, 0x7b, 0x3b, 0x3f, 0xc4, 0x3b, 0xdb, 0x6c,
	0xe3, 0xf9, 0x4b, 0xfd, 0x7a, 0xb0, 0x75, 0xf4, 0x5e, 0x57, 0x4d, 0x45, 0x82, 0x6a, 0x6f, 0xa3,
	0x9e, 0xdc, 0x4e, 0x90, 0x9f, 0x76, 0xf0, 0xfe, 0xcb, 0x43, 0x41, 0x59, 0xd5, 0x6f, 0xb6, 0x2e,
	0x9a, 0x6d, 0x63, 0x49, 0x88, 0x2e, 0x4a, 0xa9, 0x33, 0xa6, 0x4d, 0x01, 0x14, 0x98, 0xe1, 0x62,
	0x43, 0xd7, 0x3c, 0xc5, 0x27, 0xe4, 0x1a, 0x5d, 0x3a, 0x84, 0x96, 0x94, 0x03, 0x17, 0xab, 0x36,
	0x6a, 0xc0, 0x23, 0x72, 0x83, 0xa6, 0x8e, 0xb0, 0x92, 0x2d, 0x05, 0x69, 0x14, 0x5f, 0x64, 0x8f,
	0x0a, 0x7b, 0x64, 0x8a, 0x42, 0xc7, 0x10, 0x8c, 0xea, 0x3f, 0x3a, 0x1e, 0xdc, 0xff, 0xff, 0xa0,
	0x37, 0x4e, 0x07, 0xc3, 0x49, 0x61, 0x40, 0x73, 0x85, 0xcf, 0xc8, 0x0c, 0x5d, 0x0d, 0x5f, 0x4e,
	0x59, 0x01, 0x5c, 0x50, 0xe0, 0x52, 0x60, 0x7f, 0x38, 0x3f, 0x38, 0xc2, 0x39, 0xb9, 0x43, 0x33,
	0x47, 0x58, 0xc8, 0x5c, 0xad, 0x19, 0x30, 0x93, 0x6a, 0xa9, 0x14, 0x4b, 0x8d, 0x5c, 0x2e, 0xf1,
	0x24, 0xc9, 0xd0, 0x71, 0x95, 0x49, 0xf0, 0x5b, 0xa4, 0x6a, 0x4b, 0x56, 0xa3, 0xa7, 0xe0, 0x68,
	0x01, 0x1f, 0xde, 0x58, 0xe6, 0xec, 0xd3, 0x0b, 0x64, 0xce, 0xa2, 0x4d, 0x9c, 0xb4, 0x67, 0xdf,
	0x5d, 0x7a, 0xee, 0xd3, 0xce, 0xef, 0x96, 0xf3, 0xf0, 0x13, 0x00, 0x00, 0xff, 0xff, 0x20, 0x99,
	0x66, 0x4f, 0xb8, 0x01, 0x00, 0x00,
}
