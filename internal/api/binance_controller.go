package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/adshao/go-binance/v2/futures"
	"github.com/kryptomind/bidboxapi/KeyService/internal/exchange/binance"
	"github.com/kryptomind/bidboxapi/KeyService/internal/models"
	"github.com/kryptomind/bidboxapi/KeyService/internal/shared"

	"github.com/kryptomind/bidboxapi/KeyService/response"
)

func (server *Server) GetBinanceAccountDetailsData(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("email")
	if userEmail == "" {
		response.ERROR(w, http.StatusBadRequest, errors.New("email is required"))
		return
	}

	Key := models.Key{}

	userKeys, err := Key.FindKeyByEmailAndService(server.DB, "binance", userEmail)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("exchange not connected"))
		return
	}

	trandformedKeys, err := shared.DecryptUserKeys(userKeys)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, errors.New("failed to decrypt keys"))
	}

	accountResponse, err := binance.GetBinanceAccountDetails(trandformedKeys.ApiKey, trandformedKeys.Secret)
	if err != nil {

		fmt.Println(err)
		response.ERROR(w, http.StatusBadRequest, errors.New("unable to get user account details at the moment. Please try again later"))
		return
	}

	transformedResponse := binance.TransformAccountResponse(*accountResponse)

	response.JSON(w, http.StatusOK, transformedResponse)

}

func (server *Server) GetBinanceAccountPositionsData(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("email")
	if userEmail == "" {
		response.ERROR(w, http.StatusBadRequest, errors.New("email is required"))
		return
	}

	Key := models.Key{}

	userKeys, err := Key.FindKeyByEmailAndService(server.DB, "binance", userEmail)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("exchange not connected"))
		return
	}

	trandformedKeys, err := shared.DecryptUserKeys(userKeys)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, errors.New("failed to decrypt keys"))
	}

	positionsResponse, err := binance.GetBinanceAccountOpenPositions(trandformedKeys.ApiKey, trandformedKeys.Secret)
	if err != nil {
		fmt.Println(err)
		response.ERROR(w, http.StatusBadRequest, errors.New("unable to get user account details at the moment. Please try again later"))
		return
	}

	transformedResponse := binance.TransformAccountPositionsData(*positionsResponse)

	response.JSON(w, http.StatusOK, transformedResponse)

}

func (server *Server) GetBinanceOrderHistory(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("email")
	if userEmail == "" {
		response.ERROR(w, http.StatusBadRequest, errors.New("email is required"))
		return
	}

	Key := models.Key{}

	userKeys, err := Key.FindKeyByEmailAndService(server.DB, "binance", userEmail)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("exchange not connected"))
		return
	}

	trandformedKeys, err := shared.DecryptUserKeys(userKeys)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, errors.New("failed to decrypt keys"))
	}

	futures.UseTestnet = true
	BinanceClient := futures.NewClient(trandformedKeys.ApiKey, trandformedKeys.Secret)

	b, err := BinanceClient.NewListOrdersService().Do(context.Background())
	sh := make([]shared.OrderHistoryResponse, 0)
	for _, v := range b {
		var side string
		if v.Side == futures.SideTypeBuy {
			side = "long"
		} else {
			side = "short"
		}
		state := "Unknown start"
		if v.Status == "NEW"{
			state = "New"
		}else if v.Status == "PARTIALLY_FILLED"{
			state = "PartiallyFilled"
		}else if v.Status == "FILLED"{
			state = "Filled"
		}else if v.Status == "CANCELED"{
			state = "Canceled"
		}else {
			state = "Unknown"
		}
		ns := shared.OrderHistoryResponse{
			Symbol:        v.Symbol,
			OrderID:       fmt.Sprintf("%d", v.OrderID),
			Price:         v.Price,
			Qty:           v.CumQuote,
			Side:          side,
			OrderStatus:   state,
			CreatedTime:   fmt.Sprintf("%d", v.Time),
			UpdatedTime:   fmt.Sprintf("%d", v.UpdateTime),
			StopOrderType: "UNKNOWN",
			OrderType:     string(v.Type),
			CumExecQty:    "0",
			CumExecFee:    "0",
			CumExecValue:  "0",
			AvgPrice:      v.AvgPrice,
		}
		sh = append(sh, ns)

	}
	response.JSON(w, http.StatusOK, sh)

}
