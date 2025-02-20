package routes

import (
	"encoding/json"
	"net/http"
	"pstgSQL/project/handlers"
)

func SetupRoutes() {
	
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.NewBookHandler(w, r)
		case http.MethodGet:
			handlers.GetAllBooksHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"message":"Method Not Allowed!"})
		}
	})
}