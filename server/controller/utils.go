package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/l4rma/go-rest-api/server/db/entity"
)

// Parse json from request body
func parseJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

// Send response with "Content-Type" "application/json" with data and status code as arguments
func sendResponse(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	// Set header
	w.Header().Add("Content-Type", "application/json")

	// Set status code
	w.WriteHeader(status)

	if data == nil {
		return
	}

	// Set data as body
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}

// Takes (*entity.Book) as and argument, returning (entity.JsonBook)
func mapBookToJSON(b *entity.Book) entity.JsonBook {
	return entity.JsonBook{
		ID:     b.ID,
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
	}
}
