// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/cloud/ml/v1/model_service.proto

package google_cloud_ml_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import _ "go.pedge.io/pb/gogo/google/api"
import _ "go.pedge.io/pb/gogo/google/longrunning"
import google_protobuf2 "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents a machine learning solution.
//
// A model can have multiple versions, each of which is a deployed, trained
// model ready to receive prediction requests. The model itself is just a
// container.
type Model struct {
	// Required. The name specified for the model when it was created.
	//
	// The model name must be unique within the project it is created in.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The description specified for the model when it was created.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. The default version of the model. This version will be used to
	// handle prediction requests that do not specify a version.
	//
	// You can change the default version by calling
	// [projects.methods.versions.setDefault](/ml/reference/rest/v1/projects.models.versions/setDefault).
	DefaultVersion *Version `protobuf:"bytes,3,opt,name=default_version,json=defaultVersion" json:"default_version,omitempty"`
	// Optional. The list of regions where the model is going to be deployed.
	// Currently only one region per model is supported.
	// Defaults to 'us-central1' if nothing is set.
	Regions []string `protobuf:"bytes,4,rep,name=regions" json:"regions,omitempty"`
	// Optional. If true, enables StackDriver Logging for online prediction.
	// Default is false.
	OnlinePredictionLogging bool `protobuf:"varint,5,opt,name=online_prediction_logging,json=onlinePredictionLogging,proto3" json:"online_prediction_logging,omitempty"`
}

func (m *Model) Reset()                    { *m = Model{} }
func (m *Model) String() string            { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()               {}
func (*Model) Descriptor() ([]byte, []int) { return fileDescriptorModelService, []int{0} }

func (m *Model) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Model) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Model) GetDefaultVersion() *Version {
	if m != nil {
		return m.DefaultVersion
	}
	return nil
}

func (m *Model) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *Model) GetOnlinePredictionLogging() bool {
	if m != nil {
		return m.OnlinePredictionLogging
	}
	return false
}

// Represents a version of the model.
//
// Each version is a trained model deployed in the cloud, ready to handle
// prediction requests. A model can have multiple versions. You can get
// information about all of the versions of a given model by calling
// [projects.models.versions.list](/ml/reference/rest/v1/projects.models.versions/list).
type Version struct {
	// Required.The name specified for the version when it was created.
	//
	// The version name must be unique within the model it is created in.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The description specified for the version when it was created.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. If true, this version will be used to handle prediction
	// requests that do not specify a version.
	//
	// You can change the default version by calling
	// [projects.methods.versions.setDefault](/ml/reference/rest/v1/projects.models.versions/setDefault).
	IsDefault bool `protobuf:"varint,3,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
	// Required. The Google Cloud Storage location of the trained model used to
	// create the version. See the
	// [overview of model deployment](/ml/docs/concepts/deployment-overview) for
	// more informaiton.
	//
	// When passing Version to
	// [projects.models.versions.create](/ml/reference/rest/v1/projects.models.versions/create)
	// the model service uses the specified location as the source of the model.
	// Once deployed, the model version is hosted by the prediction service, so
	// this location is useful only as a historical record.
	DeploymentUri string `protobuf:"bytes,4,opt,name=deployment_uri,json=deploymentUri,proto3" json:"deployment_uri,omitempty"`
	// Output only. The time the version was created.
	CreateTime *google_protobuf2.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	// Output only. The time the version was last used for prediction.
	LastUseTime *google_protobuf2.Timestamp `protobuf:"bytes,6,opt,name=last_use_time,json=lastUseTime" json:"last_use_time,omitempty"`
	// Optional. The Google Cloud ML runtime version to use for this deployment.
	// If not set, Google Cloud ML will choose a version.
	RuntimeVersion string `protobuf:"bytes,8,opt,name=runtime_version,json=runtimeVersion,proto3" json:"runtime_version,omitempty"`
	// Optional. Manually select the number of nodes to use for serving the
	// model. If unset (i.e., by default), the number of nodes used to serve
	// the model automatically scales with traffic. However, care should be
	// taken to ramp up traffic according to the model's ability to scale. If
	// your model needs to handle bursts of traffic beyond it's ability to
	// scale, it is recommended you set this field appropriately.
	ManualScaling *ManualScaling `protobuf:"bytes,9,opt,name=manual_scaling,json=manualScaling" json:"manual_scaling,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptorModelService, []int{1} }

func (m *Version) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Version) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Version) GetIsDefault() bool {
	if m != nil {
		return m.IsDefault
	}
	return false
}

