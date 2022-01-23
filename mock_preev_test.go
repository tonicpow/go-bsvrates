package bsvrates

import (
	"context"
	"fmt"

	"github.com/mrz1836/go-preev"
)

// mockPreevValid for mocking requests
type mockPreevBase struct{}

// GetPair is a mock response
func (m *mockPreevBase) GetPair(context.Context, string) (pair *preev.Pair, err error) {
	return
}

// GetPairs is a mock response
func (m *mockPreevBase) GetPairs(context.Context) (pairList *preev.PairList, err error) {
	return
}

// GetTicker is a mock response
func (m *mockPreevBase) GetTicker(context.Context, string) (ticker *preev.Ticker, err error) {
	return
}

// GetTickerHistory is a mock response
func (m *mockPreevBase) GetTickerHistory(context.Context, string, int64, int64, int64) (tickers []*preev.Ticker, err error) {
	return
}

// GetTickers is a mock response
func (m *mockPreevBase) GetTickers(context.Context) (tickerList *preev.TickerList, err error) {
	return
}

// LastRequest is a mock response
func (m *mockPreevBase) LastRequest() *preev.LastRequest {
	return nil
}

// UserAgent is a mock response
func (m *mockPreevBase) UserAgent() string {
	return "default-user-agent"
}

// mockPreevValid for mocking requests
type mockPreevValid struct {
	mockPreevBase
}

// GetTicker is a mock response
func (m *mockPreevValid) GetTicker(_ context.Context, pairID string) (ticker *preev.Ticker, err error) {

	ticker = &preev.Ticker{
		ID:        pairID,
		Timestamp: 1593628860,
		Tx: &preev.Transaction{
			Hash:      "175d87a3656a5d745af9fe9cee6afc0297a83fb317255962c40085eb31f06a4b",
			Timestamp: 1593628871,
		},
		Prices: &preev.PriceSource{
			Ppi: &preev.Price{
				LastPrice: 159.17,
				Volume:    935279,
			},
		},
	}

	return
}

// mockPreevFailed for mocking requests
type mockPreevFailed struct {
	mockPreevBase
}

// GetPair is a mock response
func (m *mockPreevFailed) GetPair(context.Context, string) (pair *preev.Pair, err error) {
	return nil, fmt.Errorf("some error occurred")
}

// GetTicker is a mock response
func (m *mockPreevFailed) GetTicker(_ context.Context, _ string) (ticker *preev.Ticker, err error) {
	return nil, fmt.Errorf("some error occurred")
}

// GetTickers is a mock response
func (m *mockPreevFailed) GetTickers(_ context.Context) (tickerList *preev.TickerList, err error) {
	return nil, fmt.Errorf("some error occurred")
}

// GetPairs is a mock response
func (m *mockPreevFailed) GetPairs(_ context.Context) (pairList *preev.PairList, err error) {
	return nil, fmt.Errorf("some error occurred")
}
