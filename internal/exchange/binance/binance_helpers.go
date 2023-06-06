package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/kryptomind/bidboxapi/KeyService/internal/shared"
)

func GenerateBinanceSignature(params map[string]string, secretKey string) string {

	var queryString string
	for key, value := range params {
		queryString += key + "=" + url.QueryEscape(value) + "&"
	}

	queryString = strings.TrimSuffix(queryString, "&")

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(queryString))
	signature := hex.EncodeToString(mac.Sum(nil))

	return signature
}

func TransformAccountResponse(accountData AccountsResponse) shared.AccountData {
	availableBalance, err := strconv.ParseFloat(accountData.AvailableBalance, 64)
	if err != nil {
		fmt.Println("Error parsing available balance:", err)
	}

	walletBalance, err := strconv.ParseFloat(accountData.TotalWalletBalance, 64)
	if err != nil {
		fmt.Println("Error parsing total wallet balance:", err)
	}

	marginBalance, err := strconv.ParseFloat(accountData.TotalMarginBalance, 64)
	if err != nil {
		fmt.Println("Error parsing margin balance:", err)
	}

	unrealizedPl, err := strconv.ParseFloat(accountData.TotalCrossUnPnl, 64)
	if err != nil {
		fmt.Println("Error parsing unrealizedPl:", err)
	}

	transformedData := shared.AccountData{
		Available:     availableBalance,
		Equity:        walletBalance,
		MarginBalance: marginBalance,
		UnrealizedPL:  unrealizedPl,
	}

	return transformedData

}
