// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/streetview/publish/v1/rpcmessages.proto

package google_streetview_publish_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "go.pedge.io/pb/go/google/protobuf"
import google_rpc "go.pedge.io/pb/go/google/rpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Specifies which view of the `Photo` should be included in the response.
type PhotoView int32

const (
	// Server reponses do not include the download URL for the photo bytes.
	// The default value.
	PhotoView_BASIC PhotoView = 0
	// Server responses include the download URL for the photo bytes.
	PhotoView_INCLUDE_DOWNLOAD_URL PhotoView = 1
)

var PhotoView_name = map[int32]string{
	0: "BASIC",
	1: "INCLUDE_DOWNLOAD_URL",
}
var PhotoView_value = map[string]int32{
	"BASIC":                0,
	"INCLUDE_DOWNLOAD_URL": 1,
}

func (x PhotoView) String() string {
	return proto.EnumName(PhotoView_name, int32(x))
}
func (PhotoView) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Request to create a photo.
type CreatePhotoRequest struct {
	// Required. Photo to create.
	Photo *Photo `protobuf:"bytes,1,opt,name=photo" json:"photo,omitempty"`
}

func (m *CreatePhotoRequest) Reset()                    { *m = CreatePhotoRequest{} }
func (m *CreatePhotoRequest) String() string            { return proto.CompactTextString(m) }
func (*CreatePhotoRequest) ProtoMessage()               {}
func (*CreatePhotoRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreatePhotoRequest) GetPhoto() *Photo {
	if m != nil {
		return m.Photo
	}
	return nil
}

// Request to get a photo.
//
// By default
// - does not return the download URL for the photo bytes.
//
// Parameters:
// - 'view' controls if the download URL for the photo bytes will be returned.
type GetPhotoRequest struct {
	// Required. ID of the photo.
	PhotoId string `protobuf:"bytes,1,opt,name=photo_id,json=photoId" json:"photo_id,omitempty"`
	// Specifies if a download URL for the photo bytes should be returned in the
	// Photo response.
	View PhotoView `protobuf:"varint,2,opt,name=view,enum=google.streetview.publish.v1.PhotoView" json:"view,omitempty"`
}

func (m *GetPhotoRequest) Reset()                    { *m = GetPhotoRequest{} }
func (m *GetPhotoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPhotoRequest) ProtoMessage()               {}
func (*GetPhotoRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GetPhotoRequest) GetPhotoId() string {
	if m != nil {
		return m.PhotoId
	}
	return ""
}

func (m *GetPhotoRequest) GetView() PhotoView {
	if m != nil {
		return m.View
	}
	return PhotoView_BASIC
}

// Request to get one or more photos.
// By default
// - does not return the download URL for the photo bytes.
//
// Parameters:
// - 'view' controls if the download URL for the photo bytes will be returned.
type BatchGetPhotosRequest struct {
	// Required. IDs of the photos.
	PhotoIds []string `protobuf:"bytes,1,rep,name=photo_ids,json=photoIds" json:"photo_ids,omitempty"`
	// Specifies if a download URL for the photo bytes should be returned in the
	// Photo response.
	View PhotoView `protobuf:"varint,2,opt,name=view,enum=google.streetview.publish.v1.PhotoView" json:"view,omitempty"`
}

func (m *BatchGetPhotosRequest) Reset()                    { *m = BatchGetPhotosRequest{} }
func (m *BatchGetPhotosRequest) String() string            { return proto.CompactTextString(m) }
func (*BatchGetPhotosRequest) ProtoMessage()               {}
func (*BatchGetPhotosRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *BatchGetPhotosRequest) GetPhotoIds() []string {
	if m != nil {
		return m.PhotoIds
	}
	return nil
}

func (m *BatchGetPhotosRequest) GetView() PhotoView {
	if m != nil {
		return m.View
	}
	return PhotoView_BASIC
}

