package flare

var store Store
var connectors map[uint32]Connector

func Register(chain uint32, connector Connector) {
	connectors[chain] = connector
}
