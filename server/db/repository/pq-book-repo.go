package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/l4rma/go-rest-api/server/db/entity"

	_ "github.com/lib/pq"
)

// An extension on the standard "databases/sql" library
var (
	db *sqlx.DB
)

// Repository struct
type pgBookRepository struct{}

// Method to initialize a new repository
func NewPostgresRepository() BookRepository {
	// Open and connect to the database
	pg, err := sqlx.Open("postgres", pgConnStr)
	if err != nil {
		log.Fatal(err)
	}
	// Close the db when function ends
	//defer pg.Close()

	log.Println("Connected to postgres database.")

	// Create the books table in the db
	pg.MustExec(createTableBooks)

	// Set the postgres db as the db of this struct
	db = pg

	return &pgBookRepository{}
}

// Insert a book in the database
func (d *pgBookRepository) Save(book *entity.Book) (int64, error) {
	res, err := db.Exec(insertBook, book.Title, book.Author, book.Year)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()

	return id, err
}

// Retrieve a book from the db with matching ID
func (d *pgBookRepository) FindById(id int64) (*entity.Book, error) {
	book := &entity.Book{}
	err := db.Get(book, "SELECT * FROM books WHERE id=$1", id)
	if err != nil {
		return book, err
	}

	return book, nil
}

// Retrieve all books from db
func (d *pgBookRepository) FindAll() ([]*entity.Book, error) {
	var books []*entity.Book
	err := db.Select(&books, "SELECT * FROM books")
	if err != nil {
		return books, err
	}

	return books, nil
}

// Delete a book in the db with matching ID
func (d *pgBookRepository) DeleteById(id int64) error {
	_, err := db.Exec(deleteBook, id)
	if err != nil {
		return err
	}
	return err
}

// Temp function to create some dummy data for manual testing.
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
