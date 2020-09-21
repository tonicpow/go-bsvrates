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
		var err error

		response, err := c.CoinPaprika.GetPriceConversion(USDCurrencyID, CoinPaprikaQuoteID, amount)
		if response != nil && err == nil {
			sats, err := response.GetSatoshi()
			if err == nil {
				satoshis, err := strconv.ParseInt(sats, 10, 64)
				if err == nil {
					return satoshis, ProviderCoinPaprika, nil
				}
			}
		}
		// todo: log the error for sanity in case the user want's to see the failure?
	}

	if c.Providers&ProviderWhatsOnChain != 0 {
		var err error

		response, err := c.WhatsOnChain.GetExchangeRate()
		if response != nil && err == nil {
			rate, err := strconv.ParseFloat(response.Rate, 8)
			if err == nil {
				satoshis, err := ConvertPriceToSatoshis(rate, amount)
				if err == nil {
					return satoshis, ProviderWhatsOnChain, nil
				}
			}
		}
		// todo: log the error for sanity in case the user want's to see the failure?

	}

	if c.Providers&ProviderPreev != 0 {
		var err error

		response, err := c.Preev.GetTicker(PreevTickerID)
		if response != nil && err == nil {
			satoshis, err := ConvertPriceToSatoshis(response.Prices.Ppi.LastPrice, amount)
			if err == nil {
				return satoshis, ProviderPreev, nil
			}
		}
		// todo: log the error for sanity in case the user want's to see the failure?
	}

	//TODO: average all received conversions?
	return 0, 0, fmt.Errorf("unable to get conversion from providers: %s", c.Providers.Names())
}
