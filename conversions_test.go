package bsvrates

import "testing"

// TestClient_GetConversion will test the method GetConversion()
func TestClient_GetConversion(t *testing.T) {
	t.Parallel()

	// Set a valid client
	client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevValid{}, DefaultProviders)

	// Test a valid response
	satoshis, provider, err := client.GetConversion(CurrencyDollars, 1)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	} else if satoshis == 0 {
		t.Fatalf("satoshis was 0 for provider: %s", provider.Names())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Names())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Names())
}

// TestClient_GetConversionPreev will test the method GetConversion()
func TestClient_GetConversionPreev(t *testing.T) {
	t.Parallel()

	// Set a valid client
	client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevValid{}, ProviderPreev)

	// Test a valid response
	satoshis, provider, err := client.GetConversion(CurrencyDollars, 1)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	} else if satoshis == 0 {
		t.Fatalf("satoshis was 0 for provider: %s", provider.Names())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Names())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Names())
}

// TestClient_GetConversionWhatsOnChain will test the method GetConversion()
func TestClient_GetConversionWhatsOnChain(t *testing.T) {
	t.Parallel()

	// Set a valid client
	client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevValid{}, ProviderWhatsOnChain)

	// Test a valid response
	satoshis, provider, err := client.GetConversion(CurrencyDollars, 1)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	} else if satoshis == 0 {
		t.Fatalf("satoshis was 0 for provider: %s", provider.Names())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Names())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Names())
}

// TestClient_GetConversionCoinPaprika will test the method GetConversion()
func TestClient_GetConversionCoinPaprika(t *testing.T) {
	t.Parallel()

	// Set a valid client
	client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevValid{}, ProviderCoinPaprika)

	// Test a valid response
	satoshis, provider, err := client.GetConversion(CurrencyDollars, 1)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	} else if satoshis == 0 {
		t.Fatalf("satoshis was 0 for provider: %s", provider.Names())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Names())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Names())
}

// TestClient_GetConversionFailed will test the method GetConversion()
func TestClient_GetConversionFailed(t *testing.T) {
	t.Parallel()

	// Set a valid client (2 valid, 1 invalid)
	client := newMockClient(&mockWOCValid{}, &mockPaprikaFailed{}, &mockPreevValid{}, DefaultProviders)

	// Test a NON accepted currency
	_, _, rateErr := client.GetConversion(123, 1)
	if rateErr == nil {
		t.Fatalf("expected an error to occur, currency %d is not accepted", 123)
	}

	// Test a valid response (after failing on the first provider)
	satoshis, provider, err := client.GetConversion(CurrencyDollars, 1)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	} else if satoshis == 0 {
		t.Fatalf("satoshis was 0 for provider: %s", provider.Names())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Names())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Names())
}

// TestClient_GetConversionFailedPreev will test the method GetConversion()
func TestClient_GetConversionFailedPreev(t *testing.T) {
	t.Parallel()

	// Set a valid client (2 valid, 1 invalid)
	client := newMockClient(&mockWOCValid{}, &mockPaprikaFailed{}, &mockPreevFailed{}, ProviderPreev&ProviderCoinPaprika&ProviderWhatsOnChain)

	// Test a NON accepted currency
	_, _, rateErr := client.GetConversion(123, 1)
	if rateErr == nil {
		t.Fatalf("expected an error to occur, currency %d is not accepted", 123)
	}

	// Test a valid response (after failing on the first provider)
	satoshis, provider, err := client.GetConversion(CurrencyDollars, 1)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	} else if satoshis == 0 {
		t.Fatalf("satoshis was 0 for provider: %s", provider.Names())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Names())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Names())
}

// TestClient_GetConversionFailedWhatsOnChain will test the method GetConversion()
func TestClient_GetConversionFailedWhatsOnChain(t *testing.T) {
	t.Parallel()

	// Set a valid client (2 valid, 1 invalid)
	client := newMockClient(&mockWOCFailed{}, &mockPaprikaValid{}, &mockPreevFailed{}, ProviderPreev&ProviderWhatsOnChain&ProviderCoinPaprika)

	// Test a NON accepted currency
	_, _, rateErr := client.GetConversion(123, 1)
	if rateErr == nil {
		t.Fatalf("expected an error to occur, currency %d is not accepted", 123)
	}

	// Test a valid response (after failing on the first provider)
	satoshis, provider, err := client.GetConversion(CurrencyDollars, 1)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	} else if satoshis == 0 {
		t.Fatalf("satoshis was 0 for provider: %s", provider.Names())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Names())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Names())
}

// TestClient_GetConversionCustomProviders will test the method GetConversion()
func TestClient_GetConversionCustomProviders(t *testing.T) {
	t.Parallel()

	// Set a valid client
	client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevValid{}, ProviderPreev&ProviderWhatsOnChain)

	// Test a valid response
	satoshis, provider, err := client.GetConversion(CurrencyDollars, 1)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	} else if satoshis == 0 {
		t.Fatalf("satoshis was 0 for provider: %s", provider.Names())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Names())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Names())
}
