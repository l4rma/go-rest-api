package main

import (
	"fmt"
	"net/http"
	"rest-api/server"
)

func main() {
	fmt.Println("Hello world")
	server := server.New()
	
	http.HandleFunc("/", server.Router.ServeHTTP)

	port := ":8080"
	fmt.Println("Server running at: http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
