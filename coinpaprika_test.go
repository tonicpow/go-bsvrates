package bsvrates

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"testing"
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
	if req.URL.String() == coinPaprikaBaseURL+"price-converter?base_currency_id="+USDCurrencyID+"&quote_currency_id="+CoinPaprikaQuoteID+"&amount=0.01" {
		resp.StatusCode = http.StatusOK
		resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"base_currency_id":"usd-us-dollars","base_currency_name":"US Dollars","base_price_last_updated":"2020-06-28T19:05:12Z","quote_currency_id":"bsv-bitcoin-sv","quote_currency_name":"Bitcoin SV","quote_price_last_updated":"2020-06-28T19:04:16Z","amount":0.01,"price":0.000062865274346746}`)))
	}

	// Valid
	if req.URL.String() == coinPaprikaBaseURL+"price-converter?base_currency_id="+USDCurrencyID+"&quote_currency_id="+CoinPaprikaQuoteID+"&amount=1" {
		resp.StatusCode = http.StatusOK
		resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"base_currency_id":"usd-us-dollars","base_currency_name":"US Dollars","base_price_last_updated":"2020-06-28T19:07:37Z","quote_currency_id":"bsv-bitcoin-sv","quote_currency_name":"Bitcoin SV","quote_price_last_updated":"2020-06-28T19:05:52Z","amount":1,"price":0.006277681354322026}`)))
	}

	// Valid
	if req.URL.String() == coinPaprikaBaseURL+"price-converter?base_currency_id="+JPYCurrencyID+"&quote_currency_id="+CoinPaprikaQuoteID+"&amount=1" {
		resp.StatusCode = http.StatusOK
		resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(`{"base_currency_id":"jpy-japanese-yen","base_currency_name":"Japanese Yen","base_price_last_updated":"2020-06-28T19:01:09Z","quote_currency_id":"bsv-bitcoin-sv","quote_currency_name":"Bitcoin SV","quote_price_last_updated":"2020-06-28T19:10:17Z","amount":1,"price":0.00005857139480395992}`)))
	}

	// Invalid (error in request)
	if req.URL.String() == coinPaprikaBaseURL+"price-converter?base_currency_id="+USDCurrencyID+"&quote_currency_id="+CoinPaprikaQuoteID+"&amount=501" {
		resp.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(``)))
		return resp, fmt.Errorf(`http timeout`)
	}

	// Invalid (bad gateway)
	if req.URL.String() == coinPaprikaBaseURL+"price-converter?base_currency_id="+USDCurrencyID+"&quote_currency_id="+CoinPaprikaQuoteID+"&amount=502" {
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

	// Create the list of tests
	var tests = []struct {
		currency         string
		amount           float64
		expectedCurrency string
		expectedAmount   float64
	}{
		{"aud", 0.01, AUDCurrencyID, 0.01},
		{"bogus", 0, "", 0},
		{"brl", 0.01, BRLCurrencyID, 0.01},
		{"cad", 0.01, CADCurrencyID, 0.01},
		{"chf", 0.01, CHFCurrencyID, 0.01},
		{"cny", 0.01, CNYCurrencyID, 0.01},
		{"eur", 0.01, EURCurrencyID, 0.01},
		{"gbp", 0.01, GBPCurrencyID, 0.01},
		{"jpy", 0.01, JPYCurrencyID, 1},
		{"krw", 0.01, KRWCurrencyID, 1},
		{"mxn", 0.01, MXNCurrencyID, 0.01},
		{"new", 0.01, NEWCurrencyID, 0.01},
		{"nok", 0.01, NOKCurrencyID, 0.01},
		{"pln", 0.01, PLNCurrencyID, 0.01},
		{"rub", 0.01, RUBCurrencyID, 0.01},
		{"sek", 0.01, SEKCurrencyID, 0.01},
		{"try", 0.01, TRYCurrencyID, 0.01},
		{"twd", 0.01, TWDCurrencyID, 0.01},
		{"usd", 0, USDCurrencyID, 0.01},
		{"usd", 0.01, USDCurrencyID, 0.01},
		{"zar", 0.01, ZARCurrencyID, 0.01},
	}

	// Test all
	for _, test := range tests {
		if currency, amount := client.CoinPaprika.GetBaseAmountAndCurrencyID(test.currency, test.amount); currency != test.expectedCurrency {
			t.Errorf("%s Failed: [%s] inputted currency, got [%s] but expected [%s]", t.Name(), test.currency, currency, test.expectedCurrency)
		} else if amount != test.expectedAmount {
			t.Errorf("%s Failed: [%f] inputted amount, got [%f] but expected [%f]", t.Name(), test.amount, amount, test.expectedAmount)
		}
	}
}

