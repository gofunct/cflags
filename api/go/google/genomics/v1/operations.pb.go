// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/genomics/v1/operations.proto

package google_genomics_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/api"
import google_protobuf5 "github.com/golang/protobuf/ptypes/any"
import google_protobuf6 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Metadata describing an [Operation][google.longrunning.Operation].
type OperationMetadata struct {
	// The Google Cloud Project in which the job is scoped.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	// The time at which the job was submitted to the Genomics service.
	CreateTime *google_protobuf6.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	// The time at which the job began to run.
	StartTime *google_protobuf6.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// The time at which the job stopped running.
	EndTime *google_protobuf6.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// The original request that started the operation. Note that this will be in
	// current version of the API. If the operation was started with v1beta2 API
	// and a GetOperation is performed on v1 API, a v1 request will be returned.
	Request *google_protobuf5.Any `protobuf:"bytes,5,opt,name=request" json:"request,omitempty"`
	// Optional event messages that were generated during the job's execution.
	// This also contains any warnings that were generated during import
	// or export.
	Events []*OperationEvent `protobuf:"bytes,6,rep,name=events" json:"events,omitempty"`
	// This field is deprecated. Use `labels` instead. Optionally provided by the
	// caller when submitting the request that creates the operation.
	ClientId string `protobuf:"bytes,7,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Runtime metadata on this Operation.
	RuntimeMetadata *google_protobuf5.Any `protobuf:"bytes,8,opt,name=runtime_metadata,json=runtimeMetadata" json:"runtime_metadata,omitempty"`
	// Optionally provided by the caller when submitting the request that creates
	// the operation.
	Labels map[string]string `protobuf:"bytes,9,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *OperationMetadata) Reset()                    { *m = OperationMetadata{} }
func (m *OperationMetadata) String() string            { return proto.CompactTextString(m) }
func (*OperationMetadata) ProtoMessage()               {}
func (*OperationMetadata) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *OperationMetadata) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *OperationMetadata) GetCreateTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *OperationMetadata) GetStartTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *OperationMetadata) GetEndTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *OperationMetadata) GetRequest() *google_protobuf5.Any {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *OperationMetadata) GetEvents() []*OperationEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *OperationMetadata) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *OperationMetadata) GetRuntimeMetadata() *google_protobuf5.Any {
	if m != nil {
		return m.RuntimeMetadata
	}
	return nil
}

func (m *OperationMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// An event that occurred during an [Operation][google.longrunning.Operation].
type OperationEvent struct {
	// Optional time of when event started.
	StartTime *google_protobuf6.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// Optional time of when event finished. An event can have a start time and no
	// finish time. If an event has a finish time, there must be a start time.
	EndTime *google_protobuf6.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// Required description of event.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *OperationEvent) Reset()                    { *m = OperationEvent{} }
func (m *OperationEvent) String() string            { return proto.CompactTextString(m) }
func (*OperationEvent) ProtoMessage()               {}
func (*OperationEvent) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *OperationEvent) GetStartTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *OperationEvent) GetEndTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *OperationEvent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*OperationMetadata)(nil), "google.genomics.v1.OperationMetadata")
	proto.RegisterType((*OperationEvent)(nil), "google.genomics.v1.OperationEvent")
}

