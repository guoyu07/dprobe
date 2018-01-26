// Copyright 2016, Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package logging

import (
	"fmt"
	"math"
	"runtime"
	"strings"
	"time"

	gax "github.com/googleapis/gax-go"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	loggingpb "google.golang.org/genproto/googleapis/logging/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var (
	metricsParentPathTemplate = gax.MustCompilePathTemplate("projects/***REMOVED***project***REMOVED***")
	metricsMetricPathTemplate = gax.MustCompilePathTemplate("projects/***REMOVED***project***REMOVED***/metrics/***REMOVED***metric***REMOVED***")
)

// MetricsCallOptions contains the retry settings for each method of MetricsClient.
type MetricsCallOptions struct ***REMOVED***
	ListLogMetrics  []gax.CallOption
	GetLogMetric    []gax.CallOption
	CreateLogMetric []gax.CallOption
	UpdateLogMetric []gax.CallOption
	DeleteLogMetric []gax.CallOption
***REMOVED***

func defaultMetricsClientOptions() []option.ClientOption ***REMOVED***
	return []option.ClientOption***REMOVED***
		option.WithEndpoint("logging.googleapis.com:443"),
		option.WithScopes(
			"https://www.googleapis.com/auth/cloud-platform",
			"https://www.googleapis.com/auth/cloud-platform.read-only",
			"https://www.googleapis.com/auth/logging.admin",
			"https://www.googleapis.com/auth/logging.read",
			"https://www.googleapis.com/auth/logging.write",
		),
	***REMOVED***
***REMOVED***

func defaultMetricsCallOptions() *MetricsCallOptions ***REMOVED***
	retry := map[[2]string][]gax.CallOption***REMOVED***
		***REMOVED***"default", "idempotent"***REMOVED***: ***REMOVED***
			gax.WithRetry(func() gax.Retryer ***REMOVED***
				return gax.OnCodes([]codes.Code***REMOVED***
					codes.DeadlineExceeded,
					codes.Unavailable,
				***REMOVED***, gax.Backoff***REMOVED***
					Initial:    100 * time.Millisecond,
					Max:        1000 * time.Millisecond,
					Multiplier: 1.2,
				***REMOVED***)
			***REMOVED***),
		***REMOVED***,
	***REMOVED***
	return &MetricsCallOptions***REMOVED***
		ListLogMetrics:  retry[[2]string***REMOVED***"default", "idempotent"***REMOVED***],
		GetLogMetric:    retry[[2]string***REMOVED***"default", "idempotent"***REMOVED***],
		CreateLogMetric: retry[[2]string***REMOVED***"default", "non_idempotent"***REMOVED***],
		UpdateLogMetric: retry[[2]string***REMOVED***"default", "non_idempotent"***REMOVED***],
		DeleteLogMetric: retry[[2]string***REMOVED***"default", "idempotent"***REMOVED***],
	***REMOVED***
***REMOVED***

// MetricsClient is a client for interacting with Stackdriver Logging API.
type MetricsClient struct ***REMOVED***
	// The connection to the service.
	conn *grpc.ClientConn

	// The gRPC API client.
	metricsClient loggingpb.MetricsServiceV2Client

	// The call options for this service.
	CallOptions *MetricsCallOptions

	// The metadata to be sent with each request.
	metadata metadata.MD
***REMOVED***

// NewMetricsClient creates a new metrics service v2 client.
//
// Service for configuring logs-based metrics.
func NewMetricsClient(ctx context.Context, opts ...option.ClientOption) (*MetricsClient, error) ***REMOVED***
	conn, err := transport.DialGRPC(ctx, append(defaultMetricsClientOptions(), opts...)...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	c := &MetricsClient***REMOVED***
		conn:        conn,
		CallOptions: defaultMetricsCallOptions(),

		metricsClient: loggingpb.NewMetricsServiceV2Client(conn),
	***REMOVED***
	c.SetGoogleClientInfo("gax", gax.Version)
	return c, nil
***REMOVED***

// Connection returns the client's connection to the API service.
func (c *MetricsClient) Connection() *grpc.ClientConn ***REMOVED***
	return c.conn
***REMOVED***

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *MetricsClient) Close() error ***REMOVED***
	return c.conn.Close()
***REMOVED***

// SetGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *MetricsClient) SetGoogleClientInfo(name, version string) ***REMOVED***
	goVersion := strings.Replace(runtime.Version(), " ", "_", -1)
	v := fmt.Sprintf("%s/%s %s gax/%s go/%s", name, version, gapicNameVersion, gax.Version, goVersion)
	c.metadata = metadata.Pairs("x-goog-api-client", v)
***REMOVED***

// MetricsParentPath returns the path for the parent resource.
func MetricsParentPath(project string) string ***REMOVED***
	path, err := metricsParentPathTemplate.Render(map[string]string***REMOVED***
		"project": project,
	***REMOVED***)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return path
***REMOVED***

// MetricsMetricPath returns the path for the metric resource.
func MetricsMetricPath(project, metric string) string ***REMOVED***
	path, err := metricsMetricPathTemplate.Render(map[string]string***REMOVED***
		"project": project,
		"metric":  metric,
	***REMOVED***)
	if err != nil ***REMOVED***
		panic(err)
	***REMOVED***
	return path
