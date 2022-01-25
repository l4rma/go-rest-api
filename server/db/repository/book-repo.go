package repository

import (
	"log"
	"rest-api/server/db/entity"

	"github.com/jmoiron/sqlx"
)

type BookRepository interface {
	Open() error
	Close() error
	Save(book *entity.Book) (int64, error)
	//FindById(id int64) (*entity.Book, error)
	FindAll() ([]*entity.Book, error)
	//UpdateById(id int64) (*entity.Book, error)
	//DeleteById(id int64) (*entity.Book, error)
}

type repo struct {
	db *sqlx.DB
}

func NewPostgresRepository() BookRepository {
	return &repo{}
}

func (d *repo) Open() error {
	pg, err := sqlx.Open("postgres", pgConnStr)
	if err != nil {
		return err
	}
	log.Println("Connected to database.")
	
	pg.MustExec(createTableBooks)

	d.db = pg

	return nil
}

func (d *repo) Close() error {
	return d.db.Close()
}
