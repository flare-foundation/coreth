// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindThresholds(t *testing.T) {

	tests := []struct {
		size  uint
		steps uint
		min   uint
		want  []uint
	}{
		{
			size:  20,
			steps: 0,
			min:   4,
			want:  []uint{},
		},
		{
			size:  20,
			steps: 1,
			min:   4,
			want:  []uint{5},
		},
		{
			size:  20,
			steps: 2,
			min:   4,
			want:  []uint{10, 5},
		},
		{
			size:  20,
			steps: 3,
			min:   4,
			want:  []uint{15, 10, 5},
		},
		{
			size:  20,
			steps: 4,
			min:   4,
			want:  []uint{20, 15, 10, 5},
		},
	}

	for _, test := range tests {
		have := findThresholds(test.size, test.steps, test.min)
		assert.Equal(t, test.want, have)
	}
}
