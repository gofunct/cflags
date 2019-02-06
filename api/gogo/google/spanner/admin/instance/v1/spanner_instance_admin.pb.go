// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/spanner/admin/instance/v1/spanner_instance_admin.proto

package google_spanner_admin_instance_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import _ "go.pedge.io/pb/gogo/google/iam/v1"
import _ "go.pedge.io/pb/gogo/google/iam/v1"
import _ "go.pedge.io/pb/gogo/google/longrunning"
import _ "github.com/gogo/protobuf/types"
import google_protobuf3 "go.pedge.io/pb/gogo/google/protobuf"
import google_protobuf4 "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Indicates the current state of the instance.
type Instance_State int32

const (
	// Not specified.
	Instance_STATE_UNSPECIFIED Instance_State = 0
	// The instance is still being created. Resources may not be
	// available yet, and operations such as database creation may not
	// work.
	Instance_CREATING Instance_State = 1
	// The instance is fully created and ready to do work such as
	// creating databases.
	Instance_READY Instance_State = 2
)

var Instance_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "CREATING",
	2: "READY",
}
var Instance_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"CREATING":          1,
	"READY":             2,
}

func (x Instance_State) String() string {
	return proto.EnumName(Instance_State_name, int32(x))
}
func (Instance_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorSpannerInstanceAdmin, []int{1, 0}
}

// A possible configuration for a Cloud Spanner instance. Configurations
// define the geographic placement of nodes and their replication.
type InstanceConfig struct {
	// A unique identifier for the instance configuration.  Values
	// are of the form
	// `projects/<project>/instanceConfigs/[a-z][-a-z0-9]*`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The name of this instance configuration as it appears in UIs.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (m *InstanceConfig) Reset()         { *m = InstanceConfig{} }
func (m *InstanceConfig) String() string { return proto.CompactTextString(m) }
func (*InstanceConfig) ProtoMessage()    {}
func (*InstanceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerInstanceAdmin, []int{0}
}

func (m *InstanceConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InstanceConfig) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

// An isolated set of Cloud Spanner resources on which databases can be hosted.
type Instance struct {
	// Required. A unique identifier for the instance, which cannot be changed
	// after the instance is created. Values are of the form
	// `projects/<project>/instances/[a-z][-a-z0-9]*[a-z0-9]`. The final
	// segment of the name must be between 6 and 30 characters in length.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of the instance's configuration. Values are of the form
	// `projects/<project>/instanceConfigs/<configuration>`. See
	// also [InstanceConfig][google.spanner.admin.instance.v1.InstanceConfig] and
	// [ListInstanceConfigs][google.spanner.admin.instance.v1.InstanceAdmin.ListInstanceConfigs].
	Config string `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	// Required. The descriptive name for this instance as it appears in UIs.
	// Must be unique per project and between 4 and 30 characters in length.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. The number of nodes allocated to this instance. This may be zero
	// in API responses for instances that are not yet in state `READY`.
	//
	// See [the documentation](https://cloud.google.com/spanner/docs/instances#node_count)
	// for more information about nodes.
	NodeCount int32 `protobuf:"varint,5,opt,name=node_count,json=nodeCount,proto3" json:"node_count,omitempty"`
	// Output only. The current instance state. For
	// [CreateInstance][google.spanner.admin.instance.v1.InstanceAdmin.CreateInstance], the state must be
	// either omitted or set to `CREATING`. For
	// [UpdateInstance][google.spanner.admin.instance.v1.InstanceAdmin.UpdateInstance], the state must be
	// either omitted or set to `READY`.
	State Instance_State `protobuf:"varint,6,opt,name=state,proto3,enum=google.spanner.admin.instance.v1.Instance_State" json:"state,omitempty"`
	// Cloud Labels are a flexible and lightweight mechanism for organizing cloud
	// resources into groups that reflect a customer's organizational needs and
	// deployment strategies. Cloud Labels can be used to filter collections of
	// resources. They can be used to control how resource metrics are aggregated.
	// And they can be used as arguments to policy management rules (e.g. route,
	// firewall, load balancing, etc.).
	//
	//  * Label keys must be between 1 and 63 characters long and must conform to
	//    the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?`.
	//  * Label values must be between 0 and 63 characters long and must conform
	//    to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
	//  * No more than 64 labels can be associated with a given resource.
	//
	// See https://goo.gl/xmQnxf for more information on and examples of labels.
	//
	// If you plan to use labels in your own code, please note that additional
	// characters may be allowed in the future. And so you are advised to use an
	// internal label representation, such as JSON, which doesn't rely upon
	// specific characters being disallowed.  For example, representing labels
	// as the string:  name + "_" + value  would prove problematic if we were to
	// allow "_" in a future release.
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptorSpannerInstanceAdmin, []int{1} }

func (m *Instance) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Instance) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func (m *Instance) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Instance) GetNodeCount() int32 {
	if m != nil {
		return m.NodeCount
	}
	return 0
}

