// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ome/v1beta1/latlng.proto

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An object that represents a latitude/longitude pair. This is expressed as a
// pair of doubles to represent degrees latitude and degrees longitude. Unless
// specified otherwise, this must conform to the
// <a href="http://www.unoosa.org/pdf/icg/2012/template/WGS_84.pdf">WGS84
// standard</a>. Values must be within normalized ranges.
type LatLng struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *LatLng) Reset()                    { *m = LatLng{} }
func (m *LatLng) String() string            { return proto.CompactTextString(m) }
func (*LatLng) ProtoMessage()               {}
func (*LatLng) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *LatLng) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *LatLng) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func init() {
	proto.RegisterType((*LatLng)(nil), "ome.v1beta1.LatLng")
}

func init() { proto.RegisterFile("ome/v1beta1/latlng.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xcf, 0x4d, 0xd5,
	0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0xcf, 0x49, 0x2c, 0xc9, 0xc9, 0x4b, 0xd7, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0xcf, 0x4d, 0xd5, 0x83, 0xca, 0x28, 0x39, 0x71, 0xb1,
	0xf9, 0x24, 0x96, 0xf8, 0xe4, 0xa5, 0x0b, 0x49, 0x71, 0x71, 0xe4, 0x24, 0x96, 0x64, 0x96, 0x94,
	0xa6, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x06, 0xc1, 0xf9, 0x42, 0x32, 0x5c, 0x9c, 0x39,
	0xf9, 0x79, 0xe9, 0x10, 0x49, 0x26, 0xb0, 0x24, 0x42, 0xc0, 0xc9, 0x9b, 0x0b, 0xd9, 0x48, 0x27,
	0x6e, 0x1f, 0xb0, 0x6d, 0x01, 0x20, 0xcb, 0x02, 0x18, 0xa3, 0xb8, 0x91, 0x1c, 0xb2, 0x88, 0x89,
	0xd9, 0xdf, 0xd7, 0x75, 0x15, 0x13, 0xb7, 0xbf, 0xaf, 0xab, 0x5e, 0x98, 0xa1, 0x13, 0x48, 0xec,
	0x14, 0x98, 0x17, 0x03, 0xe5, 0x25, 0xb1, 0x81, 0x1d, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x59, 0x5f, 0x23, 0xeb, 0xc0, 0x00, 0x00, 0x00,
}
