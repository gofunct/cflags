// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/cloud/ml/v1beta1/operation_metadata.proto

package google_cloud_ml_v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import google_protobuf2 "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The operation type.
type OperationMetadata_OperationType int32

const (
	// Unspecified operation type.
	OperationMetadata_OPERATION_TYPE_UNSPECIFIED OperationMetadata_OperationType = 0
	// An operation to create a new version.
	OperationMetadata_CREATE_VERSION OperationMetadata_OperationType = 1
	// An operation to delete an existing version.
	OperationMetadata_DELETE_VERSION OperationMetadata_OperationType = 2
	// An operation to delete an existing model.
	OperationMetadata_DELETE_MODEL OperationMetadata_OperationType = 3
)

var OperationMetadata_OperationType_name = map[int32]string{
	0: "OPERATION_TYPE_UNSPECIFIED",
	1: "CREATE_VERSION",
	2: "DELETE_VERSION",
	3: "DELETE_MODEL",
}
var OperationMetadata_OperationType_value = map[string]int32{
	"OPERATION_TYPE_UNSPECIFIED": 0,
	"CREATE_VERSION":             1,
	"DELETE_VERSION":             2,
	"DELETE_MODEL":               3,
}

func (x OperationMetadata_OperationType) String() string {
	return proto.EnumName(OperationMetadata_OperationType_name, int32(x))
}
func (OperationMetadata_OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorOperationMetadata, []int{0, 0}
}

