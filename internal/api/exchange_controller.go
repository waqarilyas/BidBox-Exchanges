package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kryptomind/bidboxapi/KeyService/models"
	"github.com/kryptomind/bidboxapi/KeyService/response"
)

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

func (server *Server) GetExchanges(w http.ResponseWriter, r *http.Request) {

	Exchange := models.Exchanges{}

	Exchanges, err := Exchange.FindAllExchanges(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, Exchanges)
}
