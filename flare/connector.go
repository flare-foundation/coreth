package flare

type Connector interface {
	ProveDataAvailabilityPeriodFinality(ret []byte) (bool, error)
	ProvePaymentFinality(ret []byte) (bool, error)
	DisprovePaymentFinality(ret []byte) (bool, error)
}
