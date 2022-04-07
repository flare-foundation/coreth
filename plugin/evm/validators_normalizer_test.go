// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"
)

func TestNewValidatorsNormalizer(t *testing.T) {

	retrieve := &ValidatorsRetrieverMock{}

	normalize := NewValidatorsNormalizer(logging.NoLog{}, retrieve)
	assert.Equal(t, retrieve, normalize.retrieve)
}

func TestValidatorsNormalizer_ByEpoch(t *testing.T) {

	tests := []struct {
		name        string
		inWeights   []uint64
		wantWeights []uint64
		err         error
		wantErr     require.ErrorAssertionFunc
	}{
		{
			name: "single validator",
			inWeights: []uint64{
				1,
			},
			wantWeights: []uint64{
				uint64(math.MaxInt64),
			},
			wantErr: require.NoError,
		},
		{
			name: "three same weight validators",
			inWeights: []uint64{
				1,
				1,
				1,
			},
			wantWeights: []uint64{
				3074457345618258602,
				3074457345618258602,
				3074457345618258602,
			},
			wantErr: require.NoError,
		},
		{
			name: "three different weight validators",
			inWeights: []uint64{
				1,
				2,
				3,
			},
			wantWeights: []uint64{
				1537228672809129301,
				3074457345618258602,
				4611686018427387903,
			},
			wantErr: require.NoError,
		},
		{
			name: "retriever failure",
			inWeights: []uint64{
				1,
			},
			err:     errors.New("dummy error"),
			wantErr: require.Error,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			validators := make(map[ids.ShortID]uint64, len(test.inWeights))
			for i, weight := range test.inWeights {
				validators[ids.ShortID{byte(i)}] = weight
			}

			retrieve := &ValidatorsRetrieverMock{
				ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
					return validators, test.err
				},
			}

			normalize := ValidatorsNormalizer{
				log:      logging.NoLog{},
				retrieve: retrieve,
			}

			got, err := normalize.ByEpoch(0)
			test.wantErr(t, err)

			if err != nil {
				return
			}

			for i := range test.inWeights {
				gotWeight := got[ids.ShortID{byte(i)}]
				assert.Equal(t, test.wantWeights[i], gotWeight)
			}
		})
	}
}
