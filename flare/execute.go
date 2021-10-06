package flare

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Execute is used to execute a transaction as a state connector call. It takes
// the transaction's call data and the return data from its first EVM call
// execution as inputs.
func Execute(data []byte, ret []byte) error {

	// First, we check what currency we have and which state connector should
	// thus be used.
	chainID := binary.BigEndian.Uint32(ret[28:32])
	connector, ok := connectors[chainID]
	if !ok {
		return fmt.Errorf("unknown chain (id: %d)", chainID)
	}

	// Second, we double-check if we have a valid state connector call and use
	// the selector to choose the connector function we want to call.
	if len(data) < 4 {
		return fmt.Errorf("insufficient call data (%d < %d)", len(data), 4)
	}

	// If we should do a call, but we already have a result, the operation is
	// supposed to be a no-op.
	accepted, err := store.Accepted(data, ret)
	if err != nil {
		return fmt.Errorf("could not check accepted status: %w", err)
	}
	rejected, err := store.Rejected(data, ret)
	if err != nil {
		return fmt.Errorf("could not check rejected states: %w", err)
	}

	// If the call is not an active one, we return the result that we are
	// already aware of, or an error if we don't know it.
	activator := ret[88:96]
	active := !bytes.Equal(activator, ActivatorEmpty[:])
	if !active && accepted {
		switch {
		case accepted:
			return nil
		case rejected:
			return ErrRejected
		default:
			return fmt.Errorf("unknown result")
		}
	}

	// Otherwise, if the call is an active one, and it was already accepted or
	// rejected previously, this should be a no-op.
	if active && (accepted || rejected) {
		return nil
	}

	// Lastly, at this point, we should execute the function call to the API and
	// store the result accordingly.
	var valid bool
	selector := data[0:4]
	switch {
	case bytes.Equal(selector, SelectorProveAvailability[:]):
		valid, err = connector.ProveAvailability(ret)
	case bytes.Equal(selector, SelectorProvePayment[:]):
		valid, err = connector.ProvePayment(ret)
	case bytes.Equal(selector, SelectorDisprovePayment[:]):
		valid, err = connector.DisprovePayment(ret)
	default:
		return fmt.Errorf("invalid function selector (%x)", selector)
	}
	if err != nil {
		return fmt.Errorf("could not execute state connector call: %w", err)
	}

	// Finally, we store the result depending on whether we saw success or not.
	if valid {
		err = store.Accept(data, ret)
	} else {
		err = store.Reject(data, ret)
	}
	if err != nil {
		return fmt.Errorf("could not store state connector call result: %w", err)
	}

	return nil
}
