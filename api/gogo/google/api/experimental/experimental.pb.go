// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/api/experimental/experimental.proto

package google_api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Experimental service configuration. These configuration options can
// only be used by whitelisted users.
type Experimental struct {
	// Authorization configuration.
	Authorization *AuthorizationConfig `protobuf:"bytes,8,opt,name=authorization" json:"authorization,omitempty"`
}

func (m *Experimental) Reset()                    { *m = Experimental{} }
func (m *Experimental) String() string            { return proto.CompactTextString(m) }
func (*Experimental) ProtoMessage()               {}
func (*Experimental) Descriptor() ([]byte, []int) { return fileDescriptorExperimental, []int{0} }

func (m *Experimental) GetAuthorization() *AuthorizationConfig {
	if m != nil {
		return m.Authorization
	}
	return nil
}

func init() {
	proto.RegisterType((*Experimental)(nil), "google.api.Experimental")
}

func init() {
	proto.RegisterFile("google/api/experimental/experimental.proto", fileDescriptorExperimental)
}

var fileDescriptorExperimental = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4a, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xad, 0x28, 0x48, 0x2d, 0xca, 0xcc, 0x4d, 0xcd,
	0x2b, 0x49, 0xcc, 0x41, 0xe1, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x71, 0x41, 0xd4, 0xea,
	0x25, 0x16, 0x64, 0x4a, 0x19, 0xe1, 0xd2, 0x97, 0x58, 0x5a, 0x92, 0x91, 0x5f, 0x94, 0x59, 0x95,
	0x58, 0x92, 0x99, 0x9f, 0x17, 0x9f, 0x9c, 0x9f, 0x97, 0x96, 0x99, 0x0e, 0xd1, 0xaf, 0x14, 0xca,
	0xc5, 0xe3, 0x8a, 0xa4, 0x54, 0xc8, 0x95, 0x8b, 0x17, 0x45, 0xb5, 0x04, 0x87, 0x02, 0xa3, 0x06,
	0xb7, 0x91, 0xbc, 0x1e, 0xc2, 0x1e, 0x3d, 0x47, 0x64, 0x05, 0xce, 0x60, 0xd3, 0x82, 0x50, 0x75,
	0x39, 0xe9, 0x70, 0xf1, 0x25, 0xe7, 0xe7, 0x22, 0x69, 0x72, 0x12, 0x44, 0xb6, 0x26, 0x00, 0x64,
	0x77, 0x00, 0xe3, 0x22, 0x26, 0x16, 0x77, 0xc7, 0x00, 0xcf, 0x24, 0x36, 0xb0, 0x5b, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x54, 0xd7, 0xf1, 0xa1, 0xf9, 0x00, 0x00, 0x00,
}
