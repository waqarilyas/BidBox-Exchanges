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
