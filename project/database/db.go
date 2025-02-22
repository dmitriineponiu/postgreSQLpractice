package db

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {	

	err := godotenv.Load("project/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUSER := os.Getenv("DB_USER") 
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUSER, dbPass, dbName, dbHost, dbPort)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Error connecting to database ", err)
	}

	fmt.Println("Succsessfully connected to DB!")
}

func SQLScript() {

	sqlFile, err := os.Open("project/database/table_books.sql")
	if err != nil {
		log.Fatal("Error opening SQL file: ", err)
	}
	defer sqlFile.Close()

	sqlBytes, err := io.ReadAll(sqlFile)
	if err != nil {
		log.Fatal("Error reading SQL file: ", err)
	}

	_, err = DB.Exec(string(sqlBytes))
	if err != nil {
		log.Fatal("Error with SQL file: ", err)
	}
	fmt.Println("Table 'books' created successfully!")
}