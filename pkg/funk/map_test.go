package funk_test

import (
	"testing"

	"github.com/justtrackio/gosoline/pkg/funk"
	"github.com/stretchr/testify/assert"
)

func TestMergeMaps(t *testing.T) {
	tests := []struct {
		name string
		ins  []map[string]int
		out  map[string]int
	}{
		{
			name: "simple",
			ins: []map[string]int{
				{
					"one": 1,
					"two": 2,
				},
				{
					"two":   2,
					"three": 3,
				},
			},
			out: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
		},
		{
			name: "empty second",
			ins: []map[string]int{
				{
					"one": 1,
					"two": 2,
				},
			},
			out: map[string]int{
				"one": 1,
				"two": 2,
			},
		},
		{
			name: "nil second",
			ins: []map[string]int{
				{
					"one": 1,
					"two": 2,
				},
				nil,
			},
			out: map[string]int{
				"one": 1,
				"two": 2,
			},
		},
		{
			name: "nil first",
			ins: []map[string]int{
				nil,
				{
					"one": 1,
					"two": 2,
				},
				nil,
			},
			out: map[string]int{
				"one": 1,
				"two": 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := funk.MergeMaps(tt.ins...)
			assert.Equal(t, tt.out, out)
		})
	}
}
