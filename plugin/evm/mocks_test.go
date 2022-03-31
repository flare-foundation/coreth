package evm

import (
	"math/rand"

	"github.com/flare-foundation/flare/ids"
)

// validatorsTestData represents a test object that keeps records
// of validators ids and corresponding weights.
type validatorsTestData map[uint64]map[ids.ShortID]uint64

var (
	validatorsData = validatorsTestData{
		uint64(0): map[ids.ShortID]uint64{
			ids.ShortID([20]byte{202, 22, 105, 80, 240, 156, 96, 235, 127, 79, 82, 60, 159, 12, 172, 216, 171, 131, 235, 63}):   6185718675989118760,
			ids.ShortID([20]byte{182, 115, 39, 234, 78, 181, 63, 96, 108, 42, 134, 160, 148, 157, 103, 198, 77, 146, 118, 200}): 6465383072520688749,
			ids.ShortID([20]byte{20, 206, 223, 250, 40, 74, 158, 159, 174, 134, 44, 20, 207, 134, 220, 231, 63, 247, 84, 240}):  515007957320493220,
		},
		uint64(1): map[ids.ShortID]uint64{
			ids.ShortID([20]byte{231, 37, 78, 169, 39, 173, 150, 37, 20, 7, 91, 212, 235, 6, 235, 74, 120, 22, 233, 150}):     9599604527249266409,
			ids.ShortID([20]byte{236, 230, 116, 165, 56, 37, 35, 211, 216, 60, 90, 243, 59, 36, 173, 140, 109, 252, 169, 32}): 3417212044008417831,
		},
	}
)

// genericValidators creates a new ValidatoraTestData populated with random values
// and separated by epochs.
// The function takes the number of epochs to create a starting with 0th.
func genericValidators(numValidators, epochs int) validatorsTestData {
	data := validatorsTestData{}
	if numValidators == 0 {
		return data
	}

	if epochs == 0 {
		epochs = 1
	}
	if epochs > numValidators {
		epochs = numValidators
	}

	remainder := numValidators % epochs
	cuts := (numValidators - remainder) / epochs

	for i := 0; i < epochs; i++ {
		data[uint64(i)] = map[ids.ShortID]uint64{}
		for ci := 0; ci < cuts; ci++ {
			data[uint64(i)][randID()] = rand.Uint64()
		}
		if i == epochs-1 {
			for ri := 0; ri < remainder; ri++ {
				data[uint64(i)][randID()] = rand.Uint64()
			}
		}
	}

	return data
}

// randID creates a new slice of random bytes used as ids.ShortID.
func randID() ids.ShortID {
	id := ids.ShortID{}
	rand.Read(id[:])

	return id
}
