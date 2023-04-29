package binance

import "github.com/pkg/errors"

type binanceError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

var (
	ErrorBadAPIKey    = errors.New("bad api key")
	ErrorBadAPISecret = errors.New("bad api secret")
	//ErrorInvalidTimestamp = errors.New("invalid timestamp")
	//ErrorInvalidSignature = errors.New("invalid signature")
)