// TestPaprikaClient_GetPriceConversion will test the method GetPriceConversion()
func TestPaprikaClient_GetPriceConversion(t *testing.T) {
	// t.Parallel()

	// New mock client
	client := newMockPaprikaClient(&mockHTTPPaprika{})

	// Create the list of tests
	var tests = []struct {
		baseCurrency  string
		quoteCurrency string
		amount        string
		expectedPrice float64
		expectedError bool
		statusCode    int
	}{
		{USDCurrencyID, CoinPaprikaQuoteID, "0.01", 0.000062865274346746, false, http.StatusOK},
		{USDCurrencyID, CoinPaprikaQuoteID, "1", 0.006277681354322026, false, http.StatusOK},
		{USDCurrencyID, CoinPaprikaQuoteID, "501", 0, true, http.StatusBadRequest},
		{USDCurrencyID, CoinPaprikaQuoteID, "502", 0, true, http.StatusBadGateway},
		{JPYCurrencyID, CoinPaprikaQuoteID, "1", 0.00005857139480395992, false, http.StatusOK},
	}

	// Test all
	for _, test := range tests {
		if output, err := client.CoinPaprika.GetPriceConversion(test.baseCurrency, test.quoteCurrency, test.amount); err == nil && test.expectedError {
			t.Errorf("%s Failed: expected to throw an error, no error [%s] [%s] [%s] inputted", t.Name(), test.baseCurrency, test.quoteCurrency, test.amount)
		} else if err != nil && !test.expectedError {
			t.Errorf("%s Failed: [%s] [%s] [%s] inputted, received: [%v] error [%s]", t.Name(), test.baseCurrency, test.quoteCurrency, test.amount, output, err.Error())
		} else if output != nil && output.Price != test.expectedPrice && !test.expectedError {
			t.Errorf("%s Failed: [%s] [%s] [%s] inputted and [%f] expected, received: [%f]", t.Name(), test.baseCurrency, test.quoteCurrency, test.amount, test.expectedPrice, output.Price)
		} else if output != nil && output.LastRequest.StatusCode != test.statusCode {
			t.Errorf("%s Expected status code to be %d, got %d, [%s] [%s] [%s] inputted", t.Name(), test.statusCode, output.LastRequest.StatusCode, test.baseCurrency, test.quoteCurrency, test.amount)
		}
	}

}

