package evm

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare/ids"
)

func TestWithCacheSize(t *testing.T) {
	wantCacheSize := uint(42)

	cfg := CacheConfig{}
	WithCacheSize(wantCacheSize)(&cfg)

	assert.Equal(t, wantCacheSize, cfg.CacheSize)
}

type TestValidatorCache struct {
	ByEpochFunc func(uint64) (map[ids.ShortID]uint64, error)
}

func (t TestValidatorCache) ByEpoch(e uint64) (map[ids.ShortID]uint64, error) {
	return t.ByEpochFunc(e)
}

func TestValidatorsCache_ByEpoch(t *testing.T) {
	t.Skip("broken")
	t.Run("nominal case", func(t *testing.T) {
		var (
			numValidators        = 5
			numEpochs            = 2
			epoch         uint64 = 1
		)

		testValidators := genericValidators(numValidators, numEpochs)

		var calls int
		mock := TestValidatorCache{
			ByEpochFunc: func(e uint64) (map[ids.ShortID]uint64, error) {
				return testValidators[e], nil
			},
		}
		valCache := NewValidatorsCache(mock, WithCacheSize(32))
		for k, v := range testValidators {
			valCache.cache.Add(k, v)
		}
		result, err := valCache.ByEpoch(epoch)
		require.NoError(t, err)
		assert.ElementsMatch(t, result, testValidators[epoch])
		assert.Equal(t, 1, calls)
	})

	t.Run("handles missing key", func(t *testing.T) {
		var (
			numValidators        = 5
			numEpochs            = 2
			epoch         uint64 = 7
			calls         int
		)

		testValidators := genericValidators(numValidators, numEpochs)
		testEpochResult := map[ids.ShortID]uint64{}

		mock := TestValidatorCache{
			ByEpochFunc: func(uint64) (map[ids.ShortID]uint64, error) {
				return testEpochResult, nil
			},
		}
		valCache := NewValidatorsCache(mock, WithCacheSize(32))
		for k, v := range testValidators {
			valCache.cache.Add(k, v)
		}
		got, err := valCache.ByEpoch(epoch)
		require.NoError(t, err)
		assert.Empty(t, got)
		assert.ElementsMatch(t, testEpochResult, got)
		assert.Equal(t, 2, calls)
	})

	t.Run("handles empty validators map", func(t *testing.T) {
		epoch := uint64(0)
		testValidators := validatorsTestData{
			epoch: map[ids.ShortID]uint64{},
		}

		var calls int
		mock := TestValidatorCache{
			ByEpochFunc: func(e uint64) (map[ids.ShortID]uint64, error) {
				return testValidators[e], nil
			},
		}
		valCache := NewValidatorsCache(mock, WithCacheSize(32))
		for k, v := range testValidators {
			valCache.cache.Add(k, v)
		}
		got, err := valCache.ByEpoch(epoch)
		assert.NoError(t, err)
		assert.Empty(t, got)
		assert.ElementsMatch(t, testValidators[epoch], got)
		assert.Equal(t, 1, calls)
	})

	t.Run("handles failure to retrieve validator by epoch", func(t *testing.T) {
		epoch := uint64(9)

		var calls int
		mock := TestValidatorsNormalizer{
			ByEpochFunc: func(uint64) (map[ids.ShortID]uint64, error) {
				return nil, errors.New("dummy error")
			},
		}

		valCache := NewValidatorsCache(mock, WithCacheSize(32))

		_, err := valCache.ByEpoch(epoch)
		assert.Error(t, err)
		assert.Equal(t, 1, calls)
	})
}
