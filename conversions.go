package bsvrates

import (
	"context"
	"fmt"

	"github.com/mrz1836/go-whatsonchain"
)

// GetConversion will get the satoshi amount for the given currency + amount provided.
// The first provider that succeeds is the conversion that is returned
func (c *Client) GetConversion(ctx context.Context, currency Currency, amount float64) (satoshis int64, providerUsed Provider, err error) {

	// Check if currency is accepted across all providers
	if !currency.IsAccepted() {
		err = fmt.Errorf("currency [%s] is not accepted by all providers at this time", currency.Name())
		return
	}

	// Loop providers and get a conversion value
	for _, provider := range c.Providers() {
		providerUsed = provider
		switch provider {
		case ProviderCoinPaprika:
			var response *PriceConversionResponse
			if response, err = c.CoinPaprika().GetPriceConversion(
				ctx, USDCurrencyID, CoinPaprikaQuoteID, amount,
			); err == nil && response != nil {
				satoshis, err = response.GetSatoshi()
			}
		case ProviderWhatsOnChain:
			var response *whatsonchain.ExchangeRate
			if response, err = c.WhatsOnChain().GetExchangeRate(ctx); err == nil && response != nil {
				satoshis, err = ConvertPriceToSatoshis(response.Rate, amount)
			}
		case providerLast:
			err = fmt.Errorf("provider unknown")
			return
		}

		// todo: log the error for sanity in case the user want's to see the failure?

		// Did we get a satoshi value? Otherwise, keep looping
		if satoshis > 0 {
			return
		}
	}

	return
}

// todo: create a new method to get all three and then average the results
