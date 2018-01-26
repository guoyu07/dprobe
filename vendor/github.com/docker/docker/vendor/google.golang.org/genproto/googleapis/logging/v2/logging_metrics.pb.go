// Code generated by protoc-gen-go.
// source: google/logging/v2/logging_metrics.proto
// DO NOT EDIT!

package logging

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "google.golang.org/genproto/googleapis/api/distribution"
import _ "google.golang.org/genproto/googleapis/api/metric"
import google_protobuf5 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

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

var LogMetric_ApiVersion_name = map[int32]string***REMOVED***
	0: "V2",
	1: "V1",
***REMOVED***
var LogMetric_ApiVersion_value = map[string]int32***REMOVED***
	"V2": 0,
	"V1": 1,
***REMOVED***

func (x LogMetric_ApiVersion) String() string ***REMOVED***
	return proto.EnumName(LogMetric_ApiVersion_name, int32(x))
***REMOVED***
func (LogMetric_ApiVersion) EnumDescriptor() ([]byte, []int) ***REMOVED*** return fileDescriptor3, []int***REMOVED***0, 0***REMOVED*** ***REMOVED***

// Describes a logs-based metric.  The value of the metric is the
// number of log entries that match a logs filter in a given time interval.
type LogMetric struct ***REMOVED***
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
	// Output only. The API version that created or updated this metric.
	// The version also dictates the syntax of the filter expression. When a value
	// for this field is missing, the default value of V2 should be assumed.
	Version LogMetric_ApiVersion `protobuf:"varint,4,opt,name=version,enum=google.logging.v2.LogMetric_ApiVersion" json:"version,omitempty"`
***REMOVED***

func (m *LogMetric) Reset()                    ***REMOVED*** *m = LogMetric***REMOVED******REMOVED*** ***REMOVED***
func (m *LogMetric) String() string            ***REMOVED*** return proto.CompactTextString(m) ***REMOVED***
func (*LogMetric) ProtoMessage()               ***REMOVED******REMOVED***
func (*LogMetric) Descriptor() ([]byte, []int) ***REMOVED*** return fileDescriptor3, []int***REMOVED***0***REMOVED*** ***REMOVED***

func (m *LogMetric) GetName() string ***REMOVED***
	if m != nil ***REMOVED***
		return m.Name
	***REMOVED***
	return ""
***REMOVED***

func (m *LogMetric) GetDescription() string ***REMOVED***
	if m != nil ***REMOVED***
		return m.Description
	***REMOVED***
	return ""
***REMOVED***

func (m *LogMetric) GetFilter() string ***REMOVED***
	if m != nil ***REMOVED***
		return m.Filter
	***REMOVED***
	return ""
***REMOVED***

func (m *LogMetric) GetVersion() LogMetric_ApiVersion ***REMOVED***
	if m != nil ***REMOVED***
		return m.Version
	***REMOVED***
	return LogMetric_V2
***REMOVED***

// The parameters to ListLogMetrics.
type ListLogMetricsRequest struct ***REMOVED***
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
***REMOVED***

func (m *ListLogMetricsRequest) Reset()                    ***REMOVED*** *m = ListLogMetricsRequest***REMOVED******REMOVED*** ***REMOVED***
func (m *ListLogMetricsRequest) String() string            ***REMOVED*** return proto.CompactTextString(m) ***REMOVED***
func (*ListLogMetricsRequest) ProtoMessage()               ***REMOVED******REMOVED***
func (*ListLogMetricsRequest) Descriptor() ([]byte, []int) ***REMOVED*** return fileDescriptor3, []int***REMOVED***1***REMOVED*** ***REMOVED***

func (m *ListLogMetricsRequest) GetParent() string ***REMOVED***
	if m != nil ***REMOVED***
		return m.Parent
	***REMOVED***
	return ""
***REMOVED***

func (m *ListLogMetricsRequest) GetPageToken() string ***REMOVED***
	if m != nil ***REMOVED***
		return m.PageToken
	***REMOVED***
	return ""
