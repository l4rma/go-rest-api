package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/l4rma/go-rest-api/server/db/entity"
	"github.com/l4rma/go-rest-api/server/db/repository"
	"github.com/l4rma/go-rest-api/server/service"
	"github.com/stretchr/testify/assert"
)

var (
	bookRepo       repository.BookRepository = repository.NewSQLiteBookRepo()
	bookSer        service.BookService       = service.NewBookService(bookRepo)
	bookController BookController            = NewBookController(bookSer)
)

const (
	TITLE  string = "Test Title"
	AUTHOR string = "Test Author"
	YEAR   int16  = 2000
)

func TestAddBook(t *testing.T) {
	bookJson := []byte(`{"title":"TestTittel","author":"TestAuthor","year":2022}`)
	req, _ := http.NewRequest("POST", "/api/books", bytes.NewBuffer(bookJson))
	handler := http.HandlerFunc(bookController.AddBook)

	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	status := res.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Got:%v, Expected:%v", status, http.StatusOK)
	}

	var book entity.Book
	json.NewDecoder(io.Reader(res.Body)).Decode(&book)
	assert.NotNil(t, book.ID)
	assert.Equal(t, "TestTittel", book.Title)
	assert.Equal(t, "TestAuthor", book.Author)
	assert.Equal(t, int16(2022), book.Year)
}

func TestGetBook(t *testing.T) {

}
