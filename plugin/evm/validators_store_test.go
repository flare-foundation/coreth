// (c) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"encoding/binary"
	"errors"
	"io"
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/logging"
)

type ReaderMock struct {
	GetFunc func(key []byte) ([]byte, error)
	HasFunc func(key []byte) (bool, error)
}

func (r *ReaderMock) Get(key []byte) ([]byte, error) {
	return r.GetFunc(key)
}

func (r *ReaderMock) Has(key []byte) (bool, error) {
	return r.HasFunc(key)
}

type WriterMock struct {
	PutFunc    func(key []byte, value []byte) error
	DeleteFunc func(key []byte) error
}

func (w *WriterMock) Put(key []byte, value []byte) error {
	return w.PutFunc(key, value)
}

func (w *WriterMock) Delete(key []byte) error {
	return w.DeleteFunc(key)
}

type EncModeMock struct {
	MarshalFunc    func(v interface{}) ([]byte, error)
	NewEncoderFunc func(w io.Writer) *cbor.Encoder
	EncOptionsFunc func() cbor.EncOptions
}

func (e *EncModeMock) Marshal(v interface{}) ([]byte, error) {
	return e.MarshalFunc(v)
}

func (e *EncModeMock) NewEncoder(_ io.Writer) *cbor.Encoder {
	panic("don't use")
}

func (e *EncModeMock) EncOptions() cbor.EncOptions {
	panic("don't use")
}

type DecModeMock struct {
	UnmarshalFunc func(data []byte, v interface{}) error
}

func (d *DecModeMock) Unmarshal(data []byte, v interface{}) error {
	return d.UnmarshalFunc(data, v)
}

func (d *DecModeMock) Valid(data []byte) error {
	panic("don't use")
}

func (d *DecModeMock) NewDecoder(r io.Reader) *cbor.Decoder {
	panic("don't use")
}

func (d *DecModeMock) DecOptions() cbor.DecOptions {
	panic("don't use")
}

func TestNewValidatorsStore(t *testing.T) {

	read := &ReaderMock{}
	write := &WriterMock{}

	got, err := NewValidatorsStore(logging.NoLog{}, read, write)
	require.NoError(t, err)

	assert.Equal(t, read, got.read)
	assert.Equal(t, write, got.write)
	assert.NotNil(t, got.dec)
	assert.NotNil(t, got.enc)
}

func TestValidatorsStore_ByEpoch(t *testing.T) {

	t.Run("nominal case", func(t *testing.T) {
		t.Parallel()

		wantValidators := map[ids.ShortID]uint64{
			{1}: 100,
			{2}: 200,
			{3}: 300,
		}

		epoch := uint64(1337)
		data := []byte{99}

		dec := &DecModeMock{
			UnmarshalFunc: func(d []byte, v interface{}) error {
				assert.Equal(t, data, d)
				validators, ok := v.(*map[ids.ShortID]uint64)
				require.True(t, ok)
				*validators = wantValidators
				return nil
			},
		}

		read := &ReaderMock{
			GetFunc: func(key []byte) ([]byte, error) {
				e := binary.BigEndian.Uint64(key)
				assert.Equal(t, epoch, e)
				return data, nil
			},
		}

		store := ValidatorsStore{
			log:  logging.NoLog{},
			dec:  dec,
			read: read,
		}

		gotValidators, err := store.ByEpoch(epoch)
		require.NoError(t, err)

		assert.Equal(t, wantValidators, gotValidators)
	})

	t.Run("handles decoding failure", func(t *testing.T) {
		t.Parallel()

		epoch := uint64(1337)
		data := []byte{99}

		dec := &DecModeMock{
			UnmarshalFunc: func(d []byte, v interface{}) error {
				return errors.New("dummy error")
			},
		}

		read := &ReaderMock{
			GetFunc: func(key []byte) ([]byte, error) {
				return data, nil
			},
		}

		store := ValidatorsStore{
			log:  logging.NoLog{},
			dec:  dec,
			read: read,
		}

		_, err := store.ByEpoch(epoch)
		require.Error(t, err)
	})

	t.Run("handles reading failure", func(t *testing.T) {
		t.Parallel()

		epoch := uint64(1337)

		dec := &DecModeMock{
			UnmarshalFunc: func(d []byte, v interface{}) error {
				return nil
			},
		}

		read := &ReaderMock{
			GetFunc: func(key []byte) ([]byte, error) {
				return nil, errors.New("dummy error")
			},
		}

		store := ValidatorsStore{
			log:  logging.NoLog{},
			dec:  dec,
			read: read,
		}

		_, err := store.ByEpoch(epoch)
		require.Error(t, err)
	})
}

func TestValidatorsStore_Persist(t *testing.T) {

	t.Run("nominal case", func(t *testing.T) {
		t.Parallel()

		wantValidators := map[ids.ShortID]uint64{
			{1}: 100,
			{2}: 200,
			{3}: 300,
		}

		epoch := uint64(1337)
		data := []byte{99}

		enc := &EncModeMock{
			MarshalFunc: func(v interface{}) ([]byte, error) {
				gotValidators, ok := v.(map[ids.ShortID]uint64)
				require.True(t, ok)
				assert.Equal(t, wantValidators, gotValidators)
				return data, nil
			},
		}

		write := &WriterMock{
			PutFunc: func(key []byte, d []byte) error {
				e := binary.BigEndian.Uint64(key)
				assert.Equal(t, epoch, e)
				assert.Equal(t, data, d)
				return nil
			},
		}

		store := ValidatorsStore{
			log:   logging.NoLog{},
			enc:   enc,
			write: write,
		}

		err := store.Persist(epoch, wantValidators)
		require.NoError(t, err)
	})

	t.Run("handles encoding failure", func(t *testing.T) {
		t.Parallel()

		wantValidators := map[ids.ShortID]uint64{
			{1}: 100,
			{2}: 200,
			{3}: 300,
		}

		epoch := uint64(1337)

		enc := &EncModeMock{
			MarshalFunc: func(v interface{}) ([]byte, error) {
				return nil, errors.New("dummy error")
			},
		}

		write := &WriterMock{
			PutFunc: func(key []byte, d []byte) error {
				return nil
			},
		}

		store := ValidatorsStore{
			log:   logging.NoLog{},
			enc:   enc,
			write: write,
		}

		err := store.Persist(epoch, wantValidators)
		require.Error(t, err)
	})

	t.Run("handles writing failure", func(t *testing.T) {
		t.Parallel()

		wantValidators := map[ids.ShortID]uint64{
			{1}: 100,
			{2}: 200,
			{3}: 300,
		}

		epoch := uint64(1337)
		data := []byte{99}

		enc := &EncModeMock{
			MarshalFunc: func(v interface{}) ([]byte, error) {
				return data, nil
			},
		}

		write := &WriterMock{
			PutFunc: func(key []byte, d []byte) error {
				return errors.New("dummy error")
			},
		}

		store := ValidatorsStore{
			log:   logging.NoLog{},
			enc:   enc,
			write: write,
		}

		err := store.Persist(epoch, wantValidators)
		require.Error(t, err)
	})

}
