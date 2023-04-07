// SPDX-License-Identifier: AGPL-3.0-only
// Provenance-includes-location: https://github.com/cortexproject/cortex/blob/master/pkg/util/labels_test.go
// Provenance-includes-license: Apache-2.0
// Provenance-includes-copyright: The Cortex Authors.

package util

import (
	"testing"

	"github.com/prometheus/prometheus/model/labels"
	"github.com/stretchr/testify/assert"
)

func TestLabelMatchersToString(t *testing.T) {
	tests := []struct {
		input    []*labels.Matcher
		expected string
	}{
		{
			input:    nil,
			expected: "{}",
		}, {
			input: []*labels.Matcher{
				labels.MustNewMatcher(labels.MatchEqual, "foo", "bar"),
			},
			expected: `{foo="bar"}`,
		}, {
			input: []*labels.Matcher{
				labels.MustNewMatcher(labels.MatchEqual, "foo", "bar"),
				labels.MustNewMatcher(labels.MatchNotEqual, "who", "boh"),
			},
			expected: `{foo="bar",who!="boh"}`,
		}, {
			input: []*labels.Matcher{
				labels.MustNewMatcher(labels.MatchEqual, labels.MetricName, "metric"),
				labels.MustNewMatcher(labels.MatchNotEqual, "who", "boh"),
			},
			expected: `{__name__="metric",who!="boh"}`,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expected, LabelMatchersToString(tc.input))
	}
}
