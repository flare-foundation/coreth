// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"
)

type ValidatorsRetrieverMock struct {
	ByEpochFunc func(epoch uint64) (map[ids.ShortID]uint64, error)
}

func (v *ValidatorsRetrieverMock) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {
	return v.ByEpochFunc(epoch)
}

type ValidatorsPersisterMock struct {
	PersistFunc func(epoch uint64, validators map[ids.ShortID]uint64) error
}

func (v *ValidatorsPersisterMock) Persist(epoch uint64, validators map[ids.ShortID]uint64) error {
	return v.PersistFunc(epoch, validators)
}

func TestNewValidatorsTransitioner(t *testing.T) {

	validators := &ValidatorsRetrieverMock{}
	providers := &ValidatorsRetrieverMock{}
	active := &ValidatorsRetrieverMock{}
	store := &ValidatorsPersisterMock{}
	size := uint(8)

	got := NewValidatorsTransitioner(logging.NoLog{}, active, validators, providers, store, WithStepSize(size))
	require.NotNil(t, got)
	assert.Equal(t, active, got.active)
	assert.Equal(t, validators, got.validators)
	assert.Equal(t, providers, got.providers)
	assert.Equal(t, store, got.store)
	assert.Equal(t, size, got.cfg.StepSize)
}

func TestValidatorsTransitioner_ByEpoch(t *testing.T) {

	validatorIDs := []ids.ShortID{
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
	}

	weights := []uint64{
		100,
		200,
		300,
		400,
		500,
		600,
	}

	useValidators := func(begin int, end int) map[ids.ShortID]uint64 {
		validators := make(map[ids.ShortID]uint64)
		weights := weights[begin:end]
		for i, validatorID := range validatorIDs[begin:end] {
			weight := weights[i]
			validators[validatorID] = weight
		}
		return validators
	}

	returnValidators := func(begin int, end int) func(epoch uint64) (map[ids.ShortID]uint64, error) {
		return func(epoch uint64) (map[ids.ShortID]uint64, error) {
			return useValidators(begin, end), nil
		}
	}

	retrieveValidators := &ValidatorsRetrieverMock{
		ByEpochFunc: returnValidators(0, 3),
	}

	retrieveProviders := &ValidatorsRetrieverMock{
		ByEpochFunc: returnValidators(3, 6),
	}

	retrieveNothing := &ValidatorsRetrieverMock{
		ByEpochFunc: returnValidators(0, 0),
	}

	retrieveError := &ValidatorsRetrieverMock{
		ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
			return nil, fmt.Errorf("could not retrieve validators")
		},
	}

	tests := []struct {
		name               string
		epoch              uint64
		stepSize           uint
		retrieveValidators ValidatorsRetriever
		retrieveProviders  ValidatorsRetriever
		wantValidators     map[ids.ShortID]uint64
		wantErr            assert.ErrorAssertionFunc
	}{
		{
			name:               "epoch zero",
			epoch:              0,
			retrieveValidators: retrieveValidators,
			retrieveProviders:  retrieveProviders,
			wantValidators:     useValidators(0, 3),
			wantErr:            assert.NoError,
		},
		{
			name:               "no providers",
			epoch:              10,
			retrieveValidators: retrieveValidators,
			retrieveProviders:  retrieveNothing,
			wantErr:            assert.NoError,
		},
		{
			name:               "validators error",
			epoch:              10,
			retrieveValidators: retrieveError,
			retrieveProviders:  retrieveProviders,
			wantErr:            assert.Error,
		},
		{
			name:               "providers error",
			epoch:              10,
			retrieveValidators: retrieveValidators,
			retrieveProviders:  retrieveProviders,
			wantErr:            assert.Error,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			transitioner := &ValidatorsTransitioner{
				log:        logging.NoLog{},
				validators: test.retrieveValidators,
				providers:  test.retrieveProviders,
				cfg:        TransitionConfig{StepSize: test.stepSize},
			}

			gotValidators, err := transitioner.ByEpoch(test.epoch)
			test.wantErr(t, err)

			if err != nil {
				return
			}

			assert.Equal(t, gotValidators, test.wantValidators)
		})
	}
}
