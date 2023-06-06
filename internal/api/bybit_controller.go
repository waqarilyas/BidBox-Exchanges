package api

import (
	"errors"
	"net/http"

	"github.com/kryptomind/bidboxapi/KeyService/internal/exchange/bybit"
	"github.com/kryptomind/bidboxapi/KeyService/internal/models"
	"github.com/kryptomind/bidboxapi/KeyService/internal/shared"

	"github.com/kryptomind/bidboxapi/KeyService/response"
)

func (server *Server) GetBybitAccountDetails(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("email")
	if userEmail == "" {
		response.ERROR(w, http.StatusBadRequest, errors.New("email is required"))
		return
	}

	Key := models.Key{}

	userKeys, err := Key.FindKeyByEmailAndService(server.DB, "bybit", userEmail)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("exchange not connected"))
		return
	}

	trandformedKeys, err := shared.DecryptUserKeys(userKeys)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, errors.New("failed to decrypt keys"))
	}

	accountResponse, err := bybit.GetBybitAccountBalance(trandformedKeys.ApiKey, trandformedKeys.Secret)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("unable to get user positions at the moment"))
	}

	response.JSON(w, http.StatusOK, accountResponse)

}

func (server *Server) GetBybitAccountPositions(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("email")
	if userEmail == "" {
		response.ERROR(w, http.StatusBadRequest, errors.New("email is required"))
		return
	}

	Key := models.Key{}

	userKeys, err := Key.FindKeyByEmailAndService(server.DB, "bybit", userEmail)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("exchange not connected"))
		return
	}

	trandformedKeys, err := shared.DecryptUserKeys(userKeys)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, errors.New("failed to decrypt keys"))
	}

	accountResponse, err := bybit.GetBybitAccountPositions(trandformedKeys.ApiKey, trandformedKeys.Secret)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("unable to get user positions at the moment"))
	}

	trandformedData := bybit.TransformAccountPositionsResponse(*accountResponse)

	response.JSON(w, http.StatusOK, trandformedData)

}
