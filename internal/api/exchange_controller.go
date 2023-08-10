package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kryptomind/bidboxapi/KeyService/internal/models"
	"github.com/kryptomind/bidboxapi/KeyService/response"
)

// Add New Exchange godoc
// @Summary      Add New Exchange
// @Description   Add New Exchange
// @Tags         exchanges
// @Accept       json
// @Produce      json
// @Param        exchange body  models.Exchanges true  "add exchange"
// @Success      200  {object} models.Exchanges
// @Failure      422  {string}  KeyResp
// @Failure      500  {string}  KeyResp
// @Router       / [post]
func (server *Server) CreateExchanges(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	Exchange := models.Exchanges{}
	err = json.Unmarshal(body, &Exchange)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = Exchange.ValidateExchange()
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	KeyCreated, err := Exchange.SaveExchange(server.DB)

	if err != nil {

		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, KeyCreated.Id))
	response.JSON(w, http.StatusCreated, KeyCreated)
}

// Get Supported Exchanges godoc
// @Summary      Get Supported Exchanges
// @Description  Get Supported Exchanges
// @Tags         exchanges
// @Accept       json
// @Produce      json
// @Success      200  {object} []models.Exchanges
// @Failure      500  {string}  KeyResp
// @Router       /supported [get]
func (server *Server) GetExchanges(w http.ResponseWriter, r *http.Request) {

	Exchange := models.Exchanges{}

	Exchanges, err := Exchange.FindAllExchanges(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, Exchanges)
}
