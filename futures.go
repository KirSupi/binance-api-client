package binance

var (
	baseFuturesUrl = "https://fapi.binance.com"
)

type FuturesClient struct {
	base *baseClient
}

func NewFuturesClient(apiKey, apiSecret string) *FuturesClient {
	return &FuturesClient{
		base: newBaseClient(baseFuturesUrl, apiKey, apiSecret),
	}
}

func (c *FuturesClient) AccountBalance() (result []FuturesAccountBalance, err error) {
	p := Params{
		"recvWindow": "5000",
	}
	err = c.base.Get("/fapi/v2/balance", true, p, &result)
	return
}
func (c *FuturesClient) SymbolPriceTicker(ticker string) (result FuturesSymbolPriceTicker, err error) {
	p := Params{
		"recvWindow": "5000",
		"symbol":     ticker,
	}
	err = c.base.Get("/fapi/v1/ticker/price", false, p, &result)
	return
}

func (c *FuturesClient) SymbolPriceAllTickers() (result []FuturesSymbolPriceTicker, err error) {
	p := Params{
		"recvWindow": "5000",
	}
	err = c.base.Get("/fapi/v1/ticker/price", false, p, &result)
	return
}
