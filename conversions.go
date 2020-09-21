package bsvrates

import (
	"fmt"
	"strconv"
)

// GetConversion will get the satoshi amount for the given currency + amount provided.
// The first provider that succeeds is the conversion that is returned
func (c *Client) GetConversion(currency Currency, amount float64) (int64, Providers, error) {

	// Check if currency is accepted across all providers
	if !currency.IsAccepted() {
		return 0, 0, fmt.Errorf("currency [%s] is not accepted by all providers at this time", currency.Name())
	}

	// todo: serial for now, later can become a go routine group with a race across all providers
	if c.Providers&ProviderCoinPaprika != 0 {
		response, err := c.CoinPaprika.GetPriceConversion(USDCurrencyID, CoinPaprikaQuoteID, amount)
		if response != nil && err == nil {
			var satoshis int64
			if satoshis, err = response.GetSatoshi(); err == nil {
				return satoshis, ProviderCoinPaprika, nil
			}
		}
		// todo: log the error for sanity in case the user want's to see the failure?
	}

	if c.Providers&ProviderWhatsOnChain != 0 {
		response, err := c.WhatsOnChain.GetExchangeRate()
		if response != nil && err == nil {
			var rate float64
			if rate, err = strconv.ParseFloat(response.Rate, 8); err == nil {
				var satoshis int64
				if satoshis, err = ConvertPriceToSatoshis(rate, amount); err == nil {
					return satoshis, ProviderWhatsOnChain, nil
				}
			}
		}
		// todo: log the error for sanity in case the user want's to see the failure?

	}

	if c.Providers&ProviderPreev != 0 {
		response, err := c.Preev.GetTicker(PreevTickerID)
		if response != nil && err == nil {
			var satoshis int64
			if satoshis, err = ConvertPriceToSatoshis(response.Prices.Ppi.LastPrice, amount); err == nil {
				return satoshis, ProviderPreev, nil
			}
		}
		// todo: log the error for sanity in case the user want's to see the failure?
	}

	// TODO: average all received conversions?
	return 0, 0, fmt.Errorf("unable to get conversion from providers: %s", c.Providers.Names())
}
