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

type BookService interface {
	Validate(book *entity.Book) error
	Create(book *entity.Book) (int64, error)
	FindAll() ([]*entity.Book, error)
	Delete(id int64) error
}

type service struct{}

func NewBookService(repo repository.BookRepository) BookService {
	bookRepo = repo
	return &service{}
}

func (*service) Validate(book *entity.Book) error {
	if book == nil {
		err := errors.New("The book is nil")
		return err
	}
	if book.Title == "" {
		err := errors.New("The book has no title")
		return err
	}
	return nil
}
func (*service) Delete(id int64) error {
	return bookRepo.DeleteById(id)
}

func (*service) Create(book *entity.Book) (int64, error) {
	return bookRepo.Save(book)
}

func (*service) FindAll() ([]*entity.Book, error) {
	return bookRepo.FindAll()
}
