package main

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/joho/godotenv"
	"github.com/kryptomind/bidboxapi/KeyService/internal/api"
	"github.com/kryptomind/bidboxapi/KeyService/internal/database"
	"github.com/kryptomind/bidboxapi/KeyService/internal/exchange/binance"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

var server = api.Server{}
var databaseConnection = database.Server{}

func Run() {
	err := godotenv.Load()
	log := logrus.New()
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"file", "function"},
	})
	if err != nil {
		log.WithFields(logrus.Fields{
			"file":     "main.go",
			"function": "Run",
		}).Fatal("Error getting env")
	} else {
		log.WithFields(logrus.Fields{
			"file":     "main.go",
			"function": "Run",
		}).Info("Getting Values")
	}

	databaseConnection.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	server.Router = databaseConnection.Router // Assign the same router instance

	response, err := binance.GetBinanceAccountDetails("8d011d5eac34cd4bf92310bd82fb05c1926959a0a9e1e8dcfd41bbd2407bf51b", "abd53029baadf319290a07bc771505480e2801931aedc2a3a18f6afba9571314")
	if err != nil {
		fmt.Println("---- error in calling api ------", err)
	}

	fmt.Println("----- api response ------", response)

	databaseConnection.Run(":8080")

}

func main() {
	Run()

}
