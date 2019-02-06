// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/bigtable/admin/cluster/v1/bigtable_cluster_service_messages.proto

package google_bigtable_admin_cluster_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf3 "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Request message for BigtableClusterService.ListZones.
type ListZonesRequest struct {
	// The unique name of the project for which a list of supported zones is
	// requested.
	// Values are of the form projects/<project>
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ListZonesRequest) Reset()         { *m = ListZonesRequest{} }
func (m *ListZonesRequest) String() string { return proto.CompactTextString(m) }
func (*ListZonesRequest) ProtoMessage()    {}
func (*ListZonesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableClusterServiceMessages, []int{0}
}

func (m *ListZonesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Response message for BigtableClusterService.ListZones.
type ListZonesResponse struct {
	// The list of requested zones.
	Zones []*Zone `protobuf:"bytes,1,rep,name=zones" json:"zones,omitempty"`
}

func (m *ListZonesResponse) Reset()         { *m = ListZonesResponse{} }
func (m *ListZonesResponse) String() string { return proto.CompactTextString(m) }
func (*ListZonesResponse) ProtoMessage()    {}
func (*ListZonesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableClusterServiceMessages, []int{1}
}

func (m *ListZonesResponse) GetZones() []*Zone {
	if m != nil {
		return m.Zones
	}
	return nil
}

// Request message for BigtableClusterService.GetCluster.
type GetClusterRequest struct {
	// The unique name of the requested cluster.
	// Values are of the form projects/<project>/zones/<zone>/clusters/<cluster>
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetClusterRequest) Reset()         { *m = GetClusterRequest{} }
func (m *GetClusterRequest) String() string { return proto.CompactTextString(m) }
func (*GetClusterRequest) ProtoMessage()    {}
func (*GetClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableClusterServiceMessages, []int{2}
}

func (m *GetClusterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for BigtableClusterService.ListClusters.
type ListClustersRequest struct {
	// The unique name of the project for which a list of clusters is requested.
	// Values are of the form projects/<project>
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ListClustersRequest) Reset()         { *m = ListClustersRequest{} }
func (m *ListClustersRequest) String() string { return proto.CompactTextString(m) }
func (*ListClustersRequest) ProtoMessage()    {}
func (*ListClustersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableClusterServiceMessages, []int{3}
}

func (m *ListClustersRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Response message for BigtableClusterService.ListClusters.
type ListClustersResponse struct {
	// The list of requested Clusters.
	Clusters []*Cluster `protobuf:"bytes,1,rep,name=clusters" json:"clusters,omitempty"`
	// The zones for which clusters could not be retrieved.
	FailedZones []*Zone `protobuf:"bytes,2,rep,name=failed_zones,json=failedZones" json:"failed_zones,omitempty"`
}

func (m *ListClustersResponse) Reset()         { *m = ListClustersResponse{} }
func (m *ListClustersResponse) String() string { return proto.CompactTextString(m) }
func (*ListClustersResponse) ProtoMessage()    {}
func (*ListClustersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableClusterServiceMessages, []int{4}
}

func (m *ListClustersResponse) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *ListClustersResponse) GetFailedZones() []*Zone {
	if m != nil {
		return m.FailedZones
	}
	return nil
}

// Request message for BigtableClusterService.CreateCluster.
type CreateClusterRequest struct {
	// The unique name of the zone in which to create the cluster.
	// Values are of the form projects/<project>/zones/<zone>
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The id to be used when referring to the new cluster within its zone,
	// e.g. just the "test-cluster" section of the full name
	// "projects/<project>/zones/<zone>/clusters/test-cluster".
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// The cluster to create.
	// The "name", "delete_time", and "current_operation" fields must be left
	// blank.
	Cluster *Cluster `protobuf:"bytes,3,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *CreateClusterRequest) Reset()         { *m = CreateClusterRequest{} }
func (m *CreateClusterRequest) String() string { return proto.CompactTextString(m) }
func (*CreateClusterRequest) ProtoMessage()    {}
func (*CreateClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableClusterServiceMessages, []int{5}
}

func (m *CreateClusterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateClusterRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *CreateClusterRequest) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

// Metadata type for the operation returned by
// BigtableClusterService.CreateCluster.
type CreateClusterMetadata struct {
	// The request which prompted the creation of this operation.
	OriginalRequest *CreateClusterRequest `protobuf:"bytes,1,opt,name=original_request,json=originalRequest" json:"original_request,omitempty"`
	// The time at which original_request was received.
	RequestTime *google_protobuf3.Timestamp `protobuf:"bytes,2,opt,name=request_time,json=requestTime" json:"request_time,omitempty"`
	// The time at which this operation failed or was completed successfully.
	FinishTime *google_protobuf3.Timestamp `protobuf:"bytes,3,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
}

func (m *CreateClusterMetadata) Reset()         { *m = CreateClusterMetadata{} }
func (m *CreateClusterMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateClusterMetadata) ProtoMessage()    {}
func (*CreateClusterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableClusterServiceMessages, []int{6}
}

func (m *CreateClusterMetadata) GetOriginalRequest() *CreateClusterRequest {
	if m != nil {
		return m.OriginalRequest
	}
	return nil
}

func (m *CreateClusterMetadata) GetRequestTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *CreateClusterMetadata) GetFinishTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.FinishTime
	}
	return nil
}

