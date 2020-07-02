package bsvrates

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/mrz1836/go-preev"
	"github.com/mrz1836/go-whatsonchain"
)

// mockWOCValid for mocking requests
type mockWOCValid struct{}

// GetExchangeRate is a mock response
func (m *mockWOCValid) GetExchangeRate() (rate *whatsonchain.ExchangeRate, err error) {

	rate = &whatsonchain.ExchangeRate{
		Rate:     "159.01",
		Currency: CurrencyToName(CurrencyDollars),
	}

	return
}

// mockPaprikaValid for mocking requests
type mockPaprikaValid struct{}

// GetMarketPrice is a mock response
func (m *mockPaprikaValid) GetMarketPrice(coinID string) (response *TickerResponse, err error) {

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

// GetBaseAmountAndCurrencyID is a mock response
func (m *mockPaprikaValid) GetBaseAmountAndCurrencyID(currency string, amount float64) (string, float64) {

	// This is just a mock request

	return currency, 0.01
}

// GetPriceConversion is a mock response
func (m *mockPaprikaValid) GetPriceConversion(baseCurrencyID, quoteCurrencyID, amount string) (response *PriceConversionResponse, err error) {

	floatVal, _ := strconv.ParseFloat(amount, 64)
	response = &PriceConversionResponse{
		Amount:                floatVal,
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

// IsAcceptedCurrency is a mock response
func (m *mockPaprikaValid) IsAcceptedCurrency(currency string) bool {

	// This is just a mock response

	return true
}

// mockPreevValid for mocking requests
type mockPreevValid struct{}

// GetTicker is a mock response
func (m *mockPreevValid) GetTicker(pairID string) (ticker *preev.Ticker, err error) {

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

// GetPair is a mock response
func (m *mockPreevValid) GetPair(pairID string) (pair *preev.Pair, err error) {

	return
}

// GetPairs is a mock response
func (m *mockPreevValid) GetPairs() (pairList *preev.PairList, err error) {

	return
}

// GetTickers is a mock response
func (m *mockPreevValid) GetTickers() (tickerList *preev.TickerList, err error) {

	return
}

// mockPaprikaFailed for mocking requests
type mockPaprikaFailed struct{}

// GetMarketPrice is a mock response
func (m *mockPaprikaFailed) GetMarketPrice(coinID string) (response *TickerResponse, err error) {
	err = fmt.Errorf("request to paprika fails... 502")
	return
}

// GetBaseAmountAndCurrencyID is a mock response
func (m *mockPaprikaFailed) GetBaseAmountAndCurrencyID(currency string, amount float64) (string, float64) {

	return "", 0
}

// GetPriceConversion is a mock response
func (m *mockPaprikaFailed) GetPriceConversion(baseCurrencyID, quoteCurrencyID, amount string) (response *PriceConversionResponse, err error) {

	return
}

// IsAcceptedCurrency is a mock response
func (m *mockPaprikaFailed) IsAcceptedCurrency(currency string) bool {

	return false
}

// newMockClient returns a client for mocking
func newMockClient(wocClient whatsOnChainInterface, paprikaClient coinPaprikaInterface, preevClient preevInterface) *Client {
	client := NewClient(nil, nil)
	client.WhatsOnChain = wocClient
	client.CoinPaprika = paprikaClient
	client.Preev = preevClient
	return client
}

// TestClient_GetRate will test the method GetRate()
func TestClient_GetRate(t *testing.T) {

	// Set a valid client
	client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevValid{})

	// Test a valid response
	rate, provider, err := client.GetRate(CurrencyDollars)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	} else if rate == 0 {
		t.Fatalf("rate was 0 for provider: %s", provider.Name())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Name())
	}

	t.Logf("found rate: %f from provider: %s", rate, provider.Name())
}

// TestClient_GetRateFailed will test the method GetRate()
// This tests for a provider failing, but succeeding on the next provider
func TestClient_GetRateFailed(t *testing.T) {

	// Set a valid client (2 valid, 1 invalid)
	client := newMockClient(&mockWOCValid{}, &mockPaprikaFailed{}, &mockPreevValid{})

	// Test a NON accepted currency
	_, _, rateErr := client.GetRate(123)
	if rateErr == nil {
		t.Fatalf("expected an error to occur, currency: %d is not accepted", 123)
	}

	// Test a valid response (after failing on the first provider)
	rate, provider, err := client.GetRate(CurrencyDollars)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	} else if rate == 0 {
		t.Fatalf("rate was 0 for provider: %s", provider.Name())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Name())
	}

	t.Logf("found rate: %f from provider: %s", rate, provider.Name())
}

// todo: add a test where 2 providers fail

// todo: add a test where all 3 providers fail