// TestPaprikaClient_GetMarketPrice will test the method GetMarketPrice()
func TestPaprikaClient_GetMarketPrice(t *testing.T) {
	t.Parallel()

	// New mock client
	client := newMockPaprikaClient(&mockHTTPPaprika{})

	// Create the list of tests
	var tests = []struct {
		coinID        string
		expectedPrice float64
		expectedError bool
		statusCode    int
	}{
		{CoinPaprikaQuoteID, 159.190332, false, http.StatusOK},
		{"unknown", 0, true, http.StatusNotFound},
		{"error", 0, true, http.StatusBadGateway},
	}

	// Test all
	for _, test := range tests {
		if output, err := client.CoinPaprika.GetMarketPrice(test.coinID); err == nil && test.expectedError {
			t.Errorf("%s Failed: expected to throw an error, no error [%s] inputted", t.Name(), test.coinID)
		} else if err != nil && !test.expectedError {
			t.Errorf("%s Failed: [%s] inputted, received: [%v] error [%s]", t.Name(), test.coinID, output, err.Error())
		} else if output != nil && output.Quotes != nil && output.Quotes.USD.Price != test.expectedPrice && !test.expectedError {
			t.Errorf("%s Failed: [%s] inputted and [%f] expected, received: [%f]", t.Name(), test.coinID, test.expectedPrice, output.Quotes.USD.Price)
		} else if output != nil && output.LastRequest.StatusCode != test.statusCode {
			t.Errorf("%s Expected status code to be %d, got %d, [%s] inputted", t.Name(), test.statusCode, output.LastRequest.StatusCode, test.coinID)
		}
	}
}

// TestPaprikaClient_IsAcceptedCurrency will test the method IsAcceptedCurrency()
func TestPaprikaClient_IsAcceptedCurrency(t *testing.T) {
	t.Parallel()

	// New mock client
	client := newMockPaprikaClient(&mockHTTPPaprika{})

	// Create the list of tests
	var tests = []struct {
		currency string
		found    bool
	}{
		{"aud", true},
		{"brl", true},
		{"cad", true},
		{"chf", true},
		{"cny", true},
		{"eur", true},
		{"gbp", true},
		{"jpy", true},
		{"krw", true},
		{"mxn", true},
		{"new", true},
		{"nok", true},
		{"pln", true},
		{"rub", true},
		{"sek", true},
		{"try", true},
		{"twd", true},
		{"usd", true},
		{"zar", true},
		{"www", false},
		{"xxx", false},
		{"usa", false},
		{"", false},
	}

	// Test all
	for _, test := range tests {
		if found := client.CoinPaprika.IsAcceptedCurrency(test.currency); found != test.found {
			t.Errorf("%s Failed: [%s] inputted currency, found did not match expected value", t.Name(), test.currency)
		}
	}
}

// TestPriceConversionResponse_GetSatoshi will test the method GetSatoshi()
func TestPriceConversionResponse_GetSatoshi(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	var tests = []struct {
		response      PriceConversionResponse
		satoshi       string
		expectedError bool
	}{
		{PriceConversionResponse{Price: 0}, "0", false},
		{PriceConversionResponse{Price: 1}, "100000000", false},
		{PriceConversionResponse{Price: 0.01}, "1000000", false},
		{PriceConversionResponse{Price: 0.001}, "100000", false},
		{PriceConversionResponse{Price: 0.0001}, "10000", false},
		{PriceConversionResponse{Price: 0.00001}, "1000", false},
		{PriceConversionResponse{Price: 0.000001}, "100", false},
		{PriceConversionResponse{Price: 0.0000001}, "10", false},
		{PriceConversionResponse{Price: 0.00000001}, "1", false},
		{PriceConversionResponse{Price: 0.000000001}, "0", false},
		{PriceConversionResponse{Price: 45627467}, "4562746700000000", false},
		{PriceConversionResponse{Price: math.NaN()}, "", true},
		{PriceConversionResponse{Price: math.Inf(1)}, "", true},
	}

	// Test all
	for _, test := range tests {
		if satoshi, err := test.response.GetSatoshi(); err == nil && test.expectedError {
			t.Errorf("%s Failed: expected to throw an error, no error [%v] inputted", t.Name(), test.response)
		} else if err != nil && !test.expectedError {
			t.Errorf("%s Failed: [%f] inputted, received: [%s] error [%s]", t.Name(), test.response.Amount, satoshi, err.Error())
		} else if satoshi != test.satoshi && !test.expectedError {
			t.Errorf("%s Failed: [%f] inputted and [%s] expected, received: [%s]", t.Name(), test.response.Amount, test.satoshi, satoshi)
		}
	}
}
