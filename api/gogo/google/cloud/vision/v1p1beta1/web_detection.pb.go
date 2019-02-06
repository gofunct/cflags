// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/cloud/vision/v1p1beta1/web_detection.proto

package google_cloud_vision_v1p1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Relevant information for the image from the Internet.
type WebDetection struct {
	// Deduced entities from similar images on the Internet.
	WebEntities []*WebDetection_WebEntity `protobuf:"bytes,1,rep,name=web_entities,json=webEntities" json:"web_entities,omitempty"`
	// Fully matching images from the Internet.
	// Can include resized copies of the query image.
	FullMatchingImages []*WebDetection_WebImage `protobuf:"bytes,2,rep,name=full_matching_images,json=fullMatchingImages" json:"full_matching_images,omitempty"`
	// Partial matching images from the Internet.
	// Those images are similar enough to share some key-point features. For
	// example an original image will likely have partial matching for its crops.
	PartialMatchingImages []*WebDetection_WebImage `protobuf:"bytes,3,rep,name=partial_matching_images,json=partialMatchingImages" json:"partial_matching_images,omitempty"`
	// Web pages containing the matching images from the Internet.
	PagesWithMatchingImages []*WebDetection_WebPage `protobuf:"bytes,4,rep,name=pages_with_matching_images,json=pagesWithMatchingImages" json:"pages_with_matching_images,omitempty"`
	// The visually similar image results.
	VisuallySimilarImages []*WebDetection_WebImage `protobuf:"bytes,6,rep,name=visually_similar_images,json=visuallySimilarImages" json:"visually_similar_images,omitempty"`
	// Best guess text labels for the request image.
	BestGuessLabels []*WebDetection_WebLabel `protobuf:"bytes,8,rep,name=best_guess_labels,json=bestGuessLabels" json:"best_guess_labels,omitempty"`
}

func (m *WebDetection) Reset()                    { *m = WebDetection{} }
func (m *WebDetection) String() string            { return proto.CompactTextString(m) }
func (*WebDetection) ProtoMessage()               {}
func (*WebDetection) Descriptor() ([]byte, []int) { return fileDescriptorWebDetection, []int{0} }

func (m *WebDetection) GetWebEntities() []*WebDetection_WebEntity {
	if m != nil {
		return m.WebEntities
	}
	return nil
}

func (m *WebDetection) GetFullMatchingImages() []*WebDetection_WebImage {
	if m != nil {
		return m.FullMatchingImages
	}
	return nil
}

func (m *WebDetection) GetPartialMatchingImages() []*WebDetection_WebImage {
	if m != nil {
		return m.PartialMatchingImages
	}
	return nil
}

func (m *WebDetection) GetPagesWithMatchingImages() []*WebDetection_WebPage {
	if m != nil {
		return m.PagesWithMatchingImages
	}
	return nil
}

func (m *WebDetection) GetVisuallySimilarImages() []*WebDetection_WebImage {
	if m != nil {
		return m.VisuallySimilarImages
	}
	return nil
}

func (m *WebDetection) GetBestGuessLabels() []*WebDetection_WebLabel {
	if m != nil {
		return m.BestGuessLabels
	}
	return nil
}

