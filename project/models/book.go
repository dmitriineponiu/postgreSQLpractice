package models

import (
	"database/sql"	
	db "pstgSQL/project/database"
)

type Books struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	Published_year int    `json:"published_year"`
	ISBN           string `json:"isbn"`
}

func NewBook(title, author, isbn string, published_year int) error {
	query := `INSERT INTO users (title, author, published_year, isbn)
	VALUES ($1, $2, $3, $4)`
	_, err := db.DB.Exec(query, title, author, published_year)
	return err
}

func GetAllBooks(rows *sql.Rows) ([]Books, error) {
	books := []Books{}
	for rows.Next() {
		var book Books 
		if err := rows.Scan(&book.ID, &book.Title, &book.Author,
			 &book.Published_year, &book.ISBN); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}