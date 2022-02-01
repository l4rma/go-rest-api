package repository

import (
	"github.com/l4rma/go-rest-api/server/db/entity"
)

// Book reposotory. CRUD operations to database.
// Interface to abstract the db type.
type BookRepository interface {
	Save(book *entity.Book) (int64, error)
	FindById(id int64) (*entity.Book, error)
	FindAll() ([]*entity.Book, error)
	//UpdateById(id int64) (*entity.Book, error)
	DeleteById(id int64) error
	InsertDummyData(repo BookRepository)
}