// Response to batch get of photos.
type BatchGetPhotosResponse struct {
	// List of results for each individual photo requested, in the same order as
	// the request.
	Results []*PhotoResponse `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *BatchGetPhotosResponse) Reset()                    { *m = BatchGetPhotosResponse{} }
func (m *BatchGetPhotosResponse) String() string            { return proto.CompactTextString(m) }
func (*BatchGetPhotosResponse) ProtoMessage()               {}
func (*BatchGetPhotosResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *BatchGetPhotosResponse) GetResults() []*PhotoResponse {
	if m != nil {
		return m.Results
	}
	return nil
}

// Response payload for a single `Photo` in batch operations including
// `BatchGetPhotosRequest` and `BatchUpdatePhotosRequest`.
type PhotoResponse struct {
	// The status for the operation to get or update a single photo in the batch
	// request.
	Status *google_rpc.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	// The photo resource, if the request was successful.
	Photo *Photo `protobuf:"bytes,2,opt,name=photo" json:"photo,omitempty"`
}

func (m *PhotoResponse) Reset()                    { *m = PhotoResponse{} }
func (m *PhotoResponse) String() string            { return proto.CompactTextString(m) }
func (*PhotoResponse) ProtoMessage()               {}
func (*PhotoResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *PhotoResponse) GetStatus() *google_rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PhotoResponse) GetPhoto() *Photo {
	if m != nil {
		return m.Photo
	}
	return nil
}

// Request to list all photos that belong to the user sending the request.
//
// By default
// - does not return the download URL for the photo bytes.
//
// Parameters:
// - 'view' controls if the download URL for the photo bytes will be returned.
// - 'page_size' determines the maximum number of photos to return.
// - 'page_token' is the next page token value returned from a previous List
//     request, if any.
type ListPhotosRequest struct {
	// Specifies if a download URL for the photos bytes should be returned in the
	// Photos response.
	View PhotoView `protobuf:"varint,1,opt,name=view,enum=google.streetview.publish.v1.PhotoView" json:"view,omitempty"`
	// The maximum number of photos to return.
	// `page_size` must be non-negative. If `page_size` is zero or is not
	// provided, the default page size of 100 will be used.
	// The number of photos returned in the response may be less than `page_size`
	// if the number of photos that belong to the user is less than `page_size`.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// The filter expression.
	// Example: `placeId=ChIJj61dQgK6j4AR4GeTYWZsKWw`
	Filter string `protobuf:"bytes,4,opt,name=filter" json:"filter,omitempty"`
}

func (m *ListPhotosRequest) Reset()                    { *m = ListPhotosRequest{} }
func (m *ListPhotosRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPhotosRequest) ProtoMessage()               {}
func (*ListPhotosRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ListPhotosRequest) GetView() PhotoView {
	if m != nil {
		return m.View
	}
	return PhotoView_BASIC
}

func (m *ListPhotosRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListPhotosRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListPhotosRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// Response to list all photos that belong to a user.
type ListPhotosResponse struct {
	// List of photos. There will be a maximum number of items returned based on
	// the page_size field in the request.
	Photos []*Photo `protobuf:"bytes,1,rep,name=photos" json:"photos,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListPhotosResponse) Reset()                    { *m = ListPhotosResponse{} }
func (m *ListPhotosResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPhotosResponse) ProtoMessage()               {}
func (*ListPhotosResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ListPhotosResponse) GetPhotos() []*Photo {
	if m != nil {
		return m.Photos
	}
	return nil
}

func (m *ListPhotosResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request to update the metadata of a photo.
// Updating the pixels of a photo is not supported.
type UpdatePhotoRequest struct {
	// Required. Photo object containing the new metadata. Only the fields
	// specified in `update_mask` are used. If `update_mask` is not present, the
	// update applies to all fields.
	// **Note:** To update `pose.altitude`, `pose.latlngpair` has to be filled as
	// well. Otherwise, the request will fail.
	Photo *Photo `protobuf:"bytes,1,opt,name=photo" json:"photo,omitempty"`
	// Mask that identifies fields on the photo metadata to update.
	// If not present, the old Photo metadata will be entirely replaced with the
	// new Photo metadata in this request. The update fails if invalid fields are
	// specified. Multiple fields can be specified in a comma-delimited list.
	//
	// The following fields are valid:
	//
	// * `pose.heading`
	// * `pose.latlngpair`
	// * `pose.pitch`
	// * `pose.roll`
	// * `pose.level`
	// * `pose.altitude`
	// * `connections`
	// * `places`
	//
	//
	// **Note:** Repeated fields in `update_mask` mean the entire set of repeated
	// values will be replaced with the new contents. For example, if
	// `UpdatePhotoRequest.photo.update_mask` contains `connections` and
	// `UpdatePhotoRequest.photo.connections` is empty, all connections will be
	// removed.
	UpdateMask *google_protobuf2.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
}

func (m *UpdatePhotoRequest) Reset()                    { *m = UpdatePhotoRequest{} }
func (m *UpdatePhotoRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdatePhotoRequest) ProtoMessage()               {}
func (*UpdatePhotoRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *UpdatePhotoRequest) GetPhoto() *Photo {
	if m != nil {
		return m.Photo
	}
	return nil
}

func (m *UpdatePhotoRequest) GetUpdateMask() *google_protobuf2.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Request to update the metadata of photos.
// Updating the pixels of photos is not supported.
type BatchUpdatePhotosRequest struct {
	// Required. List of update photo requests.
	UpdatePhotoRequests []*UpdatePhotoRequest `protobuf:"bytes,1,rep,name=update_photo_requests,json=updatePhotoRequests" json:"update_photo_requests,omitempty"`
}

func (m *BatchUpdatePhotosRequest) Reset()                    { *m = BatchUpdatePhotosRequest{} }
func (m *BatchUpdatePhotosRequest) String() string            { return proto.CompactTextString(m) }
func (*BatchUpdatePhotosRequest) ProtoMessage()               {}
func (*BatchUpdatePhotosRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *BatchUpdatePhotosRequest) GetUpdatePhotoRequests() []*UpdatePhotoRequest {
	if m != nil {
		return m.UpdatePhotoRequests
	}
	return nil
}

// Response to batch update of metadata of one or more photos.
type BatchUpdatePhotosResponse struct {
	// List of results for each individual photo updated, in the same order as
	// the request.
	Results []*PhotoResponse `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *BatchUpdatePhotosResponse) Reset()                    { *m = BatchUpdatePhotosResponse{} }
func (m *BatchUpdatePhotosResponse) String() string            { return proto.CompactTextString(m) }
func (*BatchUpdatePhotosResponse) ProtoMessage()               {}
func (*BatchUpdatePhotosResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *BatchUpdatePhotosResponse) GetResults() []*PhotoResponse {
	if m != nil {
		return m.Results
	}
	return nil
}

// Request to delete a photo.
type DeletePhotoRequest struct {
	// Required. ID of the photo.
	PhotoId string `protobuf:"bytes,1,opt,name=photo_id,json=photoId" json:"photo_id,omitempty"`
}

func (m *DeletePhotoRequest) Reset()                    { *m = DeletePhotoRequest{} }
func (m *DeletePhotoRequest) String() string            { return proto.CompactTextString(m) }
func (*DeletePhotoRequest) ProtoMessage()               {}
func (*DeletePhotoRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *DeletePhotoRequest) GetPhotoId() string {
	if m != nil {
		return m.PhotoId
	}
	return ""
}

// Request to delete multiple photos.
type BatchDeletePhotosRequest struct {
	// Required. List of delete photo requests.
	PhotoIds []string `protobuf:"bytes,1,rep,name=photo_ids,json=photoIds" json:"photo_ids,omitempty"`
}

func (m *BatchDeletePhotosRequest) Reset()                    { *m = BatchDeletePhotosRequest{} }
func (m *BatchDeletePhotosRequest) String() string            { return proto.CompactTextString(m) }
func (*BatchDeletePhotosRequest) ProtoMessage()               {}
func (*BatchDeletePhotosRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *BatchDeletePhotosRequest) GetPhotoIds() []string {
	if m != nil {
		return m.PhotoIds
	}
	return nil
}

// Response to batch delete of one or more photos.
type BatchDeletePhotosResponse struct {
	// The status for the operation to delete a single photo in the batch request.
	Status []*google_rpc.Status `protobuf:"bytes,1,rep,name=status" json:"status,omitempty"`
}

func (m *BatchDeletePhotosResponse) Reset()                    { *m = BatchDeletePhotosResponse{} }
func (m *BatchDeletePhotosResponse) String() string            { return proto.CompactTextString(m) }
func (*BatchDeletePhotosResponse) ProtoMessage()               {}
func (*BatchDeletePhotosResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *BatchDeletePhotosResponse) GetStatus() []*google_rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*CreatePhotoRequest)(nil), "google.streetview.publish.v1.CreatePhotoRequest")
	proto.RegisterType((*GetPhotoRequest)(nil), "google.streetview.publish.v1.GetPhotoRequest")
	proto.RegisterType((*BatchGetPhotosRequest)(nil), "google.streetview.publish.v1.BatchGetPhotosRequest")
	proto.RegisterType((*BatchGetPhotosResponse)(nil), "google.streetview.publish.v1.BatchGetPhotosResponse")
	proto.RegisterType((*PhotoResponse)(nil), "google.streetview.publish.v1.PhotoResponse")
	proto.RegisterType((*ListPhotosRequest)(nil), "google.streetview.publish.v1.ListPhotosRequest")
	proto.RegisterType((*ListPhotosResponse)(nil), "google.streetview.publish.v1.ListPhotosResponse")
	proto.RegisterType((*UpdatePhotoRequest)(nil), "google.streetview.publish.v1.UpdatePhotoRequest")
	proto.RegisterType((*BatchUpdatePhotosRequest)(nil), "google.streetview.publish.v1.BatchUpdatePhotosRequest")
	proto.RegisterType((*BatchUpdatePhotosResponse)(nil), "google.streetview.publish.v1.BatchUpdatePhotosResponse")
	proto.RegisterType((*DeletePhotoRequest)(nil), "google.streetview.publish.v1.DeletePhotoRequest")
	proto.RegisterType((*BatchDeletePhotosRequest)(nil), "google.streetview.publish.v1.BatchDeletePhotosRequest")
	proto.RegisterType((*BatchDeletePhotosResponse)(nil), "google.streetview.publish.v1.BatchDeletePhotosResponse")
	proto.RegisterEnum("google.streetview.publish.v1.PhotoView", PhotoView_name, PhotoView_value)
}

func init() { proto.RegisterFile("google/streetview/publish/v1/rpcmessages.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5f, 0x73, 0xd2, 0x4e,
	0x14, 0xfd, 0xa5, 0x7f, 0x68, 0xb9, 0x9d, 0xfe, 0x5a, 0x57, 0x5b, 0x53, 0xac, 0x33, 0x4c, 0x9c,
	0x51, 0xa6, 0x3a, 0x49, 0x8b, 0x0f, 0x8e, 0xc3, 0x53, 0x81, 0x5a, 0x99, 0xa1, 0x85, 0x09, 0xa2,
	0x8f, 0x99, 0x10, 0x2e, 0x21, 0x43, 0x20, 0x69, 0x76, 0x43, 0xb5, 0x4f, 0x7e, 0x00, 0x3f, 0x86,
	0x1f, 0xd4, 0xc9, 0x66, 0xb7, 0x05, 0x8a, 0x0c, 0x6a, 0xdf, 0xd8, 0xfb, 0xe7, 0xdc, 0xb3, 0xe7,
	0xee, 0x09, 0xa0, 0xbb, 0x41, 0xe0, 0xfa, 0x68, 0x50, 0x16, 0x21, 0xb2, 0xb1, 0x87, 0xd7, 0x46,
	0x18, 0x77, 0x7c, 0x8f, 0xf6, 0x8d, 0xf1, 0x89, 0x11, 0x85, 0xce, 0x10, 0x29, 0xb5, 0x5d, 0xa4,
	0x7a, 0x18, 0x05, 0x2c, 0x20, 0x87, 0x69, 0xbd, 0x7e, 0x57, 0xaf, 0x8b, 0x7a, 0x7d, 0x7c, 0x92,
	0xcb, 0x0b, 0x34, 0x5e, 0xdb, 0x89, 0x7b, 0x46, 0xcf, 0x43, 0xbf, 0x6b, 0x0d, 0x6d, 0x3a, 0x48,
	0xfb, 0x73, 0x4f, 0x45, 0x45, 0x14, 0x3a, 0x06, 0x65, 0x36, 0x8b, 0x05, 0x70, 0xee, 0xcd, 0x62,
	0x22, 0x48, 0x83, 0x38, 0x72, 0x24, 0x0d, 0xad, 0x01, 0xa4, 0x12, 0xa1, 0xcd, 0xb0, 0xd9, 0x0f,
	0x58, 0x60, 0xe2, 0x55, 0x8c, 0x94, 0x91, 0xf7, 0xb0, 0x1e, 0x26, 0x67, 0x55, 0xc9, 0x2b, 0x85,
	0xad, 0xe2, 0x0b, 0x7d, 0x11, 0x59, 0x3d, 0x6d, 0x4d, 0x3b, 0x34, 0x0f, 0x76, 0xce, 0x91, 0x4d,
	0xa1, 0x1d, 0xc0, 0x26, 0xcf, 0x59, 0x5e, 0x97, 0x03, 0x66, 0xcd, 0x0d, 0x7e, 0xae, 0x75, 0x49,
	0x09, 0xd6, 0x12, 0x34, 0x75, 0x25, 0xaf, 0x14, 0xfe, 0x2f, 0xbe, 0x5a, 0x62, 0xce, 0x67, 0x0f,
	0xaf, 0x4d, 0xde, 0xa4, 0x5d, 0xc1, 0x5e, 0xd9, 0x66, 0x4e, 0x5f, 0xce, 0xa3, 0x72, 0xe0, 0x33,
	0xc8, 0xca, 0x81, 0x54, 0x55, 0xf2, 0xab, 0x85, 0xac, 0xb9, 0x29, 0x26, 0xd2, 0x7f, 0x1b, 0x69,
	0xc1, 0xfe, 0xec, 0x48, 0x1a, 0x06, 0x23, 0x8a, 0xe4, 0x0c, 0x36, 0x22, 0xa4, 0xb1, 0xcf, 0xd2,
	0x89, 0x5b, 0xc5, 0xd7, 0xcb, 0x88, 0x26, 0xba, 0x4d, 0xd9, 0xab, 0x8d, 0x61, 0x7b, 0x2a, 0x43,
	0x8e, 0x20, 0x93, 0xae, 0x57, 0xec, 0x82, 0x48, 0xd8, 0x28, 0x74, 0xf4, 0x16, 0xcf, 0x98, 0xa2,
	0xe2, 0x6e, 0x6d, 0x2b, 0x7f, 0xbc, 0xb6, 0x9f, 0x0a, 0x3c, 0xaa, 0x7b, 0x74, 0x46, 0x48, 0xa9,
	0x95, 0xf2, 0x17, 0x5a, 0xf1, 0x2d, 0xd8, 0x2e, 0x5a, 0xd4, 0xbb, 0x41, 0xce, 0x68, 0xdd, 0xdc,
	0x4c, 0x02, 0x2d, 0xef, 0x06, 0xc9, 0x73, 0x00, 0x9e, 0x64, 0xc1, 0x00, 0x47, 0xea, 0x2a, 0x7f,
	0x15, 0xbc, 0xfc, 0x53, 0x12, 0x20, 0xfb, 0x90, 0xe9, 0x79, 0x3e, 0xc3, 0x48, 0x5d, 0xe3, 0x29,
	0x71, 0xd2, 0xbe, 0x01, 0x99, 0x64, 0x29, 0x34, 0x2a, 0x41, 0x86, 0xdf, 0x42, 0x4a, 0xbf, 0xd4,
	0xc5, 0x45, 0x0b, 0x79, 0x09, 0x3b, 0x23, 0xfc, 0xca, 0xac, 0x09, 0x3a, 0x2b, 0x7c, 0xe6, 0x76,
	0x12, 0x6e, 0x4a, 0x4a, 0xda, 0x0f, 0x05, 0x48, 0x3b, 0xec, 0x3e, 0x9c, 0x55, 0x48, 0x09, 0xb6,
	0x62, 0x0e, 0xc8, 0x7d, 0x2d, 0x96, 0x96, 0x93, 0x00, 0xd2, 0xfa, 0xfa, 0x87, 0xc4, 0xfa, 0x17,
	0x36, 0x1d, 0x98, 0x90, 0x96, 0x27, 0xbf, 0xb5, 0xef, 0x0a, 0xa8, 0xfc, 0x29, 0x4e, 0x70, 0xba,
	0xdd, 0x5b, 0x17, 0xf6, 0x04, 0x72, 0xea, 0x83, 0x28, 0x8d, 0x4b, 0x7d, 0x8e, 0x17, 0x93, 0xbc,
	0x7f, 0x4b, 0xf3, 0x71, 0x7c, 0x2f, 0x46, 0xb5, 0x0e, 0x1c, 0xcc, 0x61, 0xf0, 0xb0, 0x7e, 0x30,
	0x80, 0x54, 0xd1, 0xc7, 0x19, 0xd1, 0x7f, 0xff, 0x45, 0xd1, 0xde, 0x09, 0x59, 0x26, 0xba, 0x96,
	0xfa, 0x2e, 0x68, 0xe7, 0xe2, 0x36, 0xd3, 0x8d, 0x73, 0x5c, 0xb8, 0xba, 0xd8, 0x85, 0x47, 0xc7,
	0x90, 0xbd, 0xb5, 0x02, 0xc9, 0xc2, 0x7a, 0xf9, 0xb4, 0x55, 0xab, 0xec, 0xfe, 0x47, 0x54, 0x78,
	0x52, 0xbb, 0xac, 0xd4, 0xdb, 0xd5, 0x33, 0xab, 0xda, 0xf8, 0x72, 0x59, 0x6f, 0x9c, 0x56, 0xad,
	0xb6, 0x59, 0xdf, 0x55, 0xca, 0x1f, 0xa1, 0xe0, 0x04, 0x43, 0x09, 0xe9, 0x62, 0xa0, 0xc7, 0xae,
	0x33, 0x5f, 0xa7, 0xf2, 0x61, 0x8b, 0x87, 0x13, 0xf0, 0x66, 0x1a, 0x35, 0x43, 0xe7, 0x42, 0xfc,
	0xb7, 0x74, 0x32, 0xfc, 0xd5, 0xbc, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x8d, 0x2b, 0xd0,
	0x8e, 0x06, 0x00, 0x00,
}
