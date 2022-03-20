package evm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/coreth/tests/mocks"
	"github.com/flare-foundation/flare/ids"
)

func TestValidatorsNormalizer_ByEpoch(t *testing.T) {
	normTestData := mocks.NewValidatorsTestDataByEpochs(9, 2)
	copyNormTestData := normTestData.Duplicate()

	validatorsCache := NewValidatorsCache(normTestData, WithCacheSize(128))
	for k, v := range normTestData {
		validatorsCache.cache.Add(k, v)
	}
	normalizer := NewValidatorsNormalizer(validatorsCache)

	t.Run("nominal case", func(t *testing.T) {
		for _, i := range []uint64{0, 1} {
			result, err := normalizer.ByEpoch(i)
			require.NoError(t, err)

			weighted := calcWeightRatio(copyNormTestData[i])
			assert.Equal(t, len(weighted), len(result))
			assert.InDeltaMapValues(t, weighted, result, 0, 0)
		}
	})

	t.Run("handles missing key", func(t *testing.T) {
		_, err := normalizer.ByEpoch(uint64(7))
		require.NoError(t, err)
	})
}

func TestCalcWeightRatio(t *testing.T) {
	validators := mocks.TestValidatorsWeightsData()
	expected := map[ids.ShortID]uint64{
		mocks.TestIdsList[2]: 515007957320493220,
		mocks.TestIdsList[1]: 6465383072520688749,
		mocks.TestIdsList[0]: 6185718675989118760,
		mocks.TestIdsList[3]: 9599604527249266409,
		mocks.TestIdsList[4]: 3417212044008417831,
	}
	weighted := calcWeightRatio(validators)
	assert.InDeltaMapValues(t, weighted, expected, 0, 0)
}
