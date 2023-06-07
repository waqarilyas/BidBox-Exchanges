package api

import (
	"errors"
	"net/http"

	"github.com/kryptomind/bidboxapi/KeyService/helpers"
	"github.com/kryptomind/bidboxapi/KeyService/internal/exchange/bitget"
	"github.com/kryptomind/bidboxapi/KeyService/internal/models"
	"github.com/kryptomind/bidboxapi/KeyService/internal/shared"

	"github.com/kryptomind/bidboxapi/KeyService/response"
)

func (server *Server) GetBitgetAccountDetailsData(w http.ResponseWriter, r *http.Request) {

	userEmail := r.URL.Query().Get("email")
	if userEmail == "" {
		response.ERROR(w, http.StatusBadRequest, errors.New("email is required"))
		return
	}

	Key := models.Key{}

	userKeys, err := Key.FindKeyByEmailAndService(server.DB, "bitget", userEmail)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("cannot fetch user keys"))
		return
	}

	api_key, err := helpers.DecryptStrings(userKeys.ApiKey)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("error while decrypting user keys"))
		return
	}

	api_secret, err := helpers.DecryptStrings(userKeys.SecretKey)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("error while decrypting user api secret"))
		return
	}

	passphrase, err := helpers.DecryptStrings(userKeys.Passphrase)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("error while decrypting user api passphrase"))
		return
	}

	accountResponse, err := bitget.GetBitgetAccountData(api_key, api_secret, passphrase)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("unable to get user accoutn details at the moment. Please try again later"))
		return
	}

	transformedResponse := bitget.TransformAccountsResponse(*accountResponse)

	response.JSON(w, http.StatusOK, transformedResponse)

}

func (server *Server) GetBitgetOpenPositions(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("email")
	if userEmail == "" {
		response.ERROR(w, http.StatusBadRequest, errors.New("email is required"))
		return
	}

	Key := models.Key{}

	userKeys, err := Key.FindKeyByEmailAndService(server.DB, "bitget", userEmail)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("unable to fetch user keys. Please check your email"))
		return
	}

	trandformedKeys, err := shared.DecryptUserKeys(userKeys)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, errors.New("failed to decrypt keys"))
	}

	positionsresponse, err := bitget.PerformBitgetPositionQuery(trandformedKeys.ApiKey, trandformedKeys.Secret, trandformedKeys.Passphrase)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("unable to get user positions at the moment"))
	}

	formattedPositions := bitget.TransformPositionsResponse(*positionsresponse)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("unable to get format user positions at the moment"))
	}

	response.JSON(w, http.StatusOK, formattedPositions)

}
