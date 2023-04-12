package binance

import "errors"

var (
	ErrorBadAPIKey        = errors.New("bad api key")
	ErrorBadAPISecret     = errors.New("bad api secret")
	ErrorInvalidTimestamp = errors.New("invalid timestamp")
	ErrorInvalidSignature = errors.New("invalid signature")
)
