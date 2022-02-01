package repository

// Create table
const createTableBooks = `
CREATE TABLE IF NOT EXISTS books
(
	id SERIAL PRIMARY KEY,
	title VARCHAR(50),
	author VARCHAR(100),
	year BIGINT
)
`

// SQL queries: add book to db, retrieve book from db, delete book in db.
var (
	insertBook = `INSERT INTO books (title, author, year) VALUES ($1, $2, $3) RETURNING id`
	getBook    = `SELECT * FROM books WHERE id = $1;`
	deleteBook = `DELETE FROM books WHERE id = $1;`
)
