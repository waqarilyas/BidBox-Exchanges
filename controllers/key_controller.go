package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/kryptomind/bidboxapi/KeyService/models"
	"github.com/kryptomind/bidboxapi/KeyService/response"
)

func (server *Server) CreateKey(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "controllers/key_controller.go",
			"function": "CreateKey",
		}).Error("Error reading request body")
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	Key := models.Key{}
	err = json.Unmarshal(body, &Key)
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "controllers/key_controller.go",
			"function": "CreateKey",
		}).Error("Error parsing json payload")
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	Key.Prepare()
	err = Key.Validate()
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "controllers/key_controller.go",
			"function": "CreateKey",
		}).Error("Validation error - ", err)
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	KeyCreated, err := Key.SaveKey(server.DB)

	if err != nil {
		log.WithFields(log.Fields{
			"file":     "controllers/key_controller.go",
			"function": "CreateKey",
		}).Error("Error saving to database")
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, KeyCreated.Uid))
	log.WithFields(log.Fields{
		"file":     "controllers/key_controller.go",
		"function": "CreateKey",
	}).Info("Key Created with Id ", KeyCreated.Uid)
	response.JSON(w, http.StatusCreated, KeyCreated)
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

	kid := mux.Vars(r)["id"] //grab the id
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
