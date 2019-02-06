// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/logging/v2/logging_metrics.proto

package google_logging_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/api"
import google_api4 "go.pedge.io/pb/go/google/api"
import google_api5 "go.pedge.io/pb/go/google/api"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Stackdriver Logging API version.
type LogMetric_ApiVersion int32

const (
	// Stackdriver Logging API v2.
	LogMetric_V2 LogMetric_ApiVersion = 0
	// Stackdriver Logging API v1.
	LogMetric_V1 LogMetric_ApiVersion = 1
)

var LogMetric_ApiVersion_name = map[int32]string{
	0: "V2",
	1: "V1",
}
var LogMetric_ApiVersion_value = map[string]int32{
	"V2": 0,
	"V1": 1,
}

func (x LogMetric_ApiVersion) String() string {
	return proto.EnumName(LogMetric_ApiVersion_name, int32(x))
}
func (LogMetric_ApiVersion) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

// Describes a logs-based metric.  The value of the metric is the
// number of log entries that match a logs filter in a given time interval.
//
// Logs-based metric can also be used to extract values from logs and create a
// a distribution of the values. The distribution records the statistics of the
// extracted values along with an optional histogram of the values as specified
// by the bucket options.
type LogMetric struct {
	// Required. The client-assigned metric identifier.
	// Examples: `"error_count"`, `"nginx/requests"`.
	//
	// Metric identifiers are limited to 100 characters and can include
	// only the following characters: `A-Z`, `a-z`, `0-9`, and the
	// special characters `_-.,+!*',()%/`.  The forward-slash character
	// (`/`) denotes a hierarchy of name pieces, and it cannot be the
	// first character of the name.
	//
	// The metric identifier in this field must not be
	// [URL-encoded](https://en.wikipedia.org/wiki/Percent-encoding).
	// However, when the metric identifier appears as the `[METRIC_ID]`
	// part of a `metric_name` API parameter, then the metric identifier
	// must be URL-encoded. Example:
	// `"projects/my-project/metrics/nginx%2Frequests"`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Optional. A description of this metric, which is used in documentation.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Required. An [advanced logs filter](/logging/docs/view/advanced_filters)
	// which is used to match log entries.
	// Example:
	//
	//     "resource.type=gae_app AND severity>=ERROR"
	//
	// The maximum length of the filter is 20000 characters.
	Filter string `protobuf:"bytes,3,opt,name=filter" json:"filter,omitempty"`
	// Optional. The metric descriptor associated with the logs-based metric.
	// If unspecified, it uses a default metric descriptor with a DELTA metric
	// kind, INT64 value type, with no labels and a unit of "1". Such a metric
	// counts the number of log entries matching the `filter` expression.
	//
	// The `name`, `type`, and `description` fields in the `metric_descriptor`
	// are output only, and is constructed using the `name` and `description`
	// field in the LogMetric.
	//
	// To create a logs-based metric that records a distribution of log values, a
	// DELTA metric kind with a DISTRIBUTION value type must be used along with
	// a `value_extractor` expression in the LogMetric.
	//
	// Each label in the metric descriptor must have a matching label
	// name as the key and an extractor expression as the value in the
	// `label_extractors` map.
	//
	// The `metric_kind` and `value_type` fields in the `metric_descriptor` cannot
	// be updated once initially configured. New labels can be added in the
	// `metric_descriptor`, but existing labels cannot be modified except for
	// their description.
	MetricDescriptor *google_api5.MetricDescriptor `protobuf:"bytes,5,opt,name=metric_descriptor,json=metricDescriptor" json:"metric_descriptor,omitempty"`
	// Optional. A `value_extractor` is required when using a distribution
	// logs-based metric to extract the values to record from a log entry.
	// Two functions are supported for value extraction: `EXTRACT(field)` or
	// `REGEXP_EXTRACT(field, regex)`. The argument are:
	//   1. field: The name of the log entry field from which the value is to be
	//      extracted.
	//   2. regex: A regular expression using the Google RE2 syntax
	//      (https://github.com/google/re2/wiki/Syntax) with a single capture
	//      group to extract data from the specified log entry field. The value
	//      of the field is converted to a string before applying the regex.
	//      It is an error to specify a regex that does not include exactly one
	//      capture group.
	//
	// The result of the extraction must be convertible to a double type, as the
	// distribution always records double values. If either the extraction or
	// the conversion to double fails, then those values are not recorded in the
	// distribution.
	//
	// Example: `REGEXP_EXTRACT(jsonPayload.request, ".*quantity=(\d+).*")`
	ValueExtractor string `protobuf:"bytes,6,opt,name=value_extractor,json=valueExtractor" json:"value_extractor,omitempty"`
	// Optional. A map from a label key string to an extractor expression which is
	// used to extract data from a log entry field and assign as the label value.
	// Each label key specified in the LabelDescriptor must have an associated
	// extractor expression in this map. The syntax of the extractor expression
	// is the same as for the `value_extractor` field.
	//
	// The extracted value is converted to the type defined in the label
	// descriptor. If the either the extraction or the type conversion fails,
	// the label will have a default value. The default value for a string
	// label is an empty string, for an integer label its 0, and for a boolean
	// label its `false`.
	//
	// Note that there are upper bounds on the maximum number of labels and the
	// number of active time series that are allowed in a project.
	LabelExtractors map[string]string `protobuf:"bytes,7,rep,name=label_extractors,json=labelExtractors" json:"label_extractors,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Optional. The `bucket_options` are required when the logs-based metric is
	// using a DISTRIBUTION value type and it describes the bucket boundaries
	// used to create a histogram of the extracted values.
	BucketOptions *google_api4.Distribution_BucketOptions `protobuf:"bytes,8,opt,name=bucket_options,json=bucketOptions" json:"bucket_options,omitempty"`
	// Deprecated. The API version that created or updated this metric.
	// The v2 format is used by default and cannot be changed.
	Version LogMetric_ApiVersion `protobuf:"varint,4,opt,name=version,enum=google.logging.v2.LogMetric_ApiVersion" json:"version,omitempty"`
}

func (m *LogMetric) Reset()                    { *m = LogMetric{} }
func (m *LogMetric) String() string            { return proto.CompactTextString(m) }
func (*LogMetric) ProtoMessage()               {}
func (*LogMetric) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *LogMetric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogMetric) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LogMetric) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *LogMetric) GetMetricDescriptor() *google_api5.MetricDescriptor {
	if m != nil {
		return m.MetricDescriptor
	}
	return nil
}

func (m *LogMetric) GetValueExtractor() string {
	if m != nil {
		return m.ValueExtractor
	}
	return ""
}

func (m *LogMetric) GetLabelExtractors() map[string]string {
	if m != nil {
		return m.LabelExtractors
	}
	return nil
}

func (m *LogMetric) GetBucketOptions() *google_api4.Distribution_BucketOptions {
	if m != nil {
		return m.BucketOptions
	}
	return nil
}

func (m *LogMetric) GetVersion() LogMetric_ApiVersion {
	if m != nil {
		return m.Version
	}
	return LogMetric_V2
}

// The parameters to ListLogMetrics.
type ListLogMetricsRequest struct {
	// Required. The name of the project containing the metrics:
	//
	//     "projects/[PROJECT_ID]"
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Optional. If present, then retrieve the next batch of results from the
	// preceding call to this method.  `pageToken` must be the value of
	// `nextPageToken` from the previous response.  The values of other method
	// parameters should be identical to those in the previous call.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// Optional. The maximum number of results to return from this request.
	// Non-positive values are ignored.  The presence of `nextPageToken` in the
	// response indicates that more results might be available.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListLogMetricsRequest) Reset()                    { *m = ListLogMetricsRequest{} }
func (m *ListLogMetricsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLogMetricsRequest) ProtoMessage()               {}
func (*ListLogMetricsRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ListLogMetricsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListLogMetricsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListLogMetricsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// Result returned from ListLogMetrics.
type ListLogMetricsResponse struct {
	// A list of logs-based metrics.
	Metrics []*LogMetric `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
	// If there might be more results than appear in this response, then
	// `nextPageToken` is included.  To get the next set of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListLogMetricsResponse) Reset()                    { *m = ListLogMetricsResponse{} }
func (m *ListLogMetricsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLogMetricsResponse) ProtoMessage()               {}
func (*ListLogMetricsResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *ListLogMetricsResponse) GetMetrics() []*LogMetric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *ListLogMetricsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The parameters to GetLogMetric.
type GetLogMetricRequest struct {
	// The resource name of the desired metric:
	//
	//     "projects/[PROJECT_ID]/metrics/[METRIC_ID]"
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
}

func (m *GetLogMetricRequest) Reset()                    { *m = GetLogMetricRequest{} }
func (m *GetLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLogMetricRequest) ProtoMessage()               {}
func (*GetLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *GetLogMetricRequest) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

// The parameters to CreateLogMetric.
type CreateLogMetricRequest struct {
	// The resource name of the project in which to create the metric:
	//
	//     "projects/[PROJECT_ID]"
	//
	// The new metric must be provided in the request.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// The new logs-based metric, which must not have an identifier that
	// already exists.
	Metric *LogMetric `protobuf:"bytes,2,opt,name=metric" json:"metric,omitempty"`
}

func (m *CreateLogMetricRequest) Reset()                    { *m = CreateLogMetricRequest{} }
func (m *CreateLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateLogMetricRequest) ProtoMessage()               {}
func (*CreateLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *CreateLogMetricRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateLogMetricRequest) GetMetric() *LogMetric {
	if m != nil {
		return m.Metric
	}
	return nil
}

// The parameters to UpdateLogMetric.
type UpdateLogMetricRequest struct {
	// The resource name of the metric to update:
	//
	//     "projects/[PROJECT_ID]/metrics/[METRIC_ID]"
	//
	// The updated metric must be provided in the request and it's
	// `name` field must be the same as `[METRIC_ID]` If the metric
	// does not exist in `[PROJECT_ID]`, then a new metric is created.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
	// The updated metric.
	Metric *LogMetric `protobuf:"bytes,2,opt,name=metric" json:"metric,omitempty"`
}

func (m *UpdateLogMetricRequest) Reset()                    { *m = UpdateLogMetricRequest{} }
func (m *UpdateLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateLogMetricRequest) ProtoMessage()               {}
func (*UpdateLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *UpdateLogMetricRequest) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

func (m *UpdateLogMetricRequest) GetMetric() *LogMetric {
	if m != nil {
		return m.Metric
	}
	return nil
}

// The parameters to DeleteLogMetric.
type DeleteLogMetricRequest struct {
	// The resource name of the metric to delete:
	//
	//     "projects/[PROJECT_ID]/metrics/[METRIC_ID]"
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
}

func (m *DeleteLogMetricRequest) Reset()                    { *m = DeleteLogMetricRequest{} }
func (m *DeleteLogMetricRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteLogMetricRequest) ProtoMessage()               {}
func (*DeleteLogMetricRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *DeleteLogMetricRequest) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

func init() {
	proto.RegisterType((*LogMetric)(nil), "google.logging.v2.LogMetric")
	proto.RegisterType((*ListLogMetricsRequest)(nil), "google.logging.v2.ListLogMetricsRequest")
	proto.RegisterType((*ListLogMetricsResponse)(nil), "google.logging.v2.ListLogMetricsResponse")
	proto.RegisterType((*GetLogMetricRequest)(nil), "google.logging.v2.GetLogMetricRequest")
	proto.RegisterType((*CreateLogMetricRequest)(nil), "google.logging.v2.CreateLogMetricRequest")
	proto.RegisterType((*UpdateLogMetricRequest)(nil), "google.logging.v2.UpdateLogMetricRequest")
	proto.RegisterType((*DeleteLogMetricRequest)(nil), "google.logging.v2.DeleteLogMetricRequest")
	proto.RegisterEnum("google.logging.v2.LogMetric_ApiVersion", LogMetric_ApiVersion_name, LogMetric_ApiVersion_value)
}

func init() { proto.RegisterFile("google/logging/v2/logging_metrics.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 812 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xc1, 0x4e, 0xeb, 0x46,
	0x14, 0xad, 0x03, 0x09, 0x8f, 0x9b, 0xbe, 0x24, 0x0c, 0x10, 0xa2, 0x10, 0x44, 0xea, 0x45, 0x12,
	0x58, 0xd8, 0xc5, 0xad, 0x50, 0x8b, 0xd4, 0x05, 0x01, 0x84, 0x2a, 0x85, 0x16, 0x99, 0x36, 0xab,
	0x48, 0x96, 0x93, 0xdc, 0x44, 0x6e, 0x1c, 0x8f, 0x6b, 0x4f, 0x22, 0xa0, 0x62, 0x53, 0xb1, 0xab,
	0xd4, 0x45, 0xfb, 0x01, 0xdd, 0xf7, 0x53, 0xba, 0x6d, 0x3f, 0xa1, 0x1f, 0xd1, 0x65, 0xe5, 0xf1,
	0x38, 0x98, 0xc4, 0x8f, 0x20, 0x56, 0xcc, 0xdc, 0x73, 0x66, 0xce, 0x99, 0x7b, 0x4f, 0x30, 0xd4,
	0x87, 0x94, 0x0e, 0x6d, 0x54, 0x6d, 0x3a, 0x1c, 0x5a, 0xce, 0x50, 0x9d, 0x6a, 0xd1, 0xd2, 0x18,
	0x23, 0xf3, 0xac, 0x9e, 0xaf, 0xb8, 0x1e, 0x65, 0x94, 0x6c, 0x84, 0x44, 0x45, 0xa0, 0xca, 0x54,
	0x2b, 0x57, 0xc4, 0x59, 0xd3, 0xb5, 0x54, 0xd3, 0x71, 0x28, 0x33, 0x99, 0x45, 0x1d, 0x71, 0xa0,
	0xbc, 0x17, 0x43, 0xfb, 0x96, 0xcf, 0x3c, 0xab, 0x3b, 0x09, 0x70, 0x01, 0xef, 0xc4, 0xe0, 0x50,
	0x49, 0x00, 0xbb, 0x02, 0xe0, 0xbb, 0xee, 0x64, 0xa0, 0xe2, 0xd8, 0x65, 0x77, 0x02, 0xac, 0xce,
	0x83, 0x03, 0x0b, 0xed, 0xbe, 0x31, 0x36, 0xfd, 0x51, 0xc8, 0x90, 0xff, 0x58, 0x85, 0xf5, 0x16,
	0x1d, 0x5e, 0xf1, 0x2b, 0x09, 0x81, 0x55, 0xc7, 0x1c, 0x63, 0x49, 0xaa, 0x4a, 0x8d, 0x75, 0x9d,
	0xaf, 0x49, 0x15, 0xb2, 0x7d, 0xf4, 0x7b, 0x9e, 0xe5, 0x06, 0x76, 0x4a, 0x29, 0x0e, 0xc5, 0x4b,
	0xa4, 0x08, 0x99, 0x81, 0x65, 0x33, 0xf4, 0x4a, 0x2b, 0x1c, 0x14, 0x3b, 0xf2, 0x35, 0x6c, 0x84,
	0x56, 0x8d, 0x88, 0x4d, 0xbd, 0x52, 0xba, 0x2a, 0x35, 0xb2, 0x5a, 0x45, 0x11, 0xfd, 0x31, 0x5d,
	0x4b, 0x09, 0xc5, 0xcf, 0x67, 0x1c, 0xbd, 0x30, 0x9e, 0xab, 0x90, 0x3a, 0xe4, 0xa7, 0xa6, 0x3d,
	0x41, 0x03, 0x6f, 0x99, 0x67, 0xf6, 0x82, 0x8b, 0x32, 0x5c, 0x2b, 0xc7, 0xcb, 0x17, 0x51, 0x95,
	0x74, 0xa0, 0x60, 0x9b, 0x5d, 0xb4, 0x9f, 0x88, 0x7e, 0x69, 0xad, 0xba, 0xd2, 0xc8, 0x6a, 0x47,
	0xca, 0xc2, 0x48, 0x94, 0xd9, 0xcb, 0x95, 0x56, 0x70, 0x68, 0x76, 0x8d, 0x7f, 0xe1, 0x30, 0xef,
	0x4e, 0xcf, 0xdb, 0xcf, 0xab, 0xe4, 0x0a, 0x72, 0xdd, 0x49, 0x6f, 0x84, 0xcc, 0xa0, 0xfc, 0xe9,
	0x7e, 0xe9, 0x1d, 0x7f, 0x4e, 0x2d, 0xfe, 0x9c, 0xf3, 0xf8, 0xf4, 0x9a, 0x9c, 0xfe, 0x6d, 0xc8,
	0xd6, 0xdf, 0x77, 0xe3, 0x5b, 0x72, 0x0a, 0x6b, 0x53, 0xf4, 0xfc, 0xa0, 0xad, 0xab, 0x55, 0xa9,
	0x91, 0xd3, 0xea, 0x2f, 0x7a, 0x3c, 0x75, 0xad, 0x76, 0x48, 0xd7, 0xa3, 0x73, 0xe5, 0x26, 0x6c,
	0x25, 0x59, 0x27, 0x05, 0x58, 0x19, 0xe1, 0x9d, 0x18, 0x64, 0xb0, 0x24, 0x5b, 0x90, 0xe6, 0xbd,
	0x12, 0x13, 0x0c, 0x37, 0x27, 0xa9, 0x2f, 0x24, 0xb9, 0x02, 0xf0, 0x74, 0x35, 0xc9, 0x40, 0xaa,
	0xad, 0x15, 0x3e, 0xe2, 0x7f, 0x8f, 0x0a, 0x92, 0x3c, 0x82, 0xed, 0x96, 0xe5, 0xb3, 0x99, 0x0d,
	0x5f, 0xc7, 0x1f, 0x27, 0xe8, 0xb3, 0x60, 0xec, 0xae, 0xe9, 0xa1, 0xc3, 0x84, 0x8a, 0xd8, 0x91,
	0x3d, 0x00, 0xd7, 0x1c, 0xa2, 0xc1, 0xe8, 0x08, 0xa3, 0xbc, 0xac, 0x07, 0x95, 0xef, 0x82, 0x02,
	0xd9, 0x05, 0xbe, 0x31, 0x7c, 0xeb, 0x1e, 0x79, 0x60, 0xd2, 0xfa, 0xbb, 0xa0, 0x70, 0x63, 0xdd,
	0xa3, 0x7c, 0x0b, 0xc5, 0x79, 0x31, 0xdf, 0xa5, 0x8e, 0x8f, 0xe4, 0x18, 0xd6, 0xc4, 0x2f, 0xac,
	0x24, 0xf1, 0x79, 0x56, 0x5e, 0xea, 0x95, 0x1e, 0x91, 0x49, 0x0d, 0xf2, 0x0e, 0xde, 0x32, 0x63,
	0xc1, 0xd2, 0xfb, 0xa0, 0x7c, 0x1d, 0xd9, 0x92, 0x8f, 0x61, 0xf3, 0x12, 0x9f, 0x84, 0xa3, 0x47,
	0xee, 0x43, 0x56, 0x64, 0x38, 0xf6, 0xc3, 0x80, 0xb0, 0xf4, 0x8d, 0x39, 0x46, 0x79, 0x00, 0xc5,
	0x33, 0x0f, 0x4d, 0x86, 0x0b, 0x47, 0x3f, 0xd4, 0x9f, 0xcf, 0x21, 0x13, 0x9e, 0xe7, 0x46, 0x96,
	0x3d, 0x44, 0x70, 0x65, 0x0a, 0xc5, 0xef, 0xdd, 0x7e, 0x92, 0xce, 0x32, 0x8b, 0x6f, 0x14, 0xfc,
	0x12, 0x8a, 0xe7, 0x68, 0xe3, 0x1b, 0x04, 0xb5, 0x7f, 0xd2, 0x50, 0x10, 0xf3, 0xbb, 0x41, 0x6f,
	0x6a, 0xf5, 0xb0, 0xad, 0x91, 0x5f, 0x25, 0xc8, 0x3d, 0x9f, 0x2d, 0x69, 0x24, 0x19, 0x49, 0xca,
	0x5a, 0xf9, 0xe0, 0x15, 0xcc, 0x30, 0x28, 0x72, 0xfd, 0xe7, 0xbf, 0xff, 0xfd, 0x3d, 0xf5, 0x09,
	0xd9, 0x0f, 0xfe, 0x39, 0xff, 0x14, 0xf6, 0xfc, 0x2b, 0xd7, 0xa3, 0x3f, 0x60, 0x8f, 0xf9, 0xea,
	0xe1, 0x83, 0x1a, 0x25, 0xe3, 0x51, 0x82, 0x8f, 0xe3, 0x23, 0x27, 0xb5, 0x04, 0x91, 0x84, 0x4c,
	0x94, 0x5f, 0xec, 0x9f, 0xac, 0x70, 0xfd, 0x06, 0xa9, 0x71, 0xfd, 0x58, 0xa3, 0x62, 0x26, 0x22,
	0x0f, 0xea, 0xe1, 0x03, 0xf9, 0x45, 0x82, 0xfc, 0x5c, 0x82, 0x48, 0xd2, 0x73, 0x93, 0x53, 0xb6,
	0xc4, 0x8c, 0xca, 0xcd, 0x1c, 0xc8, 0xcb, 0x9a, 0x71, 0x22, 0xa6, 0x4e, 0x7e, 0x93, 0x20, 0x3f,
	0x97, 0xb3, 0x44, 0x37, 0xc9, 0x59, 0x5c, 0xe2, 0xe6, 0x98, 0xbb, 0xf9, 0xb4, 0xfc, 0xca, 0xd6,
	0xcc, 0x4c, 0x3d, 0x4a, 0x90, 0x9f, 0xcb, 0x62, 0xa2, 0xa9, 0xe4, 0xbc, 0x96, 0x8b, 0x11, 0x35,
	0xfa, 0x0c, 0x2a, 0x17, 0xc1, 0x37, 0x32, 0x9a, 0xd4, 0xe1, 0x2b, 0xed, 0x34, 0x11, 0xb6, 0x7b,
	0x74, 0xbc, 0xa8, 0xdb, 0xdc, 0x6c, 0x85, 0x6b, 0x11, 0xc5, 0xeb, 0x40, 0xe6, 0x5a, 0xfa, 0x4f,
	0x92, 0xfe, 0x4c, 0xed, 0x5c, 0x86, 0xec, 0x33, 0x9b, 0x4e, 0xfa, 0x8a, 0xe0, 0x29, 0x6d, 0xed,
	0xaf, 0x08, 0xe9, 0x70, 0xa4, 0x23, 0x90, 0x4e, 0x5b, 0xeb, 0x66, 0xb8, 0xcd, 0xcf, 0xfe, 0x0f,
	0x00, 0x00, 0xff, 0xff, 0xc4, 0x31, 0x35, 0x4d, 0x6c, 0x08, 0x00, 0x00,
}
