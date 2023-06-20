package bitget

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/kryptomind/bidboxapi/KeyService/helpers"
	"github.com/kryptomind/bidboxapi/KeyService/internal/models"
	// "github.com/kryptomind/bidboxapi/AccountsService/api/helpers"
)

func GetBitgetAccountData(apiKey string, apiSecret string, passphrase string) (*AccountData, error) {
	expires := helpers.GetBitgetServerTimeStamp()

	host := "https://api.bitget.com"
	path := "/api/mix/v1/account/accounts?productType=sumcbl"
	url := host + path

	signature := GenerateBitgetSignature(apiSecret, "GET", path, expires)

	fmt.Println(passphrase)
	fmt.Println(apiSecret)
	fmt.Println(apiKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("ACCESS-KEY", apiKey)
	req.Header.Add("ACCESS-PASSPHRASE", passphrase)
	req.Header.Add("ACCESS-TIMESTAMP", expires)
	req.Header.Add("ACCESS-SIGN", signature)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())

		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return nil, err
	}

	var accountData AccountData

	err = json.Unmarshal(body, &accountData)
	if err != nil {
		return nil, err
	}

	return &accountData, nil
}

func PerformBitgetApikeyInformation(apiKey, apiSecret, passphrase string) (*ApiKeyResponse, error) {
	expires := helpers.GetBitgetServerTimeStamp()
	uri := "/api/spot/v1/account/getInfo"
	signature := GenerateBitgetSignature(apiSecret, "GET", uri, expires)

	url := fmt.Sprintf("https://api.bitget.com%s", uri)
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("ACCESS-KEY", apiKey)
	req.Header.Add("ACCESS-PASSPHRASE", passphrase)
	req.Header.Add("ACCESS-TIMESTAMP", expires)
	req.Header.Add("ACCESS-SIGN", signature)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return nil, err
	}

	var accountData ApiKeyResponse

	err = json.Unmarshal(body, &accountData)
	if err != nil {
		return nil, err
	}

	return &accountData, nil
}
func PerformBitgetPositionQuery(apiKey, apiSecret, passphrase string) (*MarginDataResponse, error) {
	expires := helpers.GetBitgetServerTimeStamp()
	uri := "/api/mix/v1/position/allPosition?productType=sumcbl"
	signature := GenerateBitgetSignature(apiSecret, "GET", uri, expires)

	url := fmt.Sprintf("https://api.bitget.com%s", uri)
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("ACCESS-KEY", apiKey)
	req.Header.Add("ACCESS-PASSPHRASE", passphrase)
	req.Header.Add("ACCESS-TIMESTAMP", expires)
	req.Header.Add("ACCESS-SIGN", signature)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return nil, err
	}

	var accountData MarginDataResponse

	err = json.Unmarshal(body, &accountData)
	if err != nil {
		return nil, err
	}

	return &accountData, nil
}

func GetOrderHistory(apiKey, apiSecret, passphrase string) (*OrderHistory, error) {
	expires := helpers.GetBitgetServerTimeStamp()
	currentTimeMillis := time.Now().UnixNano() / int64(time.Millisecond)
	currentTimeStr := fmt.Sprintf("%d", currentTimeMillis)
	uri := fmt.Sprintf("/api/mix/v1/order/historyProductType?productType=sumcbl&startTime=1609441200000&endTime=%s&pageSize=100",currentTimeStr)
	signature := GenerateBitgetSignature(apiSecret, "GET", uri, expires)

	url := fmt.Sprintf("https://api.bitget.com%s", uri)
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("ACCESS-KEY", apiKey)
	req.Header.Add("ACCESS-PASSPHRASE", passphrase)
	req.Header.Add("ACCESS-TIMESTAMP", expires)
	req.Header.Add("ACCESS-SIGN", signature)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return nil, err
	}

	var accountData OrderHistory

	err = json.Unmarshal(body, &accountData)
	if err != nil {
		return nil, err
	}

	return &accountData, nil
}

type Response struct {
	Symbol string `json:"symbol"`
	MarginCoin string `json:"marginCoin"`
	LongLeverage string `json:"longLeverage"`
	ShortLeverage string `json:"shortLeverage"`
}

// func UpdateBitgetLeverage(apiKey, apiSecret, passphrase string, leverage *models.UpdateLeverageBitget ) (error) {
// 	expires := helpers.GetBitgetServerTimeStamp()
// 	uri := fmt.Sprintf("/api/mix/v1/account/setLeverage")
// 	signature := GenerateBitgetSignature(apiSecret, "POST", uri, expires)
// 	jsonVal, err := json.Marshal(leverage)
// 	url := fmt.Sprintf("https://api.bitget.com%s", uri)
// 	method := "POST"

// 	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonVal))
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Add("ACCESS-KEY", apiKey)
// 	req.Header.Add("ACCESS-PASSPHRASE", passphrase)
// 	req.Header.Add("ACCESS-TIMESTAMP", expires)
// 	req.Header.Add("ACCESS-SIGN", signature)

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {

// 		return err
// 	}

// 	var errorResponse Response
// 		err = json.Unmarshal(body, &errorResponse)
// 		if err != nil {
// 			return err
// 		}
		
// 	return nil
// }
type UpdateLeverageResponse struct {
	Code        string `json:"code"`
	Data        struct {
		OpenCount string `json:"openCount"`
	} `json:"data"`
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
}

func UpdateBitgetLeverage(apiKey, apiSecret, passphrase string, leverage *models.UpdateLeverageBitget) ( error) {
	expires := helpers.GetBitgetServerTimeStamp()
	uri := fmt.Sprintf("/api/mix/v1/account/setLeverage")
	signature := GenerateBitgetSignature(apiSecret, "POST", uri, expires)
	jsonVal, err := json.Marshal(leverage)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("https://api.bitget.com%s", uri)
	method := "POST"

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonVal))
	if err != nil {
		return err
	}

	req.Header.Add("ACCESS-KEY", apiKey)
	req.Header.Add("ACCESS-PASSPHRASE", passphrase)
	req.Header.Add("ACCESS-TIMESTAMP", expires)
	req.Header.Add("ACCESS-SIGN", signature)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var response UpdateLeverageResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	// Handle error response from the API
	if response.Code != "00000" {
		return fmt.Errorf("API error: %s", response.Msg)
	}

	return nil
}
