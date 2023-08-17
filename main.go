package main

import (
	"os"

	"github.com/sirupsen/logrus"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/joho/godotenv"
	"github.com/kryptomind/bidboxapi/KeyService/helpers"
	"github.com/kryptomind/bidboxapi/KeyService/internal/api"
	"github.com/kryptomind/bidboxapi/KeyService/internal/database"
)

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
	server.DB = databaseConnection.DB

	helpers.InitSharedData(databaseConnection.DB)

	// response, err := bitget.PerformBitgetApikeyInformation("bg_7c52d3c7de17a4c18d8f1eb835b71158", "f4a466791e9779b55c9f15250f93747290bab56d20ee57963fc15db249638323", "thisisapassphrase")
	// if err != nil {
	// 	fmt.Println("--error calling api---", err)
	// }

	// fmt.Println("---response----", response)
	databaseConnection.Run(":8080")

}

//	@title			BidBox Exchanges Service API
//	@version		1.0
//	@description	This is BidBox Exchanges Service.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		stg-api-bidbox.kryptomind.net
// @BasePath	/exchanges
func main() {
	Run()
}
