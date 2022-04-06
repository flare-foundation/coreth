// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb"

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

	retrieveDefault := &ValidatorsRetrieverMock{}
	retrieveFTSO := &ValidatorsRetrieverMock{}
	retrieveActive := &ValidatorsRetrieverMock{}
	store := &ValidatorsPersisterMock{}
	size := uint(8)

	got := NewValidatorsTransitioner(logging.NoLog{}, retrieveDefault, retrieveFTSO, retrieveActive, store, WithStepSize(size))
	require.NotNil(t, got)
	assert.Equal(t, retrieveDefault, got.retrieveDefault)
	assert.Equal(t, retrieveFTSO, got.retrieveFTSO)
	assert.Equal(t, retrieveActive, got.retrieveActive)
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
		{7},
		{8},
		{9},
	}

	weights := []uint64{
		100,
		200,
		300,
		400,
		500,
		600,
		700,
		800,
		900,
	}

	useValidators := func(begin int, end int) map[ids.ShortID]uint64 {
		validators := make(map[ids.ShortID]uint64)
		weights := weights[begin:end]
		for i, validatorID := range validatorIDs[begin:end] {
			validators[validatorID] = weights[i]
		}
		return validators
	}

	returnValidators := func(validators map[ids.ShortID]uint64) func(epoch uint64) (map[ids.ShortID]uint64, error) {
		return func(epoch uint64) (map[ids.ShortID]uint64, error) {
			return validators, nil
		}
	}

	returnError := func() func(epoch uint64) (map[ids.ShortID]uint64, error) {
		return func(epoch uint64) (map[ids.ShortID]uint64, error) {
			return nil, errors.New("dummy error")
		}
	}

	returnNotFound := func() func(epoch uint64) (map[ids.ShortID]uint64, error) {
		return func(epoch uint64) (map[ids.ShortID]uint64, error) {
			return nil, leveldb.ErrNotFound
		}
	}

	_ = returnError

	t.Run("epoch zero returns default validators", func(t *testing.T) {

		epoch := uint64(0)
		stepSize := uint(1)

		defaultValidators := useValidators(0, 3)
		otherValidators := useValidators(3, 6)
		ignoredValidators := useValidators(6, 9)

		wantValidators := validatorIDs[0:3]

		retrieveDefault := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(defaultValidators),
		}

		retrieveFTSO := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(otherValidators),
		}

		retrieveActive := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(ignoredValidators),
		}

		store := &ValidatorsPersisterMock{
			PersistFunc: func(epoch uint64, validators map[ids.ShortID]uint64) error {
				return errors.New("dummy error")
			},
		}

		transition := ValidatorsTransitioner{
			log:             logging.NoLog{},
			retrieveDefault: retrieveDefault,
			retrieveFTSO:    retrieveFTSO,
			retrieveActive:  retrieveActive,
			store:           store,
			cfg:             TransitionConfig{StepSize: stepSize},
		}

		gotValidators, err := transition.ByEpoch(epoch)
		require.NoError(t, err)

		assert.Len(t, gotValidators, len(wantValidators))
		for _, validatorID := range wantValidators {
			assert.Contains(t, gotValidators, validatorID)
		}
	})

	t.Run("no FTSO validators returns default validators", func(t *testing.T) {

		epoch := uint64(1)
		stepSize := uint(1)

		noValidators := useValidators(0, 0)
		defaultValidators := useValidators(0, 3)
		previousValidators := useValidators(3, 6)

		wantValidators := validatorIDs[0:3]

		retrieveDefault := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(defaultValidators),
		}

		retrieveFTSO := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(noValidators),
		}

		retrieveActive := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(previousValidators),
		}

		store := &ValidatorsPersisterMock{
			PersistFunc: func(epoch uint64, validators map[ids.ShortID]uint64) error {
				return errors.New("dummy error")
			},
		}

		transition := ValidatorsTransitioner{
			log:             logging.NoLog{},
			retrieveDefault: retrieveDefault,
			retrieveFTSO:    retrieveFTSO,
			retrieveActive:  retrieveActive,
			store:           store,
			cfg:             TransitionConfig{StepSize: stepSize},
		}

		gotValidators, err := transition.ByEpoch(epoch)
		require.NoError(t, err)

		assert.Len(t, gotValidators, len(wantValidators))
		for _, validatorID := range wantValidators {
			assert.Contains(t, gotValidators, validatorID)
		}
	})

	t.Run("no previous validators recurses into self", func(t *testing.T) {

		epoch := uint64(1)
		stepSize := uint(1)

		defaultValidators := useValidators(0, 3)
		ftsoValidators := useValidators(3, 6)

		wantValidators := []ids.ShortID{
			validatorIDs[0],
			validatorIDs[1],
			validatorIDs[3],
			validatorIDs[4],
			validatorIDs[5],
		}

		retrieveDefault := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(defaultValidators),
		}

		retrieveFTSO := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(ftsoValidators),
		}

		retrieveActive := &ValidatorsRetrieverMock{
			ByEpochFunc: returnNotFound(),
		}

		store := &ValidatorsPersisterMock{
			PersistFunc: func(e uint64, _ map[ids.ShortID]uint64) error {
				assert.Equal(t, epoch, e)
				return nil
			},
		}

		transition := ValidatorsTransitioner{
			log:             logging.NoLog{},
			retrieveDefault: retrieveDefault,
			retrieveFTSO:    retrieveFTSO,
			retrieveActive:  retrieveActive,
			store:           store,
			cfg:             TransitionConfig{StepSize: stepSize},
		}

		gotValidators, err := transition.ByEpoch(epoch)
		require.NoError(t, err)

		assert.Len(t, gotValidators, len(wantValidators))
		for _, validatorID := range wantValidators {
			assert.Contains(t, gotValidators, validatorID)
		}
	})

	t.Run("no default validators in previous validators returns FTSO validators", func(t *testing.T) {

		epoch := uint64(1)
		stepSize := uint(1)

		defaultValidators := useValidators(0, 3)
		ftsoValidators := useValidators(6, 9)
		previousValidators := useValidators(3, 6)

		wantValidators := validatorIDs[6:9]

		retrieveDefault := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(defaultValidators),
		}

		retrieveFTSO := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(ftsoValidators),
		}

		retrieveActive := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(previousValidators),
		}

		store := &ValidatorsPersisterMock{
			PersistFunc: func(epoch uint64, validators map[ids.ShortID]uint64) error {
				return errors.New("dummy error")
			},
		}

		transition := ValidatorsTransitioner{
			log:             logging.NoLog{},
			retrieveDefault: retrieveDefault,
			retrieveFTSO:    retrieveFTSO,
			retrieveActive:  retrieveActive,
			store:           store,
			cfg:             TransitionConfig{StepSize: stepSize},
		}

		gotValidators, err := transition.ByEpoch(epoch)
		require.NoError(t, err)

		assert.Len(t, gotValidators, len(wantValidators))
		for _, validatorID := range wantValidators {
			assert.Contains(t, gotValidators, validatorID)
		}
	})

	t.Run("removing all default validators returns FTSO validators", func(t *testing.T) {

		epoch := uint64(1)
		stepSize := uint(1)

		defaultValidators := useValidators(0, 3)
		ftsoValidators := useValidators(5, 9)
		previousValidators := useValidators(2, 5)

		wantValidators := validatorIDs[5:9]

		retrieveDefault := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(defaultValidators),
		}

		retrieveFTSO := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(ftsoValidators),
		}

		retrieveActive := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(previousValidators),
		}

		store := &ValidatorsPersisterMock{
			PersistFunc: func(epoch uint64, validators map[ids.ShortID]uint64) error {
				return errors.New("dummy error")
			},
		}

		transition := ValidatorsTransitioner{
			log:             logging.NoLog{},
			retrieveDefault: retrieveDefault,
			retrieveFTSO:    retrieveFTSO,
			retrieveActive:  retrieveActive,
			store:           store,
			cfg:             TransitionConfig{StepSize: stepSize},
		}

		gotValidators, err := transition.ByEpoch(epoch)
		require.NoError(t, err)

		assert.Len(t, gotValidators, len(wantValidators))
		for _, validatorID := range wantValidators {
			assert.Contains(t, gotValidators, validatorID)
		}
	})

	t.Run("don't remove default validator when insufficient FTSOs", func(t *testing.T) {

		epoch := uint64(1)
		stepSize := uint(1)

		defaultValidators := useValidators(0, 3)
		ftsoValidators := useValidators(3, 5)
		previousValidators := useValidators(0, 1)

		wantValidators := []ids.ShortID{
			validatorIDs[0],
			validatorIDs[3],
			validatorIDs[4],
		}

		retrieveDefault := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(defaultValidators),
		}

		retrieveFTSO := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(ftsoValidators),
		}

		retrieveActive := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(previousValidators),
		}

		store := &ValidatorsPersisterMock{
			PersistFunc: func(e uint64, vs map[ids.ShortID]uint64) error {
				assert.Equal(t, epoch, e)
				assert.Len(t, vs, len(wantValidators))
				for _, validatorID := range wantValidators {
					assert.Contains(t, vs, validatorID)
				}
				return nil
			},
		}

		transition := ValidatorsTransitioner{
			log:             logging.NoLog{},
			retrieveDefault: retrieveDefault,
			retrieveFTSO:    retrieveFTSO,
			retrieveActive:  retrieveActive,
			store:           store,
			cfg:             TransitionConfig{StepSize: stepSize},
		}

		gotValidators, err := transition.ByEpoch(epoch)
		require.NoError(t, err)

		assert.Len(t, gotValidators, len(wantValidators))
		for _, validatorID := range wantValidators {
			assert.Contains(t, gotValidators, validatorID)
		}
	})

	t.Run("remove one default validator at start", func(t *testing.T) {

		epoch := uint64(1)
		stepSize := uint(1)

		defaultValidators := useValidators(0, 3)
		ftsoValidators := useValidators(3, 6)
		previousValidators := useValidators(0, 3)

		wantValidators := []ids.ShortID{
			validatorIDs[0],
			validatorIDs[1],
			validatorIDs[3],
			validatorIDs[4],
			validatorIDs[5],
		}

		retrieveDefault := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(defaultValidators),
		}

		retrieveFTSO := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(ftsoValidators),
		}

		retrieveActive := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(previousValidators),
		}

		store := &ValidatorsPersisterMock{
			PersistFunc: func(e uint64, vs map[ids.ShortID]uint64) error {
				assert.Equal(t, epoch, e)
				assert.Len(t, vs, len(wantValidators))
				for _, validatorID := range wantValidators {
					assert.Contains(t, vs, validatorID)
				}
				return nil
			},
		}

		transition := ValidatorsTransitioner{
			log:             logging.NoLog{},
			retrieveDefault: retrieveDefault,
			retrieveFTSO:    retrieveFTSO,
			retrieveActive:  retrieveActive,
			store:           store,
			cfg:             TransitionConfig{StepSize: stepSize},
		}

		gotValidators, err := transition.ByEpoch(epoch)
		require.NoError(t, err)

		assert.Len(t, gotValidators, len(wantValidators))
		for _, validatorID := range wantValidators {
			assert.Contains(t, gotValidators, validatorID)
		}
	})

	t.Run("remove one default validator at middle", func(t *testing.T) {

		epoch := uint64(1)
		stepSize := uint(1)

		defaultValidators := useValidators(0, 3)
		ftsoValidators := useValidators(3, 6)
		previousValidators := useValidators(0, 2)

		wantValidators := []ids.ShortID{
			validatorIDs[0],
			validatorIDs[3],
			validatorIDs[4],
			validatorIDs[5],
		}

		retrieveDefault := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(defaultValidators),
		}

		retrieveFTSO := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(ftsoValidators),
		}

		retrieveActive := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(previousValidators),
		}

		store := &ValidatorsPersisterMock{
			PersistFunc: func(e uint64, vs map[ids.ShortID]uint64) error {
				assert.Equal(t, epoch, e)
				assert.Len(t, vs, len(wantValidators))
				for _, validatorID := range wantValidators {
					assert.Contains(t, vs, validatorID)
				}
				return nil
			},
		}

		transition := ValidatorsTransitioner{
			log:             logging.NoLog{},
			retrieveDefault: retrieveDefault,
			retrieveFTSO:    retrieveFTSO,
			retrieveActive:  retrieveActive,
			store:           store,
			cfg:             TransitionConfig{StepSize: stepSize},
		}

		gotValidators, err := transition.ByEpoch(epoch)
		require.NoError(t, err)

		assert.Len(t, gotValidators, len(wantValidators))
		for _, validatorID := range wantValidators {
			assert.Contains(t, gotValidators, validatorID)
		}
	})

	t.Run("remove two default validator at start", func(t *testing.T) {

		epoch := uint64(1)
		stepSize := uint(2)

		defaultValidators := useValidators(0, 5)
		ftsoValidators := useValidators(5, 9)
		previousValidators := useValidators(0, 5)

		wantValidators := []ids.ShortID{
			validatorIDs[0],
			validatorIDs[1],
			validatorIDs[2],
			validatorIDs[5],
			validatorIDs[6],
			validatorIDs[7],
			validatorIDs[8],
		}

		retrieveDefault := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(defaultValidators),
		}

		retrieveFTSO := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(ftsoValidators),
		}

		retrieveActive := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(previousValidators),
		}

		store := &ValidatorsPersisterMock{
			PersistFunc: func(e uint64, vs map[ids.ShortID]uint64) error {
				assert.Equal(t, epoch, e)
				assert.Len(t, vs, len(wantValidators))
				for _, validatorID := range wantValidators {
					assert.Contains(t, vs, validatorID)
				}
				return nil
			},
		}

		transition := ValidatorsTransitioner{
			log:             logging.NoLog{},
			retrieveDefault: retrieveDefault,
			retrieveFTSO:    retrieveFTSO,
			retrieveActive:  retrieveActive,
			store:           store,
			cfg:             TransitionConfig{StepSize: stepSize},
		}

		gotValidators, err := transition.ByEpoch(epoch)
		require.NoError(t, err)

		assert.Len(t, gotValidators, len(wantValidators))
		for _, validatorID := range wantValidators {
			assert.Contains(t, gotValidators, validatorID)
		}
	})

	t.Run("remove two default validator at middle", func(t *testing.T) {

		epoch := uint64(1)
		stepSize := uint(2)

		defaultValidators := useValidators(0, 5)
		ftsoValidators := useValidators(5, 9)
		previousValidators := useValidators(0, 3)

		wantValidators := []ids.ShortID{
			validatorIDs[0],
			validatorIDs[5],
			validatorIDs[6],
			validatorIDs[7],
			validatorIDs[8],
		}

		retrieveDefault := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(defaultValidators),
		}

		retrieveFTSO := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(ftsoValidators),
		}

		retrieveActive := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(previousValidators),
		}

		store := &ValidatorsPersisterMock{
			PersistFunc: func(e uint64, vs map[ids.ShortID]uint64) error {
				assert.Equal(t, epoch, e)
				assert.Len(t, vs, len(wantValidators))
				for _, validatorID := range wantValidators {
					assert.Contains(t, vs, validatorID)
				}
				return nil
			},
		}

		transition := ValidatorsTransitioner{
			log:             logging.NoLog{},
			retrieveDefault: retrieveDefault,
			retrieveFTSO:    retrieveFTSO,
			retrieveActive:  retrieveActive,
			store:           store,
			cfg:             TransitionConfig{StepSize: stepSize},
		}

		gotValidators, err := transition.ByEpoch(epoch)
		require.NoError(t, err)

		assert.Len(t, gotValidators, len(wantValidators))
		for _, validatorID := range wantValidators {
			assert.Contains(t, gotValidators, validatorID)
		}
	})

	t.Run("remove one default validator when insufficient for two", func(t *testing.T) {

		epoch := uint64(1)
		stepSize := uint(2)

		defaultValidators := useValidators(0, 5)
		ftsoValidators := useValidators(5, 8)
		previousValidators := useValidators(0, 3)

		wantValidators := []ids.ShortID{
			validatorIDs[0],
			validatorIDs[1],
			validatorIDs[5],
			validatorIDs[6],
			validatorIDs[7],
		}

		retrieveDefault := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(defaultValidators),
		}

		retrieveFTSO := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(ftsoValidators),
		}

		retrieveActive := &ValidatorsRetrieverMock{
			ByEpochFunc: returnValidators(previousValidators),
		}

		store := &ValidatorsPersisterMock{
			PersistFunc: func(e uint64, vs map[ids.ShortID]uint64) error {
				assert.Equal(t, epoch, e)
				assert.Len(t, vs, len(wantValidators))
				for _, validatorID := range wantValidators {
					assert.Contains(t, vs, validatorID)
				}
				return nil
			},
		}

		transition := ValidatorsTransitioner{
			log:             logging.NoLog{},
			retrieveDefault: retrieveDefault,
			retrieveFTSO:    retrieveFTSO,
			retrieveActive:  retrieveActive,
			store:           store,
			cfg:             TransitionConfig{StepSize: stepSize},
		}

		gotValidators, err := transition.ByEpoch(epoch)
		require.NoError(t, err)

		assert.Len(t, gotValidators, len(wantValidators))
		for _, validatorID := range wantValidators {
			assert.Contains(t, gotValidators, validatorID)
		}
	})
}
