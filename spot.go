package binance

import (
	"strconv"
	"time"
)

var (
	baseSpotUrl = "https://api.binance.com"
)

type SpotClient struct {
	base *baseClient
}

func NewSpotClient(apiKey, apiSecret string) *SpotClient {
	return &SpotClient{
		base: newBaseClient(baseSpotUrl, apiKey, apiSecret),
	}
}

func (c *SpotClient) Time() (int64, error) {
	var res struct {
		Timestamp int64 `json:"serverTime"`
	}
	err := c.base.Get("/api/v3/time", false, nil, &res)
	if err != nil {
		return 0, err
	}
	return res.Timestamp, nil
}

func (c *SpotClient) TransactionHistory() (result SpotTransactionHistory, err error) {
	p := Params{
		"recvWindow": "5000",
		"asset":      "USDT",
		"startTime":  strconv.FormatInt(time.Now().Add(-time.Hour*24*30*6).UnixMilli(), 10),
	}
	err = c.base.Get("/sapi/v1/futures/transfer", true, p, &result)
	return
}
