package database

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/kryptomind/bidboxapi/KeyService/internal/api"
	"github.com/kryptomind/bidboxapi/KeyService/models"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

var routeModule = api.Server{
	Router: mux.NewRouter(),
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	var err error
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		log.Fatalf("Failed to connect to the %s database: %v", Dbdriver, err)
	}
	log.Infof("Connected to the database successfully")

	if err := server.DB.Debug().AutoMigrate(&models.Key{}).Error; err != nil {
		log.Fatalf("Failed to perform database migration: %v", err)
	}

	server.Router = routeModule.Router
	routeModule.DB = server.DB
	routeModule.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	log.Infof("Listening on port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
