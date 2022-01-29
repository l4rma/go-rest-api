package repository

const createTableBooks = `
CREATE TABLE IF NOT EXISTS books
(
	id SERIAL PRIMARY KEY,
	title VARCHAR(50),
	author VARCHAR(100),
	year BIGINT
)
`

var (
	insertBook = `INSERT INTO books (title, author, year) VALUES ($1, $2, $3) RETURNING id`
	getBook    = `SELECT * FROM books WHERE id = $1;`
	deleteBook = `DELETE FROM books WHERE id = $1;`
)
