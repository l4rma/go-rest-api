package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

}

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
	s.Router.HandleFunc("/", func(w http.ResponseWriter,r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	}).Methods("GET")
}
