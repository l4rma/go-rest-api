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
func (d *repo) FindById(id int64) (*entity.Book, error) {
	book := &entity.Book{}
	err := d.db.Get(book, "SELECT * FROM books WHERE id=$1", id)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (d *repo) FindAll() ([]*entity.Book, error) {
	var books []*entity.Book
	err := d.db.Select(&books, "SELECT * FROM books")
	if err != nil {
		return books, err
	}

	return books, nil
}

func (d *repo) DeleteById(id int64) error {
	_, err := d.db.Exec(deleteBook, id)
	if err != nil {
		return err
	}
	return err
}
