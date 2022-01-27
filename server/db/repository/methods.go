package repository

import "github.com/l4rma/go-rest-api/server/db/entity"

func (d *repo) Save(book *entity.Book) (int64, error) {
	res, err := d.db.Exec(insertBook, book.Title, book.Author, book.Year)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()

	return id, err
}

func (d *repo) FindAll() ([]*entity.Book, error) {
	var books []*entity.Book
	err := d.db.Select(&books, "SELECT * FROM books")
	if err != nil {
		return books, err
	}

	return books, nil
}