func (m *Instance) GetState() Instance_State {
	if m != nil {
		return m.State
	}
	return Instance_STATE_UNSPECIFIED
}

func (m *Instance) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// The request for [ListInstanceConfigs][google.spanner.admin.instance.v1.InstanceAdmin.ListInstanceConfigs].
type ListInstanceConfigsRequest struct {
	// Required. The name of the project for which a list of supported instance
	// configurations is requested. Values are of the form
	// `projects/<project>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Number of instance configurations to be returned in the response. If 0 or
	// less, defaults to the server's maximum allowed page size.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// If non-empty, `page_token` should contain a
	// [next_page_token][google.spanner.admin.instance.v1.ListInstanceConfigsResponse.next_page_token]
	// from a previous [ListInstanceConfigsResponse][google.spanner.admin.instance.v1.ListInstanceConfigsResponse].
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (m *ListInstanceConfigsRequest) Reset()         { *m = ListInstanceConfigsRequest{} }
func (m *ListInstanceConfigsRequest) String() string { return proto.CompactTextString(m) }
func (*ListInstanceConfigsRequest) ProtoMessage()    {}
func (*ListInstanceConfigsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerInstanceAdmin, []int{2}
}

func (m *ListInstanceConfigsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListInstanceConfigsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListInstanceConfigsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response for [ListInstanceConfigs][google.spanner.admin.instance.v1.InstanceAdmin.ListInstanceConfigs].
type ListInstanceConfigsResponse struct {
	// The list of requested instance configurations.
	InstanceConfigs []*InstanceConfig `protobuf:"bytes,1,rep,name=instance_configs,json=instanceConfigs" json:"instance_configs,omitempty"`
	// `next_page_token` can be sent in a subsequent
	// [ListInstanceConfigs][google.spanner.admin.instance.v1.InstanceAdmin.ListInstanceConfigs] call to
	// fetch more of the matching instance configurations.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *ListInstanceConfigsResponse) Reset()         { *m = ListInstanceConfigsResponse{} }
func (m *ListInstanceConfigsResponse) String() string { return proto.CompactTextString(m) }
func (*ListInstanceConfigsResponse) ProtoMessage()    {}
func (*ListInstanceConfigsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerInstanceAdmin, []int{3}
}

func (m *ListInstanceConfigsResponse) GetInstanceConfigs() []*InstanceConfig {
	if m != nil {
		return m.InstanceConfigs
	}
	return nil
}

func (m *ListInstanceConfigsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request for
// [GetInstanceConfigRequest][google.spanner.admin.instance.v1.InstanceAdmin.GetInstanceConfig].
type GetInstanceConfigRequest struct {
	// Required. The name of the requested instance configuration. Values are of
	// the form `projects/<project>/instanceConfigs/<config>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetInstanceConfigRequest) Reset()         { *m = GetInstanceConfigRequest{} }
func (m *GetInstanceConfigRequest) String() string { return proto.CompactTextString(m) }
func (*GetInstanceConfigRequest) ProtoMessage()    {}
func (*GetInstanceConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerInstanceAdmin, []int{4}
}

func (m *GetInstanceConfigRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request for [GetInstance][google.spanner.admin.instance.v1.InstanceAdmin.GetInstance].
type GetInstanceRequest struct {
	// Required. The name of the requested instance. Values are of the form
	// `projects/<project>/instances/<instance>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetInstanceRequest) Reset()         { *m = GetInstanceRequest{} }
func (m *GetInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*GetInstanceRequest) ProtoMessage()    {}
func (*GetInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerInstanceAdmin, []int{5}
}

func (m *GetInstanceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request for [CreateInstance][google.spanner.admin.instance.v1.InstanceAdmin.CreateInstance].
type CreateInstanceRequest struct {
	// Required. The name of the project in which to create the instance. Values
	// are of the form `projects/<project>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The ID of the instance to create.  Valid identifiers are of the
	// form `[a-z][-a-z0-9]*[a-z0-9]` and must be between 6 and 30 characters in
	// length.
	InstanceId string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// Required. The instance to create.  The name may be omitted, but if
	// specified must be `<parent>/instances/<instance_id>`.
	Instance *Instance `protobuf:"bytes,3,opt,name=instance" json:"instance,omitempty"`
}

func (m *CreateInstanceRequest) Reset()         { *m = CreateInstanceRequest{} }
func (m *CreateInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateInstanceRequest) ProtoMessage()    {}
func (*CreateInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerInstanceAdmin, []int{6}
}

func (m *CreateInstanceRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateInstanceRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *CreateInstanceRequest) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

// The request for [ListInstances][google.spanner.admin.instance.v1.InstanceAdmin.ListInstances].
type ListInstancesRequest struct {
	// Required. The name of the project for which a list of instances is
	// requested. Values are of the form `projects/<project>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Number of instances to be returned in the response. If 0 or less, defaults
	// to the server's maximum allowed page size.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// If non-empty, `page_token` should contain a
	// [next_page_token][google.spanner.admin.instance.v1.ListInstancesResponse.next_page_token] from a
	// previous [ListInstancesResponse][google.spanner.admin.instance.v1.ListInstancesResponse].
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// An expression for filtering the results of the request. Filter rules are
	// case insensitive. The fields eligible for filtering are:
	//
	//   * `name`
	//   * `display_name`
	//   * `labels.key` where key is the name of a label
	//
	// Some examples of using filters are:
	//
	//   * `name:*` --> The instance has a name.
	//   * `name:Howl` --> The instance's name contains the string "howl".
	//   * `name:HOWL` --> Equivalent to above.
	//   * `NAME:howl` --> Equivalent to above.
	//   * `labels.env:*` --> The instance has the label "env".
	//   * `labels.env:dev` --> The instance has the label "env" and the value of
	//                        the label contains the string "dev".
	//   * `name:howl labels.env:dev` --> The instance's name contains "howl" and
	//                                  it has the label "env" with its value
	//                                  containing "dev".
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (m *ListInstancesRequest) Reset()         { *m = ListInstancesRequest{} }
func (m *ListInstancesRequest) String() string { return proto.CompactTextString(m) }
func (*ListInstancesRequest) ProtoMessage()    {}
func (*ListInstancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerInstanceAdmin, []int{7}
}

func (m *ListInstancesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListInstancesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListInstancesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListInstancesRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// The response for [ListInstances][google.spanner.admin.instance.v1.InstanceAdmin.ListInstances].
type ListInstancesResponse struct {
	// The list of requested instances.
	Instances []*Instance `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty"`
	// `next_page_token` can be sent in a subsequent
	// [ListInstances][google.spanner.admin.instance.v1.InstanceAdmin.ListInstances] call to fetch more
	// of the matching instances.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *ListInstancesResponse) Reset()         { *m = ListInstancesResponse{} }
func (m *ListInstancesResponse) String() string { return proto.CompactTextString(m) }
func (*ListInstancesResponse) ProtoMessage()    {}
func (*ListInstancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerInstanceAdmin, []int{8}
}

func (m *ListInstancesResponse) GetInstances() []*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *ListInstancesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request for [UpdateInstance][google.spanner.admin.instance.v1.InstanceAdmin.UpdateInstance].
type UpdateInstanceRequest struct {
	// Required. The instance to update, which must always include the instance
	// name.  Otherwise, only fields mentioned in [][google.spanner.admin.instance.v1.UpdateInstanceRequest.field_mask] need be included.
	Instance *Instance `protobuf:"bytes,1,opt,name=instance" json:"instance,omitempty"`
	// Required. A mask specifying which fields in [][google.spanner.admin.instance.v1.UpdateInstanceRequest.instance] should be updated.
	// The field mask must always be specified; this prevents any future fields in
	// [][google.spanner.admin.instance.v1.Instance] from being erased accidentally by clients that do not know
	// about them.
	FieldMask *google_protobuf3.FieldMask `protobuf:"bytes,2,opt,name=field_mask,json=fieldMask" json:"field_mask,omitempty"`
}

func (m *UpdateInstanceRequest) Reset()         { *m = UpdateInstanceRequest{} }
func (m *UpdateInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateInstanceRequest) ProtoMessage()    {}
func (*UpdateInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerInstanceAdmin, []int{9}
}

func (m *UpdateInstanceRequest) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *UpdateInstanceRequest) GetFieldMask() *google_protobuf3.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

// The request for [DeleteInstance][google.spanner.admin.instance.v1.InstanceAdmin.DeleteInstance].
type DeleteInstanceRequest struct {
	// Required. The name of the instance to be deleted. Values are of the form
	// `projects/<project>/instances/<instance>`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DeleteInstanceRequest) Reset()         { *m = DeleteInstanceRequest{} }
func (m *DeleteInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteInstanceRequest) ProtoMessage()    {}
func (*DeleteInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerInstanceAdmin, []int{10}
}

func (m *DeleteInstanceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Metadata type for the operation returned by
// [CreateInstance][google.spanner.admin.instance.v1.InstanceAdmin.CreateInstance].
type CreateInstanceMetadata struct {
	// The instance being created.
	Instance *Instance `protobuf:"bytes,1,opt,name=instance" json:"instance,omitempty"`
	// The time at which the
	// [CreateInstance][google.spanner.admin.instance.v1.InstanceAdmin.CreateInstance] request was
	// received.
	StartTime *google_protobuf4.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// The time at which this operation was cancelled. If set, this operation is
	// in the process of undoing itself (which is guaranteed to succeed) and
	// cannot be cancelled again.
	CancelTime *google_protobuf4.Timestamp `protobuf:"bytes,3,opt,name=cancel_time,json=cancelTime" json:"cancel_time,omitempty"`
	// The time at which this operation failed or was completed successfully.
	EndTime *google_protobuf4.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
}

func (m *CreateInstanceMetadata) Reset()         { *m = CreateInstanceMetadata{} }
func (m *CreateInstanceMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateInstanceMetadata) ProtoMessage()    {}
func (*CreateInstanceMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerInstanceAdmin, []int{11}
}

func (m *CreateInstanceMetadata) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *CreateInstanceMetadata) GetStartTime() *google_protobuf4.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *CreateInstanceMetadata) GetCancelTime() *google_protobuf4.Timestamp {
	if m != nil {
		return m.CancelTime
	}
	return nil
}

func (m *CreateInstanceMetadata) GetEndTime() *google_protobuf4.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// Metadata type for the operation returned by
// [UpdateInstance][google.spanner.admin.instance.v1.InstanceAdmin.UpdateInstance].
type UpdateInstanceMetadata struct {
	// The desired end state of the update.
	Instance *Instance `protobuf:"bytes,1,opt,name=instance" json:"instance,omitempty"`
	// The time at which [UpdateInstance][google.spanner.admin.instance.v1.InstanceAdmin.UpdateInstance]
	// request was received.
	StartTime *google_protobuf4.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// The time at which this operation was cancelled. If set, this operation is
	// in the process of undoing itself (which is guaranteed to succeed) and
	// cannot be cancelled again.
	CancelTime *google_protobuf4.Timestamp `protobuf:"bytes,3,opt,name=cancel_time,json=cancelTime" json:"cancel_time,omitempty"`
	// The time at which this operation failed or was completed successfully.
	EndTime *google_protobuf4.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
}

func (m *UpdateInstanceMetadata) Reset()         { *m = UpdateInstanceMetadata{} }
func (m *UpdateInstanceMetadata) String() string { return proto.CompactTextString(m) }
func (*UpdateInstanceMetadata) ProtoMessage()    {}
func (*UpdateInstanceMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptorSpannerInstanceAdmin, []int{12}
}

func (m *UpdateInstanceMetadata) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *UpdateInstanceMetadata) GetStartTime() *google_protobuf4.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *UpdateInstanceMetadata) GetCancelTime() *google_protobuf4.Timestamp {
	if m != nil {
		return m.CancelTime
	}
	return nil
}

func (m *UpdateInstanceMetadata) GetEndTime() *google_protobuf4.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func init() {
	proto.RegisterType((*InstanceConfig)(nil), "google.spanner.admin.instance.v1.InstanceConfig")
	proto.RegisterType((*Instance)(nil), "google.spanner.admin.instance.v1.Instance")
	proto.RegisterType((*ListInstanceConfigsRequest)(nil), "google.spanner.admin.instance.v1.ListInstanceConfigsRequest")
	proto.RegisterType((*ListInstanceConfigsResponse)(nil), "google.spanner.admin.instance.v1.ListInstanceConfigsResponse")
	proto.RegisterType((*GetInstanceConfigRequest)(nil), "google.spanner.admin.instance.v1.GetInstanceConfigRequest")
	proto.RegisterType((*GetInstanceRequest)(nil), "google.spanner.admin.instance.v1.GetInstanceRequest")
	proto.RegisterType((*CreateInstanceRequest)(nil), "google.spanner.admin.instance.v1.CreateInstanceRequest")
	proto.RegisterType((*ListInstancesRequest)(nil), "google.spanner.admin.instance.v1.ListInstancesRequest")
	proto.RegisterType((*ListInstancesResponse)(nil), "google.spanner.admin.instance.v1.ListInstancesResponse")
	proto.RegisterType((*UpdateInstanceRequest)(nil), "google.spanner.admin.instance.v1.UpdateInstanceRequest")
	proto.RegisterType((*DeleteInstanceRequest)(nil), "google.spanner.admin.instance.v1.DeleteInstanceRequest")
	proto.RegisterType((*CreateInstanceMetadata)(nil), "google.spanner.admin.instance.v1.CreateInstanceMetadata")
	proto.RegisterType((*UpdateInstanceMetadata)(nil), "google.spanner.admin.instance.v1.UpdateInstanceMetadata")
	proto.RegisterEnum("google.spanner.admin.instance.v1.Instance_State", Instance_State_name, Instance_State_value)
}

func init() {
	proto.RegisterFile("google/spanner/admin/instance/v1/spanner_instance_admin.proto", fileDescriptorSpannerInstanceAdmin)
}

var fileDescriptorSpannerInstanceAdmin = []byte{
	// 1181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x57, 0xcf, 0x6f, 0x1b, 0xc5,
	0x17, 0xff, 0x8e, 0x53, 0xa7, 0xf1, 0x73, 0x9a, 0xa6, 0xf3, 0xad, 0x2b, 0xe3, 0x52, 0xea, 0x6e,
	0x51, 0x71, 0x0d, 0xda, 0x25, 0x86, 0xfe, 0x0a, 0xe4, 0x90, 0x3a, 0x8e, 0xb1, 0xd4, 0x86, 0x68,
	0xed, 0x56, 0x42, 0x44, 0xb2, 0x26, 0xf6, 0xc4, 0x5a, 0xb2, 0xbf, 0xd8, 0x19, 0x47, 0xa4, 0xa8,
	0x97, 0x8a, 0x03, 0x48, 0x48, 0x1c, 0x10, 0x08, 0xf5, 0x82, 0xc4, 0x11, 0x24, 0x0e, 0xfc, 0x0b,
	0xdc, 0xb8, 0xf2, 0x2f, 0x20, 0xfe, 0x0e, 0x34, 0xb3, 0x3b, 0x8e, 0x77, 0x6d, 0xc7, 0x76, 0x45,
	0x4f, 0xdc, 0x76, 0xe6, 0x7d, 0xde, 0x7b, 0x9f, 0xf9, 0xbc, 0xd9, 0xf7, 0x76, 0x61, 0xa3, 0xe7,
	0x79, 0x3d, 0x9b, 0x1a, 0xcc, 0x27, 0xae, 0x4b, 0x03, 0x83, 0x74, 0x1d, 0xcb, 0x35, 0x2c, 0x97,
	0x71, 0xe2, 0x76, 0xa8, 0x71, 0xb4, 0xa6, 0x2c, 0x6d, 0xb5, 0xd7, 0x96, 0x10, 0xdd, 0x0f, 0x3c,
	0xee, 0xe1, 0x62, 0xe8, 0xae, 0x47, 0x20, 0x3d, 0xb4, 0x29, 0xa8, 0x7e, 0xb4, 0x56, 0x78, 0x35,
	0x4a, 0x40, 0x7c, 0xcb, 0x20, 0xae, 0xeb, 0x71, 0xc2, 0x2d, 0xcf, 0x65, 0xa1, 0x7f, 0xe1, 0xb5,
	0xc8, 0x6a, 0x11, 0x47, 0xe4, 0xb2, 0x88, 0xd3, 0xf6, 0x3d, 0xdb, 0xea, 0x1c, 0x47, 0xf6, 0x42,
	0xdc, 0x1e, 0xb3, 0x5d, 0x8f, 0x6c, 0xb6, 0xe7, 0xf6, 0x82, 0xbe, 0xeb, 0x5a, 0x6e, 0xcf, 0xf0,
	0x7c, 0x1a, 0xc4, 0x12, 0x5c, 0x8e, 0x40, 0x72, 0xb5, 0xdf, 0x3f, 0x30, 0xa8, 0xe3, 0x73, 0x15,
	0xa1, 0x98, 0x34, 0x1e, 0x58, 0xd4, 0xee, 0xb6, 0x1d, 0xc2, 0x0e, 0x23, 0xc4, 0xd5, 0x24, 0x82,
	0x5b, 0x0e, 0x65, 0x9c, 0x38, 0x7e, 0x08, 0xd0, 0xea, 0xb0, 0xd2, 0x88, 0x4e, 0x5b, 0xf5, 0xdc,
	0x03, 0xab, 0x87, 0x31, 0x9c, 0x71, 0x89, 0x43, 0xf3, 0xa8, 0x88, 0x4a, 0x19, 0x53, 0x3e, 0xe3,
	0x6b, 0xb0, 0xdc, 0xb5, 0x98, 0x6f, 0x93, 0xe3, 0xb6, 0xb4, 0xa5, 0xa4, 0x2d, 0x1b, 0xed, 0xed,
	0x10, 0x87, 0x6a, 0x5f, 0x2c, 0xc0, 0x92, 0x8a, 0x34, 0x36, 0xc6, 0x25, 0x58, 0xec, 0xc8, 0x0c,
	0x91, 0x77, 0xb4, 0x1a, 0x89, 0xbd, 0x30, 0x12, 0x1b, 0x5f, 0x01, 0x70, 0xbd, 0x2e, 0x6d, 0x77,
	0xbc, 0xbe, 0xcb, 0xf3, 0xe9, 0x22, 0x2a, 0xa5, 0xcd, 0x8c, 0xd8, 0xa9, 0x8a, 0x0d, 0xbc, 0x0d,
	0x69, 0xc6, 0x09, 0xa7, 0xf9, 0xc5, 0x22, 0x2a, 0xad, 0x54, 0xde, 0xd6, 0xa7, 0x15, 0x55, 0x57,
	0x44, 0xf5, 0xa6, 0xf0, 0x33, 0x43, 0x77, 0xbc, 0x03, 0x8b, 0x36, 0xd9, 0xa7, 0x36, 0xcb, 0x9f,
	0x2d, 0x2e, 0x94, 0xb2, 0x95, 0xdb, 0x73, 0x04, 0x7a, 0x20, 0x1d, 0x6b, 0x2e, 0x0f, 0x8e, 0xcd,
	0x28, 0x4a, 0xe1, 0x1e, 0x64, 0x87, 0xb6, 0xf1, 0x2a, 0x2c, 0x1c, 0xd2, 0xe3, 0x48, 0x13, 0xf1,
	0x88, 0x2f, 0x42, 0xfa, 0x88, 0xd8, 0x7d, 0xa5, 0x67, 0xb8, 0x58, 0x4f, 0xdd, 0x45, 0xda, 0x1d,
	0x48, 0x4b, 0x6a, 0x38, 0x07, 0x17, 0x9a, 0xad, 0xcd, 0x56, 0xad, 0xfd, 0x68, 0xa7, 0xb9, 0x5b,
	0xab, 0x36, 0xb6, 0x1b, 0xb5, 0xad, 0xd5, 0xff, 0xe1, 0x65, 0x58, 0xaa, 0x9a, 0xb5, 0xcd, 0x56,
	0x63, 0xa7, 0xbe, 0x8a, 0x70, 0x06, 0xd2, 0x66, 0x6d, 0x73, 0xeb, 0xa3, 0xd5, 0x94, 0xe6, 0x43,
	0xe1, 0x81, 0xc5, 0x78, 0xbc, 0xa6, 0xcc, 0xa4, 0x9f, 0xf6, 0x29, 0xe3, 0xa2, 0x06, 0x3e, 0x09,
	0xa8, 0xcb, 0x23, 0x16, 0xd1, 0x0a, 0x5f, 0x86, 0x8c, 0x4f, 0x7a, 0xb4, 0xcd, 0xac, 0x27, 0x21,
	0x99, 0xb4, 0xb9, 0x24, 0x36, 0x9a, 0xd6, 0x13, 0xa9, 0xbe, 0x34, 0x72, 0xef, 0x90, 0xba, 0x51,
	0x79, 0x24, 0xbc, 0x25, 0x36, 0xb4, 0x9f, 0x10, 0x5c, 0x1e, 0x9b, 0x92, 0xf9, 0x9e, 0xcb, 0x28,
	0xfe, 0x18, 0x56, 0x07, 0xaf, 0x5e, 0x58, 0x72, 0x96, 0x47, 0x52, 0xdf, 0x39, 0x0a, 0x15, 0x06,
	0x35, 0xcf, 0x5b, 0xf1, 0x24, 0xf8, 0x06, 0x9c, 0x77, 0xe9, 0x67, 0xbc, 0x3d, 0x44, 0x30, 0xd4,
	0xf2, 0x9c, 0xd8, 0xde, 0x1d, 0x90, 0xd4, 0x21, 0x5f, 0xa7, 0x09, 0x8a, 0x4a, 0x94, 0x31, 0x97,
	0x55, 0x2b, 0x01, 0x1e, 0xc2, 0x9f, 0x86, 0xfc, 0x01, 0x41, 0xae, 0x1a, 0x50, 0xc2, 0x69, 0x12,
	0x3d, 0x49, 0xec, 0xab, 0x90, 0x1d, 0x08, 0x62, 0x75, 0x23, 0xbe, 0xa0, 0xb6, 0x1a, 0x5d, 0xbc,
	0x0d, 0x4b, 0x6a, 0x25, 0xe5, 0xce, 0x56, 0xca, 0xb3, 0x2b, 0x65, 0x0e, 0x7c, 0xb5, 0x67, 0x08,
	0x2e, 0x0e, 0x57, 0xe6, 0x65, 0x5e, 0x03, 0x11, 0xf3, 0xc0, 0xb2, 0x39, 0x0d, 0xf2, 0x67, 0xc2,
	0x98, 0xe1, 0x4a, 0xfb, 0x0a, 0x41, 0x2e, 0x41, 0x22, 0xba, 0x18, 0x1f, 0x40, 0x46, 0x51, 0x55,
	0x37, 0x62, 0x9e, 0x73, 0x9e, 0x38, 0xcf, 0x7c, 0x0b, 0x9e, 0x23, 0xc8, 0x3d, 0xf2, 0xbb, 0x63,
	0x6a, 0x35, 0x2c, 0x39, 0x7a, 0x71, 0xc9, 0xf1, 0x3d, 0x80, 0x93, 0x1e, 0x2c, 0x49, 0x64, 0x2b,
	0x05, 0x15, 0x49, 0x35, 0x61, 0x7d, 0x5b, 0x40, 0x1e, 0x12, 0x76, 0x68, 0x66, 0x0e, 0xd4, 0xa3,
	0xf6, 0x26, 0xe4, 0xb6, 0xa8, 0x4d, 0x47, 0xb9, 0x8d, 0xbb, 0x75, 0xdf, 0xa4, 0xe0, 0x52, 0xfc,
	0xd6, 0x3d, 0xa4, 0x9c, 0x74, 0x09, 0x27, 0xff, 0xe6, 0x51, 0x18, 0x27, 0x01, 0x6f, 0x8b, 0x91,
	0x31, 0xf1, 0x28, 0x2d, 0x35, 0x4f, 0xcc, 0x8c, 0x44, 0x8b, 0x35, 0x7e, 0x0f, 0xb2, 0x1d, 0x11,
	0xc3, 0x0e, 0x7d, 0x17, 0xa6, 0xfa, 0x42, 0x08, 0x97, 0xce, 0xb7, 0x60, 0x89, 0xba, 0xdd, 0xd0,
	0xf3, 0xcc, 0x54, 0xcf, 0xb3, 0xd4, 0xed, 0x8a, 0x95, 0x54, 0x24, 0x5e, 0xdb, 0xff, 0xb8, 0x22,
	0x95, 0xbf, 0x97, 0xe1, 0x9c, 0x3a, 0xc5, 0xa6, 0x38, 0x1f, 0xfe, 0x1d, 0xc1, 0xff, 0xc7, 0xb4,
	0x6a, 0xfc, 0xfe, 0x74, 0x39, 0x26, 0x0f, 0x95, 0xc2, 0xc6, 0x0b, 0x7a, 0x87, 0x6d, 0x40, 0x33,
	0x9e, 0xfd, 0xf9, 0xd7, 0xb7, 0xa9, 0x9b, 0xf8, 0x0d, 0xf1, 0x81, 0xf4, 0x79, 0xd8, 0x89, 0x36,
	0xfc, 0xc0, 0xfb, 0x84, 0x76, 0x38, 0x33, 0xca, 0x4f, 0x8d, 0x64, 0xcf, 0xff, 0x0d, 0xc1, 0x85,
	0x91, 0x66, 0x8e, 0xd7, 0xa7, 0xb3, 0x98, 0x34, 0x01, 0x0a, 0x73, 0x0f, 0xa2, 0x04, 0x69, 0xf1,
	0x4a, 0x0e, 0x51, 0x4e, 0x32, 0x36, 0xca, 0x4f, 0xf1, 0x2f, 0x08, 0xce, 0xc5, 0xda, 0x20, 0xbe,
	0x3d, 0x9f, 0x6c, 0x03, 0xb9, 0xef, 0xcc, 0xed, 0x17, 0x09, 0x7d, 0x53, 0x72, 0xbe, 0x8e, 0xaf,
	0x4d, 0x13, 0x9a, 0xe1, 0xe7, 0x08, 0xb2, 0x43, 0x6a, 0xe1, 0x77, 0xe7, 0x12, 0x57, 0x31, 0x9d,
	0xe3, 0x2d, 0x4b, 0x90, 0x9b, 0x24, 0xa8, 0x94, 0xf2, 0x7b, 0x04, 0x2b, 0xf1, 0xde, 0x87, 0x67,
	0xd0, 0x64, 0xec, 0x8c, 0x2e, 0x5c, 0x51, 0x8e, 0x43, 0x1f, 0xe1, 0xfa, 0x87, 0xea, 0x23, 0x5c,
	0x7b, 0x4b, 0xb2, 0xba, 0xa1, 0x4d, 0x97, 0x6c, 0x1d, 0x95, 0xf1, 0x8f, 0x08, 0x56, 0xe2, 0x2d,
	0x68, 0x16, 0x62, 0x63, 0x07, 0xd2, 0x34, 0x62, 0xb7, 0x24, 0x31, 0xa3, 0x52, 0x96, 0xc4, 0x06,
	0xe1, 0x4e, 0xd3, 0x4d, 0x30, 0xfc, 0x1a, 0xc1, 0x4a, 0x7c, 0xc8, 0xcc, 0xc2, 0x70, 0xec, 0x58,
	0x2a, 0x5c, 0x1a, 0xe9, 0x41, 0x35, 0xf1, 0x6b, 0xa2, 0x2a, 0x59, 0x9e, 0xa1, 0x92, 0x5f, 0x22,
	0x58, 0x6e, 0x52, 0xde, 0x20, 0xce, 0xae, 0xfc, 0x31, 0xc2, 0x9a, 0x8a, 0x69, 0x11, 0x47, 0x64,
	0x1e, 0x36, 0xaa, 0xbc, 0xb9, 0x04, 0x26, 0xb4, 0x6a, 0x1b, 0x32, 0xed, 0x1d, 0xad, 0x22, 0xd3,
	0x06, 0x94, 0x79, 0xfd, 0xa0, 0x33, 0x59, 0x0c, 0x36, 0x14, 0x59, 0x28, 0x23, 0xa8, 0xd4, 0x4f,
	0xa3, 0x52, 0x7f, 0x69, 0x54, 0x7a, 0x09, 0x2a, 0xbf, 0x22, 0xc0, 0x2d, 0xca, 0xe4, 0x26, 0x0d,
	0x1c, 0x8b, 0x31, 0xf1, 0x3f, 0x88, 0x4b, 0x89, 0x64, 0xa3, 0x10, 0x45, 0xeb, 0xe6, 0x0c, 0xc8,
	0xa8, 0x27, 0x54, 0x25, 0xd5, 0x0d, 0xed, 0xee, 0x6c, 0x54, 0xf9, 0x48, 0xa4, 0x75, 0x54, 0xbe,
	0xff, 0x1d, 0x82, 0xd7, 0x3b, 0x9e, 0x33, 0xf5, 0x22, 0xdd, 0x7f, 0xa5, 0x19, 0x9a, 0x62, 0x53,
	0x69, 0x57, 0x5c, 0x9f, 0x5d, 0xf4, 0x73, 0xea, 0x46, 0x3d, 0xf4, 0xaf, 0xda, 0x5e, 0xbf, 0xab,
	0x47, 0x50, 0x5d, 0x62, 0x4e, 0xfe, 0xb3, 0x1e, 0xaf, 0xfd, 0xa1, 0x80, 0x7b, 0x12, 0xb8, 0x17,
	0x01, 0xf7, 0x24, 0x70, 0x4f, 0x01, 0xf7, 0x1e, 0xaf, 0xed, 0x2f, 0xca, 0x9b, 0xf9, 0xce, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x43, 0x95, 0x56, 0xec, 0x24, 0x10, 0x00, 0x00,
}