// Entity deduced from similar images on the Internet.
type WebDetection_WebEntity struct {
	// Opaque entity ID.
	EntityId string `protobuf:"bytes,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	// Overall relevancy score for the entity.
	// Not normalized and not comparable across different image queries.
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	// Canonical description of the entity, in English.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *WebDetection_WebEntity) Reset()         { *m = WebDetection_WebEntity{} }
func (m *WebDetection_WebEntity) String() string { return proto.CompactTextString(m) }
func (*WebDetection_WebEntity) ProtoMessage()    {}
func (*WebDetection_WebEntity) Descriptor() ([]byte, []int) {
	return fileDescriptorWebDetection, []int{0, 0}
}

func (m *WebDetection_WebEntity) GetEntityId() string {
	if m != nil {
		return m.EntityId
	}
	return ""
}

func (m *WebDetection_WebEntity) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *WebDetection_WebEntity) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Metadata for online images.
type WebDetection_WebImage struct {
	// The result image URL.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// (Deprecated) Overall relevancy score for the image.
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (m *WebDetection_WebImage) Reset()         { *m = WebDetection_WebImage{} }
func (m *WebDetection_WebImage) String() string { return proto.CompactTextString(m) }
func (*WebDetection_WebImage) ProtoMessage()    {}
func (*WebDetection_WebImage) Descriptor() ([]byte, []int) {
	return fileDescriptorWebDetection, []int{0, 1}
}

func (m *WebDetection_WebImage) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *WebDetection_WebImage) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

// Metadata for web pages.
type WebDetection_WebPage struct {
	// The result web page URL.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// (Deprecated) Overall relevancy score for the web page.
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	// Title for the web page, may contain HTML markups.
	PageTitle string `protobuf:"bytes,3,opt,name=page_title,json=pageTitle,proto3" json:"page_title,omitempty"`
	// Fully matching images on the page.
	// Can include resized copies of the query image.
	FullMatchingImages []*WebDetection_WebImage `protobuf:"bytes,4,rep,name=full_matching_images,json=fullMatchingImages" json:"full_matching_images,omitempty"`
	// Partial matching images on the page.
	// Those images are similar enough to share some key-point features. For
	// example an original image will likely have partial matching for its
	// crops.
	PartialMatchingImages []*WebDetection_WebImage `protobuf:"bytes,5,rep,name=partial_matching_images,json=partialMatchingImages" json:"partial_matching_images,omitempty"`
}

func (m *WebDetection_WebPage) Reset()         { *m = WebDetection_WebPage{} }
func (m *WebDetection_WebPage) String() string { return proto.CompactTextString(m) }
func (*WebDetection_WebPage) ProtoMessage()    {}
func (*WebDetection_WebPage) Descriptor() ([]byte, []int) {
	return fileDescriptorWebDetection, []int{0, 2}
}

func (m *WebDetection_WebPage) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *WebDetection_WebPage) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *WebDetection_WebPage) GetPageTitle() string {
	if m != nil {
		return m.PageTitle
	}
	return ""
}

func (m *WebDetection_WebPage) GetFullMatchingImages() []*WebDetection_WebImage {
	if m != nil {
		return m.FullMatchingImages
	}
	return nil
}

func (m *WebDetection_WebPage) GetPartialMatchingImages() []*WebDetection_WebImage {
	if m != nil {
		return m.PartialMatchingImages
	}
	return nil
}

// Label to provide extra metadata for the web detection.
type WebDetection_WebLabel struct {
	// Label for extra metadata.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// The BCP-47 language code for `label`, such as "en-US" or "sr-Latn".
	// For more information, see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode string `protobuf:"bytes,2,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
}

func (m *WebDetection_WebLabel) Reset()         { *m = WebDetection_WebLabel{} }
func (m *WebDetection_WebLabel) String() string { return proto.CompactTextString(m) }
func (*WebDetection_WebLabel) ProtoMessage()    {}
func (*WebDetection_WebLabel) Descriptor() ([]byte, []int) {
	return fileDescriptorWebDetection, []int{0, 3}
}

func (m *WebDetection_WebLabel) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *WebDetection_WebLabel) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func init() {
	proto.RegisterType((*WebDetection)(nil), "google.cloud.vision.v1p1beta1.WebDetection")
	proto.RegisterType((*WebDetection_WebEntity)(nil), "google.cloud.vision.v1p1beta1.WebDetection.WebEntity")
	proto.RegisterType((*WebDetection_WebImage)(nil), "google.cloud.vision.v1p1beta1.WebDetection.WebImage")
	proto.RegisterType((*WebDetection_WebPage)(nil), "google.cloud.vision.v1p1beta1.WebDetection.WebPage")
	proto.RegisterType((*WebDetection_WebLabel)(nil), "google.cloud.vision.v1p1beta1.WebDetection.WebLabel")
}

func init() {
	proto.RegisterFile("google/cloud/vision/v1p1beta1/web_detection.proto", fileDescriptorWebDetection)
}

