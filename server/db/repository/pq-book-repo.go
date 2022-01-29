package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/l4rma/go-rest-api/server/db/entity"
)

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

func (d *repo) InsertDummyData(repo BookRepository) {
	books := []*entity.Book{
		{Title: "Hakkebakkeskogen", Author: "Thorbjørn Egner", Year: 1953},
		{Title: "Folk og røvere i Kardemomme by", Author: "Thorbjørn Egner", Year: 1955},
		{Title: "Refactoring", Author: "Martin Fowler", Year: 2019},
	}

	for _, book := range books {
		repo.Save(book)
		log.Printf("Inserted %v into db", book)
	}
}
