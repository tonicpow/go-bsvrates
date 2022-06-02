package bsvrates

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestClient_GetConversion will test the method GetConversion()
func TestClient_GetConversion(t *testing.T) {
	t.Parallel()

	t.Run("valid get conversion - default", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{})
		assert.NotNil(t, client)

		satoshis, provider, err := client.GetConversion(context.Background(), CurrencyDollars, 1)
		assert.NoError(t, err)
		assert.Equal(t, int64(633157), satoshis)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "CoinPaprika", provider.Name())
	})

	t.Run("valid get conversion - whats on chain", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, ProviderWhatsOnChain)
		assert.NotNil(t, client)

		satoshis, provider, err := client.GetConversion(context.Background(), CurrencyDollars, 1)
		assert.NoError(t, err)
		assert.Equal(t, int64(628892), satoshis)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "WhatsOnChain", provider.Name())
	})

	t.Run("valid get conversion - coin paprika", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, ProviderCoinPaprika)
		assert.NotNil(t, client)

		satoshis, provider, err := client.GetConversion(context.Background(), CurrencyDollars, 1)
		assert.NoError(t, err)
		assert.Equal(t, int64(633157), satoshis)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "CoinPaprika", provider.Name())
	})

	t.Run("valid get conversion - custom provider list", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaValid{}, ProviderCoinPaprika, ProviderWhatsOnChain)
		assert.NotNil(t, client)

		satoshis, provider, err := client.GetConversion(context.Background(), CurrencyDollars, 1)
		assert.NoError(t, err)
		assert.NotEqual(t, 633157, satoshis)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "CoinPaprika", provider.Name())
	})

	t.Run("non accepted currency", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaFailed{})
		assert.NotNil(t, client)

		_, _, rateErr := client.GetConversion(context.Background(), 123, 1)
		assert.Error(t, rateErr)
	})

	t.Run("failed conversion - default", func(t *testing.T) {
		client := newMockClient(&mockWOCValid{}, &mockPaprikaFailed{})
		assert.NotNil(t, client)

		satoshis, provider, err := client.GetConversion(context.Background(), CurrencyDollars, 1)
		assert.NoError(t, err)
		assert.Equal(t, int64(628892), satoshis)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "WhatsOnChain", provider.Name())
	})

	t.Run("failed conversion - whats on chain", func(t *testing.T) {
		client := newMockClient(&mockWOCFailed{}, &mockPaprikaValid{})
		assert.NotNil(t, client)

		satoshis, provider, err := client.GetConversion(context.Background(), CurrencyDollars, 1)
		assert.NoError(t, err)
		assert.Equal(t, int64(633157), satoshis)
		assert.Equal(t, true, provider.IsValid())
		assert.Equal(t, "CoinPaprika", provider.Name())
	})

	t.Run("failed conversion - all providers", func(t *testing.T) {
		client := newMockClient(&mockWOCFailed{}, &mockPaprikaFailed{})
		assert.NotNil(t, client)

		satoshis, _, err := client.GetConversion(context.Background(), CurrencyDollars, 1)
		assert.Error(t, err)
		assert.Equal(t, int64(0), satoshis)
	})
}
