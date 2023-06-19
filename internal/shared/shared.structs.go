package shared

type AccountData struct {
	Available     float64 `json:"available"`
	Equity        float64 `json:"equity"`
	MarginBalance float64 `json:"marginBalance"`
	UnrealizedPL  float64 `json:"unrealizedPl"`
}

type DecryptedKeys struct {
	ApiKey     string `json:"api_key"`
	Secret     string `json:"secret"`
	Passphrase string `json:"passphrase"`
}

type PositionsData struct {
	Symbol           string `json:"symbol"`
	MarginCoin       string `json:"marginCoin"`
	HoldSide         string `json:"holdSide"`
	Margin           string `json:"margin"`
	Available        string `json:"available"`
	Total            string `json:"total"`
	MarginMode       string `json:"marginMode"`
	HoldMode         string `json:"holdMode"`
	LiquidationPrice string `json:"liquidationPrice"`
	MarketPrice      string `json:"marketPrice"`
	EntryPrice       string `json:"entryPrice"`
	CreationTime     string `json:"creationTime"`
	UnrealizedPL     string `json:"unrealizedPl"`
	Leverage         string `json:"leverage"`
}

type ClosedPnlData struct {
    Symbol         string `json:"symbol"`
    OrderType      string `json:"orderType"`
    Leverage       string `json:"leverage"`
    UpdatedTime    string `json:"updatedTime"`
    Side           string `json:"side"`
    OrderID        string `json:"orderId"`
    ClosedPnl      string `json:"closedPnl"`
    AvgEntryPrice  string `json:"avgEntryPrice"`
    Quantity       string `json:"qty"`
    CumEntryValue  string `json:"cumEntryValue"`
    CreatedTime    string `json:"createdTime"`
    OrderPrice     string `json:"orderPrice"`
    ClosedSize     string `json:"closedSize"`
    AvgExitPrice   string `json:"avgExitPrice"`
    ExecType       string `json:"execType"`
    FillCount      string `json:"fillCount"`
    CumExitValue   string `json:"cumExitValue"`
}

type OrderHistoryResponse struct {
	OrderID           string  `json:"orderId"`
	Symbol            string  `json:"symbol"`
	Price             string  `json:"price"`
	Qty               string  `json:"qty"`
	Side              string  `json:"side"`
	OrderStatus       string  `json:"orderStatus"`
	Profit 		  	  float64  `json:"profit"`
	// CancelType        string  `json:"cancelType"`
	AvgPrice          string  `json:"avgPrice"`
	CumExecQty        string  `json:"cumExecQty"`
	CumExecValue      string  `json:"cumExecValue"`
	CumExecFee        string  `json:"cumExecFee"`
	OrderType         string  `json:"orderType"`
	StopOrderType     string  `json:"stopOrderType"`
	// OrderIv           string  `json:"orderIv"`
	CreatedTime       string  `json:"createdTime"`
	UpdatedTime       string  `json:"updatedTime"`
}

