// Copyright 2014 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prometheus

import (
	"errors"
)

// Counter is a Metric that represents a single numerical value that only ever
// goes up. That implies that it cannot be used to count items whose number can
// also go down, e.g. the number of currently running goroutines. Those
// "counters" are represented by Gauges.
//
// A Counter is typically used to count requests served, tasks completed, errors
// occurred, etc.
//
// To create Counter instances, use NewCounter.
type Counter interface ***REMOVED***
	Metric
	Collector

	// Set is used to set the Counter to an arbitrary value. It is only used
	// if you have to transfer a value from an external counter into this
	// Prometheus metric. Do not use it for regular handling of a
	// Prometheus counter (as it can be used to break the contract of
	// monotonically increasing values).
	Set(float64)
	// Inc increments the counter by 1.
	Inc()
	// Add adds the given value to the counter. It panics if the value is <
	// 0.
	Add(float64)
***REMOVED***

// CounterOpts is an alias for Opts. See there for doc comments.
type CounterOpts Opts

// NewCounter creates a new Counter based on the provided CounterOpts.
func NewCounter(opts CounterOpts) Counter ***REMOVED***
	desc := NewDesc(
		BuildFQName(opts.Namespace, opts.Subsystem, opts.Name),
		opts.Help,
		nil,
		opts.ConstLabels,
	)
	result := &counter***REMOVED***value: value***REMOVED***desc: desc, valType: CounterValue, labelPairs: desc.constLabelPairs***REMOVED******REMOVED***
	result.Init(result) // Init self-collection.
	return result
***REMOVED***

type counter struct ***REMOVED***
	value
***REMOVED***

func (c *counter) Add(v float64) ***REMOVED***
	if v < 0 ***REMOVED***
		panic(errors.New("counter cannot decrease in value"))
	***REMOVED***
	c.value.Add(v)
***REMOVED***

// CounterVec is a Collector that bundles a set of Counters that all share the
// same Desc, but have different values for their variable labels. This is used
// if you want to count the same thing partitioned by various dimensions
// (e.g. number of HTTP requests, partitioned by response code and
// method). Create instances with NewCounterVec.
//
// CounterVec embeds MetricVec. See there for a full list of methods with
// detailed documentation.
type CounterVec struct ***REMOVED***
	MetricVec
***REMOVED***

// NewCounterVec creates a new CounterVec based on the provided CounterOpts and
// partitioned by the given label names. At least one label name must be
// provided.
func NewCounterVec(opts CounterOpts, labelNames []string) *CounterVec ***REMOVED***
	desc := NewDesc(
		BuildFQName(opts.Namespace, opts.Subsystem, opts.Name),
		opts.Help,
		labelNames,
		opts.ConstLabels,
	)
	return &CounterVec***REMOVED***
		MetricVec: MetricVec***REMOVED***
			children: map[uint64]Metric***REMOVED******REMOVED***,
			desc:     desc,
			newMetric: func(lvs ...string) Metric ***REMOVED***
				result := &counter***REMOVED***value: value***REMOVED***
					desc:       desc,
					valType:    CounterValue,
					labelPairs: makeLabelPairs(desc, lvs),
				***REMOVED******REMOVED***
				result.Init(result) // Init self-collection.
				return result
			***REMOVED***,
		***REMOVED***,
	***REMOVED***
***REMOVED***

// GetMetricWithLabelValues replaces the method of the same name in
// MetricVec. The difference is that this method returns a Counter and not a
// Metric so that no type conversion is required.
func (m *CounterVec) GetMetricWithLabelValues(lvs ...string) (Counter, error) ***REMOVED***
	metric, err := m.MetricVec.GetMetricWithLabelValues(lvs...)
	if metric != nil ***REMOVED***
		return metric.(Counter), err
	***REMOVED***
	return nil, err
***REMOVED***

// GetMetricWith replaces the method of the same name in MetricVec. The
// difference is that this method returns a Counter and not a Metric so that no
// type conversion is required.
func (m *CounterVec) GetMetricWith(labels Labels) (Counter, error) ***REMOVED***
	metric, err := m.MetricVec.GetMetricWith(labels)
	if metric != nil ***REMOVED***
		return metric.(Counter), err
	***REMOVED***
	return nil, err
***REMOVED***

// WithLabelValues works as GetMetricWithLabelValues, but panics where
// GetMetricWithLabelValues would have returned an error. By not returning an
// error, WithLabelValues allows shortcuts like
//     myVec.WithLabelValues("404", "GET").Add(42)
func (m *CounterVec) WithLabelValues(lvs ...string) Counter ***REMOVED***
	return m.MetricVec.WithLabelValues(lvs...).(Counter)
***REMOVED***

// With works as GetMetricWith, but panics where GetMetricWithLabels would have
// returned an error. By not returning an error, With allows shortcuts like
//     myVec.With(Labels***REMOVED***"code": "404", "method": "GET"***REMOVED***).Add(42)
func (m *CounterVec) With(labels Labels) Counter ***REMOVED***
	return m.MetricVec.With(labels).(Counter)
***REMOVED***

// CounterFunc is a Counter whose value is determined at collect time by calling a
// provided function.
//
// To create CounterFunc instances, use NewCounterFunc.
type CounterFunc interface ***REMOVED***
	Metric
	Collector
***REMOVED***

// NewCounterFunc creates a new CounterFunc based on the provided
// CounterOpts. The value reported is determined by calling the given function
// from within the Write method. Take into account that metric collection may
// happen concurrently. If that results in concurrent calls to Write, like in
// the case where a CounterFunc is directly registered with Prometheus, the
// provided function must be concurrency-safe. The function should also honor
// the contract for a Counter (values only go up, not down), but compliance will
// not be checked.
func NewCounterFunc(opts CounterOpts, function func() float64) CounterFunc ***REMOVED***
	return newValueFunc(NewDesc(
		BuildFQName(opts.Namespace, opts.Subsystem, opts.Name),
		opts.Help,
		nil,
		opts.ConstLabels,
	), CounterValue, function)
***REMOVED***
