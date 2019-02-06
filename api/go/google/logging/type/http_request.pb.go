// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/logging/type/http_request.proto

package google_logging_type

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A common proto for logging HTTP requests. Only contains semantics
// defined by the HTTP specification. Product-specific logging
// information MUST be defined in a separate message.
type HttpRequest struct {
	// The request method. Examples: `"GET"`, `"HEAD"`, `"PUT"`, `"POST"`.
	RequestMethod string `protobuf:"bytes,1,opt,name=request_method,json=requestMethod" json:"request_method,omitempty"`
	// The scheme (http, https), the host name, the path and the query
	// portion of the URL that was requested.
	// Example: `"http://example.com/some/info?color=red"`.
	RequestUrl string `protobuf:"bytes,2,opt,name=request_url,json=requestUrl" json:"request_url,omitempty"`
	// The size of the HTTP request message in bytes, including the request
	// headers and the request body.
	RequestSize int64 `protobuf:"varint,3,opt,name=request_size,json=requestSize" json:"request_size,omitempty"`
	// The response code indicating the status of response.
	// Examples: 200, 404.
	Status int32 `protobuf:"varint,4,opt,name=status" json:"status,omitempty"`
	// The size of the HTTP response message sent back to the client, in bytes,
	// including the response headers and the response body.
	ResponseSize int64 `protobuf:"varint,5,opt,name=response_size,json=responseSize" json:"response_size,omitempty"`
	// The user agent sent by the client. Example:
	// `"Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; Q312461; .NET CLR 1.0.3705)"`.
	UserAgent string `protobuf:"bytes,6,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	// The IP address (IPv4 or IPv6) of the client that issued the HTTP
	// request. Examples: `"192.168.1.1"`, `"FE80::0202:B3FF:FE1E:8329"`.
	RemoteIp string `protobuf:"bytes,7,opt,name=remote_ip,json=remoteIp" json:"remote_ip,omitempty"`
	// The IP address (IPv4 or IPv6) of the origin server that the request was
	// sent to.
	ServerIp string `protobuf:"bytes,13,opt,name=server_ip,json=serverIp" json:"server_ip,omitempty"`
	// The referer URL of the request, as defined in
	// [HTTP/1.1 Header Field Definitions](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).
	Referer string `protobuf:"bytes,8,opt,name=referer" json:"referer,omitempty"`
	// The request processing latency on the server, from the time the request was
	// received until the response was sent.
	Latency *google_protobuf1.Duration `protobuf:"bytes,14,opt,name=latency" json:"latency,omitempty"`
	// Whether or not a cache lookup was attempted.
	CacheLookup bool `protobuf:"varint,11,opt,name=cache_lookup,json=cacheLookup" json:"cache_lookup,omitempty"`
	// Whether or not an entity was served from cache
	// (with or without validation).
	CacheHit bool `protobuf:"varint,9,opt,name=cache_hit,json=cacheHit" json:"cache_hit,omitempty"`
	// Whether or not the response was validated with the origin server before
	// being served from cache. This field is only meaningful if `cache_hit` is
	// True.
	CacheValidatedWithOriginServer bool `protobuf:"varint,10,opt,name=cache_validated_with_origin_server,json=cacheValidatedWithOriginServer" json:"cache_validated_with_origin_server,omitempty"`
	// The number of HTTP response bytes inserted into cache. Set only when a
	// cache fill was attempted.
	CacheFillBytes int64 `protobuf:"varint,12,opt,name=cache_fill_bytes,json=cacheFillBytes" json:"cache_fill_bytes,omitempty"`
	// Protocol used for the request. Examples: "HTTP/1.1", "HTTP/2", "websocket"
	Protocol string `protobuf:"bytes,15,opt,name=protocol" json:"protocol,omitempty"`
}

func (m *HttpRequest) Reset()                    { *m = HttpRequest{} }
func (m *HttpRequest) String() string            { return proto.CompactTextString(m) }
func (*HttpRequest) ProtoMessage()               {}
func (*HttpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HttpRequest) GetRequestMethod() string {
	if m != nil {
		return m.RequestMethod
	}
	return ""
}

func (m *HttpRequest) GetRequestUrl() string {
	if m != nil {
		return m.RequestUrl
	}
	return ""
}

func (m *HttpRequest) GetRequestSize() int64 {
	if m != nil {
		return m.RequestSize
	}
	return 0
}

func (m *HttpRequest) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *HttpRequest) GetResponseSize() int64 {
	if m != nil {
		return m.ResponseSize
	}
	return 0
}

func (m *HttpRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *HttpRequest) GetRemoteIp() string {
	if m != nil {
		return m.RemoteIp
	}
	return ""
}

func (m *HttpRequest) GetServerIp() string {
	if m != nil {
		return m.ServerIp
	}
	return ""
}

func (m *HttpRequest) GetReferer() string {
	if m != nil {
		return m.Referer
	}
	return ""
}

func (m *HttpRequest) GetLatency() *google_protobuf1.Duration {
	if m != nil {
		return m.Latency
	}
	return nil
}

func (m *HttpRequest) GetCacheLookup() bool {
	if m != nil {
		return m.CacheLookup
	}
	return false
}

func (m *HttpRequest) GetCacheHit() bool {
	if m != nil {
		return m.CacheHit
	}
	return false
}

func (m *HttpRequest) GetCacheValidatedWithOriginServer() bool {
	if m != nil {
		return m.CacheValidatedWithOriginServer
	}
	return false
}

func (m *HttpRequest) GetCacheFillBytes() int64 {
	if m != nil {
		return m.CacheFillBytes
	}
	return 0
}

func (m *HttpRequest) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func init() {
	proto.RegisterType((*HttpRequest)(nil), "google.logging.type.HttpRequest")
}

func init() { proto.RegisterFile("google/logging/type/http_request.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0x9e, 0x92, 0x6c, 0x0e, 0x54, 0x8b, 0x04, 0xdb, 0x00, 0x25, 0x14, 0x81, 0x7c,
	0xe5, 0x48, 0xf4, 0x09, 0x08, 0x08, 0x5a, 0x54, 0x44, 0xe5, 0x72, 0xb8, 0x89, 0x64, 0x39, 0xf1,
	0xc4, 0x5e, 0xb1, 0xf1, 0x2e, 0xeb, 0x71, 0x51, 0xfa, 0x48, 0xdc, 0xf0, 0x1e, 0x3c, 0x15, 0xf2,
	0xec, 0xae, 0xc4, 0x05, 0x97, 0xf3, 0x7f, 0xdf, 0x9f, 0x71, 0xc6, 0x66, 0x2f, 0x4b, 0xad, 0x4b,
	0x05, 0x73, 0xa5, 0xcb, 0x52, 0xd6, 0xe5, 0x1c, 0x77, 0x06, 0xe6, 0x15, 0xa2, 0xc9, 0x2c, 0xfc,
	0x68, 0xa1, 0xc1, 0xc4, 0x58, 0x8d, 0x9a, 0xdf, 0x77, 0x5e, 0xe2, 0xbd, 0xa4, 0xf3, 0xa6, 0x8f,
	0x7d, 0x39, 0x37, 0x72, 0x9e, 0xd7, 0xb5, 0xc6, 0x1c, 0xa5, 0xae, 0x1b, 0x57, 0x99, 0x9e, 0x7a,
	0x4a, 0xd3, 0xaa, 0xdd, 0xcc, 0x8b, 0xd6, 0x92, 0xe0, 0xf8, 0xd9, 0xef, 0x03, 0x36, 0xbc, 0x40,
	0x34, 0xa9, 0x5b, 0xc4, 0x5f, 0xb0, 0x89, 0xdf, 0x99, 0x6d, 0x01, 0x2b, 0x5d, 0x88, 0x68, 0x16,
	0xc5, 0x83, 0x74, 0xec, 0xd3, 0x8f, 0x14, 0xf2, 0xa7, 0x6c, 0x18, 0xb4, 0xd6, 0x2a, 0xb1, 0x47,
	0x0e, 0xf3, 0xd1, 0x17, 0xab, 0xf8, 0x33, 0x36, 0x0a, 0x42, 0x23, 0xef, 0x40, 0xec, 0xcf, 0xa2,
	0x78, 0x3f, 0x0d, 0xa5, 0x1b, 0x79, 0x07, 0xfc, 0x01, 0x3b, 0x6a, 0x30, 0xc7, 0xb6, 0x11, 0x07,
	0xb3, 0x28, 0x3e, 0x4c, 0xfd, 0xc4, 0x9f, 0xb3, 0xb1, 0x85, 0xc6, 0xe8, 0xba, 0x01, 0xd7, 0x3d,
	0xa4, 0xee, 0x28, 0x84, 0x54, 0x7e, 0xc2, 0x58, 0xdb, 0x80, 0xcd, 0xf2, 0x12, 0x6a, 0x14, 0x47,
	0xb4, 0x7f, 0xd0, 0x25, 0xaf, 0xbb, 0x80, 0x3f, 0x62, 0x03, 0x0b, 0x5b, 0x8d, 0x90, 0x49, 0x23,
	0x7a, 0x44, 0xfb, 0x2e, 0xb8, 0x34, 0x1d, 0x6c, 0xc0, 0xde, 0x82, 0xed, 0xe0, 0xd8, 0x41, 0x17,
	0x5c, 0x1a, 0x2e, 0x58, 0xcf, 0xc2, 0x06, 0x2c, 0x58, 0xd1, 0x27, 0x14, 0x46, 0x7e, 0xce, 0x7a,
	0x2a, 0x47, 0xa8, 0xd7, 0x3b, 0x31, 0x99, 0x45, 0xf1, 0xf0, 0xd5, 0x49, 0xe2, 0xdf, 0x47, 0x38,
	0x6e, 0xf2, 0xd6, 0x1f, 0x37, 0x0d, 0x66, 0x77, 0x87, 0x75, 0xbe, 0xae, 0x20, 0x53, 0x5a, 0x7f,
	0x6f, 0x8d, 0x18, 0xce, 0xa2, 0xb8, 0x9f, 0x0e, 0x29, 0xbb, 0xa2, 0xa8, 0x7b, 0x1c, 0xa7, 0x54,
	0x12, 0xc5, 0x80, 0x78, 0x9f, 0x82, 0x0b, 0x89, 0xfc, 0x03, 0x3b, 0x73, 0xf0, 0x36, 0x57, 0xb2,
	0xc8, 0x11, 0x8a, 0xec, 0xa7, 0xc4, 0x2a, 0xd3, 0x56, 0x96, 0xb2, 0xce, 0xdc, 0x63, 0x0b, 0x46,
	0xad, 0x53, 0x32, 0xbf, 0x06, 0xf1, 0x9b, 0xc4, 0xea, 0x13, 0x69, 0x37, 0x64, 0xf1, 0x98, 0x1d,
	0xbb, 0xdf, 0xda, 0x48, 0xa5, 0xb2, 0xd5, 0x0e, 0xa1, 0x11, 0x23, 0xba, 0xed, 0x84, 0xf2, 0x77,
	0x52, 0xa9, 0x45, 0x97, 0xf2, 0x29, 0xeb, 0xd3, 0x7f, 0x5a, 0x6b, 0x25, 0xee, 0xb9, 0x03, 0x85,
	0x79, 0x01, 0xec, 0xe1, 0x5a, 0x6f, 0x93, 0xff, 0x7c, 0x8a, 0x8b, 0xe3, 0x7f, 0xbe, 0xa4, 0xeb,
	0xce, 0xbf, 0x8e, 0x7e, 0xed, 0x9d, 0xbc, 0x77, 0xe6, 0x1b, 0xa5, 0xdb, 0x22, 0xb9, 0xf2, 0xfe,
	0xe7, 0x9d, 0x81, 0x3f, 0x81, 0x2d, 0x89, 0x2d, 0x3d, 0x5b, 0x76, 0x6c, 0x75, 0x44, 0x0b, 0xcf,
	0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x77, 0x87, 0x13, 0x6b, 0x1c, 0x03, 0x00, 0x00,
}
