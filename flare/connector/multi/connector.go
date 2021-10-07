package multi

import (
	"fmt"
	"strings"
	"sync"

	"github.com/rs/zerolog"
	"gitlab.com/flarenetwork/coreth/flare"
)

type Multi struct {
	log        zerolog.Logger
	connectors []flare.Connector
	cfg        Config
}

func NewMulti(log zerolog.Logger, connectors []flare.Connector, options ...Option) (*Multi, error) {

	cfg := DefaultConfig
	for _, option := range options {
		option(&cfg)
	}

	if cfg.MatchesRequired > uint(len(connectors)) {
		return nil, fmt.Errorf("insufficient number of connectors (required_matches: %d, connectors: %d)", cfg.MatchesRequired, len(connectors))
	}

	m := Multi{
		log:        log.With().Str("component", "multi_connector").Logger(),
		connectors: connectors,
		cfg:        cfg,
	}

	return &m, nil
}

func (m *Multi) ProveAvailability(ret []byte) (bool, error) {
	calls := make([]Call, 0, len(m.connectors))
	for _, connector := range m.connectors {
		calls = append(calls, connector.ProveAvailability)
	}
	return m.execute(ret, calls)
}

func (m *Multi) ProvePayment(ret []byte) (bool, error) {
	calls := make([]Call, 0, len(m.connectors))
	for _, connector := range m.connectors {
		calls = append(calls, connector.ProvePayment)
	}
	return m.execute(ret, calls)
}

func (m *Multi) DisprovePayment(ret []byte) (bool, error) {
	calls := make([]Call, 0, len(m.connectors))
	for _, connector := range m.connectors {
		calls = append(calls, connector.DisprovePayment)
	}
	return m.execute(ret, calls)
}

func (m *Multi) execute(ret []byte, calls []Call) (bool, error) {

	results := make(chan bool, len(calls))
	errors := make(chan error, len(calls))

	wg := &sync.WaitGroup{}
	for _, call := range calls {
		wg.Add(1)
		go func(call Call) {
			defer wg.Done()
			ok, err := call(ret)
			if err != nil {
				errors <- err
			}
			results <- ok
		}(call)
	}

	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()

	accepted := uint(0)
	rejected := uint(0)
	for ok := range results {
		if ok {
			accepted++
		} else {
			rejected++
		}
		if accepted >= m.cfg.MatchesRequired {
			return true, nil
		}
	}

	var messages []string
	for err := range errors {
		messages = append(messages, err.Error())
	}
	if len(messages) > 0 {
		err := fmt.Errorf(strings.Join(messages, ", "))
		return false, err
	}

	return false, nil
}
