package api

import (
	"errors"
	"fmt"
	"net/http"

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
