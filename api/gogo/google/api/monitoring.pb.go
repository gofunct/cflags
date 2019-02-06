// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/api/monitoring.proto

package google_api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Monitoring configuration of the service.
//
// The example below shows how to configure monitored resources and metrics
// for monitoring. In the example, a monitored resource and two metrics are
// defined. The `library.googleapis.com/book/returned_count` metric is sent
// to both producer and consumer projects, whereas the
// `library.googleapis.com/book/overdue_count` metric is only sent to the
// consumer project.
//
//     monitored_resources:
//     - type: library.googleapis.com/branch
//       labels:
//       - key: /city
//         description: The city where the library branch is located in.
//       - key: /name
//         description: The name of the branch.
//     metrics:
//     - name: library.googleapis.com/book/returned_count
//       metric_kind: DELTA
//       value_type: INT64
//       labels:
//       - key: /customer_id
//     - name: library.googleapis.com/book/overdue_count
//       metric_kind: GAUGE
//       value_type: INT64
//       labels:
//       - key: /customer_id
//     monitoring:
//       producer_destinations:
//       - monitored_resource: library.googleapis.com/branch
//         metrics:
//         - library.googleapis.com/book/returned_count
//       consumer_destinations:
//       - monitored_resource: library.googleapis.com/branch
//         metrics:
//         - library.googleapis.com/book/returned_count
//         - library.googleapis.com/book/overdue_count
type Monitoring struct {
	// Monitoring configurations for sending metrics to the producer project.
	// There can be multiple producer destinations, each one must have a
	// different monitored resource type. A metric can be used in at most
	// one producer destination.
	ProducerDestinations []*Monitoring_MonitoringDestination `protobuf:"bytes,1,rep,name=producer_destinations,json=producerDestinations" json:"producer_destinations,omitempty"`
	// Monitoring configurations for sending metrics to the consumer project.
	// There can be multiple consumer destinations, each one must have a
	// different monitored resource type. A metric can be used in at most
	// one consumer destination.
	ConsumerDestinations []*Monitoring_MonitoringDestination `protobuf:"bytes,2,rep,name=consumer_destinations,json=consumerDestinations" json:"consumer_destinations,omitempty"`
}

func (m *Monitoring) Reset()                    { *m = Monitoring{} }
func (m *Monitoring) String() string            { return proto.CompactTextString(m) }
func (*Monitoring) ProtoMessage()               {}
func (*Monitoring) Descriptor() ([]byte, []int) { return fileDescriptorMonitoring, []int{0} }

func (m *Monitoring) GetProducerDestinations() []*Monitoring_MonitoringDestination {
	if m != nil {
		return m.ProducerDestinations
	}
	return nil
}

func (m *Monitoring) GetConsumerDestinations() []*Monitoring_MonitoringDestination {
	if m != nil {
		return m.ConsumerDestinations
	}
	return nil
}

// Configuration of a specific monitoring destination (the producer project
// or the consumer project).
type Monitoring_MonitoringDestination struct {
	// The monitored resource type. The type must be defined in
	// [Service.monitored_resources][google.api.Service.monitored_resources] section.
	MonitoredResource string `protobuf:"bytes,1,opt,name=monitored_resource,json=monitoredResource,proto3" json:"monitored_resource,omitempty"`
	// Names of the metrics to report to this monitoring destination.
	// Each name must be defined in [Service.metrics][google.api.Service.metrics] section.
	Metrics []string `protobuf:"bytes,2,rep,name=metrics" json:"metrics,omitempty"`
}

func (m *Monitoring_MonitoringDestination) Reset()         { *m = Monitoring_MonitoringDestination{} }
func (m *Monitoring_MonitoringDestination) String() string { return proto.CompactTextString(m) }
func (*Monitoring_MonitoringDestination) ProtoMessage()    {}
func (*Monitoring_MonitoringDestination) Descriptor() ([]byte, []int) {
	return fileDescriptorMonitoring, []int{0, 0}
}

func (m *Monitoring_MonitoringDestination) GetMonitoredResource() string {
	if m != nil {
		return m.MonitoredResource
	}
	return ""
}

func (m *Monitoring_MonitoringDestination) GetMetrics() []string {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func init() {
	proto.RegisterType((*Monitoring)(nil), "google.api.Monitoring")
	proto.RegisterType((*Monitoring_MonitoringDestination)(nil), "google.api.Monitoring.MonitoringDestination")
}

func init() { proto.RegisterFile("google/api/monitoring.proto", fileDescriptorMonitoring) }

var fileDescriptorMonitoring = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0xcd, 0xcf, 0xcb, 0x2c, 0xc9, 0x2f, 0xca, 0xcc,
	0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0x48, 0xea, 0x25, 0x16, 0x64, 0x4a,
	0xc9, 0x20, 0x29, 0x4c, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xa8,
	0x54, 0xda, 0xc2, 0xc4, 0xc5, 0xe5, 0x0b, 0xd7, 0x2e, 0x94, 0xc8, 0x25, 0x5a, 0x50, 0x94, 0x9f,
	0x52, 0x9a, 0x9c, 0x5a, 0x14, 0x9f, 0x92, 0x5a, 0x5c, 0x92, 0x99, 0x07, 0x51, 0x2d, 0xc1, 0xa8,
	0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xa3, 0x87, 0x30, 0x58, 0x0f, 0xa1, 0x0d, 0x89, 0xe9, 0x82, 0xd0,
	0x14, 0x24, 0x02, 0x33, 0x0a, 0x49, 0xb0, 0x18, 0x64, 0x45, 0x72, 0x7e, 0x5e, 0x71, 0x69, 0x2e,
	0xba, 0x15, 0x4c, 0xe4, 0x58, 0x01, 0x33, 0x0a, 0xd9, 0x0a, 0xa9, 0x04, 0x2e, 0x51, 0xac, 0xca,
	0x85, 0x74, 0xb9, 0x84, 0xa0, 0x61, 0x95, 0x9a, 0x12, 0x5f, 0x94, 0x5a, 0x9c, 0x5f, 0x5a, 0x94,
	0x9c, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x24, 0x08, 0x97, 0x09, 0x82, 0x4a, 0x08, 0x49,
	0x70, 0xb1, 0xe7, 0xa6, 0x96, 0x14, 0x65, 0x26, 0x43, 0x1c, 0xc7, 0x19, 0x04, 0xe3, 0x3a, 0x69,
	0x71, 0xf1, 0x25, 0xe7, 0xe7, 0x22, 0x39, 0xd5, 0x89, 0x1f, 0x61, 0x63, 0x00, 0x28, 0x64, 0x03,
	0x18, 0x17, 0x31, 0xb1, 0xb8, 0x3b, 0x06, 0x78, 0x26, 0xb1, 0x81, 0x43, 0xda, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x24, 0x24, 0x13, 0x98, 0xb2, 0x01, 0x00, 0x00,
}
