package main

import (
	"fmt"
	"log"
	"net/http"
	db "pstgSQL/project/database"
	"pstgSQL/project/routes"
)

func main() {

	db.Connect()
	db.SQLScript()
	routes.SetupRoutes()
	fmt.Println("Server is ON!")
	// mux := http.NewServeMux()
	// http.ListenAndServe("localhost:8080", mux)
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}