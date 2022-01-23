package bsvrates

import (
	"context"

	"github.com/mrz1836/go-preev"
	"github.com/mrz1836/go-whatsonchain"
)

// ClientInterface is the BSVRate client interface
type ClientInterface interface {
	CoinPaprika() CoinPaprikaInterface
	GetConversion(ctx context.Context, currency Currency, amount float64) (satoshis int64, providerUsed Provider, err error)
	GetRate(ctx context.Context, currency Currency) (rate float64, providerUsed Provider, err error)
	Preev() preev.ClientInterface
	Providers() []Provider
	SetCoinPaprika(client CoinPaprikaInterface)
	SetPreev(client preev.ClientInterface)
	SetWhatsOnChain(client whatsonchain.ChainService)
	WhatsOnChain() whatsonchain.ChainService
}
