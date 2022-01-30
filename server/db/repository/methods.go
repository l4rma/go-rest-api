package repository

import "github.com/l4rma/go-rest-api/server/db/entity"

func (d *pgBookRepository) Save(book *entity.Book) (int64, error) {
	res, err := db.Exec(insertBook, book.Title, book.Author, book.Year)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()

	return id, err
}
func (d *pgBookRepository) FindById(id int64) (*entity.Book, error) {
	book := &entity.Book{}
	err := db.Get(book, "SELECT * FROM books WHERE id=$1", id)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (d *pgBookRepository) FindAll() ([]*entity.Book, error) {
	var books []*entity.Book
	err := db.Select(&books, "SELECT * FROM books")
	if err != nil {
		return books, err
	}

	return books, nil
}

func (d *pgBookRepository) DeleteById(id int64) error {
	_, err := db.Exec(deleteBook, id)
	if err != nil {
		return err
	}
	return err
}
