package bsvrates

import (
	"fmt"
	"strconv"

	"github.com/mrz1836/go-preev"
	"github.com/mrz1836/go-whatsonchain"
)

// GetConversion will get the satoshi amount for the given currency + amount provided.
// The first provider that succeeds is the conversion that is returned
func (c *Client) GetConversion(currency Currency, amount float64) (satoshis int64, providerUsed Provider, err error) {

	// Check if currency is accepted across all providers
	if !currency.IsAccepted() {
		err = fmt.Errorf("currency [%s] is not accepted by all providers at this time", currency.Name())
		return
	}

	// Loop providers and get a conversion value
	// todo: serial for now, later can become a go routine group with a race across all providers
	for _, provider := range c.Providers {
		providerUsed = provider
		switch provider {
		case ProviderCoinPaprika:
			var response *PriceConversionResponse
			if response, err = c.CoinPaprika.GetPriceConversion(USDCurrencyID, CoinPaprikaQuoteID, fmt.Sprintf("%.2f", amount)); err == nil && response != nil {
				var sats string
				if sats, err = response.GetSatoshi(); err == nil {
					satoshis, err = strconv.ParseInt(sats, 10, 64)
				}
			}
		case ProviderWhatsOnChain:
			var response *whatsonchain.ExchangeRate
			if response, err = c.WhatsOnChain.GetExchangeRate(); err == nil && response != nil {
				var rate float64
				if rate, err = strconv.ParseFloat(response.Rate, 8); err == nil {
					satoshis, err = ConvertPriceToSatoshis(rate, amount)
				}
			}
		case ProviderPreev:
			var response *preev.Ticker
			if response, err = c.Preev.GetTicker(PreevTickerID); err == nil && response != nil {
				satoshis, err = ConvertPriceToSatoshis(response.Prices.Ppi.LastPrice, amount)
			}
		}

		// todo: log the error for sanity in case the user want's to see the failure?

		// Did we get a satoshi value? Otherwise keep looping
		if satoshis > 0 {
			return
		}
	}

	return
}
