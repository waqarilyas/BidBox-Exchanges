package binance

type AccountsResponse struct {
	FeeTier                     int        `json:"feeTier"`
	CanTrade                    bool       `json:"canTrade"`
	CanDeposit                  bool       `json:"canDeposit"`
	CanWithdraw                 bool       `json:"canWithdraw"`
	UpdateTime                  int        `json:"updateTime"`
	MultiAssetsMargin           bool       `json:"multiAssetsMargin"`
	TotalInitialMargin          string     `json:"totalInitialMargin"`
	TotalMaintMargin            string     `json:"totalMaintMargin"`
	TotalWalletBalance          string     `json:"totalWalletBalance"`
	TotalUnrealizedProfit       string     `json:"totalUnrealizedProfit"`
	TotalMarginBalance          string     `json:"totalMarginBalance"`
	TotalPositionInitialMargin  string     `json:"totalPositionInitialMargin"`
	TotalOpenOrderInitialMargin string     `json:"totalOpenOrderInitialMargin"`
	TotalCrossWalletBalance     string     `json:"totalCrossWalletBalance"`
	TotalCrossUnPnl             string     `json:"totalCrossUnPnl"`
	AvailableBalance            string     `json:"availableBalance"`
	MaxWithdrawAmount           string     `json:"maxWithdrawAmount"`
	Assets                      []Asset    `json:"assets"`
	Positions                   []Position `json:"positions"`
}

type Asset struct {
	Asset                  string `json:"asset"`
	WalletBalance          string `json:"walletBalance"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	MarginBalance          string `json:"marginBalance"`
	MaintMargin            string `json:"maintMargin"`
	InitialMargin          string `json:"initialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	AvailableBalance       string `json:"availableBalance"`
	MarginAvailable        bool   `json:"marginAvailable"`
	UpdateTime             int    `json:"updateTime"`
}

type Position struct {
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
	PositionSide           string `json:"positionSide"`
	PositionAmt            string `json:"positionAmt"`
	Notional               string `json:"notional"`
	IsolatedWallet         string `json:"isolatedWallet"`
	UpdateTime             int    `json:"updateTime"`
	BidNotional            string `json:"bidNotional"`
	AskNotional            string `json:"askNotional"`
}

type Order struct {
	OrderID       int64  `json:"orderId"`
	Symbol        string `json:"symbol"`
	Status        string `json:"status"`
	ClientOrderID string `json:"clientOrderId"`
	Price         string `json:"price"`
	AvgPrice      string `json:"avgPrice"`
	OrigQty       string `json:"origQty"`
	ExecutedQty   string `json:"executedQty"`
	CumQuote      string `json:"cumQuote"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	ReduceOnly    bool   `json:"reduceOnly"`
	ClosePosition bool   `json:"closePosition"`
	Side          string `json:"side"`
	PositionSide  string `json:"positionSide"`
	StopPrice     string `json:"stopPrice"`
	WorkingType   string `json:"workingType"`
	PriceProtect  bool   `json:"priceProtect"`
	OrigType      string `json:"origType"`
	Time          int64  `json:"time"`
	UpdateTime    int64  `json:"updateTime"`
}

type BinanceErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
