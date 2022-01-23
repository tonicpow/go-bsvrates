package bsvrates

import (
	"context"
	"testing"

	"github.com/mrz1836/go-preev"
	"github.com/mrz1836/go-whatsonchain"
	"github.com/stretchr/testify/assert"
)

// newMockClient returns a client for mocking
func newMockClient(wocClient whatsonchain.ChainService, paprikaClient CoinPaprikaInterface,
	preevClient preev.ClientInterface, providers ...Provider) ClientInterface {
	client := NewClient(nil, nil, providers...)
	client.SetWhatsOnChain(wocClient)
	client.SetCoinPaprika(paprikaClient)
	client.SetPreev(preevClient)
	return client
}

// TestClient_GetRate will test the method GetRate()
func TestClient_GetRate(t *testing.T) {
	// t.Parallel()

	t.Run("valid get rate - default", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevValid{})
		assert.NotNil(t, client)

		rate, provider, err := client.GetRate(context.Background(), CurrencyDollars)
		assert.NoError(t, err)
		assert.Equal(t, 158.49415248, rate)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "CoinPaprika", provider.Name())
	})

	t.Run("valid get rate - preev", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevValid{}, ProviderPreev)
		assert.NotNil(t, client)

		rate, provider, err := client.GetRate(context.Background(), CurrencyDollars)
		assert.NoError(t, err)
		assert.Equal(t, 159.17, rate)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "Preev", provider.Name())
	})

	t.Run("valid get rate - whats on chain", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevValid{}, ProviderWhatsOnChain)
		assert.NotNil(t, client)

		rate, provider, err := client.GetRate(context.Background(), CurrencyDollars)
		assert.NoError(t, err)
		assert.Equal(t, 159.01, rate)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "WhatsOnChain", provider.Name())
	})

	t.Run("valid get rate - custom providers", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevValid{}, ProviderPreev, ProviderWhatsOnChain)
		assert.NotNil(t, client)

		rate, provider, err := client.GetRate(context.Background(), CurrencyDollars)
		assert.NoError(t, err)
		assert.Equal(t, 159.17, rate)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "Preev", provider.Name())
	})

	t.Run("non accepted currency", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaFailed{}, &mockPreevValid{})
		assert.NotNil(t, client)

		_, _, rateErr := client.GetRate(context.Background(), 123)
		assert.Error(t, rateErr)
	})

	t.Run("failed rate - default", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaFailed{}, &mockPreevValid{})
		assert.NotNil(t, client)

		rate, provider, err := client.GetRate(context.Background(), CurrencyDollars)
		assert.NoError(t, err)
		assert.Equal(t, 159.01, rate)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "WhatsOnChain", provider.Name())
	})

	t.Run("failed rate - preev", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, &mockPreevFailed{}, ProviderPreev&ProviderCoinPaprika)
		assert.NotNil(t, client)

		rate, provider, err := client.GetRate(context.Background(), CurrencyDollars)
		assert.NoError(t, err)
		assert.Equal(t, 158.49415248, rate)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "CoinPaprika", provider.Name())
	})

	t.Run("failed rate - whats on chain", func(t *testing.T) {
		client := newMockClient(&mockWOCFailed{}, &mockPaprikaValid{}, &mockPreevFailed{})
		assert.NotNil(t, client)

		rate, provider, err := client.GetRate(context.Background(), CurrencyDollars)
		assert.NoError(t, err)
		assert.Equal(t, 158.49415248, rate)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "CoinPaprika", provider.Name())
	})

	t.Run("failed rate - all providers", func(t *testing.T) {
		client := newMockClient(&mockWOCFailed{}, &mockPaprikaFailed{}, &mockPreevFailed{})
		assert.NotNil(t, client)

		rate, _, err := client.GetRate(context.Background(), CurrencyDollars)
		assert.Error(t, err)
		assert.Equal(t, float64(0), rate)
	})
}
