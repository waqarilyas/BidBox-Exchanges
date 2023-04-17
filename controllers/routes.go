package controllers

import "github.com/kryptomind/bidboxapi/KeyService/middleware"

func (s *Server) initializeRoutes() {

	//Keys routes
	s.Router.HandleFunc("/keys", middleware.MiddlewareJSON(s.CreateKey)).Methods("POST")
	s.Router.HandleFunc("/keys", middleware.MiddlewareJSON(s.GetKeys)).Methods("GET")
	s.Router.HandleFunc("/keys/{id}", middleware.MiddlewareJSON(s.GetKey)).Methods("GET")
	s.Router.HandleFunc("/supported-exchanges", middleware.MiddlewareJSON(s.GetExchanges)).Methods("GET")
	s.Router.HandleFunc("/exchanges", middleware.MiddlewareJSON(s.CreateExchanges)).Methods("POST")
}