func (m *Version) GetDeploymentUri() string {
	if m != nil {
		return m.DeploymentUri
	}
	return ""
}

func (m *Version) GetCreateTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Version) GetLastUseTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.LastUseTime
	}
	return nil
}

func (m *Version) GetRuntimeVersion() string {
	if m != nil {
		return m.RuntimeVersion
	}
	return ""
}

func (m *Version) GetManualScaling() *ManualScaling {
	if m != nil {
		return m.ManualScaling
	}
	return nil
}

// Options for manually scaling a model.
type ManualScaling struct {
	// The number of nodes to allocate for this model. These nodes are always up,
	// starting from the time the model is deployed, so the cost of operating
	// this model will be proportional to nodes * number of hours since
	// deployment.
	Nodes int32 `protobuf:"varint,1,opt,name=nodes,proto3" json:"nodes,omitempty"`
}

func (m *ManualScaling) Reset()                    { *m = ManualScaling{} }
func (m *ManualScaling) String() string            { return proto.CompactTextString(m) }
func (*ManualScaling) ProtoMessage()               {}
func (*ManualScaling) Descriptor() ([]byte, []int) { return fileDescriptorModelService, []int{2} }

func (m *ManualScaling) GetNodes() int32 {
	if m != nil {
		return m.Nodes
	}
	return 0
}

// Request message for the CreateModel method.
type CreateModelRequest struct {
	// Required. The project name.
	//
	// Authorization: requires `Editor` role on the specified project.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The model to create.
	Model *Model `protobuf:"bytes,2,opt,name=model" json:"model,omitempty"`
}

func (m *CreateModelRequest) Reset()                    { *m = CreateModelRequest{} }
func (m *CreateModelRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateModelRequest) ProtoMessage()               {}
func (*CreateModelRequest) Descriptor() ([]byte, []int) { return fileDescriptorModelService, []int{3} }

func (m *CreateModelRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateModelRequest) GetModel() *Model {
	if m != nil {
		return m.Model
	}
	return nil
}

