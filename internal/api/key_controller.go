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

func (server *Server) CreateKey(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	Key := models.Key{}
	err = json.Unmarshal(body, &Key)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	Key.Prepare()

	valErr := Key.Validate()
	if valErr != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, valErr)
		return
	}

	if Key.Service == "bitget" {
		_, validationError := helpers.ValidateBitgetKeys(Key.SecretKey, Key.ApiKey, Key.Passphrase)
		if validationError != nil {
			response.ERROR(w, http.StatusBadRequest, errors.New("invalid api keys credentials"))
			return
		}
	} else if Key.Service == "binance" {
		_, validationError := binance.GetBinanceAccountDetails(Key.ApiKey, Key.SecretKey)
		if validationError != nil {
			response.ERROR(w, http.StatusBadRequest, errors.New("invalid api keys credentials"))
			return
		}
	} else if Key.Service == "bybit" {
		bybitKeyInfo, validationError := bybit.GetBybitApiKeyInfo(Key.ApiKey, Key.SecretKey)
		if validationError != nil {
			response.ERROR(w, http.StatusBadRequest, errors.New("invalid api keys credentials"))
			return
		}

		permissions := bybitKeyInfo.Result[0].Permissions
		hasPermission := bybit.HasRequiredPermissions(permissions)
		if !hasPermission {
			response.ERROR(w, http.StatusBadRequest, errors.New("insufficient key permissions. Please provide permissions for 'Order', 'Position', 'ExchangeHistory'"))
			return
		}

	}

	dbRes, _ := Key.FindKeyByEmailAndService(server.DB, Key.Service, Key.UserEmail)
	if dbRes != nil {
		response.ERROR(w, http.StatusBadRequest, errors.New("api key already exists"))
		return
	}

	KeyCreated, err := Key.SaveKey(server.DB)

	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, KeyCreated.Uid))
	response.JSON(w, http.StatusOK, "Api key validated and saved successfully")
}

func (server *Server) GetKeys(w http.ResponseWriter, r *http.Request) {

	Key := models.Key{}

	Keys, err := Key.FindAllKeys(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, Keys)
}

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

type updateLeverageRequest struct {
	Emails []string `json:"emails"`
	symbol string `json:"symbol"`
	buyLeverage int `json:"buyLeverage"`

}
// type updateLeverage struct {
// 	Category     string `json:"category"`
// 	Symbol       string `json:"symbol"`
// 	BuyLeverage  string `json:"buyLeverage"`
// 	SellLeverage string `json:"sellLeverage"`
// }

func (server *Server) updateLeverage(w http.ResponseWriter, r *http.Request) {
		var req updateLeverageRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
	
		// Loop through the emails and print them
		for _, email := range req.Emails {
			fmt.Println(email)
			key := models.Key{}
			keys, err := key.FindKeyByEmail(server.DB, email)
			if err != nil {
				response.ERROR(w, http.StatusBadRequest, err)
			}
			for _, currentkey := range keys {
				fmt.Println(currentkey)
				if currentkey.Service == "bybit" {
					leverage := models.UpdateLeverage{}
					leverage.Category = "linear"
					leverage.Symbol = req.symbol
					leverage.BuyLeverage = string(req.buyLeverage)
					leverage.SellLeverage = string(req.buyLeverage)
					// bybit.UpdatebybitLeverage(currentkey.ApiKey, currentkey.SecretKey, leverage)

			}
			
		}
		
		}
}