***REMOVED***

// ListLogMetrics lists logs-based metrics.
func (c *MetricsClient) ListLogMetrics(ctx context.Context, req *loggingpb.ListLogMetricsRequest) *LogMetricIterator ***REMOVED***
	md, _ := metadata.FromContext(ctx)
	ctx = metadata.NewContext(ctx, metadata.Join(md, c.metadata))
	it := &LogMetricIterator***REMOVED******REMOVED***
	it.InternalFetch = func(pageSize int, pageToken string) ([]*loggingpb.LogMetric, string, error) ***REMOVED***
		var resp *loggingpb.ListLogMetricsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 ***REMOVED***
			req.PageSize = math.MaxInt32
		***REMOVED*** else ***REMOVED***
			req.PageSize = int32(pageSize)
		***REMOVED***
		err := gax.Invoke(ctx, func(ctx context.Context) error ***REMOVED***
			var err error
			resp, err = c.metricsClient.ListLogMetrics(ctx, req)
			return err
		***REMOVED***, c.CallOptions.ListLogMetrics...)
		if err != nil ***REMOVED***
			return nil, "", err
		***REMOVED***
		return resp.Metrics, resp.NextPageToken, nil
	***REMOVED***
	fetch := func(pageSize int, pageToken string) (string, error) ***REMOVED***
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil ***REMOVED***
			return "", err
		***REMOVED***
		it.items = append(it.items, items...)
		return nextPageToken, nil
	***REMOVED***
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	return it
***REMOVED***

// GetLogMetric gets a logs-based metric.
func (c *MetricsClient) GetLogMetric(ctx context.Context, req *loggingpb.GetLogMetricRequest) (*loggingpb.LogMetric, error) ***REMOVED***
	md, _ := metadata.FromContext(ctx)
	ctx = metadata.NewContext(ctx, metadata.Join(md, c.metadata))
	var resp *loggingpb.LogMetric
	err := gax.Invoke(ctx, func(ctx context.Context) error ***REMOVED***
		var err error
		resp, err = c.metricsClient.GetLogMetric(ctx, req)
		return err
	***REMOVED***, c.CallOptions.GetLogMetric...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return resp, nil
***REMOVED***

// CreateLogMetric creates a logs-based metric.
func (c *MetricsClient) CreateLogMetric(ctx context.Context, req *loggingpb.CreateLogMetricRequest) (*loggingpb.LogMetric, error) ***REMOVED***
	md, _ := metadata.FromContext(ctx)
	ctx = metadata.NewContext(ctx, metadata.Join(md, c.metadata))
	var resp *loggingpb.LogMetric
	err := gax.Invoke(ctx, func(ctx context.Context) error ***REMOVED***
		var err error
		resp, err = c.metricsClient.CreateLogMetric(ctx, req)
		return err
	***REMOVED***, c.CallOptions.CreateLogMetric...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return resp, nil
***REMOVED***

// UpdateLogMetric creates or updates a logs-based metric.
func (c *MetricsClient) UpdateLogMetric(ctx context.Context, req *loggingpb.UpdateLogMetricRequest) (*loggingpb.LogMetric, error) ***REMOVED***
	md, _ := metadata.FromContext(ctx)
	ctx = metadata.NewContext(ctx, metadata.Join(md, c.metadata))
	var resp *loggingpb.LogMetric
	err := gax.Invoke(ctx, func(ctx context.Context) error ***REMOVED***
		var err error
		resp, err = c.metricsClient.UpdateLogMetric(ctx, req)
		return err
	***REMOVED***, c.CallOptions.UpdateLogMetric...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return resp, nil
***REMOVED***

// DeleteLogMetric deletes a logs-based metric.
func (c *MetricsClient) DeleteLogMetric(ctx context.Context, req *loggingpb.DeleteLogMetricRequest) error ***REMOVED***
	md, _ := metadata.FromContext(ctx)
	ctx = metadata.NewContext(ctx, metadata.Join(md, c.metadata))
	err := gax.Invoke(ctx, func(ctx context.Context) error ***REMOVED***
		var err error
		_, err = c.metricsClient.DeleteLogMetric(ctx, req)
		return err
	***REMOVED***, c.CallOptions.DeleteLogMetric...)
	return err
***REMOVED***

// LogMetricIterator manages a stream of *loggingpb.LogMetric.
type LogMetricIterator struct ***REMOVED***
	items    []*loggingpb.LogMetric
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*loggingpb.LogMetric, nextPageToken string, err error)
***REMOVED***

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *LogMetricIterator) PageInfo() *iterator.PageInfo ***REMOVED***
	return it.pageInfo
***REMOVED***

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *LogMetricIterator) Next() (*loggingpb.LogMetric, error) ***REMOVED***
	var item *loggingpb.LogMetric
	if err := it.nextFunc(); err != nil ***REMOVED***
		return item, err
	***REMOVED***
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
***REMOVED***

func (it *LogMetricIterator) bufLen() int ***REMOVED***
	return len(it.items)
***REMOVED***

func (it *LogMetricIterator) takeBuf() interface***REMOVED******REMOVED*** ***REMOVED***
	b := it.items
	it.items = nil
	return b
***REMOVED***
