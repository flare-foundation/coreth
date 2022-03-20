// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"reflect"
)

type EVMReturn struct {
	values []interface{}
	err    error
}

func (e *EVMReturn) Decode(values ...interface{}) error {

	if e.err != nil {
		return e.err
	}

	if len(e.values) != len(values) {
		return fmt.Errorf("invalid number of decode values (have: %d, want: %d)", len(values), len(e.values))
	}

	for i, val := range values {

		ret := e.values[i]

		vv := reflect.ValueOf(val)
		if vv.IsNil() {
			continue
		}
		if vv.Kind() != reflect.Ptr {
			return fmt.Errorf("invalid non-pointer (index: %d, type: %T)", i, val)
		}

		iv := reflect.Indirect(vv)
		rv := reflect.ValueOf(ret)
		if iv.Kind() != rv.Kind() {
			return fmt.Errorf("invalid type for return value (have: %T, want: %T)", val, ret)
		}

		iv.Set(rv)
	}

	return nil
}
