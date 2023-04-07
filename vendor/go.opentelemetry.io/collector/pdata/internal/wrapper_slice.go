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

package internal // import "go.opentelemetry.io/collector/pdata/internal"

import (
	otlpcommon "go.opentelemetry.io/collector/pdata/internal/data/protogen/common/v1"
)

type Slice struct {
	orig *[]otlpcommon.AnyValue
}

func GetOrigSlice(ms Slice) *[]otlpcommon.AnyValue {
	return ms.orig
}

func NewSlice(orig *[]otlpcommon.AnyValue) Slice {
	return Slice{orig: orig}
}

func GenerateTestSlice() Slice {
	orig := []otlpcommon.AnyValue{}
	tv := NewSlice(&orig)
	FillTestSlice(tv)
	return tv
}

func FillTestSlice(tv Slice) {
	*tv.orig = make([]otlpcommon.AnyValue, 7)
	for i := 0; i < 7; i++ {
		FillTestValue(NewValue(&(*tv.orig)[i]))
	}
}
