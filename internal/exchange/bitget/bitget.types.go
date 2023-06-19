package bitget

type AccountData struct {
	Code        string          `json:"code"`
	Msg         string          `json:"msg"`
	RequestTime int64           `json:"requestTime"`
	Data        []AccountDetail `json:"data"`
}

type AccountDetail struct {
	MarginCoin        string `json:"marginCoin"`
	Locked            string `json:"locked"`
	Available         string `json:"available"`
	CrossMaxAvailable string `json:"crossMaxAvailable"`
	FixedMaxAvailable string `json:"fixedMaxAvailable"`
	MaxTransferOut    string `json:"maxTransferOut"`
	Equity            string `json:"equity"`
	USDTEquity        string `json:"usdtEquity"`
	BTCEquity         string `json:"btcEquity"`
	CrossRiskRate     string `json:"crossRiskRate"`
	UnrealizedPL      string `json:"unrealizedPL"`
	Bonus             string `json:"bonus"`
}

type MarginData struct {
	MarginCoin        string `json:"marginCoin"`
	Symbol            string `json:"symbol"`
	HoldSide          string `json:"holdSide"`
	OpenDelegateCount string `json:"openDelegateCount"`
	Margin            string `json:"margin"`
	Available         string `json:"available"`
	Locked            string `json:"locked"`
	Total             string `json:"total"`
	Leverage          int    `json:"leverage"`
	AchievedProfits   string `json:"achievedProfits"`
	AverageOpenPrice  string `json:"averageOpenPrice"`
	MarginMode        string `json:"marginMode"`
	HoldMode          string `json:"holdMode"`
	UnrealizedPL      string `json:"unrealizedPL"`
	LiquidationPrice  string `json:"liquidationPrice"`
	KeepMarginRate    string `json:"keepMarginRate"`
	MarketPrice       string `json:"marketPrice"`
	CTime             string `json:"cTime"`
}

type MarginDataResponse struct {
	Code        string       `json:"code"`
	Msg         string       `json:"msg"`
	RequestTime int64        `json:"requestTime"`
	Data        []MarginData `json:"data"`
}

type ApiKeyResponse struct {
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
	Data        User   `json:"data"`
}

type User struct {
	UserId           string   `json:"user_id"`
	InviterId        string   `json:"inviter_id"`
	AgentInviterCode string   `json:"agent_inviter_code"`
	Channel          string   `json:"channel"`
	Ips              string   `json:"ips"`
	Authorities      []string `json:"authorities"`
	ParentId         int64    `json:"parentId"`
	Trader           bool     `json:"trader"`
	IsSpotTrader     bool     `json:"isSpotTrader"`
}

// type OrderHistory struct {
// 	Code        string `json:"code"`
// 	Msg         string `json:"msg"`
// 	RequestTime int64  `json:"requestTime"`
// 	Data        Data   `json:"data"`
// }

// type Data struct {
// 	NextFlag bool       `json:"nextFlag"`
// 	EndID    string     `json:"endId"`
// 	List []OrderItem `json:"orderList"`
// }

// type OrderItem struct {
// 	Symbol         string  `json:"symbol"`
// 	Size           int     `json:"size"`
// 	OrderID        string  `json:"orderId"`
// 	ClientOID      string  `json:"clientOid"`
// 	FilledQty      int     `json:"filledQty"`
// 	Fee            float64 `json:"fee"`
// 	Price          float64 `json:"price"`
// 	PriceAvg       float64 `json:"priceAvg"`
// 	State          string  `json:"state"`
// 	Side           string  `json:"side"`
// 	TimeInForce    string  `json:"timeInForce"`
// 	TotalProfits   float64 `json:"totalProfits"`
// 	PosSide        string  `json:"posSide"`
// 	MarginCoin     string  `json:"marginCoin"`
// 	FilledAmount   float64 `json:"filledAmount"`
// 	OrderType      string  `json:"orderType"`
// 	Leverage       string  `json:"leverage"`
// 	MarginMode     string  `json:"marginMode"`
// 	ReduceOnly     bool    `json:"reduceOnly"`
// 	EnterPointSource string `json:"enterPointSource"`
// 	TradeSide      string  `json:"tradeSide"`
// 	HoldMode       string  `json:"holdMode"`
// 	OrderSource    string  `json:"orderSource"`
// 	CTime          string  `json:"cTime"`
// 	UTime          string  `json:"uTime"`
// }

type OrderHistory struct {
	Code string `json:"code"`
	Data struct {
		NextFlag  bool         `json:"nextFlag"`
		EndID     string       `json:"endId"`
		List []OrderItem `json:"orderList"`
	} `json:"data"`
	Msg string `json:"msg"`
}

type OrderItem struct {
	Symbol            string  `json:"symbol"`
	Size              float64     `json:"size"`
	OrderID           string  `json:"orderId"`
	ClientOID         string  `json:"clientOid"`
	FilledQty         float64     `json:"filledQty"`
	Fee               float64 `json:"fee"`
	Price             float64 `json:"price"`
	State             string  `json:"state"`
	Side              string  `json:"side"`
	TimeInForce       string  `json:"timeInForce"`
	TotalProfits      float64 `json:"totalProfits"`
	PosSide           string  `json:"posSide"`
	MarginCoin        string  `json:"marginCoin"`
	Leverage          string  `json:"leverage"`
	MarginMode        string  `json:"marginMode"`
	OrderType         string  `json:"orderType"`
	ReduceOnly        bool    `json:"reduceOnly"`
	EnterPointSource  string  `json:"enterPointSource"`
	TradeSide         string  `json:"tradeSide"`
	HoldMode          string  `json:"holdMode"`
	OrderSource       string  `json:"orderSource"`
	CTime             string   `json:"ctime"`
	UTime             string   `json:"utime"`
}
