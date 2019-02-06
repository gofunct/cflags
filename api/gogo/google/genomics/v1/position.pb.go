// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/genomics/v1/position.proto

package google_genomics_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An abstraction for referring to a genomic position, in relation to some
// already known reference. For now, represents a genomic position as a
// reference name, a base number on that reference (0-based), and a
// determination of forward or reverse strand.
type Position struct {
	// The name of the reference in whatever reference set is being used.
	ReferenceName string `protobuf:"bytes,1,opt,name=reference_name,json=referenceName,proto3" json:"reference_name,omitempty"`
	// The 0-based offset from the start of the forward strand for that reference.
	Position int64 `protobuf:"varint,2,opt,name=position,proto3" json:"position,omitempty"`
	// Whether this position is on the reverse strand, as opposed to the forward
	// strand.
	ReverseStrand bool `protobuf:"varint,3,opt,name=reverse_strand,json=reverseStrand,proto3" json:"reverse_strand,omitempty"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptorPosition, []int{0} }

func (m *Position) GetReferenceName() string {
	if m != nil {
		return m.ReferenceName
	}
	return ""
}

func (m *Position) GetPosition() int64 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *Position) GetReverseStrand() bool {
	if m != nil {
		return m.ReverseStrand
	}
	return false
}

func init() {
	proto.RegisterType((*Position)(nil), "google.genomics.v1.Position")
}

func init() { proto.RegisterFile("google/genomics/v1/position.proto", fileDescriptorPosition) }

var fileDescriptorPosition = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0x4d, 0x0a, 0xc2, 0x30,
	0x10, 0x85, 0x89, 0x05, 0xa9, 0x81, 0xba, 0xc8, 0x42, 0x4a, 0x71, 0x51, 0x05, 0xa1, 0x0b, 0x49,
	0x28, 0xde, 0xa0, 0x07, 0x90, 0x52, 0x0f, 0x50, 0x62, 0x1d, 0x4b, 0xc0, 0x64, 0x4a, 0x12, 0x7a,
	0x76, 0x97, 0xd2, 0xdf, 0x8d, 0xcb, 0x79, 0xbc, 0xf7, 0xcd, 0x47, 0x4f, 0x2d, 0x62, 0xfb, 0x01,
	0xd1, 0x82, 0x41, 0xad, 0x1a, 0x27, 0xfa, 0x5c, 0x74, 0xe8, 0x94, 0x57, 0x68, 0x78, 0x67, 0xd1,
	0x23, 0x63, 0x53, 0x85, 0x2f, 0x15, 0xde, 0xe7, 0xc9, 0x71, 0x9e, 0xc9, 0x4e, 0x09, 0x69, 0x0c,
	0x7a, 0x39, 0x0c, 0xdc, 0xb4, 0x38, 0x7b, 0x1a, 0x96, 0x33, 0x83, 0x5d, 0xe8, 0xde, 0xc2, 0x1b,
	0x2c, 0x98, 0x06, 0x6a, 0x23, 0x35, 0xc4, 0x24, 0x25, 0xd9, 0xae, 0x8a, 0xd6, 0xf4, 0x2e, 0x35,
	0xb0, 0x84, 0x86, 0xcb, 0xdb, 0x78, 0x93, 0x92, 0x2c, 0xa8, 0xd6, 0x7b, 0x42, 0xf4, 0x60, 0x1d,
	0xd4, 0xce, 0x5b, 0x69, 0x5e, 0x71, 0x90, 0x92, 0x2c, 0x1c, 0x10, 0x63, 0xfa, 0x18, 0xc3, 0xe2,
	0x4a, 0x0f, 0x0d, 0x6a, 0xfe, 0x6f, 0x5b, 0x44, 0x8b, 0x4d, 0x39, 0xe8, 0x95, 0xe4, 0x4b, 0xc8,
	0x73, 0x3b, 0xaa, 0xde, 0x7e, 0x01, 0x00, 0x00, 0xff, 0xff, 0x14, 0x56, 0x88, 0xfa, 0x01, 0x01,
	0x00, 0x00,
}
