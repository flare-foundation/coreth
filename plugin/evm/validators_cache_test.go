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

func TestNewValidatorsCache(t *testing.T) {
	retrieve := &ValidatorsRetrieverMock{}

	normalize := NewValidatorsCache(logging.NoLog{}, retrieve)
	assert.Equal(t, retrieve, normalize.retrieve)
}

func TestValidatorsCache_ByEpoch(t *testing.T) {

	t.Run("nominal case", func(t *testing.T) {

		epoch := uint64(1)

		want := map[ids.ShortID]uint64{
			{1}: 100,
			{2}: 200,
		}

		mock := &ValidatorsRetrieverMock{
			ByEpochFunc: func(e uint64) (map[ids.ShortID]uint64, error) {
				assert.Equal(t, e, epoch)
				return want, nil
			},
		}

		cache, _ := lru.New(1)

		retrieve := ValidatorsCache{
			log:      logging.NoLog{},
			retrieve: mock,
			cache:    cache,
		}

		got, err := retrieve.ByEpoch(epoch)
		require.NoError(t, err)
		assert.Equal(t, got, want)

		cached, ok := cache.Get(epoch)
		require.True(t, ok)
		assert.Equal(t, cached, want)
	})

	t.Run("returns cached validators", func(t *testing.T) {

		epoch := uint64(1)

		want := map[ids.ShortID]uint64{
			{1}: 100,
			{2}: 200,
		}

		mock := &ValidatorsRetrieverMock{
			ByEpochFunc: func(e uint64) (map[ids.ShortID]uint64, error) {
				return nil, errors.New("should not call retriever")
			},
		}

		cache, _ := lru.New(1)
		cache.Add(epoch, want)

		retrieve := ValidatorsCache{
			log:      logging.NoLog{},
			retrieve: mock,
			cache:    cache,
		}

		got, err := retrieve.ByEpoch(epoch)
		require.NoError(t, err)
		assert.Equal(t, got, want)

		cached, ok := cache.Get(epoch)
		require.True(t, ok)
		assert.Equal(t, cached, want)
	})

	t.Run("ejects least recently used entry", func(t *testing.T) {

		epoch1 := uint64(1)
		epoch2 := uint64(2)

		want := map[ids.ShortID]uint64{
			{1}: 100,
			{2}: 200,
		}

		mock := &ValidatorsRetrieverMock{
			ByEpochFunc: func(e uint64) (map[ids.ShortID]uint64, error) {
				assert.Equal(t, e, epoch2)
				return want, nil
			},
		}

		cache, _ := lru.New(1)
		cache.Add(epoch1, map[ids.ShortID]uint64{})

		retrieve := ValidatorsCache{
			log:      logging.NoLog{},
			retrieve: mock,
			cache:    cache,
		}

		got, err := retrieve.ByEpoch(epoch2)
		require.NoError(t, err)
		assert.Equal(t, want, got)

		cached, ok := cache.Get(epoch2)
		require.True(t, ok)
		assert.Equal(t, want, cached)

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
			log:      logging.NoLog{},
			retrieve: mock,
			cache:    cache,
		}

		_, err := retrieve.ByEpoch(epoch)
		require.Error(t, err)
	})
}
