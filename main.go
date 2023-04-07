package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/kryptomind/bidboxapi/KeyService/controllers"
)

var server = controllers.Server{}

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "main.go",
			"function": "Run",
		}).Fatal("Error getting env")
	} else {
		log.WithFields(log.Fields{
			"file":     "main.go",
			"function": "Run",
		}).Info("Getting Values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	server.Run(":8080")
}

func main() {
	Run()
}
