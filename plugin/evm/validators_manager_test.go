// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
	"testing"

	"github.com/flare-foundation/flare/ids"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type retrieverMock struct {
	ByEpochFunc func(epoch uint64) (map[ids.ShortID]uint64, error)
}

func (r retrieverMock) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {
	return r.ByEpochFunc(epoch)
}

func TestNewValidatorsManager(t *testing.T) {
	testDefault := make(map[ids.ShortID]uint64)
	testRetriever := &retrieverMock{}

	got := NewValidatorsManager(testDefault, testRetriever, testRetriever)

	require.NotNil(t, got)
	assert.Equal(t, testDefault, got.defaultValidators)
	assert.Equal(t, testRetriever, got.ftsoValidators)
	assert.Equal(t, testRetriever, got.activeValidators)
}

func TestValidatorsManager_DefaultValidators(t *testing.T) {
	testDefault := make(map[ids.ShortID]uint64)

	subject := &ValidatorsManager{
		defaultValidators: testDefault,
	}

	got, err := subject.DefaultValidators()
	require.NoError(t, err)
	assert.Equal(t, testDefault, got)
}

func TestValidatorsManager_FTSOValidators(t *testing.T) {
	testEpoch := uint64(1)
	testValidators := map[ids.ShortID]uint64{
		{13}: 37,
	}

	t.Run("nominal case", func(t *testing.T) {
		t.Parallel()

		testRetriever := &retrieverMock{
			ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
				assert.Equal(t, testEpoch, epoch)

				return testValidators, nil
			},
		}

		subject := &ValidatorsManager{
			ftsoValidators: testRetriever,
		}

		got, err := subject.FTSOValidators(testEpoch)
		require.NoError(t, err)
		assert.Equal(t, testValidators, got)
	})

	t.Run("handles failure to retrieve FTSO validators", func(t *testing.T) {
		t.Parallel()

		testRetriever := &retrieverMock{
			ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
				assert.Equal(t, testEpoch, epoch)

				return nil, errors.New("dummy error")
			},
		}

		subject := &ValidatorsManager{
			ftsoValidators: testRetriever,
		}

		_, err := subject.FTSOValidators(testEpoch)
		assert.Error(t, err)
	})
}

func TestValidatorsManager_ActiveValidators(t *testing.T) {
	testEpoch := uint64(1)
	testValidators := map[ids.ShortID]uint64{
		{13}: 37,
	}

	t.Run("nominal case", func(t *testing.T) {
		t.Parallel()

		testRetriever := &retrieverMock{
			ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
				assert.Equal(t, testEpoch, epoch)

				return testValidators, nil
			},
		}

		subject := &ValidatorsManager{
			activeValidators: testRetriever,
		}

		got, err := subject.ActiveValidators(testEpoch)
		require.NoError(t, err)
		assert.Equal(t, testValidators, got)
	})

	t.Run("handles failure to retrieve active validators", func(t *testing.T) {
		t.Parallel()

		testRetriever := &retrieverMock{
			ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
				assert.Equal(t, testEpoch, epoch)

				return nil, errors.New("dummy error")
			},
		}

		subject := &ValidatorsManager{
			activeValidators: testRetriever,
		}

		_, err := subject.ActiveValidators(testEpoch)
		assert.Error(t, err)
	})
}
