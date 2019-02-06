// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/billing.proto

package google_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Billing related configuration of the service.
//
// The following example shows how to configure monitored resources and metrics
// for billing:
//     monitored_resources:
//     - type: library.googleapis.com/branch
//       labels:
//       - key: /city
//         description: The city where the library branch is located in.
//       - key: /name
//         description: The name of the branch.
//     metrics:
//     - name: library.googleapis.com/book/borrowed_count
//       metric_kind: DELTA
//       value_type: INT64
//     billing:
//       consumer_destinations:
//       - monitored_resource: library.googleapis.com/branch
//         metrics:
//         - library.googleapis.com/book/borrowed_count
type Billing struct {
	// Billing configurations for sending metrics to the consumer project.
	// There can be multiple consumer destinations per service, each one must have
	// a different monitored resource type. A metric can be used in at most
	// one consumer destination.
	ConsumerDestinations []*Billing_BillingDestination `protobuf:"bytes,8,rep,name=consumer_destinations,json=consumerDestinations" json:"consumer_destinations,omitempty"`
}

func (m *Billing) Reset()                    { *m = Billing{} }
func (m *Billing) String() string            { return proto.CompactTextString(m) }
func (*Billing) ProtoMessage()               {}
func (*Billing) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Billing) GetConsumerDestinations() []*Billing_BillingDestination {
	if m != nil {
		return m.ConsumerDestinations
	}
	return nil
}

// Configuration of a specific billing destination (Currently only support
// bill against consumer project).
type Billing_BillingDestination struct {
	// The monitored resource type. The type must be defined in
	// [Service.monitored_resources][google.api.Service.monitored_resources] section.
	MonitoredResource string `protobuf:"bytes,1,opt,name=monitored_resource,json=monitoredResource" json:"monitored_resource,omitempty"`
	// Names of the metrics to report to this billing destination.
	// Each name must be defined in [Service.metrics][google.api.Service.metrics] section.
	Metrics []string `protobuf:"bytes,2,rep,name=metrics" json:"metrics,omitempty"`
}

func (m *Billing_BillingDestination) Reset()                    { *m = Billing_BillingDestination{} }
func (m *Billing_BillingDestination) String() string            { return proto.CompactTextString(m) }
func (*Billing_BillingDestination) ProtoMessage()               {}
func (*Billing_BillingDestination) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

func (m *Billing_BillingDestination) GetMonitoredResource() string {
	if m != nil {
		return m.MonitoredResource
	}
	return ""
}

func (m *Billing_BillingDestination) GetMetrics() []string {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func init() {
	proto.RegisterType((*Billing)(nil), "google.api.Billing")
	proto.RegisterType((*Billing_BillingDestination)(nil), "google.api.Billing.BillingDestination")
}

func init() { proto.RegisterFile("google/api/billing.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xca, 0xcc, 0xc9, 0xc9, 0xcc, 0x4b, 0xd7, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0xc8, 0xe8, 0x25, 0x16, 0x64, 0x4a, 0xc9, 0x20, 0xa9,
	0x4a, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xa8, 0x94, 0x12, 0x47,
	0x92, 0xcd, 0x4d, 0x2d, 0x29, 0xca, 0x4c, 0x86, 0x48, 0x28, 0x1d, 0x65, 0xe4, 0x62, 0x77, 0x82,
	0x18, 0x2a, 0x14, 0xcd, 0x25, 0x9a, 0x9c, 0x9f, 0x57, 0x5c, 0x9a, 0x9b, 0x5a, 0x14, 0x9f, 0x92,
	0x5a, 0x5c, 0x92, 0x99, 0x07, 0x31, 0x43, 0x82, 0x43, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x4d, 0x0f,
	0x61, 0x9d, 0x1e, 0x54, 0x0f, 0x8c, 0x76, 0x41, 0x28, 0x0f, 0x12, 0x81, 0x19, 0x82, 0x24, 0x58,
	0x2c, 0x15, 0xcb, 0x25, 0x84, 0xa9, 0x56, 0x48, 0x97, 0x4b, 0x28, 0x37, 0x3f, 0x2f, 0xb3, 0x24,
	0xbf, 0x28, 0x35, 0x25, 0xbe, 0x28, 0xb5, 0x38, 0xbf, 0xb4, 0x28, 0x39, 0x55, 0x82, 0x51, 0x81,
	0x51, 0x83, 0x33, 0x48, 0x10, 0x2e, 0x13, 0x04, 0x95, 0x10, 0x92, 0xe0, 0x62, 0x87, 0xb8, 0xbe,
	0x58, 0x82, 0x49, 0x81, 0x59, 0x83, 0x33, 0x08, 0xc6, 0x75, 0x52, 0xe7, 0xe2, 0x4b, 0xce, 0xcf,
	0x45, 0x72, 0xa1, 0x13, 0x0f, 0xd4, 0xba, 0x00, 0x90, 0x3f, 0x03, 0x18, 0x17, 0x31, 0xb1, 0xb8,
	0x3b, 0x06, 0x78, 0x26, 0xb1, 0x81, 0xfd, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xaf,
	0x8f, 0xe9, 0x56, 0x01, 0x00, 0x00,
}
