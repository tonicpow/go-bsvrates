package bsvrates

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// mockPaprikaBase for mocking requests
type mockPaprikaBase struct{}

// GetMarketPrice is a mock response
func (m *mockPaprikaBase) GetMarketPrice(_ context.Context, coinID string) (response *TickerResponse, err error) {
	return
}

// GetBaseAmountAndCurrencyID is a mock response
func (m *mockPaprikaBase) GetBaseAmountAndCurrencyID(currency string, _ float64) (string, float64) {
	return currency, 0.01
}

// GetPriceConversion is a mock response
func (m *mockPaprikaBase) GetPriceConversion(context.Context, string, string, float64) (response *PriceConversionResponse, err error) {
	return
}

// GetHistoricalTickers is a mock response
func (m *mockPaprikaBase) GetHistoricalTickers(context.Context, string, time.Time, time.Time, int,
	tickerQuote, tickerInterval) (response *HistoricalResponse, err error) {
	return
}

// IsAcceptedCurrency is a mock response
func (m *mockPaprikaBase) IsAcceptedCurrency(_ string) bool {
	return true
}

// mockPaprikaValid for mocking requests
type mockPaprikaValid struct {
	mockPaprikaBase
}

// GetMarketPrice is a mock response
func (m *mockPaprikaValid) GetMarketPrice(_ context.Context, coinID string) (response *TickerResponse, err error) {

	response = &TickerResponse{
		BetaValue:         1.39789,
		CirculatingSupply: 18448838,
		ID:                coinID,
		LastRequest: &lastRequest{
			Method:     http.MethodGet,
			StatusCode: http.StatusOK,
		},
		LastUpdated: "2020-07-01T18:36:56Z",
		MaxSupply:   21000000,
		Name:        "Bitcoin SV",
		Quotes: &currency{USD: &quote{
			Price:              158.49415248,
			Volume24h:          719426754.25105,
			Volume24hChange24h: -4.48,
			MarketCap:          2924031833,
		}},
		Rank:        6,
		Symbol:      "BSV",
		TotalSupply: 18448838,
	}

	return
}

// GetPriceConversion is a mock response
func (m *mockPaprikaValid) GetPriceConversion(_ context.Context, baseCurrencyID, quoteCurrencyID string, amount float64) (response *PriceConversionResponse, err error) {

	response = &PriceConversionResponse{
		Amount:                amount,
		BaseCurrencyID:        baseCurrencyID,
		BaseCurrencyName:      "US Dollars",
		BasePriceLastUpdated:  "2020-07-01T22:03:14Z",
		Price:                 0.006331560350007446,
		QuoteCurrencyID:       quoteCurrencyID,
		QuoteCurrencyName:     "Bitcoin SV",
		QuotePriceLastUpdated: "2020-07-01T22:03:14Z",
	}

	return
}

// mockPaprikaFailed for mocking requests
type mockPaprikaFailed struct {
	mockPaprikaBase //nolint:unused,structcheck // this is a mock
}

// GetMarketPrice is a mock response
func (m *mockPaprikaFailed) GetMarketPrice(_ context.Context, _ string) (response *TickerResponse, err error) {
	err = fmt.Errorf("request to paprika fails... 502")
	return
}

// GetBaseAmountAndCurrencyID is a mock response
func (m *mockPaprikaFailed) GetBaseAmountAndCurrencyID(_ string, _ float64) (string, float64) {
	return "", 0
}

// GetPriceConversion is a mock response
func (m *mockPaprikaFailed) GetPriceConversion(_ context.Context, _, _ string, _ float64) (response *PriceConversionResponse, err error) {
	return nil, fmt.Errorf("some error occurred")
}

// GetHistoricalTickers is a mock response
func (m *mockPaprikaFailed) GetHistoricalTickers(_ context.Context, _ string, _, _ time.Time, _ int,
	_ tickerQuote, _ tickerInterval) (response *HistoricalResponse, err error) {
	return nil, fmt.Errorf("some error occurred")
}

// IsAcceptedCurrency is a mock response
func (m *mockPaprikaFailed) IsAcceptedCurrency(_ string) bool {
	return false
}
