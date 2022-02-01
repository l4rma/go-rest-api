package entity

type Book struct {
	ID     int64  `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Year   int16  `db:"year"`
}

type JsonBook struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int16  `json:"year"`
}

type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int16  `json:"year"`
}