// Represents the metadata of the long-running operation.
type OperationMetadata struct {
	// The time the operation was submitted.
	CreateTime *google_protobuf2.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	// The time operation processing started.
	StartTime *google_protobuf2.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// The time operation processing completed.
	EndTime *google_protobuf2.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// Indicates whether a request to cancel this operation has been made.
	IsCancellationRequested bool `protobuf:"varint,4,opt,name=is_cancellation_requested,json=isCancellationRequested,proto3" json:"is_cancellation_requested,omitempty"`
	// The operation type.
	OperationType OperationMetadata_OperationType `protobuf:"varint,5,opt,name=operation_type,json=operationType,proto3,enum=google.cloud.ml.v1beta1.OperationMetadata_OperationType" json:"operation_type,omitempty"`
	// Contains the name of the model associated with the operation.
	ModelName string `protobuf:"bytes,6,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	// Contains the version associated with the operation.
	Version *Version `protobuf:"bytes,7,opt,name=version" json:"version,omitempty"`
}

func (m *OperationMetadata) Reset()         { *m = OperationMetadata{} }
func (m *OperationMetadata) String() string { return proto.CompactTextString(m) }
func (*OperationMetadata) ProtoMessage()    {}
func (*OperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptorOperationMetadata, []int{0}
}

func (m *OperationMetadata) GetCreateTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *OperationMetadata) GetStartTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *OperationMetadata) GetEndTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *OperationMetadata) GetIsCancellationRequested() bool {
	if m != nil {
		return m.IsCancellationRequested
	}
	return false
}

func (m *OperationMetadata) GetOperationType() OperationMetadata_OperationType {
	if m != nil {
		return m.OperationType
	}
	return OperationMetadata_OPERATION_TYPE_UNSPECIFIED
}

func (m *OperationMetadata) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *OperationMetadata) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func init() {
	proto.RegisterType((*OperationMetadata)(nil), "google.cloud.ml.v1beta1.OperationMetadata")
	proto.RegisterEnum("google.cloud.ml.v1beta1.OperationMetadata_OperationType", OperationMetadata_OperationType_name, OperationMetadata_OperationType_value)
}

func init() {
	proto.RegisterFile("google/cloud/ml/v1beta1/operation_metadata.proto", fileDescriptorOperationMetadata)
}

var fileDescriptorOperationMetadata = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x51, 0x6b, 0xdb, 0x30,
	0x14, 0x85, 0xe7, 0xb6, 0x6b, 0x9a, 0xdb, 0x35, 0x64, 0x7a, 0x58, 0xbd, 0xb0, 0xad, 0xa6, 0x4f,
	0x81, 0x81, 0xbd, 0x76, 0x0c, 0xb6, 0xf5, 0xa9, 0x4d, 0x34, 0x08, 0xb4, 0xb1, 0x51, 0xbd, 0xc2,
	0x9e, 0x8c, 0x62, 0xdf, 0x15, 0x81, 0x6d, 0x79, 0x96, 0x12, 0xe8, 0x0f, 0xda, 0xff, 0x1c, 0x96,
	0x6c, 0x9a, 0x11, 0x42, 0x1e, 0x7d, 0x74, 0xbe, 0x7b, 0x8f, 0xef, 0x81, 0x4f, 0x8f, 0x52, 0x3e,
	0xe6, 0x18, 0xa4, 0xb9, 0x5c, 0x66, 0x41, 0x91, 0x07, 0xab, 0x8b, 0x05, 0x6a, 0x7e, 0x11, 0xc8,
	0x0a, 0x6b, 0xae, 0x85, 0x2c, 0x93, 0x02, 0x35, 0xcf, 0xb8, 0xe6, 0x7e, 0x55, 0x4b, 0x2d, 0xc9,
	0xa9, 0x25, 0x7c, 0x43, 0xf8, 0x45, 0xee, 0xb7, 0xc4, 0xe8, 0x5d, 0x3b, 0x8a, 0x57, 0x22, 0xe0,
	0x65, 0x29, 0xb5, 0xc1, 0x95, 0xc5, 0x46, 0x1f, 0xb7, 0x2d, 0x2a, 0x64, 0x86, 0x79, 0xa2, 0xb0,
	0x5e, 0x89, 0x14, 0x5b, 0xf3, 0x59, 0x6b, 0x36, 0x5f, 0x8b, 0xe5, 0xef, 0x40, 0x8b, 0x02, 0x95,
	0xe6, 0x45, 0x65, 0x0d, 0xe7, 0x7f, 0x0f, 0xe0, 0x75, 0xd8, 0x25, 0xbc, 0x6b, 0x03, 0x92, 0x2b,
	0x38, 0x4e, 0x6b, 0xe4, 0x1a, 0x93, 0xc6, 0xef, 0x3a, 0x9e, 0x33, 0x3e, 0xbe, 0x1c, 0xf9, 0x6d,
	0xe0, 0x6e, 0x98, 0x1f, 0x77, 0xc3, 0x18, 0x58, 0x7b, 0x23, 0x90, 0x6f, 0x00, 0x4a, 0xf3, 0x5a,
	0x5b, 0x76, 0x6f, 0x27, 0xdb, 0x37, 0x6e, 0x83, 0x7e, 0x81, 0x23, 0x2c, 0x33, 0x0b, 0xee, 0xef,
	0x04, 0x7b, 0x58, 0x66, 0x06, 0xfb, 0x0e, 0x6f, 0x85, 0x4a, 0x52, 0x5e, 0xa6, 0x98, 0xe7, 0xf6,
	0xd6, 0x35, 0xfe, 0x59, 0xa2, 0xd2, 0x98, 0xb9, 0x07, 0x9e, 0x33, 0x3e, 0x62, 0xa7, 0x42, 0x4d,
	0xd6, 0xde, 0x59, 0xf7, 0x4c, 0x12, 0x18, 0x3c, 0x37, 0xa4, 0x9f, 0x2a, 0x74, 0x5f, 0x7a, 0xce,
	0x78, 0x70, 0xf9, 0xd5, 0xdf, 0x52, 0x8f, 0xbf, 0x71, 0xae, 0x67, 0x25, 0x7e, 0xaa, 0x90, 0x9d,
	0xc8, 0xf5, 0x4f, 0xf2, 0x1e, 0xc0, 0x36, 0x53, 0xf2, 0x02, 0xdd, 0x43, 0xcf, 0x19, 0xf7, 0x59,
	0xdf, 0x28, 0x73, 0x6e, 0xb2, 0xf7, 0x56, 0x58, 0x2b, 0x21, 0x4b, 0xb7, 0x67, 0xfe, 0xd8, 0xdb,
	0xba, 0xf8, 0xc1, 0xfa, 0x58, 0x07, 0x9c, 0x0b, 0x38, 0xf9, 0x6f, 0x35, 0xf9, 0x00, 0xa3, 0x30,
	0xa2, 0xec, 0x3a, 0x9e, 0x85, 0xf3, 0x24, 0xfe, 0x15, 0xd1, 0xe4, 0xe7, 0xfc, 0x3e, 0xa2, 0x93,
	0xd9, 0x8f, 0x19, 0x9d, 0x0e, 0x5f, 0x10, 0x02, 0x83, 0x09, 0xa3, 0xd7, 0x31, 0x4d, 0x1e, 0x28,
	0xbb, 0x9f, 0x85, 0xf3, 0xa1, 0xd3, 0x68, 0x53, 0x7a, 0x4b, 0xd7, 0xb4, 0x3d, 0x32, 0x84, 0x57,
	0xad, 0x76, 0x17, 0x4e, 0xe9, 0xed, 0x70, 0xff, 0xe6, 0x0a, 0xce, 0x52, 0x59, 0x6c, 0x44, 0xe3,
	0x95, 0xe8, 0xe2, 0xdd, 0xbc, 0xd9, 0x38, 0x4c, 0xd4, 0x74, 0x16, 0x39, 0x8b, 0x43, 0x53, 0xde,
	0xe7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x4d, 0x39, 0xb3, 0x24, 0x03, 0x00, 0x00,
}
