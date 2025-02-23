package models

import (
	"fmt"
	"log"
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
	query := `INSERT INTO book (title, author, published_year, isbn)
	VALUES ($1, $2, $3, $4)`
	_, err := db.DB.Exec(query, title, author, published_year, isbn)
	return err
}

func GetAllBooks() ([]Books, error) {
	rows, err := db.DB.Query("SELECT id, title, author, published_year, isbn FROM book")
		if err != nil {			
			log.Fatal("Error when querying data from table.\n", err)
			return nil, err
		}
		defer rows.Close()
	books := []Books{}
	for rows.Next() {
		var book Books 
		if err := rows.Scan(&book.ID, &book.Title, &book.Author,
			 &book.Published_year, &book.ISBN); err != nil {
			log.Fatal("Error scanning book.\n", err)
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func IfExistID(ID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM book WHERE id = $1)`
	var exist bool
	err := db.DB.QueryRow(query, ID).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func Updateinfo(ID int, newInfo Books) error {	
	exist, err :=IfExistID(ID); 
	if err != nil {
		return err
	}

	if !exist {
		return fmt.Errorf("Book with ID %d not found", ID)
	}

	query := `UPDATE book SET title = $1, author = $2 WHERE id = $3`
	_, err = db.DB.Exec(query, newInfo.Title, newInfo.Author, ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(ID int) error {
	exist, err :=IfExistID(ID); 
	if err != nil {
		return err
	}

	if !exist {
		return fmt.Errorf("Book with ID %d not found", ID)
	}

	query := `DELETE FROM book WHERE id = $1`
	_, err = db.DB.Exec(query, ID)
	if err != nil {		 
		return err
	}
	return nil
}