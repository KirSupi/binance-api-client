package binance

import "strconv"

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
	err = c.base.Get("/fapi/v2/balance", true, p, &result, 5)
	return
}

func (c *FuturesClient) SymbolPriceTicker(ticker string) (result FuturesSymbolPriceTicker, err error) {
	p := Params{
		"recvWindow": "5000",
		"symbol":     ticker,
	}
	err = c.base.Get("/fapi/v1/ticker/price", false, p, &result, 1)
	return
}

func (c *FuturesClient) SymbolPriceAllTickers() (result []FuturesSymbolPriceTicker, err error) {
	p := Params{}
	err = c.base.Get("/fapi/v1/ticker/price", false, p, &result, 2)
	return
}

func (c *FuturesClient) MarkPriceAll() (result []FuturesMarkPrice, err error) {
	p := Params{}
	err = c.base.Get("/fapi/v1/premiumIndex", false, p, &result, 1)
	return
}

func (c *FuturesClient) AccountInformationV2() (result FuturesAccountInformationV2, err error) {
	p := Params{}
	err = c.base.Get("/fapi/v2/account", true, p, &result, 5)
	return
}

func (c *FuturesClient) GetIncomeHistory(params FuturesGetIncomeHistoryParams) (result []FuturesGetIncomeHistoryItem, err error) {
	p := Params{}
	if params.Symbol != nil {
		p["symbol"] = *params.Symbol
	}
	if params.IncomeType != nil {
		p["incomeType"] = string(*params.IncomeType)
	}
	if params.StartTime != nil {
		p["startTime"] = strconv.FormatInt(*params.StartTime, 10)
	}
	if params.EndTime != nil {
		p["endTime"] = strconv.FormatInt(*params.EndTime, 10)
	}
	if params.Limit != nil {
		if *params.Limit < 1 || *params.Limit > 1000 {
			*params.Limit = 1000
		}
		p["limit"] = strconv.Itoa(*params.Limit)
	}
	err = c.base.Get("/fapi/v1/income", true, p, &result, 30)
	return
}

func (c *FuturesClient) AccountTradeList(params FuturesAccountTradeListParams) (result []FuturesAccountTradeListItem, err error) {
	p := Params{
		"symbol": params.Symbol,
	}
	if params.OrderId != nil {
		p["orderId"] = strconv.FormatInt(*params.OrderId, 10)
	}
	if params.StartTime != nil {
		p["startTime"] = strconv.FormatInt(*params.StartTime, 10)
	}
	if params.EndTime != nil {
		p["endTime"] = strconv.FormatInt(*params.EndTime, 10)
	}
	if params.OrderId != nil {
		p["fromId"] = strconv.FormatInt(*params.FromId, 10)
	}
	if params.Limit != nil {
		if *params.Limit < 1 || *params.Limit > 1000 {
			*params.Limit = 1000
		}
		p["limit"] = strconv.Itoa(*params.Limit)
	}
	if params.OrderId != nil {
		p["recvWindow"] = strconv.FormatInt(*params.RecvWindow, 10)
	}
	err = c.base.Get("/fapi/v1/userTrades", true, p, &result, 30)
	return
}