// Metadata type for the operation returned by
// BigtableClusterService.UpdateCluster.
type UpdateClusterMetadata struct {
	// The request which prompted the creation of this operation.
	OriginalRequest *Cluster `protobuf:"bytes,1,opt,name=original_request,json=originalRequest" json:"original_request,omitempty"`
	// The time at which original_request was received.
	RequestTime *google_protobuf3.Timestamp `protobuf:"bytes,2,opt,name=request_time,json=requestTime" json:"request_time,omitempty"`
	// The time at which this operation was cancelled. If set, this operation is
	// in the process of undoing itself (which is guaranteed to succeed) and
	// cannot be cancelled again.
	CancelTime *google_protobuf3.Timestamp `protobuf:"bytes,3,opt,name=cancel_time,json=cancelTime" json:"cancel_time,omitempty"`
	// The time at which this operation failed or was completed successfully.
	FinishTime *google_protobuf3.Timestamp `protobuf:"bytes,4,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
}

func (m *UpdateClusterMetadata) Reset()         { *m = UpdateClusterMetadata{} }
func (m *UpdateClusterMetadata) String() string { return proto.CompactTextString(m) }
func (*UpdateClusterMetadata) ProtoMessage()    {}
func (*UpdateClusterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableClusterServiceMessages, []int{7}
}

func (m *UpdateClusterMetadata) GetOriginalRequest() *Cluster {
	if m != nil {
		return m.OriginalRequest
	}
	return nil
}

func (m *UpdateClusterMetadata) GetRequestTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *UpdateClusterMetadata) GetCancelTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.CancelTime
	}
	return nil
}

func (m *UpdateClusterMetadata) GetFinishTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.FinishTime
	}
	return nil
}

// Request message for BigtableClusterService.DeleteCluster.
type DeleteClusterRequest struct {
	// The unique name of the cluster to be deleted.
	// Values are of the form projects/<project>/zones/<zone>/clusters/<cluster>
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DeleteClusterRequest) Reset()         { *m = DeleteClusterRequest{} }
func (m *DeleteClusterRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteClusterRequest) ProtoMessage()    {}
func (*DeleteClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableClusterServiceMessages, []int{8}
}

func (m *DeleteClusterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for BigtableClusterService.UndeleteCluster.
type UndeleteClusterRequest struct {
	// The unique name of the cluster to be un-deleted.
	// Values are of the form projects/<project>/zones/<zone>/clusters/<cluster>
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *UndeleteClusterRequest) Reset()         { *m = UndeleteClusterRequest{} }
func (m *UndeleteClusterRequest) String() string { return proto.CompactTextString(m) }
func (*UndeleteClusterRequest) ProtoMessage()    {}
func (*UndeleteClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableClusterServiceMessages, []int{9}
}

func (m *UndeleteClusterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Metadata type for the operation returned by
// BigtableClusterService.UndeleteCluster.
type UndeleteClusterMetadata struct {
	// The time at which the original request was received.
	RequestTime *google_protobuf3.Timestamp `protobuf:"bytes,1,opt,name=request_time,json=requestTime" json:"request_time,omitempty"`
	// The time at which this operation failed or was completed successfully.
	FinishTime *google_protobuf3.Timestamp `protobuf:"bytes,2,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
}

func (m *UndeleteClusterMetadata) Reset()         { *m = UndeleteClusterMetadata{} }
func (m *UndeleteClusterMetadata) String() string { return proto.CompactTextString(m) }
func (*UndeleteClusterMetadata) ProtoMessage()    {}
func (*UndeleteClusterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableClusterServiceMessages, []int{10}
}

func (m *UndeleteClusterMetadata) GetRequestTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *UndeleteClusterMetadata) GetFinishTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.FinishTime
	}
	return nil
}

