package bsvrates

import "testing"

// TestClient_GetConversion will test the method GetConversion()
func TestClient_GetConversion(t *testing.T) {

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

// TestClient_GetConversionFailed will test the method GetConversion()
func TestClient_GetConversionFailed(t *testing.T) {

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

// TestClient_GetConversionCustomProviders will test the method GetConversion()
func TestClient_GetConversionCustomProviders(t *testing.T) {

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
