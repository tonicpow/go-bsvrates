package bsvrates

import (
	"context"
	"errors"

	"github.com/mrz1836/go-whatsonchain"
)

// mockWOCBase is the base
type mockWOCBase struct{}

// GetChainInfo is a mock response
func (m *mockWOCBase) GetChainInfo(context.Context) (chainInfo *whatsonchain.ChainInfo, err error) {
	return
}

// GetCirculatingSupply is a mock response
func (m *mockWOCBase) GetCirculatingSupply(context.Context) (supply float64, err error) {
	return
}

// GetExchangeRate is a mock response
func (m *mockWOCBase) GetExchangeRate(context.Context) (rate *whatsonchain.ExchangeRate, err error) {
	return
}

// mockWOCValid for mocking requests
type mockWOCValid struct {
	mockWOCBase
}

// GetExchangeRate is a mock response
func (m *mockWOCValid) GetExchangeRate(_ context.Context) (rate *whatsonchain.ExchangeRate, err error) {

	rate = &whatsonchain.ExchangeRate{
		Rate:     "159.01",
		Currency: CurrencyToName(CurrencyDollars),
	}

	return
}

// mockWOCFailed for mocking requests
type mockWOCFailed struct {
	mockWOCBase
}

// GetExchangeRate is a mock response
func (m *mockWOCFailed) GetExchangeRate(_ context.Context) (rate *whatsonchain.ExchangeRate, err error) {

	return nil, errors.New("some error occurred")
}
