package handlers

import (
	"encoding/json"
	"net/http"
	"pstgSQL/project/models"
	"strconv"
	"strings"
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
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		books, err := models.GetAllBooks()
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"Error":"Error when writing to variable"})
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
	}
	w.WriteHeader(http.StatusBadRequest)
}

func UpdateBookInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		strID := strings.TrimPrefix(r.URL.Path, "/books/")
		ID, err := strconv.Atoi(strID)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"Error":"ID must be int!"})
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var newInfo models.Books
		err = json.NewDecoder(r.Body).Decode(&newInfo)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"Error:":"Error in decode body!"})
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		err = models.Updateinfo(ID, newInfo)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"Error":"Error in updating info!"})
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message":"Book was successfully updated!"})
	}
	w.WriteHeader(http.StatusBadRequest)
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		strID := strings.TrimPrefix(r.URL.Path, "/books/")
		ID, err := strconv.Atoi(strID)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"Error":"ID must be int!"})
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = models.DeleteBook(ID)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"Error":"Error in deletung book!"})
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message":"Book deleted successfully!"})
	}
	w.WriteHeader(http.StatusBadRequest)
}