var fileDescriptorWebDetection = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0x87, 0x95, 0x76, 0x1b, 0xcd, 0x69, 0x11, 0xcc, 0x1a, 0x5a, 0x14, 0x98, 0x54, 0xe0, 0xa6,
	0x57, 0x89, 0xba, 0xc1, 0x15, 0x77, 0x83, 0x09, 0x4d, 0x02, 0xa9, 0x0a, 0x48, 0xe3, 0x2e, 0x73,
	0x12, 0x2f, 0x3d, 0x92, 0x1b, 0x47, 0xb1, 0xd3, 0xaa, 0x2f, 0xc0, 0xa3, 0xf0, 0x8c, 0x5c, 0xa2,
	0xe3, 0x24, 0x53, 0x45, 0x61, 0x62, 0x14, 0x71, 0x67, 0x1f, 0xe9, 0xf7, 0x7d, 0xf6, 0xf1, 0x1f,
	0x98, 0xe6, 0x4a, 0xe5, 0x52, 0x84, 0xa9, 0x54, 0x75, 0x16, 0x2e, 0x51, 0xa3, 0x2a, 0xc2, 0xe5,
	0xb4, 0x9c, 0x26, 0xc2, 0xf0, 0x69, 0xb8, 0x12, 0x49, 0x9c, 0x09, 0x23, 0x52, 0x83, 0xaa, 0x08,
	0xca, 0x4a, 0x19, 0xc5, 0x4e, 0x9a, 0x48, 0x60, 0x23, 0x41, 0x13, 0x09, 0x6e, 0x23, 0xfe, 0xb3,
	0x96, 0xc8, 0x4b, 0x0c, 0x79, 0x51, 0x28, 0xc3, 0x29, 0xab, 0x9b, 0xf0, 0x8b, 0xaf, 0x2e, 0x8c,
	0xae, 0x44, 0xf2, 0xae, 0x63, 0xb2, 0x2f, 0x30, 0x22, 0x89, 0x28, 0x0c, 0x1a, 0x14, 0xda, 0x73,
	0xc6, 0xfd, 0xc9, 0xf0, 0xf4, 0x75, 0x70, 0xa7, 0x24, 0xd8, 0x44, 0xd0, 0xe4, 0x82, 0xe2, 0xeb,
	0x68, 0xb8, 0x6a, 0x87, 0x28, 0x34, 0xbb, 0x81, 0xa3, 0x9b, 0x5a, 0xca, 0x78, 0xc1, 0x4d, 0x3a,
	0xc7, 0x22, 0x8f, 0x71, 0xc1, 0x73, 0xa1, 0xbd, 0x9e, 0x35, 0xbc, 0xba, 0xa7, 0xe1, 0x92, 0xc2,
	0x11, 0x23, 0xe2, 0xc7, 0x16, 0x68, 0x4b, 0x9a, 0x49, 0x38, 0x2e, 0x79, 0x65, 0x90, 0x6f, 0xab,
	0xfa, 0x3b, 0xa8, 0x9e, 0xb4, 0xd0, 0x9f, 0x6c, 0x25, 0xf8, 0x25, 0x0d, 0xe2, 0x15, 0x9a, 0xf9,
	0x96, 0x70, 0xcf, 0x0a, 0xcf, 0xee, 0x29, 0x9c, 0x91, 0xef, 0xd8, 0x62, 0xaf, 0xd0, 0xcc, 0xb7,
	0xf7, 0xb7, 0x44, 0x5d, 0x73, 0x29, 0xd7, 0xb1, 0xc6, 0x05, 0x4a, 0x5e, 0x75, 0xba, 0x83, 0x5d,
	0xf6, 0xd7, 0x41, 0x3f, 0x35, 0xcc, 0xd6, 0x76, 0x0d, 0x87, 0x89, 0xd0, 0x26, 0xce, 0x6b, 0xa1,
	0x75, 0x2c, 0x79, 0x22, 0xa4, 0xf6, 0x06, 0x7f, 0xe5, 0xf9, 0x40, 0xe1, 0xe8, 0x11, 0xe1, 0xde,
	0x13, 0xcd, 0xce, 0xb5, 0x7f, 0x0d, 0xee, 0xed, 0x8d, 0x61, 0x4f, 0xc1, 0xb5, 0x57, 0x6f, 0x1d,
	0x63, 0xe6, 0x39, 0x63, 0x67, 0xe2, 0x46, 0x83, 0xa6, 0x70, 0x99, 0xb1, 0x23, 0xd8, 0xd7, 0xa9,
	0xaa, 0x84, 0xd7, 0x1b, 0x3b, 0x93, 0x5e, 0xd4, 0x4c, 0xd8, 0x18, 0x86, 0x99, 0xd0, 0x69, 0x85,
	0x25, 0x89, 0xbc, 0xbe, 0x0d, 0x6d, 0x96, 0xfc, 0x53, 0x18, 0x74, 0xdb, 0x64, 0x8f, 0xa1, 0x5f,
	0x57, 0xb2, 0x45, 0xd3, 0xf0, 0xd7, 0x54, 0xff, 0x5b, 0x0f, 0x1e, 0xb4, 0x47, 0xf1, 0xa7, 0x19,
	0x76, 0x02, 0x40, 0x87, 0x16, 0x1b, 0x34, 0x52, 0xb4, 0x0b, 0x71, 0xa9, 0xf2, 0x99, 0x0a, 0xbf,
	0x7d, 0x00, 0x7b, 0xff, 0xef, 0x01, 0xec, 0xff, 0xf3, 0x07, 0xe0, 0x5f, 0xd8, 0xe6, 0xda, 0xb3,
	0xa4, 0xb6, 0xd8, 0x1b, 0xd2, 0xb6, 0xaa, 0x99, 0xb0, 0x97, 0xf0, 0x50, 0xf2, 0x22, 0xaf, 0xa9,
	0x35, 0xa9, 0xca, 0x9a, 0xa6, 0xb9, 0xd1, 0xa8, 0x2b, 0xbe, 0x55, 0x99, 0x38, 0x7f, 0x03, 0xcf,
	0x53, 0xb5, 0xb8, 0x7b, 0x61, 0xe7, 0x87, 0x9b, 0x2b, 0x9b, 0xd1, 0x07, 0x36, 0x73, 0xbe, 0x3b,
	0x4e, 0x72, 0x60, 0x3f, 0xb3, 0xb3, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x60, 0x3c, 0x7f, 0x58,
	0x3e, 0x05, 0x00, 0x00,
}
