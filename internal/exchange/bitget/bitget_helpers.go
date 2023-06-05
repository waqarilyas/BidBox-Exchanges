package bitget

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/kryptomind/bidboxapi/KeyService/internal/shared"
)

func GenerateBitgetSignature(apiSecret string, method string, uri string, timestamp string) string {
	message := fmt.Sprintf("%s%s%s", timestamp, method, uri)
	hmac := hmac.New(sha256.New, []byte(apiSecret))
	hmac.Write([]byte(message))
	signature := base64.StdEncoding.EncodeToString(hmac.Sum(nil))
	return signature
}

func TransformAccountsResponse(accountsData AccountData) shared.AccountData {
	totalAvailable := 0.0
	totalUnrealizedPL := 0.0
	totalEquity := 0.0

	for _, marginData := range accountsData.Data {
		available, err := strconv.ParseFloat(marginData.Available, 64)
		if err != nil {
			fmt.Println("Error parsing available balance:", err)
		}
		totalAvailable += available
		equity, err := strconv.ParseFloat(marginData.USDTEquity, 64)
		if err != nil {
			fmt.Println("Error parsing equity:", err)
		}
		totalEquity += equity

		unrealizedPL, err := strconv.ParseFloat(marginData.UnrealizedPL, 64)
		if err != nil {
			fmt.Println("Error parsing unrealized PL:", err)
		}
		totalUnrealizedPL += unrealizedPL
	}

	return shared.AccountData{
		Available:     totalAvailable,
		MarginBalance: totalEquity,
		UnrealizedPL:  totalUnrealizedPL,
	}

}

func TransformPositionsResponse(positionsData MarginDataResponse) []shared.PositionsData {
	positions := positionsData.Data
	var formattedPositions []shared.PositionsData
	for _, position := range positions {
		if position.Available != "0" {
			pos := shared.PositionsData{
				MarginCoin:       position.MarginCoin,
				Symbol:           position.Symbol,
				HoldSide:         position.HoldSide,
				Margin:           position.Margin,
				Available:        position.Available,
				Total:            position.Total,
				MarginMode:       position.MarginMode,
				HoldMode:         position.HoldMode,
				LiquidationPrice: position.LiquidationPrice,
				MarketPrice:      position.MarketPrice,
				CreationTime:     position.CTime,
			}
			formattedPositions = append(formattedPositions, pos)
		}
	}
	return formattedPositions

}
