package binance

import (
	"binance-api-client/base"
)

var (
	baseFuturesUrl = "https://fapi.binance.com"
)

type FuturesClient struct {
	base *base.Client
}

func NewFuturesClient(apiKey, apiSecret string) *FuturesClient {
	return &FuturesClient{
		base: base.NewClient(baseFuturesUrl, apiKey, apiSecret),
	}
}

func (c *FuturesClient) AccountBalance() (result []FuturesAccountBalance, err error) {
	p := base.Params{
		"recvWindow": "5000",
	}
	err = c.base.Get("/fapi/v2/balance", true, p, &result)
	return
}
func (c *FuturesClient) SymbolPriceTicker(ticker *string) (result []FuturesSymbolPriceTicker, err error) {
	p := base.Params{
		"recvWindow": "5000",
	}
	if ticker != nil {
		p["symbol"] = *ticker
	}
	err = c.base.Get("/fapi/v1/ticker/price", true, p, &result)
	return
}
