package bsvrates

import (
	"testing"
)

// TestProvider_IsValid will test the method IsValid()
func TestProvider_IsValid(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	var tests = []struct {
		input    Provider
		expected bool
	}{
		{0, false},
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{ProviderWhatsOnChain, true},
		{ProviderCoinPaprika, true},
		{ProviderPreev, true},
		{providerLast, false},
	}

	// Test all
	for _, test := range tests {
		if isValid := test.input.IsValid(); isValid != test.expected {
			t.Errorf("%s Failed: IsValid returned a unexpected result: %v", t.Name(), isValid)
		}
	}
}

// TestProvider_Name will test the method Name()
func TestProvider_Name(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	var tests = []struct {
		input    Provider
		expected string
	}{
		{0, ""},
		{1, "WhatsOnChain"},
		{2, "CoinPaprika"},
		{3, "Preev"},
		{4, ""},
		{ProviderWhatsOnChain, "WhatsOnChain"},
		{ProviderCoinPaprika, "CoinPaprika"},
		{ProviderPreev, "Preev"},
		{providerLast, ""},
	}

	// Test all
	for _, test := range tests {
		if name := test.input.Name(); name != test.expected {
			t.Errorf("%s Failed: Name returned a unexpected result: %s", t.Name(), name)
		}
	}
}

// TestProviderToName will test the method ProviderToName()
func TestProviderToName(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	var tests = []struct {
		input    Provider
		expected string
	}{
		{0, ""},
		{1, "WhatsOnChain"},
		{2, "CoinPaprika"},
		{3, "Preev"},
		{4, ""},
		{ProviderWhatsOnChain, "WhatsOnChain"},
		{ProviderCoinPaprika, "CoinPaprika"},
		{ProviderPreev, "Preev"},
		{providerLast, ""},
	}

	// Test all
	for _, test := range tests {
		if name := ProviderToName(test.input); name != test.expected {
			t.Errorf("%s Failed: Name returned a unexpected result: %s", t.Name(), name)
		}
	}
}

// TestCurrency_IsValid will test the method IsValid()
func TestCurrency_IsValid(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	var tests = []struct {
		input    Currency
		expected bool
	}{
		{0, false},
		{1, true},
		{2, true},
		{3, false},
		{CurrencyDollars, true},
		{CurrencyBitcoin, true},
		{currencyLast, false},
	}

	// Test all
	for _, test := range tests {
		if isValid := test.input.IsValid(); isValid != test.expected {
			t.Errorf("%s Failed: IsValid returned a unexpected result: %v", t.Name(), isValid)
		}
	}
}

// TestCurrency_Name will test the method Name()
func TestCurrency_Name(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	var tests = []struct {
		input    Currency
		expected string
	}{
		{0, ""},
		{1, "usd"},
		{2, "bsv"},
		{3, ""},
		{CurrencyDollars, "usd"},
		{CurrencyBitcoin, "bsv"},
		{currencyLast, ""},
	}

	// Test all
	for _, test := range tests {
		if name := test.input.Name(); name != test.expected {
			t.Errorf("%s Failed: Name returned a unexpected result: %s", t.Name(), name)
		}
	}
}

// TestCurrencyToName will test the method CurrencyToName()
func TestCurrencyToName(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	var tests = []struct {
		input    Currency
		expected string
	}{
		{0, ""},
		{1, "usd"},
		{2, "bsv"},
		{3, ""},
		{CurrencyDollars, "usd"},
		{CurrencyBitcoin, "bsv"},
		{currencyLast, ""},
	}

	// Test all
	for _, test := range tests {
		if name := CurrencyToName(test.input); name != test.expected {
			t.Errorf("%s Failed: Name returned a unexpected result: %s", t.Name(), name)
		}
	}
}

// TestCurrencyFromName will test the method CurrencyFromName()
func TestCurrencyFromName(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	var tests = []struct {
		input    string
		expected Currency
	}{
		{"", CurrencyDollars},
		{"usd", CurrencyDollars},
		{"bsv", CurrencyBitcoin},
		{"bogus", CurrencyDollars},
	}

	// Test all
	for _, test := range tests {
		if currency := CurrencyFromName(test.input); currency != test.expected {
			t.Errorf("%s Failed: Currency returned a unexpected result: %d", t.Name(), currency)
		}
	}
}

// TestCurrency_IsAccepted will test the method IsAccepted()
func TestCurrency_IsAccepted(t *testing.T) {
	t.Parallel()

	// Create the list of tests
	var tests = []struct {
		input    Currency
		expected bool
	}{
		{0, false},
		{1, true},
		{2, false},
		{3, false},
		{CurrencyDollars, true},
		{CurrencyBitcoin, false},
		{currencyLast, false},
	}

	// Test all
	for _, test := range tests {
		if isValid := test.input.IsAccepted(); isValid != test.expected {
			t.Errorf("%s Failed: IsAccepted returned a unexpected result: %v", t.Name(), isValid)
		}
	}
}
