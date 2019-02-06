// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/genomics/v1/datasets.proto

package google_genomics_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import _ "go.pedge.io/pb/gogo/google/iam/v1"
import _ "go.pedge.io/pb/gogo/google/iam/v1"
import _ "github.com/gogo/protobuf/types"
import google_protobuf2 "go.pedge.io/pb/gogo/google/protobuf"
import google_protobuf6 "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A Dataset is a collection of genomic data.
//
// For more genomics resource definitions, see [Fundamentals of Google
// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
type Dataset struct {
	// The server-generated dataset ID, unique across all datasets.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The Google Cloud project ID that this dataset belongs to.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The dataset name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The time this dataset was created, in seconds from the epoch.
	CreateTime *google_protobuf6.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
}

func (m *Dataset) Reset()                    { *m = Dataset{} }
func (m *Dataset) String() string            { return proto.CompactTextString(m) }
func (*Dataset) ProtoMessage()               {}
func (*Dataset) Descriptor() ([]byte, []int) { return fileDescriptorDatasets, []int{0} }

func (m *Dataset) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Dataset) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Dataset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Dataset) GetCreateTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

// The dataset list request.
type ListDatasetsRequest struct {
	// Required. The Google Cloud project ID to list datasets for.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The maximum number of results to return in a single page. If unspecified,
	// defaults to 50. The maximum value is 1024.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `nextPageToken` from the previous response.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (m *ListDatasetsRequest) Reset()                    { *m = ListDatasetsRequest{} }
func (m *ListDatasetsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDatasetsRequest) ProtoMessage()               {}
func (*ListDatasetsRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatasets, []int{1} }

func (m *ListDatasetsRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *ListDatasetsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListDatasetsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The dataset list response.
type ListDatasetsResponse struct {
	// The list of matching Datasets.
	Datasets []*Dataset `protobuf:"bytes,1,rep,name=datasets" json:"datasets,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *ListDatasetsResponse) Reset()                    { *m = ListDatasetsResponse{} }
func (m *ListDatasetsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDatasetsResponse) ProtoMessage()               {}
func (*ListDatasetsResponse) Descriptor() ([]byte, []int) { return fileDescriptorDatasets, []int{2} }

func (m *ListDatasetsResponse) GetDatasets() []*Dataset {
	if m != nil {
		return m.Datasets
	}
	return nil
}

func (m *ListDatasetsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type CreateDatasetRequest struct {
	// The dataset to be created. Must contain projectId and name.
	Dataset *Dataset `protobuf:"bytes,1,opt,name=dataset" json:"dataset,omitempty"`
}

func (m *CreateDatasetRequest) Reset()                    { *m = CreateDatasetRequest{} }
func (m *CreateDatasetRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDatasetRequest) ProtoMessage()               {}
func (*CreateDatasetRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatasets, []int{3} }

func (m *CreateDatasetRequest) GetDataset() *Dataset {
	if m != nil {
		return m.Dataset
	}
	return nil
}

type UpdateDatasetRequest struct {
	// The ID of the dataset to be updated.
	DatasetId string `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// The new dataset data.
	Dataset *Dataset `protobuf:"bytes,2,opt,name=dataset" json:"dataset,omitempty"`
	// An optional mask specifying which fields to update. At this time, the only
	// mutable field is [name][google.genomics.v1.Dataset.name]. The only
	// acceptable value is "name". If unspecified, all mutable fields will be
	// updated.
	UpdateMask *google_protobuf2.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
}

func (m *UpdateDatasetRequest) Reset()                    { *m = UpdateDatasetRequest{} }
func (m *UpdateDatasetRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateDatasetRequest) ProtoMessage()               {}
func (*UpdateDatasetRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatasets, []int{4} }

func (m *UpdateDatasetRequest) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *UpdateDatasetRequest) GetDataset() *Dataset {
	if m != nil {
		return m.Dataset
	}
	return nil
}

func (m *UpdateDatasetRequest) GetUpdateMask() *google_protobuf2.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type DeleteDatasetRequest struct {
	// The ID of the dataset to be deleted.
	DatasetId string `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
}

func (m *DeleteDatasetRequest) Reset()                    { *m = DeleteDatasetRequest{} }
func (m *DeleteDatasetRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteDatasetRequest) ProtoMessage()               {}
func (*DeleteDatasetRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatasets, []int{5} }

func (m *DeleteDatasetRequest) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

type UndeleteDatasetRequest struct {
	// The ID of the dataset to be undeleted.
	DatasetId string `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
}

func (m *UndeleteDatasetRequest) Reset()                    { *m = UndeleteDatasetRequest{} }
func (m *UndeleteDatasetRequest) String() string            { return proto.CompactTextString(m) }
func (*UndeleteDatasetRequest) ProtoMessage()               {}
func (*UndeleteDatasetRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatasets, []int{6} }

func (m *UndeleteDatasetRequest) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

type GetDatasetRequest struct {
	// The ID of the dataset.
	DatasetId string `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
}

func (m *GetDatasetRequest) Reset()                    { *m = GetDatasetRequest{} }
func (m *GetDatasetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDatasetRequest) ProtoMessage()               {}
func (*GetDatasetRequest) Descriptor() ([]byte, []int) { return fileDescriptorDatasets, []int{7} }

func (m *GetDatasetRequest) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func init() {
	proto.RegisterType((*Dataset)(nil), "google.genomics.v1.Dataset")
	proto.RegisterType((*ListDatasetsRequest)(nil), "google.genomics.v1.ListDatasetsRequest")
	proto.RegisterType((*ListDatasetsResponse)(nil), "google.genomics.v1.ListDatasetsResponse")
	proto.RegisterType((*CreateDatasetRequest)(nil), "google.genomics.v1.CreateDatasetRequest")
	proto.RegisterType((*UpdateDatasetRequest)(nil), "google.genomics.v1.UpdateDatasetRequest")
	proto.RegisterType((*DeleteDatasetRequest)(nil), "google.genomics.v1.DeleteDatasetRequest")
	proto.RegisterType((*UndeleteDatasetRequest)(nil), "google.genomics.v1.UndeleteDatasetRequest")
	proto.RegisterType((*GetDatasetRequest)(nil), "google.genomics.v1.GetDatasetRequest")
}

func init() { proto.RegisterFile("google/genomics/v1/datasets.proto", fileDescriptorDatasets) }

var fileDescriptorDatasets = []byte{
	// 760 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5f, 0x4f, 0x13, 0x4b,
	0x14, 0xcf, 0x14, 0xee, 0x85, 0x1e, 0x28, 0xdc, 0x3b, 0xb7, 0x17, 0x6b, 0x2b, 0x5a, 0x26, 0x0a,
	0xb5, 0x62, 0x9b, 0xd6, 0x10, 0x92, 0x1a, 0x5f, 0x10, 0x25, 0x24, 0x92, 0x34, 0x05, 0x7c, 0x6d,
	0x86, 0xee, 0xd0, 0x8c, 0x74, 0xff, 0xd8, 0x99, 0x82, 0x82, 0xbc, 0xf0, 0xc6, 0xb3, 0x1f, 0xc0,
	0xc4, 0x37, 0x3f, 0x8f, 0x5f, 0xc1, 0x0f, 0xe1, 0xa3, 0x99, 0xd9, 0xd9, 0x76, 0xbb, 0x5d, 0x0a,
	0x18, 0xdf, 0xda, 0x73, 0x7e, 0xe7, 0xfc, 0x7e, 0xe7, 0x9c, 0xdf, 0xee, 0xc2, 0x52, 0xdb, 0x75,
	0xdb, 0x1d, 0x56, 0x6e, 0x33, 0xc7, 0xb5, 0x79, 0x4b, 0x94, 0x8f, 0x2b, 0x65, 0x8b, 0x4a, 0x2a,
	0x98, 0x14, 0x25, 0xaf, 0xeb, 0x4a, 0x17, 0x63, 0x1f, 0x52, 0x0a, 0x20, 0xa5, 0xe3, 0x4a, 0xf6,
	0x9e, 0x29, 0xa3, 0x1e, 0x2f, 0x53, 0xc7, 0x71, 0x25, 0x95, 0xdc, 0x75, 0x4c, 0x45, 0xf6, 0xbe,
	0xc9, 0x72, 0x6a, 0xab, 0x7e, 0x9c, 0xda, 0x4d, 0xcf, 0xed, 0xf0, 0xd6, 0x47, 0x93, 0xcf, 0x0e,
	0xe7, 0x87, 0x72, 0x39, 0x93, 0xd3, 0xff, 0x0e, 0x7a, 0x87, 0x65, 0x66, 0x7b, 0x32, 0x48, 0xe6,
	0xa3, 0xc9, 0x43, 0xce, 0x3a, 0x56, 0xd3, 0xa6, 0xe2, 0xc8, 0x20, 0x1e, 0x44, 0x11, 0x92, 0xdb,
	0x4c, 0x48, 0x6a, 0x7b, 0x3e, 0x80, 0x5c, 0x22, 0x98, 0xda, 0xf4, 0x07, 0xc4, 0x73, 0x90, 0xe0,
	0x56, 0x06, 0xe5, 0x51, 0x21, 0xd9, 0x48, 0x70, 0x0b, 0x2f, 0x02, 0x78, 0x5d, 0xf7, 0x1d, 0x6b,
	0xc9, 0x26, 0xb7, 0x32, 0x09, 0x1d, 0x4f, 0x9a, 0xc8, 0xb6, 0x85, 0x31, 0x4c, 0x3a, 0xd4, 0x66,
	0x99, 0x09, 0x9d, 0xd0, 0xbf, 0xf1, 0x73, 0x98, 0x69, 0x75, 0x19, 0x95, 0xac, 0xa9, 0x88, 0x32,
	0x93, 0x79, 0x54, 0x98, 0xa9, 0x66, 0x4b, 0x66, 0x65, 0x81, 0x8a, 0xd2, 0x5e, 0xa0, 0xa2, 0x01,
	0x3e, 0x5c, 0x05, 0x88, 0x07, 0xff, 0xbd, 0xe1, 0x42, 0x1a, 0x39, 0xa2, 0xc1, 0xde, 0xf7, 0x98,
	0x90, 0x11, 0x19, 0x28, 0x2a, 0x23, 0x07, 0x49, 0x8f, 0xb6, 0x59, 0x53, 0xf0, 0x53, 0xa6, 0x45,
	0xfe, 0xd5, 0x98, 0x56, 0x81, 0x5d, 0x7e, 0xca, 0x74, 0xad, 0x4a, 0x4a, 0xf7, 0x88, 0x39, 0x46,
	0xa9, 0x86, 0xef, 0xa9, 0x00, 0x39, 0x81, 0xf4, 0x30, 0xa3, 0xf0, 0x5c, 0x47, 0x30, 0xbc, 0x0e,
	0xd3, 0xc1, 0xd5, 0x33, 0x28, 0x3f, 0x51, 0x98, 0xa9, 0xe6, 0x4a, 0xa3, 0x67, 0x2f, 0x99, 0xba,
	0x46, 0x1f, 0x8c, 0x97, 0x61, 0xde, 0x61, 0x1f, 0x64, 0x33, 0x44, 0xea, 0xef, 0x2d, 0xa5, 0xc2,
	0xf5, 0x3e, 0xf1, 0x0e, 0xa4, 0x5f, 0xea, 0xc1, 0x83, 0x16, 0x66, 0xd6, 0x35, 0x98, 0x32, 0xbd,
	0xf4, 0xa0, 0xd7, 0xf0, 0x06, 0x58, 0xf2, 0x0d, 0x41, 0x7a, 0xdf, 0xb3, 0x46, 0xfb, 0x2d, 0x02,
	0x18, 0x4c, 0x68, 0x77, 0x26, 0xb2, 0x6d, 0x85, 0xe9, 0x12, 0x37, 0xa7, 0x53, 0x57, 0xee, 0x69,
	0x36, 0x6d, 0x35, 0xbd, 0xd6, 0xb8, 0x2b, 0xbf, 0x56, 0x6e, 0xdc, 0xa1, 0xe2, 0xa8, 0x01, 0x3e,
	0x5c, 0xfd, 0x26, 0x6b, 0x90, 0xde, 0x64, 0x1d, 0x76, 0x4b, 0xa9, 0x64, 0x1d, 0x16, 0xf6, 0x1d,
	0xeb, 0x37, 0x0a, 0xab, 0xf0, 0xef, 0x16, 0x93, 0xb7, 0xaa, 0xa9, 0x7e, 0x49, 0xc2, 0x3f, 0xa6,
	0x62, 0x97, 0x75, 0x8f, 0x79, 0x8b, 0xbd, 0xad, 0xe0, 0x13, 0x98, 0x0d, 0x9b, 0x05, 0xaf, 0xc4,
	0xed, 0x2a, 0xc6, 0xc0, 0xd9, 0xc2, 0xf5, 0x40, 0xdf, 0x77, 0x24, 0x7d, 0xf1, 0xfd, 0xc7, 0xe7,
	0xc4, 0x1c, 0x9e, 0x0d, 0xbf, 0x77, 0x70, 0x0f, 0x52, 0x43, 0x66, 0xc1, 0xb1, 0x0d, 0xe3, 0xfc,
	0x94, 0x1d, 0x77, 0x4f, 0xb2, 0xa8, 0xd9, 0xee, 0x90, 0x21, 0xb6, 0x5a, 0xff, 0xca, 0x02, 0x60,
	0xb0, 0x38, 0xfc, 0x28, 0xae, 0xd3, 0xc8, 0x62, 0xc7, 0x13, 0x2e, 0x69, 0xc2, 0x1c, 0xbe, 0x1b,
	0x26, 0x2c, 0x9f, 0x0d, 0x2e, 0x71, 0x8e, 0x2f, 0x10, 0xa4, 0x86, 0x9c, 0x1c, 0x3f, 0x6c, 0x9c,
	0xd9, 0xc7, 0x73, 0x17, 0x35, 0xf7, 0xc3, 0xea, 0xd5, 0xdc, 0x83, 0xc9, 0x25, 0xa4, 0x86, 0x2c,
	0x1a, 0xaf, 0x21, 0xce, 0xc5, 0xd9, 0x85, 0x91, 0xa7, 0xe0, 0x95, 0x7a, 0x61, 0x07, 0xa3, 0x17,
	0xc7, 0x8c, 0x7e, 0x89, 0x60, 0x3e, 0x62, 0x71, 0x5c, 0x8c, 0x1d, 0x3e, 0xf6, 0x39, 0x18, 0x3f,
	0xfe, 0x53, 0xcd, 0xbf, 0x42, 0xc8, 0xd5, 0xe3, 0xf7, 0x4c, 0xdb, 0x1a, 0x2a, 0xe2, 0x4f, 0x30,
	0xbb, 0xcb, 0xe4, 0x36, 0xb5, 0xeb, 0xfa, 0x63, 0x84, 0x49, 0xd0, 0x9b, 0x53, 0x5b, 0xb5, 0x0d,
	0x27, 0x03, 0xfe, 0xff, 0x23, 0x18, 0x3f, 0x4b, 0x2a, 0x9a, 0xf9, 0x09, 0x59, 0x56, 0xcc, 0x67,
	0x5d, 0x26, 0xdc, 0x5e, 0xb7, 0xc5, 0x5e, 0xf4, 0x35, 0x14, 0xcf, 0x6b, 0x22, 0xd4, 0xcd, 0xb0,
	0x6f, 0x8d, 0x63, 0xdf, 0xfa, 0xa3, 0xec, 0xed, 0x08, 0xfb, 0x57, 0x04, 0x78, 0x8f, 0x09, 0x1d,
	0x64, 0x5d, 0x9b, 0x0b, 0xa1, 0xbe, 0xe5, 0x03, 0x0f, 0x18, 0x82, 0x51, 0x48, 0x20, 0xe5, 0xf1,
	0x0d, 0x90, 0xe6, 0x81, 0x5f, 0xd7, 0xf2, 0x2a, 0x64, 0xf5, 0x6a, 0x79, 0x72, 0xa4, 0xba, 0x86,
	0x8a, 0x1b, 0xab, 0xb0, 0xd0, 0x72, 0xed, 0x98, 0x8b, 0x6f, 0xa4, 0x82, 0xb7, 0x4a, 0x5d, 0x39,
	0xb0, 0x8e, 0x7e, 0x22, 0x74, 0xf0, 0xb7, 0x76, 0xe3, 0xb3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x28, 0xdc, 0x2d, 0xf7, 0xdf, 0x08, 0x00, 0x00,
}
