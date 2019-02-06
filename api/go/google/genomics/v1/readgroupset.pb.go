// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/genomics/v1/readgroupset.proto

package google_genomics_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/api"
import google_protobuf3 "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A read group set is a logical collection of read groups, which are
// collections of reads produced by a sequencer. A read group set typically
// models reads corresponding to one sample, sequenced one way, and aligned one
// way.
//
// * A read group set belongs to one dataset.
// * A read group belongs to one read group set.
// * A read belongs to one read group.
//
// For more genomics resource definitions, see [Fundamentals of Google
// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
type ReadGroupSet struct {
	// The server-generated read group set ID, unique for all read group sets.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The dataset to which this read group set belongs.
	DatasetId string `protobuf:"bytes,2,opt,name=dataset_id,json=datasetId" json:"dataset_id,omitempty"`
	// The reference set to which the reads in this read group set are aligned.
	ReferenceSetId string `protobuf:"bytes,3,opt,name=reference_set_id,json=referenceSetId" json:"reference_set_id,omitempty"`
	// The read group set name. By default this will be initialized to the sample
	// name of the sequenced data contained in this set.
	Name string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	// The filename of the original source file for this read group set, if any.
	Filename string `protobuf:"bytes,5,opt,name=filename" json:"filename,omitempty"`
	// The read groups in this set. There are typically 1-10 read groups in a read
	// group set.
	ReadGroups []*ReadGroup `protobuf:"bytes,6,rep,name=read_groups,json=readGroups" json:"read_groups,omitempty"`
	// A map of additional read group set information.
	Info map[string]*google_protobuf3.ListValue `protobuf:"bytes,7,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ReadGroupSet) Reset()                    { *m = ReadGroupSet{} }
func (m *ReadGroupSet) String() string            { return proto.CompactTextString(m) }
func (*ReadGroupSet) ProtoMessage()               {}
func (*ReadGroupSet) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *ReadGroupSet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadGroupSet) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *ReadGroupSet) GetReferenceSetId() string {
	if m != nil {
		return m.ReferenceSetId
	}
	return ""
}

func (m *ReadGroupSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReadGroupSet) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *ReadGroupSet) GetReadGroups() []*ReadGroup {
	if m != nil {
		return m.ReadGroups
	}
	return nil
}

func (m *ReadGroupSet) GetInfo() map[string]*google_protobuf3.ListValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*ReadGroupSet)(nil), "google.genomics.v1.ReadGroupSet")
}

func init() { proto.RegisterFile("google/genomics/v1/readgroupset.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x6b, 0xc2, 0x40,
	0x10, 0xc5, 0x49, 0xfc, 0xd3, 0x3a, 0x16, 0xb1, 0x7b, 0x28, 0x21, 0x54, 0x10, 0xa1, 0x20, 0x3d,
	0x24, 0x6a, 0x2f, 0xa5, 0x07, 0x0f, 0x42, 0x29, 0x42, 0x0f, 0x92, 0x40, 0xaf, 0xb2, 0x66, 0x27,
	0x61, 0xa9, 0xee, 0xca, 0xee, 0x46, 0xf0, 0x3b, 0xf7, 0x03, 0xf4, 0x58, 0xb2, 0xf9, 0x43, 0xa1,
	0xc5, 0xdb, 0xcc, 0xcb, 0xef, 0xbd, 0xcc, 0x3e, 0x78, 0xc8, 0xa4, 0xcc, 0xf6, 0x18, 0x66, 0x28,
	0xe4, 0x81, 0x27, 0x3a, 0x3c, 0xcd, 0x43, 0x85, 0x94, 0x65, 0x4a, 0xe6, 0x47, 0x8d, 0x26, 0x38,
	0x2a, 0x69, 0x24, 0x21, 0x25, 0x16, 0xd4, 0x58, 0x70, 0x9a, 0xfb, 0xf7, 0x95, 0x95, 0x1e, 0x79,
	0x48, 0x85, 0x90, 0x86, 0x1a, 0x2e, 0x85, 0x2e, 0x1d, 0xfe, 0xe4, 0x52, 0x70, 0xc5, 0xd4, 0x09,
	0x76, 0xdb, 0xe5, 0x69, 0xa8, 0x8d, 0xca, 0x93, 0xea, 0x9f, 0x93, 0x2f, 0x17, 0x6e, 0x22, 0xa4,
	0xec, 0xad, 0x70, 0xc4, 0x68, 0xc8, 0x00, 0x5c, 0xce, 0x3c, 0x67, 0xec, 0x4c, 0x7b, 0x91, 0xcb,
	0x19, 0x19, 0x01, 0x30, 0x6a, 0xa8, 0x46, 0xb3, 0xe5, 0xcc, 0x73, 0xad, 0xde, 0xab, 0x94, 0x35,
	0x23, 0x53, 0x18, 0x2a, 0x4c, 0x51, 0xa1, 0x48, 0x70, 0x5b, 0x41, 0x2d, 0x0b, 0x0d, 0x1a, 0x3d,
	0xb6, 0x24, 0x81, 0xb6, 0xa0, 0x07, 0xf4, 0xda, 0xf6, 0xab, 0x9d, 0x89, 0x0f, 0xd7, 0x29, 0xdf,
	0xa3, 0xd5, 0x3b, 0x56, 0x6f, 0x76, 0xb2, 0x84, 0x7e, 0xf1, 0x94, 0x6d, 0x59, 0x92, 0xd7, 0x1d,
	0xb7, 0xa6, 0xfd, 0xc5, 0x28, 0xf8, 0xdb, 0x51, 0xd0, 0xdc, 0x1f, 0x81, 0xaa, 0x47, 0x4d, 0x96,
	0xd0, 0xe6, 0x22, 0x95, 0xde, 0x95, 0x35, 0x3e, 0x5e, 0x34, 0xc6, 0x68, 0x82, 0xb5, 0x48, 0xe5,
	0xab, 0x30, 0xea, 0x1c, 0x59, 0x9f, 0x1f, 0x43, 0xaf, 0x91, 0xc8, 0x10, 0x5a, 0x9f, 0x78, 0xae,
	0x6a, 0x29, 0x46, 0x32, 0x83, 0xce, 0x89, 0xee, 0x73, 0xb4, 0x95, 0xf4, 0x17, 0x7e, 0x9d, 0x5f,
	0xd7, 0x1c, 0xbc, 0x73, 0x6d, 0x3e, 0x0a, 0x22, 0x2a, 0xc1, 0x17, 0xf7, 0xd9, 0x59, 0xcd, 0xe0,
	0x2e, 0x91, 0x87, 0x7f, 0x6e, 0x59, 0xdd, 0xfe, 0x3e, 0x66, 0x53, 0x84, 0x6c, 0x9c, 0x6f, 0xc7,
	0xd9, 0x75, 0x6d, 0xe0, 0xd3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x9d, 0xbe, 0x9c, 0x44,
	0x02, 0x00, 0x00,
}
