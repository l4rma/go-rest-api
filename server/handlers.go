package server

import (
	"fmt"
	"net/http"
)

func (s *Server) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	}
}
