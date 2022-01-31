package repository

import (
	"database/sql"
	"log"
	"os"

	"github.com/l4rma/go-rest-api/server/db/entity"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteRepo struct{}

func NewSQLiteBookRepo() BookRepository {
	os.Remove("./books.db")

	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(createTableBooks)
	if err != nil {
		log.Printf("Error:%v, Query:%v", err, createTableBooks)
	}

	return &sqliteRepo{}
}

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

func (*sqliteRepo) FindById(id int64) (*entity.Book, error) {
	return &entity.Book{}, nil
}

func (*sqliteRepo) FindAll() ([]*entity.Book, error) {
	return []*entity.Book{{Title: "TestTittel", Author: "TestAuthor", Year: 2022}}, nil
}
func (*sqliteRepo) DeleteById(id int64) error {
	return nil
}
func (*sqliteRepo) InsertDummyData(repo BookRepository) {

}
