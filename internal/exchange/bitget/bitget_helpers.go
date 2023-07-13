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
	usdtPnl := 0.0

	// positionUsdUpl:=
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
		marginEquity, err := strconv.ParseFloat(marginData.Equity, 64)
		if err != nil {
			fmt.Println("Error parsing equity:", err)
		}

		totalEquity += equity

		marginRate := equity / marginEquity
		marginPnl := (marginEquity - available) * marginRate
		usdtPnl += marginPnl

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
		UsdtPnl:       usdtPnl,
	}

}

func TransformPositionsResponse(positionsData MarginDataResponse) []shared.PositionsData {
	positions := positionsData.Data
	var formattedPositions []shared.PositionsData
	for _, position := range positions {

		if position.Available != "0" {

			strLeverage := fmt.Sprintf("%v", position.Leverage)

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
				UnrealizedPL:     position.UnrealizedPL,
				Leverage:         strLeverage,
				EntryPrice:       position.AverageOpenPrice,
			}
			formattedPositions = append(formattedPositions, pos)
		}
	}
	return formattedPositions

}

func TransformOrderHistoryResponse(positionData OrderHistory) []shared.OrderHistoryResponse {
	positionsList := positionData.Data.List

	var formattedPositions []shared.OrderHistoryResponse

	for _, position := range positionsList {
		price := strconv.FormatFloat(position.Price, 'f', -1, 64)
		// priceAvg := strconv.FormatFloat(position.PriceAvg, 'f', -1, 64)
		qty := strconv.FormatFloat(position.FilledQty, 'f', -1, 64)
		// cTime := strconv.FormatInt(position.CTime, 10)
		fee := strconv.FormatFloat(position.Fee, 'f', -1, 64)
		// createdTime := strconv.FormatInt(position.CTime, 10)
		state := "position.State"
		if position.State == "init" {
			state = "Created"
		} else if position.State == "new" {
			state = "New"
		} else if position.State == "partially_filled" {
			state = "PartiallyFilled"
		} else if position.State == "filled" {
			state = "Filled"
		} else if position.State == "canceled" {
			state = "Cancelled"
		} else {
			state = "Unknown"
		}
		pos := shared.OrderHistoryResponse{
			Symbol:        position.Symbol,
			OrderID:       position.OrderID,
			Side:          position.Side,
			Price:         price,
			Profit:        position.TotalProfits,
			Qty:           qty,
			OrderStatus:   state,
			AvgPrice:      "0",
			CumExecQty:    "0",
			CumExecValue:  "0",
			CumExecFee:    fee,
			OrderType:     position.OrderType,
			StopOrderType: "UNKNOWN",
			// OrderIv: position.OrderIv,
			CreatedTime: position.CTime,
			UpdatedTime: position.UTime,
		}

		formattedPositions = append(formattedPositions, pos)
	}

	return formattedPositions
}
