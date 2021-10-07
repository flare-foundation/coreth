package ripple

import (
	"github.com/ybbus/jsonrpc/v2"
)

type RPCClient interface {
	Call(method string, params ...interface{}) (*jsonrpc.RPCResponse, error)
}