***REMOVED***

func (m *ListLogMetricsRequest) GetPageSize() int32 ***REMOVED***
	if m != nil ***REMOVED***
		return m.PageSize
	***REMOVED***
	return 0
***REMOVED***

// Result returned from ListLogMetrics.
type ListLogMetricsResponse struct ***REMOVED***
	// A list of logs-based metrics.
	Metrics []*LogMetric `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
	// If there might be more results than appear in this response, then
	// `nextPageToken` is included.  To get the next set of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
***REMOVED***

func (m *ListLogMetricsResponse) Reset()                    ***REMOVED*** *m = ListLogMetricsResponse***REMOVED******REMOVED*** ***REMOVED***
func (m *ListLogMetricsResponse) String() string            ***REMOVED*** return proto.CompactTextString(m) ***REMOVED***
func (*ListLogMetricsResponse) ProtoMessage()               ***REMOVED******REMOVED***
func (*ListLogMetricsResponse) Descriptor() ([]byte, []int) ***REMOVED*** return fileDescriptor3, []int***REMOVED***2***REMOVED*** ***REMOVED***

func (m *ListLogMetricsResponse) GetMetrics() []*LogMetric ***REMOVED***
	if m != nil ***REMOVED***
		return m.Metrics
	***REMOVED***
	return nil
***REMOVED***

func (m *ListLogMetricsResponse) GetNextPageToken() string ***REMOVED***
	if m != nil ***REMOVED***
		return m.NextPageToken
	***REMOVED***
	return ""
***REMOVED***

// The parameters to GetLogMetric.
type GetLogMetricRequest struct ***REMOVED***
	// The resource name of the desired metric:
	//
	//     "projects/[PROJECT_ID]/metrics/[METRIC_ID]"
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
***REMOVED***

func (m *GetLogMetricRequest) Reset()                    ***REMOVED*** *m = GetLogMetricRequest***REMOVED******REMOVED*** ***REMOVED***
func (m *GetLogMetricRequest) String() string            ***REMOVED*** return proto.CompactTextString(m) ***REMOVED***
func (*GetLogMetricRequest) ProtoMessage()               ***REMOVED******REMOVED***
func (*GetLogMetricRequest) Descriptor() ([]byte, []int) ***REMOVED*** return fileDescriptor3, []int***REMOVED***3***REMOVED*** ***REMOVED***

func (m *GetLogMetricRequest) GetMetricName() string ***REMOVED***
	if m != nil ***REMOVED***
		return m.MetricName
	***REMOVED***
	return ""
***REMOVED***

// The parameters to CreateLogMetric.
type CreateLogMetricRequest struct ***REMOVED***
	// The resource name of the project in which to create the metric:
	//
	//     "projects/[PROJECT_ID]"
	//
	// The new metric must be provided in the request.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// The new logs-based metric, which must not have an identifier that
	// already exists.
	Metric *LogMetric `protobuf:"bytes,2,opt,name=metric" json:"metric,omitempty"`
***REMOVED***

func (m *CreateLogMetricRequest) Reset()                    ***REMOVED*** *m = CreateLogMetricRequest***REMOVED******REMOVED*** ***REMOVED***
func (m *CreateLogMetricRequest) String() string            ***REMOVED*** return proto.CompactTextString(m) ***REMOVED***
func (*CreateLogMetricRequest) ProtoMessage()               ***REMOVED******REMOVED***
func (*CreateLogMetricRequest) Descriptor() ([]byte, []int) ***REMOVED*** return fileDescriptor3, []int***REMOVED***4***REMOVED*** ***REMOVED***

func (m *CreateLogMetricRequest) GetParent() string ***REMOVED***
	if m != nil ***REMOVED***
		return m.Parent
	***REMOVED***
	return ""
***REMOVED***

func (m *CreateLogMetricRequest) GetMetric() *LogMetric ***REMOVED***
	if m != nil ***REMOVED***
		return m.Metric
	***REMOVED***
	return nil
***REMOVED***

// The parameters to UpdateLogMetric.
type UpdateLogMetricRequest struct ***REMOVED***
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
***REMOVED***

func (m *UpdateLogMetricRequest) Reset()                    ***REMOVED*** *m = UpdateLogMetricRequest***REMOVED******REMOVED*** ***REMOVED***
func (m *UpdateLogMetricRequest) String() string            ***REMOVED*** return proto.CompactTextString(m) ***REMOVED***
func (*UpdateLogMetricRequest) ProtoMessage()               ***REMOVED******REMOVED***
func (*UpdateLogMetricRequest) Descriptor() ([]byte, []int) ***REMOVED*** return fileDescriptor3, []int***REMOVED***5***REMOVED*** ***REMOVED***

func (m *UpdateLogMetricRequest) GetMetricName() string ***REMOVED***
	if m != nil ***REMOVED***
		return m.MetricName
	***REMOVED***
	return ""
***REMOVED***

func (m *UpdateLogMetricRequest) GetMetric() *LogMetric ***REMOVED***
	if m != nil ***REMOVED***
		return m.Metric
	***REMOVED***
	return nil
***REMOVED***

// The parameters to DeleteLogMetric.
type DeleteLogMetricRequest struct ***REMOVED***
	// The resource name of the metric to delete:
	//
	//     "projects/[PROJECT_ID]/metrics/[METRIC_ID]"
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName" json:"metric_name,omitempty"`
***REMOVED***

