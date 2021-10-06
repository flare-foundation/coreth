package bitcoin

import (
	"fmt"
	"math"
	"sync"

	"github.com/rs/zerolog"

	"gitlab.com/flarenetwork/coreth/flare/connector/bitcoin"
)

type MultiClient struct {
	log     zerolog.Logger
	clients []bitcoin.API
	cfg     MultiConfig
}

func NewMultiClient(log zerolog.Logger, clients []bitcoin.API, options ...MultiOption) (*MultiClient, error) {

	cfg := DefaultMultiConfig
	for _, option := range options {
		option(&cfg)
	}

	if uint(len(clients)) > cfg.RequiredMatches {
		return nil, fmt.Errorf("insufficient number of clients for required matches (clients: %d, matches: %d)", len(clients), cfg.RequiredMatches)
	}

	m := MultiClient{
		log:     log.With().Str("component", "bitcoin_multi_client").Logger(),
		clients: clients,
		cfg:     cfg,
	}

	return &m, nil
}

func (m *MultiClient) Block(hash [32]byte) (*bitcoin.Block, error) {

	results := make(chan *bitcoin.Block, len(m.clients))
	errors := make(chan error, len(m.clients))
	done := make(chan struct{})
	wg := &sync.WaitGroup{}
	go func() {
		wg.Wait()
		close(results)
		close(errors)
		close(done)
		for range results {
		}
		for err := range errors {
			m.log.Warn().Err(err).Msg("could not get block from client")
		}
	}()

	for _, client := range m.clients {
		wg.Add(1)
		go func(client bitcoin.API) {
			block, err := client.Block(hash)
			if err != nil {
				errors <- err
				return
			}
			results <- block
		}(client)
	}

	var selected *bitcoin.Block
	candidates := make(map[uint64][]*bitcoin.Block)
RequestLoop:
	for {
		select {
		case <-done:
			return nil, fmt.Errorf("insufficient number of matching blocks")
		case block := <-results:
			if block.Hash != hash {
				m.log.Warn().
					Hex("want", block.Hash[:]).
					Hex("have", hash[:]).
					Msg("block hash does not match requested hash")
				continue RequestLoop
			}
			candidates[block.Height] = append(candidates[block.Height], block)
			if uint(len(candidates[block.Height])) >= m.cfg.RequiredMatches {
				selected = block
				break RequestLoop
			}
		}
	}

	for height, blocks := range candidates {
		if height == selected.Height {
			continue
		}
		m.log.Warn().
			Uint64("want", selected.Height).
			Uint64("have", height).
			Int("blocks", len(blocks)).
			Msg("received block(s) with mismatching height")
	}

	lowest := uint64(math.MaxUint64)
	blocks := candidates[selected.Height]
	for _, block := range blocks {
		if block.Confirmations < lowest {
			lowest = block.Confirmations
		}
	}

	selected.Confirmations = lowest

	return selected, nil
}

func (m *MultiClient) Transaction(hash [32]byte, index uint8) (*bitcoin.Transaction, error) {

	results := make(chan *bitcoin.Transaction, len(m.clients))
	errors := make(chan error, len(m.clients))
	done := make(chan struct{})
	wg := &sync.WaitGroup{}
	go func() {
		wg.Wait()
		close(results)
		close(errors)
		close(done)
		for range results {
		}
		for err := range errors {
			m.log.Warn().Err(err).Msg("could not get block from client")
		}
	}()

	for _, client := range m.clients {
		wg.Add(1)
		go func(client bitcoin.API) {
			transaction, err := client.Transaction(hash, index)
			if err != nil {
				errors <- err
				return
			}
			results <- transaction
		}(client)
	}

	var selected *bitcoin.Transaction
	type Key struct {
		block     [32]byte
		recipient string
	}
	candidates := make(map[Key][]*bitcoin.Transaction)
RequestLoop:
	for {
		select {
		case <-done:
			return nil, fmt.Errorf("insufficient number of matching transactions")
		case transaction := <-results:
			if transaction.Hash != hash {
				m.log.Warn().
					Hex("want", transaction.Hash[:]).
					Hex("have", hash[:]).
					Msg("transaction hash does not match requested hash")
				continue RequestLoop
			}
			if transaction.Index != index {
				m.log.Warn().
					Uint8("want", index).
					Uint8("have", transaction.Index).
					Msg("transaction index does not match requested index")
				continue RequestLoop
			}
			key := Key{block: transaction.Block, recipient: transaction.Recipient}
			candidates[key] = append(candidates[key], transaction)
			if uint(len(candidates[key])) >= m.cfg.RequiredMatches {
				selected = transaction
				break RequestLoop
			}
		}
	}

	for key, blocks := range candidates {
		if key.block == selected.Block && key.recipient == selected.Recipient {
			continue
		}
		m.log.Warn().
			Hex("want_block", selected.Block[:]).
			Hex("have_block", key.block[:]).
			Str("want_recipient", selected.Recipient).
			Str("have_recipient", key.recipient).
			Int("blocks", len(blocks)).Msg("received transaction(s) with mismatching block and/or recipient")
	}

	lowest := uint64(math.MaxUint64)
	key := Key{block: selected.Block, recipient: selected.Recipient}
	transactions := candidates[key]
	for _, transaction := range transactions {
		if transaction.Confirmations < lowest {
			lowest = transaction.Confirmations
		}
	}

	selected.Confirmations = lowest

	return selected, nil
}
