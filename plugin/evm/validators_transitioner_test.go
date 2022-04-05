package evm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/coreth/params"
	"github.com/flare-foundation/flare/ids"
)

type testValidatorRetriever struct {
	ByEpochFunc func(epoch uint64) (map[ids.ShortID]uint64, error)
}

func (t *testValidatorRetriever) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {
	return t.ByEpochFunc(epoch)
}

func TestValidatorsTransitioner_ByEpoch(t *testing.T) {

	t.Run("nominal case 1: includes check for recursion and number of calls in the absence of cache and compare weights", func(t *testing.T) {
		t.Parallel()

		calls := 0
		testEpoch := uint64(10)
		wantWeight := uint64(45)
		providers := fakeProviders(testEpoch)
		validators := fakeValidators(testEpoch)

		mock := &testValidatorRetriever{
			ByEpochFunc: func(uint64) (map[ids.ShortID]uint64, error) {
				calls++
				return providers, nil
			},
		}

		validatorsMock := &testValidatorRetriever{
			ByEpochFunc: func(uint64) (map[ids.ShortID]uint64, error) {
				return validators, nil
			},
		}

		validatorsTransitioner := NewValidatorsTransitioner(validatorsMock, mock)

		got, err := validatorsTransitioner.ByEpoch(testEpoch)
		require.NoError(t, err)
		assert.Equal(t, providers, got)
		assert.Equal(t, 10, calls)

		// calls is expected to increase by 10 because there is no cache
		got, err = validatorsTransitioner.ByEpoch(testEpoch)
		require.NoError(t, err)
		assert.Equal(t, providers, got)
		assert.Equal(t, 20, calls)

		mock.ByEpochFunc = func(uint64) (map[ids.ShortID]uint64, error) {
			calls++
			return providers, nil
		}

		got, err = validatorsTransitioner.ByEpoch(testEpoch + 1)
		require.NoError(t, err)
		assert.Equal(t, providers, got)
		assert.Equal(t, 31, calls)

		var totalWeight uint64
		for _, u := range got {
			totalWeight += u
		}
		assert.Equal(t, wantWeight, totalWeight)
	})

	t.Run("nominal case 2: use ValidatorsDefault", func(t *testing.T) {
		t.Parallel()

		calls := 0
		testEpoch := uint64(100)
		providers := fakeProviders(testEpoch)

		mock := &testValidatorRetriever{
			ByEpochFunc: func(uint64) (map[ids.ShortID]uint64, error) {
				calls++
				return providers, nil
			},
		}

		validatorsDefault, err := NewValidatorsDefault(params.CostonChainID)
		require.NoError(t, err)

		validatorsTransitioner := NewValidatorsTransitioner(validatorsDefault, mock)

		got, err := validatorsTransitioner.ByEpoch(testEpoch)
		require.NoError(t, err)
		assert.Equal(t, providers, got)
		assert.Equal(t, 100, calls)

		// calls is expected to increase by 1 only because the cache would be used
		got, err = validatorsTransitioner.ByEpoch(testEpoch)
		require.NoError(t, err)
		assert.Equal(t, providers, got)
		assert.Equal(t, 200, calls)

		mock.ByEpochFunc = func(uint64) (map[ids.ShortID]uint64, error) {
			calls++
			return providers, nil
		}

		got, err = validatorsTransitioner.ByEpoch(testEpoch + 1)
		require.NoError(t, err)
		assert.Equal(t, providers, got)
		assert.Equal(t, 301, calls)
	})

	t.Run("epoch less than 1", func(t *testing.T) {
		t.Parallel()

		testEpoch := uint64(0)
		validators := fakeProviders(testEpoch)

		mock := &testValidatorRetriever{
			ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
				assert.Equal(t, testEpoch, epoch)
				return nil, nil
			},
		}

		validatorsMock := &testValidatorRetriever{
			ByEpochFunc: func(uint64) (map[ids.ShortID]uint64, error) {
				return validators, nil
			},
		}

		validatorsTransitioner := NewValidatorsTransitioner(validatorsMock, mock)

		got, err := validatorsTransitioner.ByEpoch(testEpoch)
		require.NoError(t, err)
		assert.Equal(t, validators, got)
	})

	t.Run("case for non nil error", func(t *testing.T) {
		t.Parallel()

		mock := &testValidatorRetriever{
			ByEpochFunc: func(uint64) (map[ids.ShortID]uint64, error) {
				return nil, fmt.Errorf("dummy error")
			},
		}

		validatorsMock := &testValidatorRetriever{
			ByEpochFunc: func(uint64) (map[ids.ShortID]uint64, error) {
				return nil, fmt.Errorf("dummy error")
			},
		}

		validatorsTransitioner := NewValidatorsTransitioner(validatorsMock, mock)

		_, err := validatorsTransitioner.ByEpoch(1)
		require.Error(t, err)
	})

	t.Run("case for nil default validators", func(t *testing.T) {
		t.Parallel()

		epoch := uint64(10)
		providers := fakeProviders(epoch)

		mock := &testValidatorRetriever{
			ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
				return providers, nil
			},
		}

		validatorsMock := &testValidatorRetriever{
			ByEpochFunc: func(uint64) (map[ids.ShortID]uint64, error) {
				return map[ids.ShortID]uint64{}, nil
			},
		}

		validatorsTransitioner := NewValidatorsTransitioner(validatorsMock, mock)

		got, err := validatorsTransitioner.ByEpoch(epoch)
		require.NoError(t, err)
		assert.Equal(t, providers, got)
	})
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
