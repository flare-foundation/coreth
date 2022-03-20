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
	callsCount := 0
	testEpochsMock := &testEpochs{
		DetailsFunc: func(epoch uint64) (EpochDetails, error) {
			callsCount++
			epochDetails1 := EpochDetails{
				StartTime: big.NewInt(time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()).Uint64(),
				EndTime:   big.NewInt(time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC).Unix()).Uint64(),
			}
			epochDetails2 := EpochDetails{
				StartTime: big.NewInt(time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC).Unix()).Uint64(),
				EndTime:   big.NewInt(time.Date(2022, time.January, 3, 0, 0, 0, 0, time.UTC).Unix()).Uint64(),
			}
			epochDetails3 := EpochDetails{
				StartTime: big.NewInt(time.Date(2022, time.January, 3, 0, 0, 0, 0, time.UTC).Unix()).Uint64(),
				EndTime:   big.NewInt(time.Date(2022, time.January, 4, 0, 0, 0, 0, time.UTC).Unix()).Uint64(),
			}
			epochDetails4 := EpochDetails{
				StartTime: big.NewInt(time.Date(2022, time.January, 4, 0, 0, 0, 0, time.UTC).Unix()).Uint64(),
				EndTime:   big.NewInt(time.Date(2022, time.December, 4, 0, 0, 0, 0, time.UTC).Unix()).Uint64(),
			}
			if epoch < 0 {
				return EpochDetails{}, fmt.Errorf("Invalid epoch")
			}
			switch epoch {
			case 0:
				return epochDetails1, nil
			case 1:
				return epochDetails2, nil
			case 2:
				return epochDetails3, nil
			default:
				return epochDetails4, nil
			}
		},
	}

	epochManager := NewEpochsManager(testEpochsMock)

	timeFeb := big.NewInt(time.Date(2022, time.February, 26, 17, 0, 0, 0, time.UTC).Unix()).Uint64()
	epoch, err := epochManager.ByTimestamp(timeFeb)
	require.NoError(t, err)
	assert.Equal(t, uint64(3), epoch)
	assert.Equal(t, 4, callsCount)

	timeJan1 := big.NewInt(time.Date(2022, time.January, 1, 17, 0, 0, 0, time.UTC).Unix()).Uint64()
	epoch, err = epochManager.ByTimestamp(timeJan1)
	require.NoError(t, err)
	assert.Equal(t, uint64(0), epoch)

	timeJan2 := big.NewInt(time.Date(2022, time.January, 2, 17, 0, 0, 0, time.UTC).Unix()).Uint64()
	epoch, err = epochManager.ByTimestamp(timeJan2)
	require.NoError(t, err)
	assert.Equal(t, uint64(1), epoch)
}
