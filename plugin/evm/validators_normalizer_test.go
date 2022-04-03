package evm

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"
)

type TestValidatorsNormalizer struct {
	ByEpochFunc         func(uint64) (map[ids.ShortID]uint64, error)
	CalcWeightRatioFunc func(map[ids.ShortID]uint64) map[ids.ShortID]uint64
}

func (t TestValidatorsNormalizer) ByEpoch(e uint64) (map[ids.ShortID]uint64, error) {
	return t.ByEpochFunc(e)
}

func TestValidatorsNormalizer_ByEpoch(t *testing.T) {
	t.Run("nominal case", func(t *testing.T) {
		testValidators := validatorsData
		testWeightRatios := validatorsData

		mock := TestValidatorsNormalizer{
			ByEpochFunc: func(e uint64) (map[ids.ShortID]uint64, error) {
				return testValidators[e], nil
			},
			CalcWeightRatioFunc: func(m map[ids.ShortID]uint64) map[ids.ShortID]uint64 {
				return m
			},
		}
		valCache := NewValidatorsCache(mock, WithCacheSize(128))
		for k, v := range testValidators {
			valCache.cache.Add(k, v)
		}
		valNormalizer := NewValidatorsNormalizer(&logging.Log{}, valCache)
		for i := 0; i < len(validatorsData); i++ {
			got, err := valNormalizer.ByEpoch(uint64(i))
			require.NoError(t, err)

			for key := range testWeightRatios[uint64(i)] {
				assert.Contains(t, got, key)
				assert.Equal(t, got[key], uint64(0))
			}

			assert.Len(t, got, len(testWeightRatios[uint64(i)]))
		}
	})

	t.Run("handles missing key", func(t *testing.T) {
		epoch := uint64(7)
		testValidators := validatorsData
		testWeightRatios := map[ids.ShortID]uint64{}

		mock := TestValidatorsNormalizer{
			ByEpochFunc: func(_ uint64) (map[ids.ShortID]uint64, error) {
				return testWeightRatios, nil
			},
		}
		valCache := NewValidatorsCache(mock, WithCacheSize(128))
		for k, v := range testValidators {
			valCache.cache.Add(k, v)
		}
		valNormalizer := NewValidatorsNormalizer(&logging.Log{}, valCache)
		got, err := valNormalizer.ByEpoch(epoch)
		require.NoError(t, err)
		assert.Empty(t, testWeightRatios)
		assert.Empty(t, got)
	})

	t.Run("handles empty validators map", func(t *testing.T) {
		epoch := uint64(0)
		testValidators := validatorsTestData{
			epoch: map[ids.ShortID]uint64{},
		}

		mock := TestValidatorsNormalizer{
			ByEpochFunc: func(e uint64) (map[ids.ShortID]uint64, error) {
				return testValidators[e], nil
			},
		}
		valCache := NewValidatorsCache(mock, WithCacheSize(128))
		for k, v := range testValidators {
			valCache.cache.Add(k, v)
		}
		valNormalizer := NewValidatorsNormalizer(&logging.Log{}, valCache)

		got, err := valNormalizer.ByEpoch(epoch)
		assert.NoError(t, err)
		assert.Empty(t, testValidators[epoch])
		assert.Empty(t, got)
	})

	t.Run("handles failure to retrieve validators by epoch", func(t *testing.T) {
		mock := TestValidatorsNormalizer{
			ByEpochFunc: func(uint64) (map[ids.ShortID]uint64, error) {
				return nil, errors.New("dummy error")
			},
		}
		valCache := NewValidatorsCache(mock, WithCacheSize(128))
		valNormalizer := NewValidatorsNormalizer(&logging.Log{}, valCache)

		_, err := valNormalizer.ByEpoch(uint64(0))
		assert.Error(t, err)
	})
}

func TestValidatorsNormalizer_CalcWeightRatio(t *testing.T) {
	t.Run("nominal case", func(t *testing.T) {
		testValidators := validatorsData

		mock := TestValidatorsNormalizer{
			CalcWeightRatioFunc: func(vals map[ids.ShortID]uint64) map[ids.ShortID]uint64 {
				for i := range vals {
					vals[i] = uint64(0)
				}
				return vals
			},
		}
		valCache := NewValidatorsCache(mock, WithCacheSize(128))
		for k, v := range testValidators {
			valCache.cache.Add(k, v)
		}
		valNormalizer := NewValidatorsNormalizer(&logging.Log{}, valCache)
		got := valNormalizer.calcWeightRatio(testValidators[uint64(0)])

		for key := range testValidators[uint64(0)] {
			assert.Contains(t, got, key)
			assert.Equal(t, got[key], uint64(0))
		}
		assert.Len(t, got, len(testValidators[uint64(0)]))
	})
}
