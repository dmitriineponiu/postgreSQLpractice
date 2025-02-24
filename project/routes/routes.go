package routes

import (
	"encoding/json"
	"net/http"
	"pstgSQL/project/handlers"
	"pstgSQL/project/handlers/authorization"
	jwtmiddleware "pstgSQL/project/middleware/jwtMiddleware"
)

func SetupRoutes() {

	http.HandleFunc("/register", authorization.RegisterHandler)
	http.HandleFunc("/login", authorization.LoginHandler)

	http.Handle("/books", jwtmiddleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.NewBookHandler(w, r)
		case http.MethodGet:
			handlers.GetAllBooksHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"message":"Method Not Allowed!"})
		}
	})))

	http.Handle("/books/", jwtmiddleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			handlers.UpdateBookInfoHandler(w, r)
		case http.MethodDelete:
			handlers.DeleteBookHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"message":"Method Not Allowed!"})
		}
	})))
	
	// http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodPost:
	// 		handlers.NewBookHandler(w, r)
	// 	case http.MethodGet:
	// 		handlers.GetAllBooksHandler(w, r)
	// 	default:
	// 		w.WriteHeader(http.StatusMethodNotAllowed)
	// 		json.NewEncoder(w).Encode(map[string]string{"message":"Method Not Allowed!"})
	// 	}
	// })

	// http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodPut:
	// 		handlers.UpdateBookInfoHandler(w, r)
	// 	case http.MethodDelete:
	// 		handlers.DeleteBookHandler(w, r)
	// 	default:
	// 		w.WriteHeader(http.StatusMethodNotAllowed)
	// 		json.NewEncoder(w).Encode(map[string]string{"message":"Method Not Allowed!"})
	// 	}
	// })
}