package flare

import (
	"bytes"
)

// Validate is used to check whether a transaction is a valid state connector
// transaction. It takes the transaction's call data as input.
func Validate(data []byte) bool {

	// The call data should at least have a length of 4. The four first bytes
	// encode the function on the smart contract that is called. Without it,
	// the transaction is a simple value transfer which doesn't need to be
	// evaluated.
	if len(data) < 4 {
		return false
	}

	// Currently, there are three functions in the state connector smart
	// contract for which we want to execute the transaction as a state
	// connector call. Any other transaction can be disregarded.
	selector := data[0:4]
	switch {
	case bytes.Equal(selector, SelectorProveAvailability[:]):
		return true
	case bytes.Equal(selector, SelectorProvePayment[:]):
		return true
	case bytes.Equal(selector, SelectorDisprovePayment[:]):
		return true
	default:
		return false
	}
}