// Metadata type for operations initiated by the V2 BigtableAdmin service.
// More complete information for such operations is available via the V2 API.
type V2OperationMetadata struct {
}

func (m *V2OperationMetadata) Reset()         { *m = V2OperationMetadata{} }
func (m *V2OperationMetadata) String() string { return proto.CompactTextString(m) }
func (*V2OperationMetadata) ProtoMessage()    {}
func (*V2OperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableClusterServiceMessages, []int{11}
}

func init() {
	proto.RegisterType((*ListZonesRequest)(nil), "google.bigtable.admin.cluster.v1.ListZonesRequest")
	proto.RegisterType((*ListZonesResponse)(nil), "google.bigtable.admin.cluster.v1.ListZonesResponse")
	proto.RegisterType((*GetClusterRequest)(nil), "google.bigtable.admin.cluster.v1.GetClusterRequest")
	proto.RegisterType((*ListClustersRequest)(nil), "google.bigtable.admin.cluster.v1.ListClustersRequest")
	proto.RegisterType((*ListClustersResponse)(nil), "google.bigtable.admin.cluster.v1.ListClustersResponse")
	proto.RegisterType((*CreateClusterRequest)(nil), "google.bigtable.admin.cluster.v1.CreateClusterRequest")
	proto.RegisterType((*CreateClusterMetadata)(nil), "google.bigtable.admin.cluster.v1.CreateClusterMetadata")
	proto.RegisterType((*UpdateClusterMetadata)(nil), "google.bigtable.admin.cluster.v1.UpdateClusterMetadata")
	proto.RegisterType((*DeleteClusterRequest)(nil), "google.bigtable.admin.cluster.v1.DeleteClusterRequest")
	proto.RegisterType((*UndeleteClusterRequest)(nil), "google.bigtable.admin.cluster.v1.UndeleteClusterRequest")
	proto.RegisterType((*UndeleteClusterMetadata)(nil), "google.bigtable.admin.cluster.v1.UndeleteClusterMetadata")
	proto.RegisterType((*V2OperationMetadata)(nil), "google.bigtable.admin.cluster.v1.V2OperationMetadata")
}

func init() {
	proto.RegisterFile("google/bigtable/admin/cluster/v1/bigtable_cluster_service_messages.proto", fileDescriptorBigtableClusterServiceMessages)
}

