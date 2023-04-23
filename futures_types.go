package binance

type FuturesAccountBalance struct {
	AccountAlias       string `json:"accountAlias"`
	Asset              string `json:"asset"`
	Balance            string `json:"balance"`
	CrossWalletBalance string `json:"crossWalletBalance"`
	CrossUnPnl         string `json:"crossUnPnl"`
	AvailableBalance   string `json:"availableBalance"`
	MaxWithdrawAmount  string `json:"maxWithdrawAmount"`
	MarginAvailable    bool   `json:"marginAvailable"`
	UpdateTime         int64  `json:"updateTime"`
}

type FuturesSymbolPriceTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Time   int64  `json:"time"`
}

type FuturesMarkPrice struct {
	Symbol               string `json:"symbol"`
	MarkPrice            string `json:"markPrice"`
	IndexPrice           string `json:"indexPrice"`
	EstimatedSettlePrice string `json:"estimatedSettlePrice"`
	LastFundingRate      string `json:"lastFundingRate"`
	NextFundingTime      int64  `json:"nextFundingTime"`
	InterestRate         string `json:"interestRate"`
	Time                 int64  `json:"time"`
}

type FuturesAccountInformationV2 struct {
	FeeTier                     int                                   `json:"feeTier"`
	CanTrade                    bool                                  `json:"canTrade"`
	CanDeposit                  bool                                  `json:"canDeposit"`
	CanWithdraw                 bool                                  `json:"canWithdraw"`
	UpdateTime                  int                                   `json:"updateTime"`
	MultiAssetsMargin           bool                                  `json:"multiAssetsMargin"`
	TotalInitialMargin          string                                `json:"totalInitialMargin"`
	TotalMaintMargin            string                                `json:"totalMaintMargin"`
	TotalWalletBalance          string                                `json:"totalWalletBalance"`
	TotalUnrealizedProfit       string                                `json:"totalUnrealizedProfit"`
	TotalMarginBalance          string                                `json:"totalMarginBalance"`
	TotalPositionInitialMargin  string                                `json:"totalPositionInitialMargin"`
	TotalOpenOrderInitialMargin string                                `json:"totalOpenOrderInitialMargin"`
	TotalCrossWalletBalance     string                                `json:"totalCrossWalletBalance"`
	TotalCrossUnPnl             string                                `json:"totalCrossUnPnl"`
	AvailableBalance            string                                `json:"availableBalance"`
	MaxWithdrawAmount           string                                `json:"maxWithdrawAmount"`
	Assets                      []FuturesAccountInformationV2Asset    `json:"assets"`
	Positions                   []FuturesAccountInformationV2Position `json:"positions"`
}
type FuturesAccountInformationV2Asset struct {
	Asset                  string `json:"asset"`
	WalletBalance          string `json:"walletBalance"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	MarginBalance          string `json:"marginBalance"`
	MaintMargin            string `json:"maintMargin"`
	InitialMargin          string `json:"initialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	AvailableBalance       string `json:"availableBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	MarginAvailable        bool   `json:"marginAvailable"`
	UpdateTime             int64  `json:"updateTime"`
}

type FuturesAccountInformationV2Position struct {
	Symbol                 string `json:"symbol"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	Leverage               string `json:"leverage"`
	Isolated               bool   `json:"isolated"`
	EntryPrice             string `json:"entryPrice"`
	MaxNotional            string `json:"maxNotional"`
	BidNotional            string `json:"bidNotional"`
	AskNotional            string `json:"askNotional"`
	PositionSide           string `json:"positionSide"`
	PositionAmt            string `json:"positionAmt"`
	UpdateTime             int    `json:"updateTime"`
}
