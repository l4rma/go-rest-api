package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/l4rma/go-rest-api/server/db/entity"

	_ "github.com/lib/pq"
)

var (
	db *sqlx.DB
)

type pgBookRepository struct{}

func NewPostgresRepository() BookRepository {
	pg, err := sqlx.Open("postgres", pgConnStr)
	if err != nil {
		log.Fatal(err)
	}
	//defer pg.Close()

	log.Println("Connected to postgres database.")

	pg.MustExec(createTableBooks)

	db = pg

	return &pgBookRepository{}
}

func (d *pgBookRepository) InsertDummyData(repo BookRepository) {
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
