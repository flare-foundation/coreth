package evm

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/coreth/tests/mocks"

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

func (t TestValidatorsNormalizer) calcWeightRatio(vals map[ids.ShortID]uint64) map[ids.ShortID]uint64 {
	return t.CalcWeightRatioFunc(vals)
}

func TestValidatorsNormalizer_ByEpoch(t *testing.T) {
	t.Run("nominal case", func(t *testing.T) {
		var (
			testValidators     = mocks.ValidatorsData
			testValidatorsCopy = testValidators.Duplicate()
			calls              int
		)

		mock := TestValidatorsNormalizer{
			ByEpochFunc: func(e uint64) (map[ids.ShortID]uint64, error) {
				return testValidators[e], nil
			},
			CalcWeightRatioFunc: func(vals map[ids.ShortID]uint64) map[ids.ShortID]uint64 {
				for i, _ := range vals {
					vals[i] = uint64(0)
				}
				calls++
				return vals
			},
		}

		valCache := NewValidatorsCache(mock, WithCacheSize(128))
		for k, v := range testValidators {
			valCache.cache.Add(k, v)
		}
		valNormalizer := NewValidatorsNormalizer(&logging.Log{}, valCache)

		for i := 0; i < len(mocks.ValidatorsData); i++ {
			expected := mock.CalcWeightRatioFunc(testValidatorsCopy[uint64(i)])

			result, err := valNormalizer.ByEpoch(uint64(i))
			require.NoError(t, err)

			assert.InDeltaMapValues(t, expected, result, 0, 0)
		}
		assert.Equal(t, 2, calls)
	})

	t.Run("handles missing key", func(t *testing.T) {
		var (
			testValidators        = mocks.ValidatorsData
			epoch          uint64 = 7
			calls          int
		)

		mock := TestValidatorsNormalizer{
			ByEpochFunc: func(e uint64) (map[ids.ShortID]uint64, error) {
				calls++
				return map[ids.ShortID]uint64{}, nil
			},
		}
		expected, _ := mock.ByEpoch(epoch)

		valCache := NewValidatorsCache(mock, WithCacheSize(128))
		for k, v := range testValidators {
			valCache.cache.Add(k, v)
		}
		valNormalizer := NewValidatorsNormalizer(&logging.Log{}, valCache)

		result, err := valNormalizer.ByEpoch(epoch)
		require.NoError(t, err)
		assert.InDeltaMapValues(t, expected, result, 0, 0)
		assert.Len(t, result, 0)
		assert.Equal(t, 2, calls)
	})

	t.Run("handles empty validators map", func(t *testing.T) {
		var (
			epoch          uint64 = 0
			testValidators        = mocks.ValidatorsTestData{
				epoch: map[ids.ShortID]uint64{},
			}
			calls int
		)

		mock := TestValidatorsNormalizer{
			ByEpochFunc: func(uint64) (map[ids.ShortID]uint64, error) {
				calls++
				return testValidators[epoch], nil
			},
		}
		expected, _ := mock.ByEpoch(epoch)

		valCache := NewValidatorsCache(mock, WithCacheSize(128))
		for k, v := range testValidators {
			valCache.cache.Add(k, v)
		}
		valNormalizer := NewValidatorsNormalizer(&logging.Log{}, valCache)

		result, err := valNormalizer.ByEpoch(epoch)
		assert.NoError(t, err)
		assert.InDeltaMapValues(t, expected, result, 0, 0)
		assert.Len(t, result, 0)
		assert.Equal(t, 1, calls)
	})

	t.Run("handles failure to retrieve validators by epoch", func(t *testing.T) {
		var (
			epoch uint64 = 0
			calls int
		)

		mock := TestValidatorsNormalizer{
			ByEpochFunc: func(uint64) (map[ids.ShortID]uint64, error) {
				calls++
				return nil, errors.New("dummy error")
			},
		}

		valCache := NewValidatorsCache(mock, WithCacheSize(128))
		valNormalizer := NewValidatorsNormalizer(&logging.Log{}, valCache)

		_, err := valNormalizer.ByEpoch(epoch)
		assert.Error(t, err)
		assert.Equal(t, 1, calls)
	})
}

func TestValidatorsNormalizer_CalcWeightRatio(t *testing.T) {
	t.Run("nominal case", func(t *testing.T) {
		var (
			testValidators     = mocks.ValidatorsData
			testValidatorsCopy = testValidators.Duplicate()
			calls              int
		)

		mock := TestValidatorsNormalizer{
			CalcWeightRatioFunc: func(vals map[ids.ShortID]uint64) map[ids.ShortID]uint64 {
				for i, _ := range vals {
					vals[i] = uint64(0)
				}
				calls++
				return vals
			},
		}
		expected := mock.calcWeightRatio(testValidatorsCopy[uint64(0)])
		valCache := NewValidatorsCache(mock, WithCacheSize(128))
		for k, v := range testValidators {
			valCache.cache.Add(k, v)
		}
		valNormalizer := NewValidatorsNormalizer(&logging.Log{}, valCache)
		result := valNormalizer.calcWeightRatio(testValidators[uint64(0)])

		assert.InDeltaMapValues(t, expected, result, 0, 0)
		assert.Equal(t, 1, calls)
	})
}
