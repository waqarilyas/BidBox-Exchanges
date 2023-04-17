package controllers

import "github.com/kryptomind/bidboxapi/KeyService/middleware"

func (r *Server) initializeRoutes() {

	s := r.Router.PathPrefix("/exchanges").Subrouter()

	s.HandleFunc("/", middleware.MiddlewareJSON(r.Home)).Methods("GET")

	//Keys routes
	s.HandleFunc("/keys", middleware.MiddlewareJSON(r.CreateKey)).Methods("POST")
	s.HandleFunc("/keys", middleware.MiddlewareJSON(r.GetKeys)).Methods("GET")
	s.HandleFunc("/keys/{id}", middleware.MiddlewareJSON(r.GetKey)).Methods("GET")
	s.HandleFunc("/supported", middleware.MiddlewareJSON(r.GetExchanges)).Methods("GET")
	s.HandleFunc("/", middleware.MiddlewareJSON(r.CreateExchanges)).Methods("POST")
}
