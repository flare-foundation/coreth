package evm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/coreth/tests/mocks"
)

func TestWithCacheSize(t *testing.T) {
	wantCacheSize := uint(42)

	cfg := CacheConfig{}
	WithCacheSize(wantCacheSize)(&cfg)

	assert.Equal(t, wantCacheSize, cfg.CacheSize)
}

func TestValidatorsCache_ByEpoch(t *testing.T) {
	cacheTestData := mocks.NewValidatorsTestDataByEpochs(6, 2)

	validatorsCache := NewValidatorsCache(cacheTestData, WithCacheSize(32))
	for k, v := range cacheTestData {
		validatorsCache.cache.Add(k, v)
	}

	t.Run("nominal case", func(t *testing.T) {
		for _, i := range []uint64{0, 1} {
			result, err := validatorsCache.ByEpoch(i)
			require.NoError(t, err)

			want := cacheTestData[i]
			assert.Equal(t, want, result)
		}
	})

	t.Run("handles missing key", func(t *testing.T) {
		_, err := validatorsCache.ByEpoch(uint64(7))
		require.NoError(t, err)
	})
}
