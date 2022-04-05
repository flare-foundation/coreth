// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare/ids"
)

type retrieverMock struct {
	ByEpochFunc func(epoch uint64) (map[ids.ShortID]uint64, error)
}

func (r retrieverMock) ByEpoch(epoch uint64) (map[ids.ShortID]uint64, error) {
	return r.ByEpochFunc(epoch)
}

func TestNewValidatorsManager(t *testing.T) {
	testRetriever := &retrieverMock{}

	got := NewValidatorsManager(testRetriever, testRetriever, testRetriever)

	require.NotNil(t, got)
	assert.Equal(t, testRetriever, got.defaultValidators)
	assert.Equal(t, testRetriever, got.ftsoValidators)
	assert.Equal(t, testRetriever, got.activeValidators)
}

func TestValidatorsManager_DefaultValidators(t *testing.T) {
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
			defaultValidators: testRetriever,
		}

		got, err := subject.DefaultValidators(testEpoch)
		require.NoError(t, err)
		assert.Equal(t, testValidators, got)
	})

	t.Run("handles failure to retrieve default validators", func(t *testing.T) {
		t.Parallel()

		testRetriever := &retrieverMock{
			ByEpochFunc: func(epoch uint64) (map[ids.ShortID]uint64, error) {
				assert.Equal(t, testEpoch, epoch)

				return nil, errors.New("dummy error")
			},
		}

		subject := &ValidatorsManager{
			defaultValidators: testRetriever,
		}

		_, err := subject.DefaultValidators(testEpoch)
		assert.Error(t, err)
	})
}
