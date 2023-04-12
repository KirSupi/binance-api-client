package binance

type SpotTransaction struct {
	ID        int64  `json:"tranId"`
	Timestamp int64  `json:"timestamp"`
	Asset     string `json:"asset"`
	Amount    string `json:"amount"`
	Type      int    `json:"type"`
	Status    string `json:"status"`
}

type SpotTransactionHistory struct {
	Total int               `json:"total"`
	Rows  []SpotTransaction `json:"rows"`
}