var fileDescriptorBigtableClusterServiceMessages = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xd5, 0x26, 0xe5, 0xa3, 0xe3, 0x4a, 0xb4, 0x6e, 0x02, 0x51, 0x24, 0x44, 0x64, 0x50, 0x69,
	0x11, 0xb2, 0xd5, 0x20, 0x71, 0xa1, 0x5c, 0x12, 0x10, 0x54, 0x22, 0xa2, 0x98, 0x96, 0x03, 0x17,
	0x6b, 0x63, 0x4f, 0xcc, 0x4a, 0xb6, 0xd7, 0x78, 0x37, 0x39, 0xf0, 0x23, 0xb8, 0xf1, 0x17, 0x10,
	0xbf, 0x90, 0x33, 0xb2, 0x77, 0x37, 0xa2, 0x55, 0x88, 0x63, 0xa1, 0xde, 0xec, 0x99, 0xf7, 0x66,
	0xde, 0x1b, 0x3f, 0xc9, 0xf0, 0x36, 0xe6, 0x3c, 0x4e, 0xd0, 0x9b, 0xb2, 0x58, 0xd2, 0x69, 0x82,
	0x1e, 0x8d, 0x52, 0x96, 0x79, 0x61, 0x32, 0x17, 0x12, 0x0b, 0x6f, 0x71, 0xbc, 0xec, 0x04, 0xba,
	0x16, 0x08, 0x2c, 0x16, 0x2c, 0xc4, 0x20, 0x45, 0x21, 0x68, 0x8c, 0xc2, 0xcd, 0x0b, 0x2e, 0xb9,
	0x3d, 0x50, 0x93, 0x5c, 0x83, 0x77, 0xab, 0x49, 0xae, 0x66, 0xb9, 0x8b, 0xe3, 0xfe, 0x49, 0xf3,
	0x5d, 0x11, 0x95, 0x54, 0xcd, 0xef, 0x3f, 0xd0, 0xec, 0xea, 0x6d, 0x3a, 0x9f, 0x79, 0x92, 0xa5,
	0x28, 0x24, 0x4d, 0x73, 0x05, 0x70, 0x0e, 0x60, 0xf7, 0x1d, 0x13, 0xf2, 0x33, 0xcf, 0x50, 0xf8,
	0xf8, 0x75, 0x8e, 0x42, 0xda, 0x36, 0x6c, 0x65, 0x34, 0xc5, 0x1e, 0x19, 0x90, 0xc3, 0x6d, 0xbf,
	0x7a, 0x76, 0x3e, 0xc0, 0xde, 0x5f, 0x38, 0x91, 0xf3, 0x4c, 0xa0, 0x7d, 0x02, 0x37, 0xbe, 0x95,
	0x85, 0x1e, 0x19, 0xb4, 0x0f, 0xad, 0xe1, 0x81, 0x5b, 0xe7, 0xc6, 0x2d, 0xf9, 0xbe, 0x22, 0x39,
	0x8f, 0x61, 0xef, 0x0d, 0xca, 0xb1, 0x6a, 0xae, 0xdb, 0x7d, 0x04, 0xfb, 0xe5, 0x6e, 0x8d, 0x5c,
	0x2b, 0xf3, 0x17, 0x81, 0xce, 0x65, 0xac, 0x96, 0xfa, 0x1a, 0x6e, 0x6b, 0x19, 0x46, 0xed, 0x51,
	0xbd, 0x5a, 0xa3, 0x6d, 0x49, 0xb5, 0x4f, 0x61, 0x67, 0x46, 0x59, 0x82, 0x51, 0xa0, 0x8c, 0xb7,
	0x1a, 0x19, 0xb7, 0x14, 0xb7, 0x3a, 0xa2, 0xf3, 0x9d, 0x40, 0x67, 0x5c, 0x20, 0x95, 0x58, 0x7f,
	0x02, 0xfb, 0x3e, 0x80, 0xf9, 0xba, 0x2c, 0xea, 0xb5, 0xaa, 0xce, 0xb6, 0xae, 0x9c, 0x46, 0xf6,
	0x18, 0x6e, 0xe9, 0x97, 0x5e, 0x7b, 0x40, 0x9a, 0x99, 0x33, 0x4c, 0xe7, 0x37, 0x81, 0xee, 0x25,
	0x41, 0x13, 0x94, 0xb4, 0xcc, 0x92, 0x4d, 0x61, 0x97, 0x17, 0x2c, 0x66, 0x19, 0x4d, 0x82, 0x42,
	0xa9, 0xac, 0xd4, 0x59, 0xc3, 0xe7, 0x1b, 0xec, 0x59, 0xe1, 0xd1, 0xbf, 0x63, 0xe6, 0x19, 0xd3,
	0x2f, 0x61, 0x47, 0x4f, 0x0e, 0xca, 0x88, 0x56, 0x16, 0xad, 0x61, 0xdf, 0x8c, 0x37, 0xf9, 0x75,
	0xcf, 0x4d, 0x7e, 0x7d, 0x4b, 0xe3, 0xcb, 0x8a, 0xfd, 0x02, 0xac, 0x19, 0xcb, 0x98, 0xf8, 0xa2,
	0xd8, 0xed, 0x5a, 0x36, 0x28, 0x78, 0x59, 0x70, 0x7e, 0xb6, 0xa0, 0x7b, 0x91, 0x47, 0x2b, 0x8c,
	0x9f, 0xff, 0xd3, 0x78, 0x83, 0x03, 0x5f, 0x83, 0xd7, 0x90, 0x66, 0x21, 0x26, 0x1b, 0x7b, 0x55,
	0xf0, 0x55, 0x87, 0xda, 0x6a, 0x74, 0xa8, 0x27, 0xd0, 0x79, 0x85, 0x09, 0x6e, 0x92, 0x58, 0xe7,
	0x29, 0xdc, 0xbd, 0xc8, 0xa2, 0x4d, 0xd1, 0x3f, 0x08, 0xdc, 0xbb, 0x02, 0x5f, 0x7e, 0x84, 0xab,
	0xe7, 0x22, 0xff, 0x15, 0x8d, 0x56, 0x23, 0xc7, 0x5d, 0xd8, 0xff, 0x34, 0x7c, 0x9f, 0x63, 0x41,
	0x25, 0xe3, 0x99, 0x91, 0x34, 0x9a, 0xc0, 0xa3, 0x90, 0xa7, 0xb5, 0x11, 0x18, 0x3d, 0x1c, 0xe9,
	0x96, 0xf6, 0xf4, 0x51, 0xfd, 0x05, 0x26, 0xfa, 0x27, 0x70, 0x56, 0x2e, 0x3f, 0x23, 0xd3, 0x9b,
	0x95, 0x8a, 0x67, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x98, 0x20, 0xd5, 0x11, 0x58, 0x06, 0x00,
	0x00,
}
