// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/protobuf/api.proto

package google_protobuf

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Api is a light-weight descriptor for an API Interface.
//
// Interfaces are also described as "protocol buffer services" in some contexts,
// such as by the "service" keyword in a .proto file, but they are different
// from API Services, which represent a concrete implementation of an interface
// as opposed to simply a description of methods and bindings. They are also
// sometimes simply referred to as "APIs" in other contexts, such as the name of
// this message itself. See https://cloud.google.com/apis/design/glossary for
// detailed terminology.
type Api struct {
	// The fully qualified name of this interface, including package name
	// followed by the interface's simple name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The methods of this interface, in unspecified order.
	Methods []*Method `protobuf:"bytes,2,rep,name=methods" json:"methods,omitempty"`
	// Any metadata attached to the interface.
	Options []*Option `protobuf:"bytes,3,rep,name=options" json:"options,omitempty"`
	// A version string for this interface. If specified, must have the form
	// `major-version.minor-version`, as in `1.10`. If the minor version is
	// omitted, it defaults to zero. If the entire version field is empty, the
	// major version is derived from the package name, as outlined below. If the
	// field is not empty, the version in the package name will be verified to be
	// consistent with what is provided here.
	//
	// The versioning schema uses [semantic
	// versioning](http://semver.org) where the major version number
	// indicates a breaking change and the minor version an additive,
	// non-breaking change. Both version numbers are signals to users
	// what to expect from different versions, and should be carefully
	// chosen based on the product plan.
	//
	// The major version is also reflected in the package name of the
	// interface, which must end in `v<major-version>`, as in
	// `google.feature.v1`. For major versions 0 and 1, the suffix can
	// be omitted. Zero major versions must only be used for
	// experimental, non-GA interfaces.
	//
	//
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// Source context for the protocol buffer service represented by this
	// message.
	SourceContext *SourceContext `protobuf:"bytes,5,opt,name=source_context,json=sourceContext" json:"source_context,omitempty"`
	// Included interfaces. See [Mixin][].
	Mixins []*Mixin `protobuf:"bytes,6,rep,name=mixins" json:"mixins,omitempty"`
	// The source syntax of the service.
	Syntax Syntax `protobuf:"varint,7,opt,name=syntax,proto3,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
}

func (m *Api) Reset()                    { *m = Api{} }
func (m *Api) String() string            { return proto.CompactTextString(m) }
func (*Api) ProtoMessage()               {}
func (*Api) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{0} }

func (m *Api) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Api) GetMethods() []*Method {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *Api) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Api) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Api) GetSourceContext() *SourceContext {
	if m != nil {
		return m.SourceContext
	}
	return nil
}

func (m *Api) GetMixins() []*Mixin {
	if m != nil {
		return m.Mixins
	}
	return nil
}

func (m *Api) GetSyntax() Syntax {
	if m != nil {
		return m.Syntax
	}
	return Syntax_SYNTAX_PROTO2
}

// Method represents a method of an API interface.
type Method struct {
	// The simple name of this method.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A URL of the input message type.
	RequestTypeUrl string `protobuf:"bytes,2,opt,name=request_type_url,json=requestTypeUrl,proto3" json:"request_type_url,omitempty"`
	// If true, the request is streamed.
	RequestStreaming bool `protobuf:"varint,3,opt,name=request_streaming,json=requestStreaming,proto3" json:"request_streaming,omitempty"`
	// The URL of the output message type.
	ResponseTypeUrl string `protobuf:"bytes,4,opt,name=response_type_url,json=responseTypeUrl,proto3" json:"response_type_url,omitempty"`
	// If true, the response is streamed.
	ResponseStreaming bool `protobuf:"varint,5,opt,name=response_streaming,json=responseStreaming,proto3" json:"response_streaming,omitempty"`
	// Any metadata attached to the method.
	Options []*Option `protobuf:"bytes,6,rep,name=options" json:"options,omitempty"`
	// The source syntax of this method.
	Syntax Syntax `protobuf:"varint,7,opt,name=syntax,proto3,enum=google.protobuf.Syntax" json:"syntax,omitempty"`
}

func (m *Method) Reset()                    { *m = Method{} }
func (m *Method) String() string            { return proto.CompactTextString(m) }
func (*Method) ProtoMessage()               {}
func (*Method) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{1} }

func (m *Method) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Method) GetRequestTypeUrl() string {
	if m != nil {
		return m.RequestTypeUrl
	}
	return ""
}

func (m *Method) GetRequestStreaming() bool {
	if m != nil {
		return m.RequestStreaming
	}
	return false
}

func (m *Method) GetResponseTypeUrl() string {
	if m != nil {
		return m.ResponseTypeUrl
	}
	return ""
}

func (m *Method) GetResponseStreaming() bool {
	if m != nil {
		return m.ResponseStreaming
	}
	return false
}

func (m *Method) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Method) GetSyntax() Syntax {
	if m != nil {
		return m.Syntax
	}
	return Syntax_SYNTAX_PROTO2
}

// Declares an API Interface to be included in this interface. The including
// interface must redeclare all the methods from the included interface, but
// documentation and options are inherited as follows:
//
// - If after comment and whitespace stripping, the documentation
//   string of the redeclared method is empty, it will be inherited
//   from the original method.
//
// - Each annotation belonging to the service config (http,
//   visibility) which is not set in the redeclared method will be
//   inherited.
//
// - If an http annotation is inherited, the path pattern will be
//   modified as follows. Any version prefix will be replaced by the
//   version of the including interface plus the [root][] path if
//   specified.
//
// Example of a simple mixin:
//
//     package google.acl.v1;
//     service AccessControl {
//       // Get the underlying ACL object.
//       rpc GetAcl(GetAclRequest) returns (Acl) {
//         option (google.api.http).get = "/v1/{resource=**}:getAcl";
//       }
//     }
//
//     package google.storage.v2;
//     service Storage {
//       rpc GetAcl(GetAclRequest) returns (Acl);
//
//       // Get a data record.
//       rpc GetData(GetDataRequest) returns (Data) {
//         option (google.api.http).get = "/v2/{resource=**}";
//       }
//     }
//
// Example of a mixin configuration:
//
//     apis:
//     - name: google.storage.v2.Storage
//       mixins:
//       - name: google.acl.v1.AccessControl
//
// The mixin construct implies that all methods in `AccessControl` are
// also declared with same name and request/response types in
// `Storage`. A documentation generator or annotation processor will
// see the effective `Storage.GetAcl` method after inherting
// documentation and annotations as follows:
//
//     service Storage {
//       // Get the underlying ACL object.
//       rpc GetAcl(GetAclRequest) returns (Acl) {
//         option (google.api.http).get = "/v2/{resource=**}:getAcl";
//       }
//       ...
//     }
//
// Note how the version in the path pattern changed from `v1` to `v2`.
//
// If the `root` field in the mixin is specified, it should be a
// relative path under which inherited HTTP paths are placed. Example:
//
//     apis:
//     - name: google.storage.v2.Storage
//       mixins:
//       - name: google.acl.v1.AccessControl
//         root: acls
//
// This implies the following inherited HTTP annotation:
//
//     service Storage {
//       // Get the underlying ACL object.
//       rpc GetAcl(GetAclRequest) returns (Acl) {
//         option (google.api.http).get = "/v2/acls/{resource=**}:getAcl";
//       }
//       ...
//     }
type Mixin struct {
	// The fully qualified name of the interface which is included.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// If non-empty specifies a path under which inherited HTTP paths
	// are rooted.
	Root string `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
}

