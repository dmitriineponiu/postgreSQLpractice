package handlers

import (
	"encoding/json"
	"net/http"
	db "pstgSQL/project/database"
	"pstgSQL/project/models"
)

func NewBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var newBook models.Books
		err := json.NewDecoder(r.Body).Decode(&newBook)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"Error:":"Error in decode body!"})
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = models.NewBook(newBook.Title, newBook.Author, newBook.ISBN, newBook.Published_year)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"Error":"Error adding data to table"})
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "User added successfully!"})
	}
	w.WriteHeader(http.StatusBadRequest)
}

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		rows, err := db.DB.Query("SELECT id, title, author, publishef_year, isbn")
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"Error":"Error when retrieving information from a table"})
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		books, err := models.GetAllBooks(rows)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"Error":"Error when writing to variable"})
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	}
}