func init() { proto.RegisterFile("google/genomics/v1/operations.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xe5, 0x76, 0x6b, 0x9b, 0x17, 0x89, 0x0d, 0x6b, 0x42, 0xa1, 0x80, 0xa8, 0xca, 0xa5,
	0x27, 0x47, 0x1d, 0x42, 0x62, 0xe3, 0x80, 0x98, 0xb4, 0x43, 0x25, 0x10, 0x53, 0xc4, 0xbd, 0x72,
	0x93, 0x47, 0x15, 0x48, 0xec, 0x60, 0x3b, 0x91, 0xf2, 0x7d, 0xf8, 0x02, 0x7c, 0x3b, 0x8e, 0x28,
	0xb6, 0x53, 0x95, 0x0d, 0xad, 0x68, 0xb7, 0xf8, 0xbd, 0xff, 0xcf, 0xef, 0x9f, 0xf7, 0x37, 0xbc,
	0xda, 0x4a, 0xb9, 0x2d, 0x30, 0xde, 0xa2, 0x90, 0x65, 0x9e, 0xea, 0xb8, 0x59, 0xc6, 0xb2, 0x42,
	0xc5, 0x4d, 0x2e, 0x85, 0x66, 0x95, 0x92, 0x46, 0x52, 0xea, 0x44, 0xac, 0x17, 0xb1, 0x66, 0x39,
	0x7d, 0xee, 0x41, 0x5e, 0xe5, 0x31, 0x17, 0x42, 0x9a, 0x7d, 0x62, 0xfa, 0xd4, 0x77, 0xed, 0x69,
	0x53, 0x7f, 0x8d, 0xb9, 0x68, 0x7d, 0xeb, 0xe5, 0xed, 0x96, 0xc9, 0x4b, 0xd4, 0x86, 0x97, 0x95,
	0x13, 0xcc, 0x7f, 0x1d, 0xc1, 0xe3, 0xcf, 0xbd, 0x85, 0x4f, 0x68, 0x78, 0xc6, 0x0d, 0xa7, 0x2f,
	0x00, 0x2a, 0x25, 0xbf, 0x61, 0x6a, 0xd6, 0x79, 0x16, 0x91, 0x19, 0x59, 0x04, 0x49, 0xe0, 0x2b,
	0xab, 0x8c, 0xbe, 0x83, 0x30, 0x55, 0xc8, 0x0d, 0xae, 0xbb, 0xeb, 0xa2, 0xc1, 0x8c, 0x2c, 0xc2,
	0xf3, 0x29, 0xf3, 0xc6, 0xfb, 0x59, 0xec, 0x4b, 0x3f, 0x2b, 0x01, 0x27, 0xef, 0x0a, 0xf4, 0x02,
	0x40, 0x1b, 0xae, 0x8c, 0x63, 0x87, 0x07, 0xd9, 0xc0, 0xaa, 0x2d, 0xfa, 0x06, 0x26, 0x28, 0x32,
	0x07, 0x1e, 0x1d, 0x04, 0xc7, 0x28, 0x32, 0x8b, 0x31, 0x18, 0x2b, 0xfc, 0x51, 0xa3, 0x36, 0xd1,
	0xb1, 0xa5, 0xce, 0xee, 0x50, 0x1f, 0x44, 0x9b, 0xf4, 0x22, 0x7a, 0x09, 0x23, 0x6c, 0x50, 0x18,
	0x1d, 0x8d, 0x66, 0xc3, 0x45, 0x78, 0x3e, 0x67, 0x77, 0x23, 0x61, 0xbb, 0xa5, 0x5d, 0x77, 0xd2,
	0xc4, 0x13, 0xf4, 0x19, 0x04, 0x69, 0x91, 0xa3, 0xb0, 0x8b, 0x1b, 0xdb, 0xc5, 0x4d, 0x5c, 0x61,
	0x95, 0xd1, 0xf7, 0x70, 0xaa, 0x6a, 0xd1, 0xd9, 0x5f, 0x97, 0x7e, 0xd5, 0xd1, 0xe4, 0x1e, 0x47,
	0x27, 0x5e, 0xbd, 0xcb, 0x65, 0x05, 0xa3, 0x82, 0x6f, 0xb0, 0xd0, 0x51, 0x60, 0x9d, 0x2d, 0xef,
	0x75, 0xd6, 0x63, 0xec, 0xa3, 0x65, 0xae, 0x85, 0x51, 0x6d, 0xe2, 0x2f, 0x98, 0x5e, 0x40, 0xb8,
	0x57, 0xa6, 0xa7, 0x30, 0xfc, 0x8e, 0xad, 0x8f, 0xba, 0xfb, 0xa4, 0x67, 0x70, 0xdc, 0xf0, 0xa2,
	0x76, 0xf1, 0x06, 0x89, 0x3b, 0x5c, 0x0e, 0xde, 0x92, 0xf9, 0x4f, 0x02, 0x8f, 0xfe, 0xfe, 0xfd,
	0x5b, 0xa1, 0x92, 0x87, 0x86, 0x3a, 0xf8, 0xff, 0x50, 0x67, 0x10, 0x66, 0xa8, 0x53, 0x95, 0x57,
	0x9d, 0x0b, 0xfb, 0x8e, 0x82, 0x64, 0xbf, 0x74, 0xc5, 0xe0, 0x49, 0x2a, 0xcb, 0x7f, 0x6c, 0xe8,
	0xea, 0x64, 0xe7, 0x5e, 0xdf, 0x74, 0x23, 0x6e, 0xc8, 0x6f, 0x42, 0x36, 0x23, 0x3b, 0xee, 0xf5,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xd7, 0xc3, 0x75, 0xa6, 0x03, 0x00, 0x00,
}
