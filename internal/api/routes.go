package api

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/kryptomind/bidboxapi/KeyService/middleware"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (r *Server) InitializeRoutes() {
	s := r.Router.PathPrefix("/exchanges").Subrouter()
	s.HandleFunc("/", middleware.MiddlewareJSON(r.Home)).Methods("GET")

	// Keys routes
	s.HandleFunc("/keys", middleware.MiddlewareJSON(r.CreateKey)).Methods("POST")
	s.HandleFunc("/keys", middleware.MiddlewareJSON(r.GetKeys)).Methods("GET")
	s.HandleFunc("/keys/{id}", middleware.MiddlewareJSON(r.GetKey)).Methods("GET")
	s.HandleFunc("/supported", middleware.MiddlewareJSON(r.GetExchanges)).Methods("GET")
	s.HandleFunc("/", middleware.MiddlewareJSON(r.CreateExchanges)).Methods("POST")

	// Bitget Data Rputes
	s.HandleFunc("/bitget/account", middleware.MiddlewareJSON(r.GetBitgetAccountDetailsData)).Methods("GET")
	s.HandleFunc("/bitget/positions", middleware.MiddlewareJSON(r.GetBitgetOpenPositions)).Methods("GET")

}
