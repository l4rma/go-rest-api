package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/l4rma/go-rest-api/server/db/entity"
)

func parseJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func sendResponse(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}

func mapBookToJSON(b *entity.Book) entity.JsonBook {
	return entity.JsonBook{
		ID:     b.ID,
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
	}
}
