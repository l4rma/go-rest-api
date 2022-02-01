package repository

import (
	"database/sql"
	"log"
	"os"

	"github.com/l4rma/go-rest-api/server/db/entity"

	_ "github.com/mattn/go-sqlite3"
)

// SQLite repository struct
type sqliteRepo struct{}

// Method to initializa new SQLite Bookrepository
func NewSQLiteBookRepo() BookRepository {
	// Remove old db
	os.Remove("./books.db")

	// Open db and create file to store data
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create books table in db
	_, err = db.Exec(createTableBooks)
	if err != nil {
		log.Printf("Error:%v, Query:%v", err, createTableBooks)
	}

	return &sqliteRepo{}
}

// Add a book to the database
func (*sqliteRepo) Save(book *entity.Book) (int64, error) {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare(insertBook)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(book.Title, book.Author, book.Year)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return 0, err
}

// TODO: Implement method
func (*sqliteRepo) FindById(id int64) (*entity.Book, error) {
	return &entity.Book{}, nil
}

// TODO: Implement method
func (*sqliteRepo) FindAll() ([]*entity.Book, error) {
	return []*entity.Book{{Title: "TestTittel", Author: "TestAuthor", Year: 2022}}, nil
}

// TODO: Implement method
func (*sqliteRepo) DeleteById(id int64) error {
	return nil
}

// TODO: Implement method (or remove from interface)
func (*sqliteRepo) InsertDummyData(repo BookRepository) {

}
