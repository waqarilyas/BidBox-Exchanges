package api

import (
	"errors"
	"net/http"

	"github.com/kryptomind/bidboxapi/KeyService/internal/exchange/bybit"
	"github.com/kryptomind/bidboxapi/KeyService/internal/models"
	"github.com/kryptomind/bidboxapi/KeyService/internal/shared"

	"github.com/kryptomind/bidboxapi/KeyService/response"
)

// Get Bybit Account Details godoc
// @Summary      Get Bybit Account Details
// @Description  Get Bybit Account Details
// @Tags         bybit
// @Accept       json
// @Produce      json
// @Param        email query  string true  "email"
// @Success      200  {object} shared.AccountData
// @Failure      400  {string}  bad request
// @Failure      500  {string}  bad request
// @Router       /bybit/account [get]
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

// Get Bybit Positions godoc
// @Summary      Get Bybit Positions
// @Description  Get Bybit Positions
// @Tags         bybit
// @Accept       json
// @Produce      json
// @Param        email query  string true  "email"
// @Success      200  {object} bybit.PositionsResponse
// @Failure      400  {string}  bad request
// @Failure      500  {string}  bad request
// @Router       /bybit/positions [get]
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

// Get Bybit Closed PnL godoc
// @Summary      Get Bybit Closed PnL
// @Description  Get Bybit Closed PnL
// @Tags         bybit
// @Accept       json
// @Produce      json
// @Param        email query  string true  "email"
// @Success      200  {object} []shared.ClosedPnlData
// @Failure      400  {string}  bad request
// @Failure      500  {string}  bad request
// @Router       /bybit/closedpnl [get]
func (server *Server) GetBybitCloseProfit_Loss(w http.ResponseWriter, r *http.Request) {
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

	accountResponse, err := bybit.GetBybitCloseProfit_Loss(trandformedKeys.ApiKey, trandformedKeys.Secret)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("unable to get user positions at the moment"))
	}

	trandformedData := bybit.TransformAccountClosedResponse(*accountResponse)

	response.JSON(w, http.StatusOK, trandformedData)

}

func (server *Server) GetBybitOrderHistory(w http.ResponseWriter, r *http.Request) {
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

	accountResponse, err := bybit.GetOrderHistory(trandformedKeys.ApiKey, trandformedKeys.Secret)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("unable to get user positions at the moment"))
	}

	trandformedData := bybit.TransformOrderHistoryResponse(*accountResponse)

	response.JSON(w, http.StatusOK, trandformedData)

}
