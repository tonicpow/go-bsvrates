package bsvrates

import (
	"fmt"
	"net/http"
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

// mockWOCFailed for mocking requests
type mockWOCFailed struct{}

// GetExchangeRate is a mock response
func (m *mockWOCFailed) GetExchangeRate() (rate *whatsonchain.ExchangeRate, err error) {

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
func (m *mockPaprikaValid) GetPriceConversion(baseCurrencyID, quoteCurrencyID string, amount float64) (response *PriceConversionResponse, err error) {

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
func (m *mockPaprikaFailed) GetPriceConversion(baseCurrencyID, quoteCurrencyID string, amount float64) (response *PriceConversionResponse, err error) {

	return nil, fmt.Errorf("some error occurred")
}

// IsAcceptedCurrency is a mock response
func (m *mockPaprikaFailed) IsAcceptedCurrency(currency string) bool {

	return false
}

// mockPreevFailed for mocking requests
type mockPreevFailed struct{}

// GetPair is a mock response
func (m *mockPreevFailed) GetPair(pairID string) (pair *preev.Pair, err error) {

	return nil, fmt.Errorf("some error occurred")
}

// GetTicker is a mock response
func (m *mockPreevFailed) GetTicker(pairID string) (ticker *preev.Ticker, err error) {

	return nil, fmt.Errorf("some error occurred")
}

// GetTickers is a mock response
func (m *mockPreevFailed) GetTickers() (tickerList *preev.TickerList, err error) {

	return nil, fmt.Errorf("some error occurred")
}

// GetPairs is a mock response
func (m *mockPreevFailed) GetPairs() (pairList *preev.PairList, err error) {

	return nil, fmt.Errorf("some error occurred")
}

// newMockClient returns a client for mocking
func newMockClient(wocClient whatsOnChainInterface, paprikaClient coinPaprikaInterface, preevClient preevInterface, providers ...Provider) *Client {
	client := NewClient(nil, nil, providers...)
	client.WhatsOnChain = wocClient
	client.CoinPaprika = paprikaClient
	client.Preev = preevClient
	return client
}

// TestClient_GetRate will test the method GetRate()
func TestClient_GetRate(t *testing.T) {
	t.Parallel()

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

// TestClient_GetRatePreev will test the method GetRate()
func TestClient_GetRatePreev(t *testing.T) {
	t.Parallel()

	// Set a valid client
	client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevValid{}, ProviderPreev)

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

// TestClient_GetRateWhatsOnChain will test the method GetRate()
func TestClient_GetRateWhatsOnChain(t *testing.T) {
	t.Parallel()

	// Set a valid client
	client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevValid{}, ProviderWhatsOnChain)

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
	t.Parallel()

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

// TestClient_GetRateFailedPreev will test the method GetRate()
// This tests for a provider failing, but succeeding on the next provider
func TestClient_GetRateFailedPreev(t *testing.T) {
	t.Parallel()

	// Set a valid client (2 valid, 1 invalid)
	client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevFailed{}, ProviderPreev&ProviderCoinPaprika)

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

// TestClient_GetRateFailedWhatsOnChain will test the method GetRate()
// This tests for a provider failing, but succeeding on the next provider
func TestClient_GetRateFailedWhatsOnChain(t *testing.T) {
	t.Parallel()

	// Set a valid client (1 valid, 2 invalid)
	client := newMockClient(&mockWOCFailed{}, &mockPaprikaValid{}, &mockPreevFailed{})

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

// TestClient_GetRateFailedAll will test the method GetRate()
// This tests for a provider failing on all providers
func TestClient_GetRateFailedAll(t *testing.T) {
	t.Parallel()

	// Set a valid client (3 invalid)
	client := newMockClient(&mockWOCFailed{}, &mockPaprikaFailed{}, &mockPreevFailed{})

	// Test a NON accepted currency
	_, _, rateErr := client.GetRate(123)
	if rateErr == nil {
		t.Fatalf("expected an error to occur, currency: %d is not accepted", 123)
	}

	// Test a valid response (after failing on the first provider)
	rate, _, err := client.GetRate(CurrencyDollars)
	if err == nil {
		t.Fatalf("error expected but got nil")
	} else if rate != 0 {
		t.Fatalf("rate should be zero but was: %f", rate)
	}
}

// TestClient_GetRateCustomProviders will test the method GetRate()
func TestClient_GetRateCustomProviders(t *testing.T) {
	t.Parallel()

	// Set a valid client
	client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevValid{}, ProviderPreev, ProviderWhatsOnChain)

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
