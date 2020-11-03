package bsvrates

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mockHTTPPaprika for mocking requests
type mockHTTPPaprika struct{}

// Do is a mock http request
func (m *mockHTTPPaprika) Do(req *http.Request) (*http.Response, error) {
	resp := new(http.Response)
	resp.StatusCode = http.StatusBadRequest

	// No req found
	if req == nil {
		return resp, fmt.Errorf("missing request")
	}

	// Valid
	if req.URL.String() == coinPaprikaBaseURL+"price-converter?base_currency_id="+USDCurrencyID+"&quote_currency_id="+CoinPaprikaQuoteID+"&amount=0.010000" {
		resp.StatusCode = http.StatusOK
		resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"base_currency_id":"usd-us-dollars","base_currency_name":"US Dollars","base_price_last_updated":"2020-06-28T19:05:12Z","quote_currency_id":"bsv-bitcoin-sv","quote_currency_name":"Bitcoin SV","quote_price_last_updated":"2020-06-28T19:04:16Z","amount":0.01,"price":0.000062865274346746}`)))
	}

	// Valid
	if req.URL.String() == coinPaprikaBaseURL+"price-converter?base_currency_id="+USDCurrencyID+"&quote_currency_id="+CoinPaprikaQuoteID+"&amount=1.000000" {
		resp.StatusCode = http.StatusOK
		resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"base_currency_id":"usd-us-dollars","base_currency_name":"US Dollars","base_price_last_updated":"2020-06-28T19:07:37Z","quote_currency_id":"bsv-bitcoin-sv","quote_currency_name":"Bitcoin SV","quote_price_last_updated":"2020-06-28T19:05:52Z","amount":1,"price":0.006277681354322026}`)))
	}

	// Valid
	if req.URL.String() == coinPaprikaBaseURL+"price-converter?base_currency_id="+JPYCurrencyID+"&quote_currency_id="+CoinPaprikaQuoteID+"&amount=1.000000" {
		resp.StatusCode = http.StatusOK
		resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"base_currency_id":"jpy-japanese-yen","base_currency_name":"Japanese Yen","base_price_last_updated":"2020-06-28T19:01:09Z","quote_currency_id":"bsv-bitcoin-sv","quote_currency_name":"Bitcoin SV","quote_price_last_updated":"2020-06-28T19:10:17Z","amount":1,"price":0.00005857139480395992}`)))
	}

	// Invalid (error in request)
	if req.URL.String() == coinPaprikaBaseURL+"price-converter?base_currency_id="+USDCurrencyID+"&quote_currency_id="+CoinPaprikaQuoteID+"&amount=501.000000" {
		resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(``)))
		return resp, fmt.Errorf(`http timeout`)
	}

	// Invalid (bad gateway)
	if req.URL.String() == coinPaprikaBaseURL+"price-converter?base_currency_id="+USDCurrencyID+"&quote_currency_id="+CoinPaprikaQuoteID+"&amount=502.000000" {
		resp.StatusCode = http.StatusBadGateway
		resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(``)))
		return resp, nil
	}

	//
	// Get Market Price
	//

	// Valid
	if req.URL.String() == coinPaprikaBaseURL+"tickers/"+CoinPaprikaQuoteID {
		resp.StatusCode = http.StatusOK
		resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"id":"bsv-bitcoin-sv","name":"Bitcoin SV","symbol":"BSV","rank":6,"circulating_supply":18446975,"total_supply":18446975,"max_supply":21000000,"beta_value":1.39127,"last_updated":"2020-06-29T16:03:48Z","quotes":{"USD":{"price":159.190332,"volume_24h":790903664.43817,"volume_24h_change_24h":-18.74,"market_cap":2936580074,"market_cap_change_24h":-0.65,"percent_change_15m":0.12,"percent_change_30m":0.36,"percent_change_1h":0.59,"percent_change_6h":0.1,"percent_change_12h":-0.05,"percent_change_24h":-0.65,"percent_change_7d":-9.04,"percent_change_30d":-18.47,"percent_change_1y":-20.35,"ath_price":439.73278365,"ath_date":"2020-01-14T23:36:24Z","percent_from_price_ath":-63.8}}}`)))
	}

	// Invalid
	if req.URL.String() == coinPaprikaBaseURL+"tickers/unknown" {
		resp.StatusCode = http.StatusNotFound
		resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"error":"id not found"}`)))
	}

	// Invalid
	if req.URL.String() == coinPaprikaBaseURL+"tickers/error" {
		resp.StatusCode = http.StatusBadGateway
		resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(``)))
		return resp, fmt.Errorf(`http bad gateway`)
	}

	// Default is valid
	return resp, nil
}

// newMockPaprikaClient returns a client for mocking (using a custom HTTP interface)
func newMockPaprikaClient(httpClient httpInterface) *Client {
	client := NewClient(nil, nil)
	cp := createPaprikaClient(nil, nil)
	cp.HTTPClient = httpClient
	client.CoinPaprika = cp
	return client
}

// TestPaprikaClient_GetBaseAmountAndCurrencyID will test the method GetBaseAmountAndCurrencyID()
func TestPaprikaClient_GetBaseAmountAndCurrencyID(t *testing.T) {
	t.Parallel()

	// New mock client
	client := newMockPaprikaClient(&mockHTTPPaprika{})

	var tests = []struct {
		testCase         string
		currency         string
		amount           float64
		expectedCurrency string
		expectedAmount   float64
	}{
		{"aud", "aud", 0.01, AUDCurrencyID, 0.01},
		{"unknown currency", "bogus", 0, "", 0},
		{"brl", "brl", 0.01, BRLCurrencyID, 0.01},
		{"cad", "cad", 0.01, CADCurrencyID, 0.01},
		{"chf", "chf", 0.01, CHFCurrencyID, 0.01},
		{"cny", "cny", 0.01, CNYCurrencyID, 0.01},
		{"eur", "eur", 0.01, EURCurrencyID, 0.01},
		{"gbp", "gbp", 0.01, GBPCurrencyID, 0.01},
		{"jpy", "jpy", 0.01, JPYCurrencyID, 1},
		{"krw", "krw", 0.01, KRWCurrencyID, 1},
		{"mxn", "mxn", 0.01, MXNCurrencyID, 0.01},
		{"new", "new", 0.01, NEWCurrencyID, 0.01},
		{"nok", "nok", 0.01, NOKCurrencyID, 0.01},
		{"pln", "pln", 0.01, PLNCurrencyID, 0.01},
		{"rub", "rub", 0.01, RUBCurrencyID, 0.01},
		{"sek", "sek", 0.01, SEKCurrencyID, 0.01},
		{"try", "try", 0.01, TRYCurrencyID, 0.01},
		{"twd", "twd", 0.01, TWDCurrencyID, 0.01},
		{"usd zero", usd, 0, USDCurrencyID, 0.01},
		{"usd penny", usd, 0.01, USDCurrencyID, 0.01},
		{"zar", "zar", 0.01, ZARCurrencyID, 0.01},
	}
	for _, test := range tests {
		t.Run(test.testCase, func(t *testing.T) {
			currencyName, amount := client.CoinPaprika.GetBaseAmountAndCurrencyID(test.currency, test.amount)
			assert.Equal(t, test.expectedAmount, amount)
			assert.Equal(t, test.expectedCurrency, currencyName)
		})
	}
}

// TestPaprikaClient_GetPriceConversion will test the method GetPriceConversion()
func TestPaprikaClient_GetPriceConversion(t *testing.T) {
	t.Parallel()

	// New mock client
	client := newMockPaprikaClient(&mockHTTPPaprika{})

	t.Run("test valid cases", func(t *testing.T) {
		var tests = []struct {
			testCase           string
			baseCurrency       string
			quoteCurrency      string
			amount             float64
			expectedPrice      float64
			expectedStatusCode int
		}{
			{
				"valid usd penny",
				USDCurrencyID,
				CoinPaprikaQuoteID,
				0.01,
				0.000062865274346746,
				http.StatusOK,
			},
			{
				"valid usd dollar",
				USDCurrencyID,
				CoinPaprikaQuoteID,
				1,
				0.006277681354322026,
				http.StatusOK,
			},
			{
				"valid jpy",
				JPYCurrencyID,
				CoinPaprikaQuoteID,
				1,
				0.00005857139480395992,
				http.StatusOK,
			},
		}
		for _, test := range tests {
			t.Run(test.testCase, func(t *testing.T) {
				output, err := client.CoinPaprika.GetPriceConversion(test.baseCurrency, test.quoteCurrency, test.amount)
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, test.expectedPrice, output.Price)
				assert.Equal(t, test.expectedStatusCode, output.LastRequest.StatusCode)
			})
		}
	})

	t.Run("invalid cases", func(t *testing.T) {
		var tests = []struct {
			testCase           string
			baseCurrency       string
			quoteCurrency      string
			amount             float64
			expectedStatusCode int
		}{
			{
				"bad request response",
				USDCurrencyID,
				CoinPaprikaQuoteID,
				501,
				http.StatusBadRequest,
			},
			{
				"bad gateway response",
				USDCurrencyID,
				CoinPaprikaQuoteID,
				502,
				http.StatusBadGateway,
			},
		}
		for _, test := range tests {
			t.Run(test.testCase, func(t *testing.T) {
				output, err := client.CoinPaprika.GetPriceConversion(test.baseCurrency, test.quoteCurrency, test.amount)
				assert.Error(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, test.expectedStatusCode, output.LastRequest.StatusCode)
			})
		}
	})
}

// TestPaprikaClient_GetMarketPrice will test the method GetMarketPrice()
func TestPaprikaClient_GetMarketPrice(t *testing.T) {
	t.Parallel()

	// New mock client
	client := newMockPaprikaClient(&mockHTTPPaprika{})

	t.Run("test valid cases", func(t *testing.T) {
		var tests = []struct {
			testCase           string
			coinID             string
			expectedPrice      float64
			expectedStatusCode int
		}{
			{
				"valid quote",
				CoinPaprikaQuoteID,
				159.190332,
				http.StatusOK,
			},
		}
		for _, test := range tests {
			t.Run(test.testCase, func(t *testing.T) {
				output, err := client.CoinPaprika.GetMarketPrice(test.coinID)
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, test.expectedPrice, output.Quotes.USD.Price)
				assert.Equal(t, test.expectedStatusCode, output.LastRequest.StatusCode)
			})
		}
	})

	t.Run("test invalid cases", func(t *testing.T) {
		var tests = []struct {
			testCase           string
			coinID             string
			expectedStatusCode int
		}{
			{
				"unknown coin id",
				"unknown",
				http.StatusNotFound,
			},
			{
				"bad gateway response",
				"error",
				http.StatusBadGateway,
			},
		}
		for _, test := range tests {
			t.Run(test.testCase, func(t *testing.T) {
				output, err := client.CoinPaprika.GetMarketPrice(test.coinID)
				assert.Error(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, test.expectedStatusCode, output.LastRequest.StatusCode)
			})
		}
	})
}

// TestPaprikaClient_IsAcceptedCurrency will test the method IsAcceptedCurrency()
func TestPaprikaClient_IsAcceptedCurrency(t *testing.T) {
	t.Parallel()

	// New mock client
	client := newMockPaprikaClient(&mockHTTPPaprika{})

	var tests = []struct {
		testCase      string
		currency      string
		expectedFound bool
	}{
		{"aud", "aud", true},
		{"brl", "brl", true},
		{"cad", "cad", true},
		{"chf", "chf", true},
		{"cny", "cny", true},
		{"eur", "eur", true},
		{"gbp", "gbp", true},
		{"jpy", "jpy", true},
		{"krw", "krw", true},
		{"mxn", "mxn", true},
		{"new", "new", true},
		{"nok", "nok", true},
		{"pln", "pln", true},
		{"rub", "rub", true},
		{"sek", "sek", true},
		{"try", "try", true},
		{"twd", "twd", true},
		{"usd", usd, true},
		{"USD", "USD", true},
		{"zar", "zar", true},
		{"www", "www", false},
		{"xxx", "xxx", false},
		{"usa", "usa", false},
		{"no currency", "", false},
		{"unknown currency", "unknown", false},
	}
	for _, test := range tests {
		t.Run(test.testCase, func(t *testing.T) {
			found := client.CoinPaprika.IsAcceptedCurrency(test.currency)
			assert.Equal(t, test.expectedFound, found)
		})
	}
}

// TestPriceConversionResponse_GetSatoshi will test the method GetSatoshi()
func TestPriceConversionResponse_GetSatoshi(t *testing.T) {
	t.Parallel()

	t.Run("test valid cases", func(t *testing.T) {
		var tests = []struct {
			testCase        string
			response        PriceConversionResponse
			expectedSatoshi int64
		}{
			{"zero", PriceConversionResponse{Price: 0}, 0},
			{"1", PriceConversionResponse{Price: 1}, 100000000},
			{"one decimal place", PriceConversionResponse{Price: 0.1}, 10000000},
			{"two decimal places", PriceConversionResponse{Price: 0.01}, 1000000},
			{"three decimal places", PriceConversionResponse{Price: 0.001}, 100000},
			{"four decimal places", PriceConversionResponse{Price: 0.0001}, 10000},
			{"five decimal places", PriceConversionResponse{Price: 0.00001}, 1000},
			{"six decimal places", PriceConversionResponse{Price: 0.000001}, 100},
			{"seven decimal places", PriceConversionResponse{Price: 0.0000001}, 10},
			{"eight decimal places", PriceConversionResponse{Price: 0.00000001}, 1},
			{"nine decimal places", PriceConversionResponse{Price: 0.000000001}, 1},
			{"random digits", PriceConversionResponse{Price: 45627467}, 4562746700000000},
		}
		for _, test := range tests {
			t.Run(test.testCase, func(t *testing.T) {
				satoshi, err := test.response.GetSatoshi()
				assert.NoError(t, err)
				assert.Equal(t, test.expectedSatoshi, satoshi)
			})
		}
	})

	t.Run("test invalid cases", func(t *testing.T) {
		var tests = []struct {
			testCase        string
			response        PriceConversionResponse
			expectedSatoshi int64
		}{
			{"price is nan", PriceConversionResponse{Price: math.NaN()}, 0},
			{"price is inf", PriceConversionResponse{Price: math.Inf(1)}, 0},
		}
		for _, test := range tests {
			t.Run(test.testCase, func(t *testing.T) {
				satoshi, err := test.response.GetSatoshi()
				assert.Error(t, err)
				assert.Equal(t, test.expectedSatoshi, satoshi)
			})
		}
	})
}
