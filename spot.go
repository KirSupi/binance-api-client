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
	err := c.base.Get("/api/v3/time", false, nil, &res, 1)
	if err != nil {
		return 0, err
	}
	return res.Timestamp, nil
}

func (c *SpotClient) TransactionHistory(asset string, startTime *int64) (result SpotTransactionHistory, err error) {
	if startTime == nil {
		startTime = new(int64)
		*startTime = time.Now().Add(-time.Hour * 24 * 30 * 6).UnixMilli()
	}
	p := Params{
		"recvWindow": "5000",
		"asset":      asset,
		"startTime":  strconv.FormatInt(*startTime, 10),
	}
	err = c.base.Get("/sapi/v1/futures/transfer", true, p, &result, 10)
	return
}