func (m *DeleteLogMetricRequest) Reset()                    ***REMOVED*** *m = DeleteLogMetricRequest***REMOVED******REMOVED*** ***REMOVED***
func (m *DeleteLogMetricRequest) String() string            ***REMOVED*** return proto.CompactTextString(m) ***REMOVED***
func (*DeleteLogMetricRequest) ProtoMessage()               ***REMOVED******REMOVED***
func (*DeleteLogMetricRequest) Descriptor() ([]byte, []int) ***REMOVED*** return fileDescriptor3, []int***REMOVED***6***REMOVED*** ***REMOVED***

func (m *DeleteLogMetricRequest) GetMetricName() string ***REMOVED***
	if m != nil ***REMOVED***
		return m.MetricName
	***REMOVED***
	return ""
***REMOVED***

func init() ***REMOVED***
	proto.RegisterType((*LogMetric)(nil), "google.logging.v2.LogMetric")
	proto.RegisterType((*ListLogMetricsRequest)(nil), "google.logging.v2.ListLogMetricsRequest")
	proto.RegisterType((*ListLogMetricsResponse)(nil), "google.logging.v2.ListLogMetricsResponse")
	proto.RegisterType((*GetLogMetricRequest)(nil), "google.logging.v2.GetLogMetricRequest")
	proto.RegisterType((*CreateLogMetricRequest)(nil), "google.logging.v2.CreateLogMetricRequest")
	proto.RegisterType((*UpdateLogMetricRequest)(nil), "google.logging.v2.UpdateLogMetricRequest")
	proto.RegisterType((*DeleteLogMetricRequest)(nil), "google.logging.v2.DeleteLogMetricRequest")
	proto.RegisterEnum("google.logging.v2.LogMetric_ApiVersion", LogMetric_ApiVersion_name, LogMetric_ApiVersion_value)
***REMOVED***

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MetricsServiceV2 service

type MetricsServiceV2Client interface ***REMOVED***
	// Lists logs-based metrics.
	ListLogMetrics(ctx context.Context, in *ListLogMetricsRequest, opts ...grpc.CallOption) (*ListLogMetricsResponse, error)
	// Gets a logs-based metric.
	GetLogMetric(ctx context.Context, in *GetLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error)
	// Creates a logs-based metric.
	CreateLogMetric(ctx context.Context, in *CreateLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error)
	// Creates or updates a logs-based metric.
	UpdateLogMetric(ctx context.Context, in *UpdateLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error)
	// Deletes a logs-based metric.
	DeleteLogMetric(ctx context.Context, in *DeleteLogMetricRequest, opts ...grpc.CallOption) (*google_protobuf5.Empty, error)
