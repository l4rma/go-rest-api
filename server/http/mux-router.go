package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Router struct
type muxRouter struct{}

// Mux router
var (
	muxDispatcher = mux.NewRouter()
)

// Initializa new router
func NewMuxRouter() Router {
	return &muxRouter{}
}

// Routing GET requests
func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

// Routing POST requests
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

// Routing DELETE requests
func (*muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods(http.MethodDelete)
}

// Start server, listening on port port
func (*muxRouter) SERVE(port string) {
	log.Printf("Mux HTTP server running on: http://localhost%v", port)
	http.ListenAndServe(port, muxDispatcher)
}
