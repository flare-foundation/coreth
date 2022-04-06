package evm

import (
	"errors"
	"testing"

	lru "github.com/hashicorp/golang-lru"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"
)

func TestWithCacheSize(t *testing.T) {
	wantCacheSize := uint(42)

	cfg := CacheConfig{}
	WithCacheSize(wantCacheSize)(&cfg)

	assert.Equal(t, wantCacheSize, cfg.CacheSize)
}

func TestValidatorsCache_ByEpoch(t *testing.T) {

	t.Run("nominal case", func(t *testing.T) {

		epoch := uint64(1)

		testValidators := map[ids.ShortID]uint64{
			{1}: 100,
			{2}: 200,
		}

		mock := &ValidatorsRetrieverMock{
			ByEpochFunc: func(e uint64) (map[ids.ShortID]uint64, error) {
				assert.Equal(t, e, epoch)
				return testValidators, nil
			},
		}

		cache, _ := lru.New(1)

		retrieve := ValidatorsCache{
			log:        logging.NoLog{},
			validators: mock,
			cache:      cache,
		}

		got, err := retrieve.ByEpoch(epoch)
		require.NoError(t, err)
		assert.Equal(t, got, testValidators)

		cached, ok := cache.Get(epoch)
		require.True(t, ok)
		assert.Equal(t, cached, testValidators)
	})

	t.Run("returns cached validators", func(t *testing.T) {

		epoch := uint64(1)

		testValidators := map[ids.ShortID]uint64{
			{1}: 100,
			{2}: 200,
		}

		mock := &ValidatorsRetrieverMock{
			ByEpochFunc: func(e uint64) (map[ids.ShortID]uint64, error) {
				return nil, errors.New("should not call retriever")
			},
		}

		cache, _ := lru.New(1)
		cache.Add(epoch, testValidators)

		retrieve := ValidatorsCache{
			log:        logging.NoLog{},
			validators: mock,
			cache:      cache,
		}

		got, err := retrieve.ByEpoch(epoch)
		require.NoError(t, err)
		assert.Equal(t, got, testValidators)

		cached, ok := cache.Get(epoch)
		require.True(t, ok)
		assert.Equal(t, cached, testValidators)
	})

	t.Run("ejects least recently used entry", func(t *testing.T) {

		epoch1 := uint64(1)
		epoch2 := uint64(2)

		testValidators := map[ids.ShortID]uint64{
			{1}: 100,
			{2}: 200,
		}

		mock := &ValidatorsRetrieverMock{
			ByEpochFunc: func(e uint64) (map[ids.ShortID]uint64, error) {
				assert.Equal(t, e, epoch2)
				return testValidators, nil
			},
		}

		cache, _ := lru.New(1)
		cache.Add(epoch1, map[ids.ShortID]uint64{})

		retrieve := ValidatorsCache{
			log:        logging.NoLog{},
			validators: mock,
			cache:      cache,
		}

		got, err := retrieve.ByEpoch(epoch2)
		require.NoError(t, err)
		assert.Equal(t, got, testValidators)

		cached, ok := cache.Get(epoch2)
		require.True(t, ok)
		assert.Equal(t, cached, testValidators)

		_, ok = cache.Get(epoch1)
		require.False(t, ok)
	})

	t.Run("handles retrieve failure correctly", func(t *testing.T) {

		epoch := uint64(1)

		mock := &ValidatorsRetrieverMock{
			ByEpochFunc: func(e uint64) (map[ids.ShortID]uint64, error) {
				return nil, errors.New("dummy error")
			},
		}

		cache, _ := lru.New(1)

		retrieve := ValidatorsCache{
			log:        logging.NoLog{},
			validators: mock,
			cache:      cache,
		}

		_, err := retrieve.ByEpoch(epoch)
		require.Error(t, err)
	})
}
