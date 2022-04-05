// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"
)

type testValidatorRetriever struct {
	ByEpochFunc func(epoch uint64) (map[ids.ShortID]uint64, error)
}

func (t *testValidatorRetriever) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {
	return t.ByEpochFunc(epoch)
}

func TestNewValidatorsTransitioner(t *testing.T) {
	providers := fakeProviders(12)
	validators := fakeValidators(12)
	mockValidators := &testValidatorRetriever{
		ByEpochFunc: func(uint64) (map[ids.ShortID]uint64, error) {
			return validators, nil
		},
	}
	mockProviders := &testValidatorRetriever{
		ByEpochFunc: func(uint64) (map[ids.ShortID]uint64, error) {
			return providers, nil
		},
	}

	got := NewValidatorsTransitioner(logging.NoLog{}, mockValidators, mockProviders)
	require.NotNil(t, got)
	assert.Equal(t, mockValidators, got.validators)
	assert.Equal(t, mockProviders, got.providers)
}

func TestValidatorsTransitioner_ByEpoch(t *testing.T) {

	validatorIDs := []ids.ShortID{
		{1},
		{2},
		{3},
	}

	validators := map[ids.ShortID]uint64{
		validatorIDs[0]: 1,
		validatorIDs[1]: 2,
		validatorIDs[2]: 3,
	}

	providerIDs := []ids.ShortID{
		{4},
		{5},
		{6},
	}

	providers := map[ids.ShortID]uint64{
		providerIDs[0]: 4,
		providerIDs[1]: 5,
		providerIDs[2]: 6,
	}

	mockProvider := &testValidatorRetriever{
		ByEpochFunc: func(uint64) (map[ids.ShortID]uint64, error) {
			return providers, nil
		},
	}

	tests := []struct {
		name         string
		steps        []Step
		epoch        uint64
		validatorIDs []ids.ShortID
	}{
		{
			name:         "no steps zero epoch",
			steps:        []Step{},
			epoch:        0,
			validatorIDs: validatorIDs,
		},
		{
			name:         "no steps non-zero epoch",
			steps:        []Step{},
			epoch:        1,
			validatorIDs: validatorIDs,
		},
		{
			name:  "one step epoch before",
			epoch: 4,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
			},
			validatorIDs: validatorIDs,
		},
		{
			name:  "one step epoch exact",
			epoch: 5,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
			},
			validatorIDs: append(validatorIDs[:2], providerIDs[0]),
		},
		{
			name:  "one step epoch after",
			epoch: 6,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
			},
			validatorIDs: append(validatorIDs[:2], providerIDs[0]),
		},
		{
			name:  "two steps epoch before first",
			epoch: 4,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
				{Epoch: 10, Cutoff: 1},
			},
			validatorIDs: validatorIDs,
		},
		{
			name:  "two steps epoch exact first",
			epoch: 5,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
				{Epoch: 10, Cutoff: 1},
			},
			validatorIDs: append(validatorIDs[:2], providerIDs[0]),
		},
		{
			name:  "two steps epoch after first",
			epoch: 6,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
				{Epoch: 10, Cutoff: 1},
			},
			validatorIDs: append(validatorIDs[:2], providerIDs[0]),
		},
		{
			name:  "two steps epoch before second",
			epoch: 9,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
				{Epoch: 10, Cutoff: 1},
			},
			validatorIDs: append(validatorIDs[:2], providerIDs[0]),
		},
		{
			name:  "two steps epoch exact second",
			epoch: 10,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
				{Epoch: 10, Cutoff: 1},
			},
			validatorIDs: append(validatorIDs[:1], providerIDs[0], providerIDs[1]),
		},
		{
			name:  "two steps epoch after second",
			epoch: 11,
			steps: []Step{
				{Epoch: 5, Cutoff: 2},
				{Epoch: 10, Cutoff: 1},
			},
			validatorIDs: append(validatorIDs[:1], providerIDs[0], providerIDs[1]),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			v := &ValidatorsDefault{
				validators: validators,
				steps:      test.steps,
			}

			transitioner := NewValidatorsTransitioner(logging.NoLog{}, v, mockProvider, WithStepSize(1))

			got, err := transitioner.ByEpoch(test.epoch)
			require.NoError(t, err)

			assert.Len(t, got, len(test.validatorIDs))
			for _, validatorID := range test.validatorIDs {
				assert.Contains(t, got, validatorID)
			}
		})
	}
}

func fakeProviders(epoch uint64) map[ids.ShortID]uint64 {

	providers := make(map[ids.ShortID]uint64)
	for i := 0; i < int(epoch); i++ {
		providers[ids.ShortID{byte(i)}] = uint64(i)
	}

	return providers
}

// fakeValidators gives a different set of validators compared to fakeProviders() to have no overlap for now
func fakeValidators(epoch uint64) map[ids.ShortID]uint64 {

	providers := make(map[ids.ShortID]uint64)
	for i := 0; i < int(epoch); i++ {
		providers[ids.ShortID{byte(i + int(epoch))}] = uint64(i)
	}

	return providers
}
