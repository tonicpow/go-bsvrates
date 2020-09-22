package bsvrates

import "testing"

// TestClient_GetConversion will test the method GetConversion()
func TestClient_GetConversion(t *testing.T) {
	t.Parallel()

	// Set a valid client
	client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevValid{})

	// Test a valid response
	satoshis, provider, err := client.GetConversion(CurrencyDollars, 1)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	} else if satoshis == 0 {
		t.Fatalf("satoshis was 0 for provider: %s", provider.Name())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Name())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Name())
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
		t.Fatalf("satoshis was 0 for provider: %s", provider.Name())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Name())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Name())
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
		t.Fatalf("satoshis was 0 for provider: %s", provider.Name())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Name())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Name())
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
		t.Fatalf("satoshis was 0 for provider: %s", provider.Name())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Name())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Name())
}

// TestClient_GetConversionFailed will test the method GetConversion()
func TestClient_GetConversionFailed(t *testing.T) {
	t.Parallel()

	// Set a valid client (2 valid, 1 invalid)
	client := newMockClient(&mockWOCValid{}, &mockPaprikaFailed{}, &mockPreevValid{})

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
		t.Fatalf("satoshis was 0 for provider: %s", provider.Name())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Name())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Name())
}

// TestClient_GetConversionFailedPreev will test the method GetConversion()
func TestClient_GetConversionFailedPreev(t *testing.T) {
	t.Parallel()

	// Set a valid client (1 valid, 2 invalid)
	client := newMockClient(&mockWOCValid{}, &mockPaprikaFailed{}, &mockPreevFailed{})

	// Test a NON accepted currency
	_, _, rateErr := client.GetConversion(123, 1)
	if rateErr == nil {
		t.Fatalf("expected an error to occur, currency %d is not accepted", 123)
	}

	// Test a valid response (after failing on the first providers)
	satoshis, provider, err := client.GetConversion(CurrencyDollars, 1)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	} else if satoshis == 0 {
		t.Fatalf("satoshis was 0 for provider: %s", provider.Name())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Name())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Name())
}

// TestClient_GetConversionFailedWhatsOnChain will test the method GetConversion()
func TestClient_GetConversionFailedWhatsOnChain(t *testing.T) {
	t.Parallel()

	// Set a valid client (1 valid, 2 invalid)
	client := newMockClient(&mockWOCFailed{}, &mockPaprikaValid{}, &mockPreevFailed{})

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
		t.Fatalf("satoshis was 0 for provider: %s", provider.Name())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Name())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Name())
}

// TestClient_GetConversionFailedAll will test the method GetConversion()
func TestClient_GetConversionFailedAll(t *testing.T) {
	t.Parallel()

	// Set a valid client (1 valid, 2 invalid)
	client := newMockClient(&mockWOCFailed{}, &mockPaprikaFailed{}, &mockPreevFailed{})

	// Test a NON accepted currency
	_, _, rateErr := client.GetConversion(123, 1)
	if rateErr == nil {
		t.Fatalf("expected an error to occur, currency %d is not accepted", 123)
	}

	// Test a valid response (after failing on the first provider)
	satoshis, _, err := client.GetConversion(CurrencyDollars, 1)
	if err == nil {
		t.Fatalf("error was expected but got nil")
	} else if satoshis != 0 {
		t.Fatalf("satoshis should be zero but was %d", satoshis)
	}
}

// TestClient_GetConversionCustomProviders will test the method GetConversion()
func TestClient_GetConversionCustomProviders(t *testing.T) {
	t.Parallel()

	// Set a valid client
	client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevValid{}, ProviderPreev, ProviderWhatsOnChain)

	// Test a valid response
	satoshis, provider, err := client.GetConversion(CurrencyDollars, 1)
	if err != nil {
		t.Fatalf("error occurred: %s", err.Error())
	} else if satoshis == 0 {
		t.Fatalf("satoshis was 0 for provider: %s", provider.Name())
	} else if !provider.IsValid() {
		t.Fatalf("provider: %s was invalid", provider.Name())
	}

	t.Logf("found satoshis: %d from provider: %s", satoshis, provider.Name())
}
