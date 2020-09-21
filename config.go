package bsvrates

import "strings"

const (

	// version is the current package version
	version = "v0.0.1"

	// defaultUserAgent is the default user agent for all requests
	defaultUserAgent string = "go-bsvrates: " + version

	// CoinPaprikaQuoteID is the id for CoinPaprika (BSV)
	CoinPaprikaQuoteID = "bsv-bitcoin-sv"

	// PreevTickerID is the id for Preev (BSV)
	PreevTickerID = "12eLTxv1vyUeJtp5zqWbqpdWvfLdZ7dGf8"
)

var (
	// DefaultProviders uses all three (Coinpaprika, WhatsOnChain, and Preev) by default.
	DefaultProviders = Providers(ProviderCoinPaprika | ProviderWhatsOnChain | ProviderPreev)
)

// Providers is a provider for rates or prices
type Providers uint8

// Providers bitstring for the different available rate providers.
const (
	ProviderWhatsOnChain Providers = 1 << iota // 1 << 0 (which is 00000001 = 1)
	ProviderCoinPaprika                        // 1 << 1 (which is 00000010 = 2)
	ProviderPreev                              // 1 << 2 (which is 00000100 = 4)

	providerMask = (1 << iota) - 1 // will mask all used bits 00000111 (= 7)
)

// IsValid tests if the provider is valid or not
func (p Providers) IsValid() bool {
	return p&providerMask != 0
}

// Names will return the display name for the given provider
func (p Providers) Names() []string {
	names := []string{}

	if p&ProviderWhatsOnChain != 0 {
		names = append(names, "WhatsOnChain")
	}
	if p&ProviderCoinPaprika != 0 {
		names = append(names, "CoinPaprika")
	}
	if p&ProviderPreev != 0 {
		names = append(names, "Preev")
	}

	return names
}

// ProviderToNames helper function to convert the provider value to it's associated name
func ProviderToNames(provider Providers) []string {
	return provider.Names()
}

// Currency is a valid currency for rates or prices
type Currency uint8

// Currency constants for the different available currencies.
// Leave the start and last constants in place
const (
	_               Currency = iota
	CurrencyDollars          = 1
	CurrencyBitcoin          = 2

	currencyLast = iota
)

// IsValid tests if the provider is valid or not
func (c Currency) IsValid() bool {
	return c >= CurrencyDollars && c < currencyLast
}

// IsAccepted tests if the currency is accepted by all providers
func (c Currency) IsAccepted() bool {
	if c == CurrencyDollars {
		return true
	}
	return false
}

// Name will return the display name for the given currency
func (c Currency) Name() string {
	switch c {
	case CurrencyDollars:
		return "usd"
	case CurrencyBitcoin:
		return "bsv"
	default:
		return ""
	}
}

// CurrencyToName helper function to convert the currency value to it's associated name
func CurrencyToName(currency Currency) string {
	return currency.Name()
}

// CurrencyFromName helper function to convert the name into it's Currency type
func CurrencyFromName(name string) Currency {
	switch strings.ToLower(name) {
	case "usd":
		return CurrencyDollars
	case "bsv":
		return CurrencyBitcoin
	default:
		return CurrencyDollars
	}
}
