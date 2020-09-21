/*
Package bsvrates brings multiple providers into one place to obtain the current BSV exchange rate
*/
package bsvrates

import (
	"fmt"
	"strconv"
)

// GetRate will get a BSV->Currency rate from the list of providers.
// The first provider that succeeds is the rate that is returned
func (c *Client) GetRate(currency Currency) (float64, Providers, error) {

	// Check if currency is accepted across all providers
	if !currency.IsAccepted() {
		return 0, 0, fmt.Errorf("currency [%s] is not accepted by all providers at this time", currency.Name())
	}

	// Provider: CoinPaprika
	if c.Providers&ProviderCoinPaprika != 0 {
		response, err := c.CoinPaprika.GetMarketPrice(CoinPaprikaQuoteID)
		if response != nil && err == nil {
			rate := response.Quotes.USD.Price
			return rate, ProviderCoinPaprika, nil
		}
		// todo: log the error for sanity in case the user want's to see the failure?
	}

	// Provider: WhatsOnChain
	if c.Providers&ProviderWhatsOnChain != 0 {
		response, err := c.WhatsOnChain.GetExchangeRate()
		if response != nil && err == nil {
			var rate float64
			if rate, err = strconv.ParseFloat(response.Rate, 8); err == nil {
				return rate, ProviderWhatsOnChain, err
			}
		}
		// todo: log the error for sanity in case the user want's to see the failure?
	}

	// Provider: Preev
	if c.Providers&ProviderPreev != 0 {
		response, err := c.Preev.GetTicker(PreevTickerID)
		if response != nil && err == nil {
			rate := response.Prices.Ppi.LastPrice
			return rate, ProviderPreev, nil
		}
		// todo: log the error for sanity in case the user want's to see the failure?
	}

	// Return an error if all providers failed
	return 0, 0, fmt.Errorf("unable to get rate from providers: %s", c.Providers.Names())
}

// todo: create a new method to get all three and then average the results
