package bsvrates

import (
	"context"
	"testing"

	"github.com/mrz1836/go-whatsonchain"
	"github.com/stretchr/testify/assert"
)

// newMockClient returns a client for mocking
func newMockClient(wocClient whatsonchain.ChainService, paprikaClient CoinPaprikaInterface,
	providers ...Provider) ClientInterface {
	client := NewClient(nil, nil, providers...)
	client.SetWhatsOnChain(wocClient)
	client.SetCoinPaprika(paprikaClient)
	return client
}

// TestClient_GetRate will test the method GetRate()
func TestClient_GetRate(t *testing.T) {
	// t.Parallel()

	t.Run("valid get rate - default", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{})
		assert.NotNil(t, client)

		rate, provider, err := client.GetRate(context.Background(), CurrencyDollars)
		assert.NoError(t, err)
		assert.Equal(t, 158.49415248, rate)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "CoinPaprika", provider.Name())
	})

	t.Run("valid get rate - whats on chain", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, ProviderWhatsOnChain)
		assert.NotNil(t, client)

		rate, provider, err := client.GetRate(context.Background(), CurrencyDollars)
		assert.NoError(t, err)
		assert.Equal(t, 159.01, rate)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "WhatsOnChain", provider.Name())
	})

	t.Run("valid get rate - custom providers", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, ProviderCoinPaprika, ProviderWhatsOnChain)
		assert.NotNil(t, client)

		rate, provider, err := client.GetRate(context.Background(), CurrencyDollars)
		assert.NoError(t, err)
		assert.Equal(t, 158.49415248, rate)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "CoinPaprika", provider.Name())
	})

	t.Run("non accepted currency", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaFailed{})
		assert.NotNil(t, client)

		_, _, rateErr := client.GetRate(context.Background(), 123)
		assert.Error(t, rateErr)
	})

	t.Run("failed rate - default", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaFailed{})
		assert.NotNil(t, client)

		rate, provider, err := client.GetRate(context.Background(), CurrencyDollars)
		assert.NoError(t, err)
		assert.Equal(t, 159.01, rate)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "WhatsOnChain", provider.Name())
	})

	t.Run("failed rate - whats on chain", func(t *testing.T) {
		client := newMockClient(&mockWOCFailed{}, &mockPaprikaValid{})
		assert.NotNil(t, client)

		rate, provider, err := client.GetRate(context.Background(), CurrencyDollars)
		assert.NoError(t, err)
		assert.Equal(t, 158.49415248, rate)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "CoinPaprika", provider.Name())
	})

	t.Run("failed rate - all providers", func(t *testing.T) {
		client := newMockClient(&mockWOCFailed{}, &mockPaprikaFailed{})
		assert.NotNil(t, client)

		rate, _, err := client.GetRate(context.Background(), CurrencyDollars)
		assert.Error(t, err)
		assert.Equal(t, float64(0), rate)
	})
}
