package controller

import (
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
	books, err := bookService.FindAll()
	if err != nil {
		log.Printf("Failed to get books, Error:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting books"))
		return
	}
	var data = make([]entity.JsonBook, len(books))
	for i, book := range books {
		data[i] = mapBookToJSON(book)
	}

	sendResponse(w, r, data, http.StatusOK)
}

func (*controller) AddBook(w http.ResponseWriter, r *http.Request) {
	// Get JSON from request body
	reqBook := entity.BookRequest{}
	err := parseJson(w, r, &reqBook)
	if err != nil {
		log.Printf("Cannot parse post body. err=%v\n", err)
		sendResponse(w, r, nil, http.StatusBadRequest)
		return
	}

	// Create entity
	book := &entity.Book{
		ID:		0,
		Title:	reqBook.Title,
		Author:	reqBook.Author,
		Year:	reqBook.Year,
	}

	// Validate entity
	err = bookService.Validate(book)
	if err != nil {
		log.Printf("Failed to validate book, Error:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error validating book"))
		return
	}

	//Save entity in database
	bookService.Create(book)	
	data := mapBookToJSON(book)

	sendResponse(w, r, data, http.StatusOK)
}


