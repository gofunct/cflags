// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/httpbody.proto

package google_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Message that represents an arbitrary HTTP body. It should only be used for
// payload formats that can't be represented as JSON, such as raw binary or
// an HTML page.
//
//
// This message can be used both in streaming and non-streaming API methods in
// the request as well as the response.
//
// It can be used as a top-level request field, which is convenient if one
// wants to extract parameters from either the URL or HTTP template into the
// request fields and also want access to the raw HTTP body.
//
// Example:
//
//     message GetResourceRequest {
//       // A unique request id.
//       string request_id = 1;
//
//       // The raw HTTP body is bound to this field.
//       google.api.HttpBody http_body = 2;
//     }
//
//     service ResourceService {
//       rpc GetResource(GetResourceRequest) returns (google.api.HttpBody);
//       rpc UpdateResource(google.api.HttpBody) returns (google.protobuf.Empty);
//     }
//
// Example with streaming methods:
//
//     service CaldavService {
//       rpc GetCalendar(stream google.api.HttpBody)
//         returns (stream google.api.HttpBody);
//       rpc UpdateCalendar(stream google.api.HttpBody)
//         returns (stream google.api.HttpBody);
//     }
//
// Use of this type only changes how the request and response bodies are
// handled, all other features will continue to work unchanged.
type HttpBody struct {
	// The HTTP Content-Type string representing the content type of the body.
	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	// HTTP body binary data.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *HttpBody) Reset()                    { *m = HttpBody{} }
func (m *HttpBody) String() string            { return proto.CompactTextString(m) }
func (*HttpBody) ProtoMessage()               {}
func (*HttpBody) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *HttpBody) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *HttpBody) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*HttpBody)(nil), "google.api.HttpBody")
}

func init() { proto.RegisterFile("google/api/httpbody.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0x28, 0x29, 0x29, 0x48, 0xca, 0x4f, 0xa9, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0x48, 0xe9, 0x25, 0x16, 0x64, 0x2a, 0x39, 0x72,
	0x71, 0x78, 0x94, 0x94, 0x14, 0x38, 0xe5, 0xa7, 0x54, 0x0a, 0x29, 0x72, 0xf1, 0x24, 0xe7, 0xe7,
	0x95, 0xa4, 0xe6, 0x95, 0xc4, 0x97, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0x71, 0x43, 0xc5, 0x42, 0x2a, 0x0b, 0x52, 0x85, 0x84, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x12, 0x25,
	0x98, 0x14, 0x18, 0x35, 0x78, 0x82, 0xc0, 0x6c, 0x27, 0x0d, 0x2e, 0xbe, 0xe4, 0xfc, 0x5c, 0x3d,
	0x84, 0xa1, 0x4e, 0xbc, 0x30, 0x23, 0x03, 0x40, 0xf6, 0x05, 0x30, 0x2e, 0x62, 0x62, 0x71, 0x77,
	0x0c, 0xf0, 0x4c, 0x62, 0x03, 0xdb, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xcf, 0xd4,
	0x9d, 0x9c, 0x00, 0x00, 0x00,
}
