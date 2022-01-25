package server

import (
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	// Database
}

func New() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.initRoutes()
	return s
}

func (s *Server) initRoutes() {
	s.Router.HandleFunc("/", s.IndexHandler()).Methods("GET")
}
