package bitget

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kryptomind/bidboxapi/KeyService/helpers"
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
	uri := fmt.Sprintf("/api/mix/v1/order/historyProductType?productType=sumcbl&startTime=1609441200000&endTime=%s&pageSize=100", currentTimeStr)
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

func PerformBitgetHistoricalPositionQuery(apiKey, apiSecret, passphrase string) (*PositionHistory, error) {

	expires := helpers.GetBitgetServerTimeStamp()

	currentTime := time.Now()
	threeMonthsAgo := currentTime.AddDate(0, -3, 0)

	// Check if the difference between currentTime and threeMonthsAgo is more than 90 days
	if currentTime.Sub(threeMonthsAgo) > (90 * 24 * time.Hour) {
		// If the difference is more than 90 days, set threeMonthsAgo to be exactly 90 days before the current date
		threeMonthsAgo = currentTime.Add(-90 * 24 * time.Hour)
	}

	currentTimeMillis := currentTime.UnixNano() / int64(time.Millisecond)
	threeMonthsMillis := threeMonthsAgo.UnixNano() / int64(time.Millisecond)

	currentTimeStr := fmt.Sprintf("%d", currentTimeMillis)
	threeMonthsStr := fmt.Sprintf("%d", threeMonthsMillis)
	uri := fmt.Sprintf("/api/mix/v1/position/history-position?productType=sumcbl&startTime=%s&endTime=%s&pageSize=99", threeMonthsStr, currentTimeStr)

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

	var accountData PositionHistory

	err = json.Unmarshal(body, &accountData)
	if err != nil {
		return nil, err
	}

	return &accountData, nil
}
