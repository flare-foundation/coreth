package evm

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testEpochs struct {
	DetailsFunc func(epoch uint64) (EpochDetails, error)
}

func (t *testEpochs) Details(epoch uint64) (EpochDetails, error) {
	return t.DetailsFunc(epoch)
}

func TestByTimestamp(t *testing.T) {
	epochs := fakeEpochs(5)

	t.Run("nominal case", func(t *testing.T) {
		callsCount := 0
		mock := &testEpochs{
			DetailsFunc: func(epoch uint64) (EpochDetails, error) {
				callsCount++
				return epochs[int(epoch)], nil
			},
		}

		epochManager := NewEpochsManager(mock)

		// Request last known epoch. Expect 4 iterations forwards before we reach it, and therefore 4 calls to Details.
		epoch, err := epochManager.ByTimestamp(epochs[4].StartTime)
		require.NoError(t, err)
		assert.Equal(t, uint64(3), epoch)
		assert.Equal(t, 4, callsCount)

		// Request first epoch. Expect 4 iteration backwards this time, so a total of 8 calls now.
		epoch, err = epochManager.ByTimestamp(epochs[0].StartTime)
		require.NoError(t, err)
		assert.Equal(t, uint64(0), epoch)
		assert.Equal(t, 8, callsCount)

		// Request central epoch. Expect 2 more iterations forwards.
		epoch, err = epochManager.ByTimestamp(epochs[2].StartTime)
		require.NoError(t, err)
		assert.Equal(t, uint64(1), epoch)
		assert.Equal(t, 10, callsCount)
	})

	t.Run("request first epoch", func(t *testing.T) {
		callsCount := 0
		mock := &testEpochs{
			DetailsFunc: func(epoch uint64) (EpochDetails, error) {
				callsCount++
				return epochs[int(epoch)], nil
			},
		}

		epochManager := NewEpochsManager(mock)

		// Request first known epoch, should immediately retrieve the epoch without any extra iterations.
		epoch, err := epochManager.ByTimestamp(epochs[0].StartTime)
		require.NoError(t, err)
		assert.Equal(t, uint64(0), epoch)
		assert.Equal(t, 1, callsCount)
	})

	t.Run("request last epoch", func(t *testing.T) {
		callsCount := 0
		mock := &testEpochs{
			DetailsFunc: func(epoch uint64) (EpochDetails, error) {
				callsCount++
				return epochs[int(epoch)], nil
			},
		}

		epochManager := NewEpochsManager(mock)

		// Request last known epoch. Expect 4 iterations forwards before we reach it, and therefore 4 calls to Details.
		epoch, err := epochManager.ByTimestamp(epochs[4].StartTime)
		require.NoError(t, err)
		assert.Equal(t, uint64(3), epoch)
		assert.Equal(t, 4, callsCount)
	})

	t.Run("request between first and second epoch", func(t *testing.T) {
		callsCount := 0
		mock := &testEpochs{
			DetailsFunc: func(epoch uint64) (EpochDetails, error) {
				callsCount++
				return epochs[int(epoch)], nil
			},
		}

		epochManager := NewEpochsManager(mock)

		// Request 1 hour after the start of the first epoch. Since the epochs last 1 day each in this test, we should get the first one here.
		epoch, err := epochManager.ByTimestamp(epochs[0].StartTime+3600)
		require.NoError(t, err)
		assert.Equal(t, uint64(0), epoch)
		assert.Equal(t, 1, callsCount)
	})

	t.Run("handles failure to retrieve epoch", func(t *testing.T) {
		mock := &testEpochs{
			DetailsFunc: func(epoch uint64) (EpochDetails, error) {
					return EpochDetails{}, fmt.Errorf("dummy error")
			},
		}

		epochManager := NewEpochsManager(mock)

		// Request timestamp that is beyond the last known timestamp.
		_, err := epochManager.ByTimestamp(0)
		assert.Error(t, err)
	})
}

func fakeEpochs(n int) []EpochDetails {
	var ee []EpochDetails
	for i := 0; i < n; i++ {
		ee = append(ee, EpochDetails{
			StartTime: big.NewInt(time.Date(2022, time.January, i, 0, 0, 0, 0, time.UTC).Unix()).Uint64(),
			EndTime:   big.NewInt(time.Date(2022, time.January, i+1, 0, 0, 0, 0, time.UTC).Unix()).Uint64(),
		})
	}
	return ee
}