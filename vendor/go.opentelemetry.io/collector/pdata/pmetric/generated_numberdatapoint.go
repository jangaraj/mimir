// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// NumberDataPoint is a single data point in a timeseries that describes the time-varying value of a number metric.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewNumberDataPoint function to create new instances.
// Important: zero-initialized instance is not valid for use.
type NumberDataPoint struct {
	orig *otlpmetrics.NumberDataPoint
}

func newNumberDataPoint(orig *otlpmetrics.NumberDataPoint) NumberDataPoint {
	return NumberDataPoint{orig}
}

// NewNumberDataPoint creates a new empty NumberDataPoint.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewNumberDataPoint() NumberDataPoint {
	return newNumberDataPoint(&otlpmetrics.NumberDataPoint{})
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms NumberDataPoint) MoveTo(dest NumberDataPoint) {
	*dest.orig = *ms.orig
	*ms.orig = otlpmetrics.NumberDataPoint{}
}

// Attributes returns the Attributes associated with this NumberDataPoint.
func (ms NumberDataPoint) Attributes() pcommon.Map {
	return pcommon.Map(internal.NewMap(&ms.orig.Attributes))
}

// StartTimestamp returns the starttimestamp associated with this NumberDataPoint.
func (ms NumberDataPoint) StartTimestamp() pcommon.Timestamp {
	return pcommon.Timestamp(ms.orig.StartTimeUnixNano)
}

// SetStartTimestamp replaces the starttimestamp associated with this NumberDataPoint.
func (ms NumberDataPoint) SetStartTimestamp(v pcommon.Timestamp) {
	ms.orig.StartTimeUnixNano = uint64(v)
}

// Timestamp returns the timestamp associated with this NumberDataPoint.
func (ms NumberDataPoint) Timestamp() pcommon.Timestamp {
	return pcommon.Timestamp(ms.orig.TimeUnixNano)
}

// SetTimestamp replaces the timestamp associated with this NumberDataPoint.
func (ms NumberDataPoint) SetTimestamp(v pcommon.Timestamp) {
	ms.orig.TimeUnixNano = uint64(v)
}

// ValueType returns the type of the value for this NumberDataPoint.
// Calling this function on zero-initialized NumberDataPoint will cause a panic.
func (ms NumberDataPoint) ValueType() NumberDataPointValueType {
	switch ms.orig.Value.(type) {
	case *otlpmetrics.NumberDataPoint_AsDouble:
		return NumberDataPointValueTypeDouble
	case *otlpmetrics.NumberDataPoint_AsInt:
		return NumberDataPointValueTypeInt
	}
	return NumberDataPointValueTypeEmpty
}

// DoubleValue returns the double associated with this NumberDataPoint.
func (ms NumberDataPoint) DoubleValue() float64 {
	return ms.orig.GetAsDouble()
}

// SetDoubleValue replaces the double associated with this NumberDataPoint.
func (ms NumberDataPoint) SetDoubleValue(v float64) {
	ms.orig.Value = &otlpmetrics.NumberDataPoint_AsDouble{
		AsDouble: v,
	}
}

// IntValue returns the int associated with this NumberDataPoint.
func (ms NumberDataPoint) IntValue() int64 {
	return ms.orig.GetAsInt()
}

// SetIntValue replaces the int associated with this NumberDataPoint.
func (ms NumberDataPoint) SetIntValue(v int64) {
	ms.orig.Value = &otlpmetrics.NumberDataPoint_AsInt{
		AsInt: v,
	}
}

// Exemplars returns the Exemplars associated with this NumberDataPoint.
func (ms NumberDataPoint) Exemplars() ExemplarSlice {
	return newExemplarSlice(&ms.orig.Exemplars)
}

// Flags returns the flags associated with this NumberDataPoint.
func (ms NumberDataPoint) Flags() DataPointFlags {
	return DataPointFlags(ms.orig.Flags)
}

// SetFlags replaces the flags associated with this NumberDataPoint.
func (ms NumberDataPoint) SetFlags(v DataPointFlags) {
	ms.orig.Flags = uint32(v)
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms NumberDataPoint) CopyTo(dest NumberDataPoint) {
	ms.Attributes().CopyTo(dest.Attributes())
	dest.SetStartTimestamp(ms.StartTimestamp())
	dest.SetTimestamp(ms.Timestamp())
	switch ms.ValueType() {
	case NumberDataPointValueTypeDouble:
		dest.SetDoubleValue(ms.DoubleValue())
	case NumberDataPointValueTypeInt:
		dest.SetIntValue(ms.IntValue())
	}

	ms.Exemplars().CopyTo(dest.Exemplars())
	dest.SetFlags(ms.Flags())
}
