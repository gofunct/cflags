// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/common/common.proto

package pbcommon

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type DoubleValues struct {
	Value []float64 `protobuf:"fixed64,1,rep,packed,name=value" json:"value,omitempty"`
}

func (m *DoubleValues) Reset()                    { *m = DoubleValues{} }
func (m *DoubleValues) String() string            { return proto.CompactTextString(m) }
func (*DoubleValues) ProtoMessage()               {}
func (*DoubleValues) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{0} }

func (m *DoubleValues) GetValue() []float64 {
	if m != nil {
		return m.Value
	}
	return nil
}

type FloatValues struct {
	Value []float32 `protobuf:"fixed32,1,rep,packed,name=value" json:"value,omitempty"`
}

func (m *FloatValues) Reset()                    { *m = FloatValues{} }
func (m *FloatValues) String() string            { return proto.CompactTextString(m) }
func (*FloatValues) ProtoMessage()               {}
func (*FloatValues) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{1} }

func (m *FloatValues) GetValue() []float32 {
	if m != nil {
		return m.Value
	}
	return nil
}

type Int64Values struct {
	Value []int64 `protobuf:"varint,1,rep,packed,name=value" json:"value,omitempty"`
}

func (m *Int64Values) Reset()                    { *m = Int64Values{} }
func (m *Int64Values) String() string            { return proto.CompactTextString(m) }
func (*Int64Values) ProtoMessage()               {}
func (*Int64Values) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{2} }

func (m *Int64Values) GetValue() []int64 {
	if m != nil {
		return m.Value
	}
	return nil
}

type UInt64Values struct {
	Value []uint64 `protobuf:"varint,1,rep,packed,name=value" json:"value,omitempty"`
}

func (m *UInt64Values) Reset()                    { *m = UInt64Values{} }
func (m *UInt64Values) String() string            { return proto.CompactTextString(m) }
func (*UInt64Values) ProtoMessage()               {}
func (*UInt64Values) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{3} }

func (m *UInt64Values) GetValue() []uint64 {
	if m != nil {
		return m.Value
	}
	return nil
}

type Int32Values struct {
	Value []int32 `protobuf:"varint,1,rep,packed,name=value" json:"value,omitempty"`
}

func (m *Int32Values) Reset()                    { *m = Int32Values{} }
func (m *Int32Values) String() string            { return proto.CompactTextString(m) }
func (*Int32Values) ProtoMessage()               {}
func (*Int32Values) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{4} }

func (m *Int32Values) GetValue() []int32 {
	if m != nil {
		return m.Value
	}
	return nil
}

type UInt32Values struct {
	Value []uint32 `protobuf:"varint,1,rep,packed,name=value" json:"value,omitempty"`
}

func (m *UInt32Values) Reset()                    { *m = UInt32Values{} }
func (m *UInt32Values) String() string            { return proto.CompactTextString(m) }
func (*UInt32Values) ProtoMessage()               {}
func (*UInt32Values) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{5} }

func (m *UInt32Values) GetValue() []uint32 {
	if m != nil {
		return m.Value
	}
	return nil
}

type BoolValues struct {
	Value []bool `protobuf:"varint,1,rep,packed,name=value" json:"value,omitempty"`
}

func (m *BoolValues) Reset()                    { *m = BoolValues{} }
func (m *BoolValues) String() string            { return proto.CompactTextString(m) }
func (*BoolValues) ProtoMessage()               {}
func (*BoolValues) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{6} }

func (m *BoolValues) GetValue() []bool {
	if m != nil {
		return m.Value
	}
	return nil
}

type StringValues struct {
	Value []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
}

func (m *StringValues) Reset()                    { *m = StringValues{} }
func (m *StringValues) String() string            { return proto.CompactTextString(m) }
func (*StringValues) ProtoMessage()               {}
func (*StringValues) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{7} }

func (m *StringValues) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type BytesValues struct {
	Value [][]byte `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
}

func (m *BytesValues) Reset()                    { *m = BytesValues{} }
func (m *BytesValues) String() string            { return proto.CompactTextString(m) }
func (*BytesValues) ProtoMessage()               {}
func (*BytesValues) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{8} }

func (m *BytesValues) GetValue() [][]byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*DoubleValues)(nil), "pb.common.DoubleValues")
	proto.RegisterType((*FloatValues)(nil), "pb.common.FloatValues")
	proto.RegisterType((*Int64Values)(nil), "pb.common.Int64Values")
	proto.RegisterType((*UInt64Values)(nil), "pb.common.UInt64Values")
	proto.RegisterType((*Int32Values)(nil), "pb.common.Int32Values")
	proto.RegisterType((*UInt32Values)(nil), "pb.common.UInt32Values")
	proto.RegisterType((*BoolValues)(nil), "pb.common.BoolValues")
	proto.RegisterType((*StringValues)(nil), "pb.common.StringValues")
	proto.RegisterType((*BytesValues)(nil), "pb.common.BytesValues")
}

func init() { proto.RegisterFile("pb/common/common.proto", fileDescriptorCommon) }

var fileDescriptorCommon = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x48, 0xd2, 0x4f,
	0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x52, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x9c, 0x05,
	0x49, 0x7a, 0x10, 0x01, 0x25, 0x15, 0x2e, 0x1e, 0x97, 0xfc, 0xd2, 0xa4, 0x9c, 0xd4, 0xb0, 0xc4,
	0x9c, 0xd2, 0xd4, 0x62, 0x21, 0x11, 0x2e, 0xd6, 0x32, 0x10, 0x4b, 0x82, 0x51, 0x81, 0x59, 0x83,
	0x31, 0x08, 0xc2, 0x51, 0x52, 0xe6, 0xe2, 0x76, 0xcb, 0xc9, 0x4f, 0x2c, 0xc1, 0xa6, 0x88, 0x09,
	0x49, 0x91, 0x67, 0x5e, 0x89, 0x99, 0x09, 0x36, 0x45, 0xcc, 0x30, 0x45, 0x2a, 0x5c, 0x3c, 0xa1,
	0x38, 0x55, 0xb1, 0xa0, 0x1a, 0x65, 0x6c, 0x84, 0x4d, 0x11, 0x2b, 0x9a, 0x51, 0xd8, 0x55, 0xf1,
	0xc2, 0x54, 0x29, 0x71, 0x71, 0x39, 0xe5, 0xe7, 0xe7, 0x60, 0x53, 0xc3, 0x81, 0x64, 0x52, 0x70,
	0x49, 0x51, 0x66, 0x5e, 0x3a, 0x36, 0x55, 0x9c, 0x48, 0x8e, 0x72, 0xaa, 0x2c, 0x49, 0x2d, 0xc6,
	0xa6, 0x88, 0x07, 0xaa, 0xc8, 0x89, 0x2b, 0x8a, 0xa3, 0x20, 0x09, 0x12, 0xb6, 0x49, 0x6c, 0xe0,
	0xd0, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3d, 0xad, 0x3f, 0x4b, 0x87, 0x01, 0x00, 0x00,
}
