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

	// todo: serial for now, later can become a go routine group with a race across all providers
	if c.Providers&ProviderCoinPaprika != 0 {
		response, err := c.CoinPaprika.GetMarketPrice(CoinPaprikaQuoteID)
		if response != nil && err == nil {
			rate := response.Quotes.USD.Price
			return rate, ProviderCoinPaprika, nil
		}
		// todo: log the error for sanity in case the user want's to see the failure?
	}

	if c.Providers&ProviderWhatsOnChain != 0 {
		var err error

		response, err := c.WhatsOnChain.GetExchangeRate()
		if response != nil && err == nil {
			rate, err := strconv.ParseFloat(response.Rate, 8)
			if err == nil {
				return rate, ProviderWhatsOnChain, err
			}
		}
		// todo: log the error for sanity in case the user want's to see the failure?
	}

	if c.Providers&ProviderPreev != 0 {
		response, err := c.Preev.GetTicker(PreevTickerID)
		if response != nil && err == nil {
			rate := response.Prices.Ppi.LastPrice
			return rate, ProviderPreev, nil
		}
		// todo: log the error for sanity in case the user want's to see the failure?
	}

	// TODO: average all received rates?
	return 0, 0, fmt.Errorf("unable to get rate from providers: %s", c.Providers.Names())

}
