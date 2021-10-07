package flare

func Register(chain uint32, connector Connector) {
	system.register(chain, connector)
}

func (s *System) register(chain uint32, connector Connector) {
	s.connectors[chain] = connector
}
