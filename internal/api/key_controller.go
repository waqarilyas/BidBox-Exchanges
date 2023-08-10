package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/kryptomind/bidboxapi/KeyService/helpers"
	"github.com/kryptomind/bidboxapi/KeyService/internal/exchange/binance"
	"github.com/kryptomind/bidboxapi/KeyService/internal/exchange/bybit"
	"github.com/kryptomind/bidboxapi/KeyService/internal/models"

	"github.com/kryptomind/bidboxapi/KeyService/response"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, "Exchanges Service")
}

type KeyResp struct {
	status  int
	message string
}

// Create Key godoc
// @Summary      Create Key
// @Description  Create Key
// @Tags         keys
// @Accept       json
// @Produce      json
// @Param        key body  models.Key true  "create key"
// @Success      200  {object} KeyResp
// @Failure      400  {object}  KeyResp
// @Failure      422  {object}  KeyResp
// @Router       /keys [post]
func (server *Server) CreateKey(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	Key := models.Key{}
	keyResp := KeyResp{}
	err = json.Unmarshal(body, &Key)
	if err != nil {
		keyResp.message = err.Error()
		keyResp.status = http.StatusUnprocessableEntity
		response.JSON(w, http.StatusUnprocessableEntity, keyResp)
		return
	}

	Key.Prepare()

	valErr := Key.Validate()
	if valErr != nil {
		keyResp.message = err.Error()
		keyResp.status = http.StatusUnprocessableEntity
		response.JSON(w, http.StatusUnprocessableEntity, keyResp)
		return
	}

	if Key.Service == "bitget" {
		_, validationError := helpers.ValidateBitgetKeys(Key.SecretKey, Key.ApiKey, Key.Passphrase)
		if validationError != nil {
			keyResp.message = "invalid api keys credentials"
			keyResp.status = http.StatusBadRequest
			response.JSON(w, http.StatusBadRequest, keyResp)
			return
		}
	} else if Key.Service == "binance" {
		_, validationError := binance.GetBinanceAccountDetails(Key.ApiKey, Key.SecretKey)
		if validationError != nil {
			keyResp.message = "invalid api keys credentials"
			keyResp.status = http.StatusBadRequest
			response.JSON(w, http.StatusBadRequest, keyResp)
			return
		}
	} else if Key.Service == "bybit" {
		bybitKeyInfo, validationError := bybit.GetBybitApiKeyInfo(Key.ApiKey, Key.SecretKey)
		if validationError != nil {
			keyResp.message = "invalid api keys credentials"
			keyResp.status = http.StatusBadRequest
			response.JSON(w, http.StatusBadRequest, keyResp)
			return
		}

		permissions := bybitKeyInfo.Result[0].Permissions
		hasPermission := bybit.HasRequiredPermissions(permissions)
		if !hasPermission {
			keyResp.message = "insufficient key permissions. Please provide permissions for 'Order', 'Position', 'ExchangeHistory' and 'DerivativesTrade'"
			keyResp.status = http.StatusBadRequest
			response.JSON(w, http.StatusBadRequest, keyResp)
			return
		}

	}

	dbRes, _ := Key.FindKeyByEmailAndService(server.DB, Key.Service, Key.UserEmail)
	if dbRes != nil {
		keyResp.message = "api key already exists"
		keyResp.status = http.StatusBadRequest
		response.JSON(w, http.StatusBadRequest, keyResp)
		return
	}

	KeyCreated, err := Key.SaveKey(server.DB)

	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, KeyCreated.Uid))
	keyResp.message = "Api key validated and saved successfully"
	keyResp.status = http.StatusOK
	response.JSON(w, http.StatusOK, keyResp)
}

// Get Key godoc
// @Summary      Get Key
// @Description  Get Key
// @Tags         keys
// @Accept       json
// @Produce      json
// @Success      200  {object} models.Key
// @Failure      500  {string} res server error
// @Router       /keys [get]
func (server *Server) GetKeys(w http.ResponseWriter, r *http.Request) {

	Key := models.Key{}

	Keys, err := Key.FindAllKeys(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, Keys)
}

// Get Key By Id godoc
// @Summary      Get Key By Id
// @Description  Get Key By Id
// @Tags         keys
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Key ID"
// @Success      200  {object} models.Key
// @Failure      400  {string} resp bad request
// @Router       /keys/{id} [get]
func (server *Server) GetKey(w http.ResponseWriter, r *http.Request) {

	kid := mux.Vars(r)["id"]
	new_kid, err := uuid.Parse(kid)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("invalid key id"))
		return
	}
	Key := models.Key{}
	KeyGotten, err := Key.FindKeyById(server.DB, new_kid)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	response.JSON(w, http.StatusOK, KeyGotten)
}
