package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
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

	floatUpl, _ := strconv.ParseFloat(accountData.TotalCrossUnPnl, 64)

	transformedData := shared.AccountData{
		Available:     availableBalance,
		Equity:        walletBalance,
		MarginBalance: marginBalance,
		UnrealizedPL:  unrealizedPl,
		UsdtPnl:       floatUpl,
	}

	return transformedData

}

func TransformAccountPositionsData(data []Position) []shared.PositionsData {

	var formattedPositions []shared.PositionsData

	for _, position := range data {

		positionAmount, err := strconv.ParseFloat(position.PositionAmt, 64)
		if err != nil {
			fmt.Println("Error parsing position amount", err)
		}

		if positionAmount > 0 || positionAmount < 0 {

			leverage, err := strconv.ParseFloat(position.Leverage, 64)
			if err != nil {
				fmt.Println("Error parsing leverage", err)
			}

			entryPrice, err := strconv.ParseFloat(position.EntryPrice, 64)
			if err != nil {
				fmt.Println("Error parsing leverage", err)
			}

			marginAmount := (positionAmount * entryPrice) / leverage
			marginAmount = math.Abs(marginAmount)

			positionSide := "short"
			holdMode := "single_hold"

			if positionAmount > 0 {
				positionSide = "long"
			}

			if position.PositionSide == "BOTH" {
				holdMode = "double_hold"
			}

			updateTime := fmt.Sprintf("%d", position.UpdateTime)
			marginAmountStr := strconv.FormatFloat(marginAmount, 'f', -1, 64)

			pos := shared.PositionsData{
				MarginCoin:       position.Symbol,
				Symbol:           position.Symbol,
				HoldSide:         positionSide,
				Margin:           marginAmountStr,
				Available:        position.IsolatedWallet,
				Total:            position.PositionAmt,
				MarginMode:       position.MarginType,
				HoldMode:         holdMode,
				LiquidationPrice: position.LiquidationPrice,
				MarketPrice:      position.MarkPrice,
				EntryPrice:       position.EntryPrice,
				CreationTime:     updateTime,
				UnrealizedPL:     position.UnrealizedProfit,
				Leverage:         position.Leverage,
			}
			formattedPositions = append(formattedPositions, pos)
		}

	}

	return formattedPositions
}
