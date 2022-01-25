package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-api/server/db/entity"
	"rest-api/server/service"
)

type controller struct{}

var (
	bookService service.BookService
)

type BookController interface {
	GetBooks(w http.ResponseWriter, r *http.Request)
	AddBook(w http.ResponseWriter, r *http.Request)
}

func NewBookController(service service.BookService) BookController {
	bookService = service
	return &controller{}
}

func (*controller) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books, err := bookService.FindAll()
	if err != nil {
		log.Printf("Failed to get books, Error:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting books"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func (*controller) AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book entity.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Printf("Failed to add book, Error:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error adding book"))
		return
	}
	err = bookService.Validate(&book)
	if err != nil {
		log.Printf("Failed to validate book, Error:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error validating book"))
		return
	}

	bookService.Create(&book)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}


