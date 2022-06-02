package bsvrates

import (
	"context"

	"github.com/mrz1836/go-whatsonchain"
)

// RateService is the rate methods
type RateService interface {
	GetConversion(ctx context.Context, currency Currency, amount float64) (satoshis int64, providerUsed Provider, err error)
	GetRate(ctx context.Context, currency Currency) (rate float64, providerUsed Provider, err error)
}

// ClientInterface is the BSVRate client interface
type ClientInterface interface {
	RateService
	CoinPaprika() CoinPaprikaInterface
	Providers() []Provider
	SetCoinPaprika(client CoinPaprikaInterface)
	SetWhatsOnChain(client whatsonchain.ChainService)
	WhatsOnChain() whatsonchain.ChainService
}
