package helpers

type bitgetServerTimeStampResponse struct {
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	RequestTime int    `json:"requestTime"`
	Data        string `json:"data"`
}