***REMOVED***

type metricsServiceV2Client struct ***REMOVED***
	cc *grpc.ClientConn
***REMOVED***

func NewMetricsServiceV2Client(cc *grpc.ClientConn) MetricsServiceV2Client ***REMOVED***
	return &metricsServiceV2Client***REMOVED***cc***REMOVED***
***REMOVED***

func (c *metricsServiceV2Client) ListLogMetrics(ctx context.Context, in *ListLogMetricsRequest, opts ...grpc.CallOption) (*ListLogMetricsResponse, error) ***REMOVED***
	out := new(ListLogMetricsResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.MetricsServiceV2/ListLogMetrics", in, out, c.cc, opts...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return out, nil
***REMOVED***

func (c *metricsServiceV2Client) GetLogMetric(ctx context.Context, in *GetLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error) ***REMOVED***
	out := new(LogMetric)
	err := grpc.Invoke(ctx, "/google.logging.v2.MetricsServiceV2/GetLogMetric", in, out, c.cc, opts...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return out, nil
***REMOVED***

func (c *metricsServiceV2Client) CreateLogMetric(ctx context.Context, in *CreateLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error) ***REMOVED***
	out := new(LogMetric)
	err := grpc.Invoke(ctx, "/google.logging.v2.MetricsServiceV2/CreateLogMetric", in, out, c.cc, opts...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return out, nil
***REMOVED***

func (c *metricsServiceV2Client) UpdateLogMetric(ctx context.Context, in *UpdateLogMetricRequest, opts ...grpc.CallOption) (*LogMetric, error) ***REMOVED***
	out := new(LogMetric)
	err := grpc.Invoke(ctx, "/google.logging.v2.MetricsServiceV2/UpdateLogMetric", in, out, c.cc, opts...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return out, nil
***REMOVED***

func (c *metricsServiceV2Client) DeleteLogMetric(ctx context.Context, in *DeleteLogMetricRequest, opts ...grpc.CallOption) (*google_protobuf5.Empty, error) ***REMOVED***
	out := new(google_protobuf5.Empty)
	err := grpc.Invoke(ctx, "/google.logging.v2.MetricsServiceV2/DeleteLogMetric", in, out, c.cc, opts...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return out, nil
***REMOVED***

// Server API for MetricsServiceV2 service

type MetricsServiceV2Server interface ***REMOVED***
	// Lists logs-based metrics.
	ListLogMetrics(context.Context, *ListLogMetricsRequest) (*ListLogMetricsResponse, error)
	// Gets a logs-based metric.
	GetLogMetric(context.Context, *GetLogMetricRequest) (*LogMetric, error)
	// Creates a logs-based metric.
	CreateLogMetric(context.Context, *CreateLogMetricRequest) (*LogMetric, error)
	// Creates or updates a logs-based metric.
	UpdateLogMetric(context.Context, *UpdateLogMetricRequest) (*LogMetric, error)
	// Deletes a logs-based metric.
	DeleteLogMetric(context.Context, *DeleteLogMetricRequest) (*google_protobuf5.Empty, error)
***REMOVED***

func RegisterMetricsServiceV2Server(s *grpc.Server, srv MetricsServiceV2Server) ***REMOVED***
	s.RegisterService(&_MetricsServiceV2_serviceDesc, srv)
***REMOVED***

func _MetricsServiceV2_ListLogMetrics_Handler(srv interface***REMOVED******REMOVED***, ctx context.Context, dec func(interface***REMOVED******REMOVED***) error, interceptor grpc.UnaryServerInterceptor) (interface***REMOVED******REMOVED***, error) ***REMOVED***
	in := new(ListLogMetricsRequest)
	if err := dec(in); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	if interceptor == nil ***REMOVED***
		return srv.(MetricsServiceV2Server).ListLogMetrics(ctx, in)
	***REMOVED***
	info := &grpc.UnaryServerInfo***REMOVED***
		Server:     srv,
		FullMethod: "/google.logging.v2.MetricsServiceV2/ListLogMetrics",
	***REMOVED***
	handler := func(ctx context.Context, req interface***REMOVED******REMOVED***) (interface***REMOVED******REMOVED***, error) ***REMOVED***
		return srv.(MetricsServiceV2Server).ListLogMetrics(ctx, req.(*ListLogMetricsRequest))
	***REMOVED***
	return interceptor(ctx, in, info, handler)
***REMOVED***

func _MetricsServiceV2_GetLogMetric_Handler(srv interface***REMOVED******REMOVED***, ctx context.Context, dec func(interface***REMOVED******REMOVED***) error, interceptor grpc.UnaryServerInterceptor) (interface***REMOVED******REMOVED***, error) ***REMOVED***
	in := new(GetLogMetricRequest)
	if err := dec(in); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	if interceptor == nil ***REMOVED***
		return srv.(MetricsServiceV2Server).GetLogMetric(ctx, in)
	***REMOVED***
	info := &grpc.UnaryServerInfo***REMOVED***
		Server:     srv,
		FullMethod: "/google.logging.v2.MetricsServiceV2/GetLogMetric",
	***REMOVED***
	handler := func(ctx context.Context, req interface***REMOVED******REMOVED***) (interface***REMOVED******REMOVED***, error) ***REMOVED***
		return srv.(MetricsServiceV2Server).GetLogMetric(ctx, req.(*GetLogMetricRequest))
	***REMOVED***
	return interceptor(ctx, in, info, handler)
***REMOVED***

func _MetricsServiceV2_CreateLogMetric_Handler(srv interface***REMOVED******REMOVED***, ctx context.Context, dec func(interface***REMOVED******REMOVED***) error, interceptor grpc.UnaryServerInterceptor) (interface***REMOVED******REMOVED***, error) ***REMOVED***
	in := new(CreateLogMetricRequest)
	if err := dec(in); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	if interceptor == nil ***REMOVED***
		return srv.(MetricsServiceV2Server).CreateLogMetric(ctx, in)
	***REMOVED***
	info := &grpc.UnaryServerInfo***REMOVED***
		Server:     srv,
		FullMethod: "/google.logging.v2.MetricsServiceV2/CreateLogMetric",
	***REMOVED***
	handler := func(ctx context.Context, req interface***REMOVED******REMOVED***) (interface***REMOVED******REMOVED***, error) ***REMOVED***
		return srv.(MetricsServiceV2Server).CreateLogMetric(ctx, req.(*CreateLogMetricRequest))
	***REMOVED***
	return interceptor(ctx, in, info, handler)
***REMOVED***

func _MetricsServiceV2_UpdateLogMetric_Handler(srv interface***REMOVED******REMOVED***, ctx context.Context, dec func(interface***REMOVED******REMOVED***) error, interceptor grpc.UnaryServerInterceptor) (interface***REMOVED******REMOVED***, error) ***REMOVED***
	in := new(UpdateLogMetricRequest)
	if err := dec(in); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	if interceptor == nil ***REMOVED***
		return srv.(MetricsServiceV2Server).UpdateLogMetric(ctx, in)
	***REMOVED***
	info := &grpc.UnaryServerInfo***REMOVED***
		Server:     srv,
		FullMethod: "/google.logging.v2.MetricsServiceV2/UpdateLogMetric",
	***REMOVED***
	handler := func(ctx context.Context, req interface***REMOVED******REMOVED***) (interface***REMOVED******REMOVED***, error) ***REMOVED***
		return srv.(MetricsServiceV2Server).UpdateLogMetric(ctx, req.(*UpdateLogMetricRequest))
	***REMOVED***
	return interceptor(ctx, in, info, handler)
***REMOVED***

func _MetricsServiceV2_DeleteLogMetric_Handler(srv interface***REMOVED******REMOVED***, ctx context.Context, dec func(interface***REMOVED******REMOVED***) error, interceptor grpc.UnaryServerInterceptor) (interface***REMOVED******REMOVED***, error) ***REMOVED***
	in := new(DeleteLogMetricRequest)
	if err := dec(in); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	if interceptor == nil ***REMOVED***
		return srv.(MetricsServiceV2Server).DeleteLogMetric(ctx, in)
	***REMOVED***
	info := &grpc.UnaryServerInfo***REMOVED***
		Server:     srv,
		FullMethod: "/google.logging.v2.MetricsServiceV2/DeleteLogMetric",
	***REMOVED***
	handler := func(ctx context.Context, req interface***REMOVED******REMOVED***) (interface***REMOVED******REMOVED***, error) ***REMOVED***
		return srv.(MetricsServiceV2Server).DeleteLogMetric(ctx, req.(*DeleteLogMetricRequest))
	***REMOVED***
	return interceptor(ctx, in, info, handler)
***REMOVED***

var _MetricsServiceV2_serviceDesc = grpc.ServiceDesc***REMOVED***
	ServiceName: "google.logging.v2.MetricsServiceV2",
	HandlerType: (*MetricsServiceV2Server)(nil),
	Methods: []grpc.MethodDesc***REMOVED***
		***REMOVED***
			MethodName: "ListLogMetrics",
			Handler:    _MetricsServiceV2_ListLogMetrics_Handler,
		***REMOVED***,
		***REMOVED***
			MethodName: "GetLogMetric",
			Handler:    _MetricsServiceV2_GetLogMetric_Handler,
		***REMOVED***,
		***REMOVED***
			MethodName: "CreateLogMetric",
			Handler:    _MetricsServiceV2_CreateLogMetric_Handler,
		***REMOVED***,
		***REMOVED***
			MethodName: "UpdateLogMetric",
			Handler:    _MetricsServiceV2_UpdateLogMetric_Handler,
		***REMOVED***,
		***REMOVED***
			MethodName: "DeleteLogMetric",
			Handler:    _MetricsServiceV2_DeleteLogMetric_Handler,
		***REMOVED***,
	***REMOVED***,
	Streams:  []grpc.StreamDesc***REMOVED******REMOVED***,
	Metadata: "google/logging/v2/logging_metrics.proto",
***REMOVED***

func init() ***REMOVED*** proto.RegisterFile("google/logging/v2/logging_metrics.proto", fileDescriptor3) ***REMOVED***

var fileDescriptor3 = []byte***REMOVED***
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x4f, 0x13, 0x41,
	0x14, 0x77, 0x0a, 0x14, 0x79, 0x28, 0xc5, 0x21, 0x94, 0xa6, 0x40, 0xa8, 0x7b, 0x80, 0xc2, 0x61,
	0x57, 0x57, 0x43, 0xfc, 0x13, 0x0f, 0x80, 0x86, 0x0b, 0x1a, 0x52, 0xb4, 0x07, 0x2f, 0xcd, 0xd2,
	0x3e, 0x36, 0x23, 0xed, 0xce, 0xba, 0x33, 0x6d, 0x10, 0xc3, 0x85, 0x70, 0x33, 0xf1, 0xa0, 0xdf,
	0xc2, 0x8f, 0xe0, 0xd7, 0xd0, 0x8f, 0xe0, 0x07, 0x31, 0x3b, 0x33, 0x0b, 0x6b, 0xbb, 0xd2, 0x86,
	0x53, 0x67, 0xde, 0xef, 0xbd, 0xfd, 0xfd, 0xde, 0x7b, 0xbf, 0x4c, 0x61, 0xcd, 0xe7, 0xdc, 0x6f,
	0xa3, 0xd3, 0xe6, 0xbe, 0xcf, 0x02, 0xdf, 0xe9, 0xb9, 0xc9, 0xb1, 0xd1, 0x41, 0x19, 0xb1, 0xa6,
	0xb0, 0xc3, 0x88, 0x4b, 0x4e, 0xef, 0xe9, 0x44, 0xdb, 0xa0, 0x76, 0xcf, 0x2d, 0x2f, 0x99, 0x5a,
	0x2f, 0x64, 0x8e, 0x17, 0x04, 0x5c, 0x7a, 0x92, 0xf1, 0xc0, 0x14, 0x94, 0x97, 0x53, 0x68, 0x8b,
	0x09, 0x19, 0xb1, 0xc3, 0x6e, 0x8c, 0x1b, 0x78, 0x21, 0x05, 0x6b, 0x26, 0x03, 0x2c, 0x1a, 0x40,
	0xdd, 0x0e, 0xbb, 0x47, 0x0e, 0x76, 0x42, 0xf9, 0x49, 0x83, 0xd6, 0x4f, 0x02, 0x53, 0x7b, 0xdc,
	0x7f, 0xad, 0x0a, 0x28, 0x85, 0xf1, 0xc0, 0xeb, 0x60, 0x89, 0x54, 0x48, 0x75, 0xaa, 0xa6, 0xce,
	0xb4, 0x02, 0xd3, 0x2d, 0x14, 0xcd, 0x88, 0x85, 0x31, 0x59, 0x29, 0xa7, 0xa0, 0x74, 0x88, 0x16,
	0x21, 0x7f, 0xc4, 0xda, 0x12, 0xa3, 0xd2, 0x98, 0x02, 0xcd, 0x8d, 0x6e, 0xc1, 0x64, 0x0f, 0x23,
	0x11, 0x57, 0x8d, 0x57, 0x48, 0x75, 0xc6, 0x5d, 0xb3, 0x07, 0x7a, 0xb6, 0x2f, 0xc9, 0xed, 0xad,
	0x90, 0xd5, 0x75, 0x7a, 0x2d, 0xa9, 0xb3, 0x96, 0x00, 0xae, 0xc2, 0x34, 0x0f, 0xb9, 0xba, 0x3b,
	0x7b, 0x4b, 0xfd, 0x3e, 0x9c, 0x25, 0xd6, 0x31, 0xcc, 0xef, 0x31, 0x21, 0x2f, 0x3f, 0x21, 0x6a,
	0xf8, 0xb1, 0x8b, 0x42, 0xc6, 0x8a, 0x42, 0x2f, 0xc2, 0x40, 0x9a, 0x4e, 0xcc, 0x8d, 0x2e, 0x03,
	0x84, 0x9e, 0x8f, 0x0d, 0xc9, 0x8f, 0x31, 0x69, 0x65, 0x2a, 0x8e, 0xbc, 0x8d, 0x03, 0x74, 0x11,
	0xd4, 0xa5, 0x21, 0xd8, 0x29, 0xaa, 0x5e, 0x26, 0x6a, 0xb7, 0xe3, 0xc0, 0x01, 0x3b, 0x45, 0xeb,
	0x04, 0x8a, 0xfd, 0x64, 0x22, 0xe4, 0x81, 0x40, 0xba, 0x09, 0x93, 0x66, 0xb5, 0x25, 0x52, 0x19,
	0xab, 0x4e, 0xbb, 0x4b, 0xd7, 0xf5, 0x59, 0x4b, 0x92, 0xe9, 0x2a, 0x14, 0x02, 0x3c, 0x91, 0x8d,
	0x01, 0x49, 0x77, 0xe3, 0xf0, 0x7e, 0x22, 0xcb, 0xda, 0x84, 0xb9, 0x5d, 0xbc, 0x22, 0x4e, 0x9a,
	0x5c, 0x81, 0x69, 0xfd, 0xa5, 0x46, 0x6a, 0x67, 0xa0, 0x43, 0x6f, 0xbc, 0x0e, 0x5a, 0x47, 0x50,
	0xdc, 0x89, 0xd0, 0x93, 0x38, 0x50, 0xfa, 0xbf, 0xf9, 0x3c, 0x86, 0xbc, 0xae, 0x57, 0x42, 0x86,
	0x35, 0x62, 0x72, 0x2d, 0x0e, 0xc5, 0x77, 0x61, 0x2b, 0x8b, 0x67, 0x98, 0xc4, 0x1b, 0x12, 0x3e,
	0x85, 0xe2, 0x4b, 0x6c, 0xe3, 0x0d, 0x08, 0xdd, 0xdf, 0x13, 0x30, 0x6b, 0xf6, 0x77, 0x80, 0x51,
	0x8f, 0x35, 0xb1, 0xee, 0xd2, 0xaf, 0x04, 0x66, 0xfe, 0xdd, 0x2d, 0xad, 0x66, 0x09, 0xc9, 0xf2,
	0x5a, 0x79, 0x7d, 0x84, 0x4c, 0x6d, 0x14, 0x6b, 0xed, 0xfc, 0xd7, 0x9f, 0xef, 0xb9, 0xfb, 0x74,
	0x25, 0x7e, 0x15, 0x3e, 0xeb, 0x99, 0xbf, 0x08, 0x23, 0xfe, 0x01, 0x9b, 0x52, 0x38, 0x1b, 0x67,
	0x4e, 0xe2, 0x8c, 0x0b, 0x02, 0x77, 0xd2, 0x2b, 0xa7, 0xab, 0x19, 0x24, 0x19, 0x9e, 0x28, 0x5f,
	0x3b, 0x3f, 0xcb, 0x56, 0xfc, 0x55, 0xba, 0xaa, 0xf8, 0x53, 0x83, 0x4a, 0x89, 0x48, 0x34, 0x38,
	0x1b, 0x67, 0xf4, 0x0b, 0x81, 0x42, 0x9f, 0x83, 0x68, 0x56, 0xbb, 0xd9, 0x2e, 0x1b, 0x22, 0xc6,
	0x51, 0x62, 0xd6, 0xad, 0x61, 0xc3, 0x78, 0x66, 0xb6, 0x4e, 0xbf, 0x11, 0x28, 0xf4, 0xf9, 0x2c,
	0x53, 0x4d, 0xb6, 0x17, 0x87, 0xa8, 0xd9, 0x54, 0x6a, 0x1e, 0x94, 0x47, 0x1c, 0xcd, 0xa5, 0xa8,
	0x0b, 0x02, 0x85, 0x3e, 0x2f, 0x66, 0x8a, 0xca, 0xf6, 0x6b, 0xb9, 0x98, 0xa4, 0x26, 0x8f, 0xb3,
	0xfd, 0x2a, 0x7e, 0x9c, 0x93, 0x4d, 0x6d, 0x8c, 0x28, 0x67, 0xfb, 0x9c, 0xc0, 0x7c, 0x93, 0x77,
	0x06, 0x89, 0xb7, 0xe7, 0xf6, 0xf4, 0xd9, 0x78, 0x71, 0x3f, 0xe6, 0xd9, 0x27, 0xef, 0x9f, 0x98,
	0x4c, 0x9f, 0xb7, 0xbd, 0xc0, 0xb7, 0x79, 0xe4, 0x3b, 0x3e, 0x06, 0x4a, 0x85, 0xa3, 0x21, 0x2f,
	0x64, 0x22, 0xf5, 0x2f, 0xf6, 0xdc, 0x1c, 0x7f, 0xe4, 0x16, 0x76, 0x75, 0xe9, 0x4e, 0x9b, 0x77,
	0x5b, 0xb6, 0xf9, 0xbc, 0x5d, 0x77, 0x0f, 0xf3, 0xaa, 0xfc, 0xd1, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x22, 0xd3, 0x7f, 0x20, 0x03, 0x07, 0x00, 0x00,
***REMOVED***