// Request message for the ListModels method.
type ListModelsRequest struct {
	// Required. The name of the project whose models are to be listed.
	//
	// Authorization: requires `Viewer` role on the specified project.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. A page token to request the next page of results.
	//
	// You get the token from the `next_page_token` field of the response from
	// the previous call.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. The number of models to retrieve per "page" of results. If there
	// are more remaining results than this number, the response message will
	// contain a valid value in the `next_page_token` field.
	//
	// The default value is 20, and the maximum page size is 100.
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (m *ListModelsRequest) Reset()                    { *m = ListModelsRequest{} }
func (m *ListModelsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListModelsRequest) ProtoMessage()               {}
func (*ListModelsRequest) Descriptor() ([]byte, []int) { return fileDescriptorModelService, []int{4} }

func (m *ListModelsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListModelsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListModelsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// Response message for the ListModels method.
type ListModelsResponse struct {
	// The list of models.
	Models []*Model `protobuf:"bytes,1,rep,name=models" json:"models,omitempty"`
	// Optional. Pass this token as the `page_token` field of the request for a
	// subsequent call.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *ListModelsResponse) Reset()                    { *m = ListModelsResponse{} }
func (m *ListModelsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListModelsResponse) ProtoMessage()               {}
func (*ListModelsResponse) Descriptor() ([]byte, []int) { return fileDescriptorModelService, []int{5} }

func (m *ListModelsResponse) GetModels() []*Model {
	if m != nil {
		return m.Models
	}
	return nil
}

func (m *ListModelsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request message for the GetModel method.
type GetModelRequest struct {
	// Required. The name of the model.
	//
	// Authorization: requires `Viewer` role on the parent project.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetModelRequest) Reset()                    { *m = GetModelRequest{} }
func (m *GetModelRequest) String() string            { return proto.CompactTextString(m) }
func (*GetModelRequest) ProtoMessage()               {}
func (*GetModelRequest) Descriptor() ([]byte, []int) { return fileDescriptorModelService, []int{6} }

func (m *GetModelRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for the DeleteModel method.
type DeleteModelRequest struct {
	// Required. The name of the model.
	//
	// Authorization: requires `Editor` role on the parent project.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DeleteModelRequest) Reset()                    { *m = DeleteModelRequest{} }
func (m *DeleteModelRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteModelRequest) ProtoMessage()               {}
func (*DeleteModelRequest) Descriptor() ([]byte, []int) { return fileDescriptorModelService, []int{7} }

func (m *DeleteModelRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Uploads the provided trained model version to Cloud Machine Learning.
type CreateVersionRequest struct {
	// Required. The name of the model.
	//
	// Authorization: requires `Editor` role on the parent project.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The version details.
	Version *Version `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *CreateVersionRequest) Reset()                    { *m = CreateVersionRequest{} }
func (m *CreateVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateVersionRequest) ProtoMessage()               {}
func (*CreateVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptorModelService, []int{8} }

func (m *CreateVersionRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateVersionRequest) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

// Request message for the ListVersions method.
type ListVersionsRequest struct {
	// Required. The name of the model for which to list the version.
	//
	// Authorization: requires `Viewer` role on the parent project.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. A page token to request the next page of results.
	//
	// You get the token from the `next_page_token` field of the response from
	// the previous call.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. The number of versions to retrieve per "page" of results. If
	// there are more remaining results than this number, the response message
	// will contain a valid value in the `next_page_token` field.
	//
	// The default value is 20, and the maximum page size is 100.
	PageSize int32 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (m *ListVersionsRequest) Reset()                    { *m = ListVersionsRequest{} }
func (m *ListVersionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListVersionsRequest) ProtoMessage()               {}
func (*ListVersionsRequest) Descriptor() ([]byte, []int) { return fileDescriptorModelService, []int{9} }

func (m *ListVersionsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListVersionsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListVersionsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// Response message for the ListVersions method.
type ListVersionsResponse struct {
	// The list of versions.
	Versions []*Version `protobuf:"bytes,1,rep,name=versions" json:"versions,omitempty"`
	// Optional. Pass this token as the `page_token` field of the request for a
	// subsequent call.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *ListVersionsResponse) Reset()         { *m = ListVersionsResponse{} }
func (m *ListVersionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListVersionsResponse) ProtoMessage()    {}
func (*ListVersionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorModelService, []int{10}
}

func (m *ListVersionsResponse) GetVersions() []*Version {
	if m != nil {
		return m.Versions
	}
	return nil
}

func (m *ListVersionsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request message for the GetVersion method.
type GetVersionRequest struct {
	// Required. The name of the version.
	//
	// Authorization: requires `Viewer` role on the parent project.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetVersionRequest) Reset()                    { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()               {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptorModelService, []int{11} }

func (m *GetVersionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for the DeleteVerionRequest method.
type DeleteVersionRequest struct {
	// Required. The name of the version. You can get the names of all the
	// versions of a model by calling
	// [projects.models.versions.list](/ml/reference/rest/v1/projects.models.versions/list).
	//
	// Authorization: requires `Editor` role on the parent project.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DeleteVersionRequest) Reset()         { *m = DeleteVersionRequest{} }
func (m *DeleteVersionRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteVersionRequest) ProtoMessage()    {}
func (*DeleteVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorModelService, []int{12}
}

func (m *DeleteVersionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for the SetDefaultVersion request.
type SetDefaultVersionRequest struct {
	// Required. The name of the version to make the default for the model. You
	// can get the names of all the versions of a model by calling
	// [projects.models.versions.list](/ml/reference/rest/v1/projects.models.versions/list).
	//
	// Authorization: requires `Editor` role on the parent project.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *SetDefaultVersionRequest) Reset()         { *m = SetDefaultVersionRequest{} }
func (m *SetDefaultVersionRequest) String() string { return proto.CompactTextString(m) }
func (*SetDefaultVersionRequest) ProtoMessage()    {}
func (*SetDefaultVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorModelService, []int{13}
}

func (m *SetDefaultVersionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Model)(nil), "google.cloud.ml.v1.Model")
	proto.RegisterType((*Version)(nil), "google.cloud.ml.v1.Version")
	proto.RegisterType((*ManualScaling)(nil), "google.cloud.ml.v1.ManualScaling")
	proto.RegisterType((*CreateModelRequest)(nil), "google.cloud.ml.v1.CreateModelRequest")
	proto.RegisterType((*ListModelsRequest)(nil), "google.cloud.ml.v1.ListModelsRequest")
	proto.RegisterType((*ListModelsResponse)(nil), "google.cloud.ml.v1.ListModelsResponse")
	proto.RegisterType((*GetModelRequest)(nil), "google.cloud.ml.v1.GetModelRequest")
	proto.RegisterType((*DeleteModelRequest)(nil), "google.cloud.ml.v1.DeleteModelRequest")
	proto.RegisterType((*CreateVersionRequest)(nil), "google.cloud.ml.v1.CreateVersionRequest")
	proto.RegisterType((*ListVersionsRequest)(nil), "google.cloud.ml.v1.ListVersionsRequest")
	proto.RegisterType((*ListVersionsResponse)(nil), "google.cloud.ml.v1.ListVersionsResponse")
	proto.RegisterType((*GetVersionRequest)(nil), "google.cloud.ml.v1.GetVersionRequest")
	proto.RegisterType((*DeleteVersionRequest)(nil), "google.cloud.ml.v1.DeleteVersionRequest")
	proto.RegisterType((*SetDefaultVersionRequest)(nil), "google.cloud.ml.v1.SetDefaultVersionRequest")
}

func init() { proto.RegisterFile("google/cloud/ml/v1/model_service.proto", fileDescriptorModelService) }

var fileDescriptorModelService = []byte{
	// 969 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0x96, 0xdb, 0xa6, 0x4d, 0x5e, 0x36, 0xad, 0x3a, 0x14, 0xc8, 0x66, 0x29, 0x04, 0xaf, 0x9a,
	0x86, 0x00, 0xb6, 0x52, 0x16, 0x21, 0x82, 0x00, 0x69, 0xa9, 0xb4, 0x1c, 0x76, 0x45, 0xe5, 0xee,
	0x72, 0x43, 0x96, 0x37, 0x99, 0x35, 0x03, 0xf6, 0x8c, 0xf1, 0x8c, 0x03, 0x2c, 0xac, 0x90, 0xe0,
	0xc8, 0x11, 0xee, 0xfc, 0x51, 0x9c, 0xb8, 0x73, 0xe3, 0xcc, 0x1d, 0xcd, 0x0f, 0xa7, 0x76, 0xe2,
	0xc4, 0x05, 0x89, 0x5b, 0xe6, 0xf9, 0x9b, 0x79, 0xdf, 0xbc, 0xef, 0x7d, 0x6f, 0x02, 0x83, 0x90,
	0xb1, 0x30, 0xc2, 0xee, 0x34, 0x62, 0xd9, 0xcc, 0x8d, 0x23, 0x77, 0x3e, 0x76, 0x63, 0x36, 0xc3,
	0x91, 0xcf, 0x71, 0x3a, 0x27, 0x53, 0xec, 0x24, 0x29, 0x13, 0x0c, 0x21, 0x8d, 0x73, 0x14, 0xce,
	0x89, 0x23, 0x67, 0x3e, 0xee, 0xbd, 0x64, 0xf6, 0x06, 0x09, 0x71, 0x03, 0x4a, 0x99, 0x08, 0x04,
	0x61, 0x94, 0xeb, 0x1d, 0xbd, 0xe7, 0x8b, 0x5f, 0x33, 0xf1, 0xb9, 0x09, 0xdf, 0x36, 0xe1, 0x88,
	0xd1, 0x30, 0xcd, 0x28, 0x25, 0x34, 0x74, 0x59, 0x82, 0xd3, 0xd2, 0xde, 0x57, 0x0c, 0x48, 0xad,
	0x1e, 0x67, 0x4f, 0x5c, 0x41, 0x62, 0xcc, 0x45, 0x10, 0x27, 0x1a, 0x60, 0xff, 0x61, 0x41, 0xe3,
	0x81, 0xa4, 0x89, 0x10, 0xec, 0xd0, 0x20, 0xc6, 0x5d, 0xab, 0x6f, 0x0d, 0x5b, 0x9e, 0xfa, 0x8d,
	0xfa, 0xd0, 0x9e, 0x61, 0x3e, 0x4d, 0x49, 0x22, 0x0f, 0xed, 0x6e, 0xa9, 0x4f, 0xc5, 0x10, 0x3a,
	0x87, 0x83, 0x19, 0x7e, 0x12, 0x64, 0x91, 0xf0, 0xe7, 0x38, 0xe5, 0x12, 0xb5, 0xdd, 0xb7, 0x86,
	0xed, 0xb3, 0x5b, 0xce, 0xea, 0x45, 0x9d, 0x4f, 0x35, 0xc4, 0xdb, 0x37, 0x7b, 0xcc, 0x1a, 0x75,
	0x61, 0x2f, 0xc5, 0xa1, 0xe4, 0xdd, 0xdd, 0xe9, 0x6f, 0x0f, 0x5b, 0x5e, 0xbe, 0x44, 0x13, 0xb8,
	0xc9, 0x68, 0x44, 0x28, 0xf6, 0x93, 0x14, 0xcf, 0xc8, 0x54, 0x26, 0xf5, 0x23, 0x16, 0x86, 0x84,
	0x86, 0xdd, 0x46, 0xdf, 0x1a, 0x36, 0xbd, 0x17, 0x35, 0xe0, 0x62, 0xf1, 0xfd, 0xbe, 0xfe, 0x6c,
	0xff, 0xbd, 0x05, 0x7b, 0x79, 0x86, 0xff, 0x76, 0xbb, 0x63, 0x00, 0xc2, 0x7d, 0x43, 0x56, 0x5d,
	0xac, 0xe9, 0xb5, 0x08, 0x3f, 0xd7, 0x01, 0x74, 0x02, 0xfb, 0x33, 0x9c, 0x44, 0xec, 0xdb, 0x18,
	0x53, 0xe1, 0x67, 0x29, 0xe9, 0xee, 0xa8, 0x33, 0x3a, 0x57, 0xd1, 0x47, 0x29, 0x41, 0xef, 0x41,
	0x7b, 0x9a, 0xe2, 0x40, 0x60, 0x5f, 0x56, 0x5f, 0xb1, 0x6e, 0x9f, 0xf5, 0xf2, 0xfa, 0xe4, 0xd2,
	0x38, 0x0f, 0x73, 0x69, 0x3c, 0xd0, 0x70, 0x19, 0x40, 0x1f, 0x40, 0x27, 0x0a, 0xb8, 0xf0, 0x33,
	0x6e, 0xb6, 0xef, 0xd6, 0x6e, 0x6f, 0xcb, 0x0d, 0x8f, 0xb8, 0xde, 0x7f, 0x0a, 0x07, 0x69, 0x46,
	0xe5, 0xce, 0x85, 0x40, 0x4d, 0x45, 0x72, 0xdf, 0x84, 0xf3, 0x0a, 0x7d, 0x0c, 0xfb, 0x71, 0x40,
	0xb3, 0x20, 0xf2, 0xf9, 0x34, 0x88, 0x64, 0x79, 0x5b, 0x2a, 0xd3, 0xab, 0x55, 0x42, 0x3e, 0x50,
	0xc8, 0x4b, 0x0d, 0xf4, 0x3a, 0x71, 0x71, 0x69, 0x9f, 0x40, 0xa7, 0xf4, 0x1d, 0x1d, 0x41, 0x83,
	0xb2, 0x19, 0xe6, 0xaa, 0xfa, 0x0d, 0x4f, 0x2f, 0xec, 0xcf, 0x00, 0x7d, 0xa4, 0xee, 0xa9, 0xfa,
	0xcf, 0xc3, 0x5f, 0x65, 0x98, 0x0b, 0xf4, 0x02, 0xec, 0x26, 0x41, 0x8a, 0xa9, 0x30, 0x52, 0x99,
	0x15, 0x72, 0xa1, 0xa1, 0xec, 0xa4, 0x64, 0x6a, 0x9f, 0xdd, 0xac, 0x64, 0xa5, 0x0e, 0xd2, 0x38,
	0x3b, 0x84, 0xc3, 0xfb, 0x84, 0x0b, 0x15, 0xe3, 0x75, 0xa7, 0x1f, 0x03, 0x24, 0x41, 0x88, 0x7d,
	0xc1, 0xbe, 0xc4, 0xd4, 0xa8, 0xd8, 0x92, 0x91, 0x87, 0x32, 0x80, 0x6e, 0x81, 0x5a, 0xf8, 0x9c,
	0x3c, 0xd5, 0xfa, 0x35, 0xbc, 0xa6, 0x0c, 0x5c, 0x92, 0xa7, 0xd8, 0x66, 0x80, 0x8a, 0x89, 0x78,
	0xc2, 0x28, 0xc7, 0x68, 0x0c, 0xbb, 0x8a, 0x87, 0xbc, 0xf4, 0xf6, 0x66, 0xc2, 0x06, 0x88, 0x06,
	0x70, 0x40, 0xf1, 0x37, 0xc2, 0x2f, 0x30, 0xd1, 0x3d, 0xd9, 0x91, 0xe1, 0x8b, 0x9c, 0x8d, 0x7d,
	0x02, 0x07, 0xf7, 0xb0, 0x28, 0x55, 0xad, 0xa2, 0xbd, 0xed, 0x21, 0xa0, 0x73, 0x1c, 0xe1, 0xa5,
	0xfa, 0x56, 0x21, 0x31, 0x1c, 0x69, 0x25, 0x72, 0x7f, 0xd6, 0x54, 0xeb, 0x6d, 0xd8, 0xcb, 0x7b,
	0x69, 0xab, 0xde, 0xec, 0x39, 0xd6, 0x26, 0xf0, 0x9c, 0x2c, 0x94, 0x89, 0xff, 0xaf, 0x9a, 0x7c,
	0x0d, 0x47, 0xe5, 0x54, 0x46, 0x95, 0x77, 0xa0, 0x69, 0xd8, 0xe4, 0xba, 0x6c, 0xa4, 0xbe, 0x00,
	0x5f, 0x5b, 0x9b, 0x53, 0x38, 0xbc, 0x87, 0xc5, 0x52, 0x1d, 0xab, 0x6a, 0x3e, 0x82, 0x23, 0xad,
	0xce, 0x35, 0xb0, 0x0e, 0x74, 0x2f, 0xb1, 0x38, 0x2f, 0xcd, 0xcc, 0x0d, 0xf8, 0xb3, 0xbf, 0x5a,
	0x70, 0x43, 0x89, 0x7e, 0xa9, 0x9f, 0x1e, 0xf4, 0x03, 0xb4, 0x0b, 0x56, 0x43, 0x83, 0xaa, 0x3b,
	0xaf, 0x7a, 0xb1, 0xb7, 0xbe, 0x67, 0xed, 0x37, 0x7f, 0xfc, 0xfd, 0xcf, 0x5f, 0xb6, 0x4e, 0xed,
	0x97, 0xe5, 0x3b, 0xf7, 0x9d, 0x56, 0xec, 0xfd, 0x24, 0x65, 0x5f, 0xe0, 0xa9, 0xe0, 0xee, 0xe8,
	0x99, 0x7e, 0xfb, 0xf8, 0x44, 0x9b, 0x11, 0xfd, 0x64, 0x01, 0x5c, 0x99, 0x04, 0x9d, 0x54, 0x1d,
	0xbc, 0xe2, 0xd6, 0xde, 0xa0, 0x0e, 0xa6, 0x55, 0xb5, 0x07, 0x8a, 0x4c, 0x1f, 0xd5, 0x90, 0x41,
	0x29, 0x34, 0x73, 0xe3, 0xa0, 0xdb, 0x55, 0x67, 0x2f, 0xd9, 0x6a, 0x53, 0x01, 0xca, 0x39, 0x65,
	0xd9, 0x0b, 0x19, 0x4d, 0x42, 0x77, 0xf4, 0x0c, 0x7d, 0x0f, 0xed, 0x82, 0x0b, 0xab, 0x4b, 0xbf,
	0x6a, 0xd3, 0xde, 0x71, 0x8e, 0x2b, 0x3c, 0xef, 0xce, 0x27, 0xf9, 0xf3, 0x9e, 0x67, 0x1f, 0xd5,
	0x65, 0xff, 0xd5, 0x82, 0x4e, 0xc9, 0xda, 0x68, 0xb8, 0x5e, 0xfb, 0x72, 0x67, 0xd5, 0x51, 0x98,
	0x28, 0x0a, 0x77, 0xec, 0xd7, 0xaa, 0x8b, 0x7e, 0x45, 0xc2, 0xcd, 0x4d, 0x34, 0xc9, 0x27, 0x81,
	0xa4, 0x75, 0xa3, 0xe8, 0x4f, 0x74, 0xba, 0x4e, 0xe9, 0xa5, 0x61, 0xd1, 0x1b, 0xd6, 0x03, 0x4d,
	0x53, 0x8c, 0x15, 0xbf, 0xd7, 0xd1, 0xf5, 0xf9, 0xa9, 0x2e, 0xbd, 0x72, 0x6f, 0x75, 0x97, 0xae,
	0xb8, 0xbb, 0xb7, 0x69, 0x82, 0x2c, 0xb1, 0x58, 0x27, 0xd4, 0x82, 0x82, 0xd4, 0xec, 0x67, 0x0b,
	0x3a, 0xa5, 0xd1, 0x50, 0xad, 0x59, 0xd5, 0xf4, 0xa8, 0xd3, 0xcc, 0xb0, 0x19, 0xfd, 0x0b, 0x36,
	0xbf, 0x59, 0x70, 0xb8, 0x32, 0x7c, 0xd0, 0x1b, 0x55, 0x8c, 0xd6, 0xcd, 0xa8, 0xcd, 0x15, 0xfa,
	0x50, 0x71, 0x7a, 0xd7, 0xbe, 0x73, 0x6d, 0x4e, 0x13, 0xbe, 0x48, 0x34, 0xb1, 0x46, 0x77, 0xc7,
	0xd0, 0x9b, 0xb2, 0x78, 0x25, 0x45, 0x90, 0x10, 0x67, 0x3e, 0xbe, 0x7b, 0x58, 0x1c, 0x84, 0x17,
	0xf2, 0xcf, 0xd2, 0x85, 0xf5, 0x78, 0x57, 0xfd, 0x6b, 0x7a, 0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x23, 0xdf, 0x93, 0xdd, 0xb5, 0x0b, 0x00, 0x00,
}
