package service

import (
	"errors"

	"github.com/l4rma/go-rest-api/server/db/entity"
	"github.com/l4rma/go-rest-api/server/db/repository"
	repo "github.com/l4rma/go-rest-api/server/db/repository"
)

var (
	bookRepo repo.BookRepository
)

// BookService interface
type BookService interface {
	Validate(book *entity.Book) error
	Create(book *entity.Book) (int64, error)
	FindAll() ([]*entity.Book, error)
	FindbyId(id int64) (*entity.Book, error)
	Delete(id int64) error
}

// Service struct
type service struct{}

// Initializa new service
func NewBookService(repo repository.BookRepository) BookService {
	bookRepo = repo
	return &service{}
}

// Method to validate a book
func (*service) Validate(book *entity.Book) error {
	// Check if book is not nil
	if book == nil {
		err := errors.New("The book is nil")
		return err
	}
	// Check if book has a title
	if book.Title == "" {
		err := errors.New("The book has no title")
		return err
	}
	return nil
}

// Delete book from database
func (*service) Delete(id int64) error {
	return bookRepo.DeleteById(id)
}

// Add book to database
func (*service) Create(book *entity.Book) (int64, error) {
	return bookRepo.Save(book)
}

// Retrieve book with mathing ID
func (*service) FindbyId(id int64) (*entity.Book, error) {
	return bookRepo.FindById(id)
}

// Retrieve all books from database
func (*service) FindAll() ([]*entity.Book, error) {
	return bookRepo.FindAll()
}