func (m *Mixin) Reset()                    { *m = Mixin{} }
func (m *Mixin) String() string            { return proto.CompactTextString(m) }
func (*Mixin) ProtoMessage()               {}
func (*Mixin) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{2} }

func (m *Mixin) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Mixin) GetRoot() string {
	if m != nil {
		return m.Root
	}
	return ""
}

func init() {
	proto.RegisterType((*Api)(nil), "google.protobuf.Api")
	proto.RegisterType((*Method)(nil), "google.protobuf.Method")
	proto.RegisterType((*Mixin)(nil), "google.protobuf.Mixin")
}

func init() { proto.RegisterFile("google/protobuf/api.proto", fileDescriptorApi) }

var fileDescriptorApi = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xd1, 0x6e, 0xda, 0x30,
	0x14, 0x86, 0x95, 0x04, 0x02, 0x33, 0x1a, 0x6c, 0x9e, 0xb4, 0x79, 0x5c, 0xa0, 0x08, 0xed, 0x22,
	0xda, 0xb4, 0x44, 0x63, 0x4f, 0x00, 0xd3, 0xc4, 0xa4, 0x69, 0x5a, 0x14, 0x5a, 0xf5, 0x12, 0x05,
	0xea, 0x52, 0x4b, 0x89, 0xed, 0xda, 0x4e, 0x0b, 0xaf, 0xd3, 0xcb, 0x5e, 0xf6, 0x0d, 0xfa, 0x66,
	0x55, 0x9c, 0x18, 0x68, 0xa0, 0x52, 0x7b, 0xe7, 0xe3, 0xff, 0x3b, 0xbf, 0x7d, 0xfe, 0x03, 0x3e,
	0xaf, 0x18, 0x5b, 0xa5, 0x38, 0xe4, 0x82, 0x29, 0xb6, 0xc8, 0x2f, 0xc2, 0x84, 0x93, 0x40, 0x17,
	0xb0, 0x57, 0x4a, 0x81, 0x91, 0xfa, 0x5f, 0xea, 0xac, 0x64, 0xb9, 0x58, 0xe2, 0xf9, 0x92, 0x51,
	0x85, 0xd7, 0xaa, 0x04, 0xfb, 0xfd, 0x3a, 0xa5, 0x36, 0xbc, 0x32, 0x19, 0x3e, 0xd8, 0xc0, 0x19,
	0x73, 0x02, 0x21, 0x68, 0xd0, 0x24, 0xc3, 0xc8, 0xf2, 0x2c, 0xff, 0x4d, 0xac, 0xcf, 0xf0, 0x07,
	0x68, 0x65, 0x58, 0x5d, 0xb2, 0x73, 0x89, 0x6c, 0xcf, 0xf1, 0x3b, 0xa3, 0x4f, 0x41, 0xed, 0x03,
	0xc1, 0x3f, 0xad, 0xc7, 0x86, 0x2b, 0x5a, 0x18, 0x57, 0x84, 0x51, 0x89, 0x9c, 0x67, 0x5a, 0xfe,
	0x6b, 0x3d, 0x36, 0x1c, 0x44, 0xa0, 0x75, 0x8d, 0x85, 0x24, 0x8c, 0xa2, 0x86, 0x7e, 0xdc, 0x94,
	0xf0, 0x37, 0xe8, 0x3e, 0x9d, 0x07, 0x35, 0x3d, 0xcb, 0xef, 0x8c, 0x06, 0x07, 0x9e, 0x33, 0x8d,
	0xfd, 0x2a, 0xa9, 0xf8, 0xad, 0xdc, 0x2f, 0x61, 0x00, 0xdc, 0x8c, 0xac, 0x09, 0x95, 0xc8, 0xd5,
	0x5f, 0xfa, 0x78, 0x38, 0x45, 0x21, 0xc7, 0x15, 0x05, 0x43, 0xe0, 0xca, 0x0d, 0x55, 0xc9, 0x1a,
	0xb5, 0x3c, 0xcb, 0xef, 0x1e, 0x19, 0x61, 0xa6, 0xe5, 0xb8, 0xc2, 0x86, 0xf7, 0x36, 0x70, 0xcb,
	0x20, 0x8e, 0xc6, 0xe8, 0x83, 0x77, 0x02, 0x5f, 0xe5, 0x58, 0xaa, 0x79, 0x11, 0xfc, 0x3c, 0x17,
	0x29, 0xb2, 0xb5, 0xde, 0xad, 0xee, 0x4f, 0x36, 0x1c, 0x9f, 0x8a, 0x14, 0x7e, 0x03, 0xef, 0x0d,
	0x29, 0x95, 0xc0, 0x49, 0x46, 0xe8, 0x0a, 0x39, 0x9e, 0xe5, 0xb7, 0x63, 0x63, 0x31, 0x33, 0xf7,
	0xf0, 0x6b, 0x01, 0x4b, 0xce, 0xa8, 0xc4, 0x3b, 0xdf, 0x32, 0xc1, 0x9e, 0x11, 0x8c, 0xf1, 0x77,
	0x00, 0xb7, 0xec, 0xce, 0xb9, 0xa9, 0x9d, 0xb7, 0x2e, 0x3b, 0xeb, 0xbd, 0x2d, 0xba, 0x2f, 0xdc,
	0xe2, 0xab, 0x43, 0x0b, 0x41, 0x53, 0xc7, 0x7e, 0x34, 0x32, 0x08, 0x1a, 0x82, 0x31, 0x55, 0xc5,
	0xa4, 0xcf, 0x93, 0x3f, 0xe0, 0xc3, 0x92, 0x65, 0x75, 0xdb, 0x49, 0x7b, 0xcc, 0x49, 0x54, 0x14,
	0x91, 0x75, 0x6b, 0x3b, 0xd3, 0x68, 0x72, 0x67, 0x0f, 0xa6, 0x25, 0x13, 0x99, 0xa7, 0xcf, 0x70,
	0x9a, 0xfe, 0xa5, 0xec, 0x86, 0x16, 0x79, 0xc8, 0x85, 0xab, 0x9b, 0x7f, 0x3e, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x7b, 0x9c, 0xed, 0x00, 0x6a, 0x03, 0x00, 0x00,
}
