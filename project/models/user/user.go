package user

import (
	"fmt"
	"log"
	"os"
	db "pstgSQL/project/database"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil 
}

func RegisterUser(username, password string) error {
	// fmt.Println("Checking username:", username)

	// Проверяем, существует ли пользователь
    var exists bool
    queryCheck := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
    err := db.DB.QueryRow(queryCheck, username).Scan(&exists)
    if err != nil {
        return err
    }

	// fmt.Println("User exists:", exists)

    if exists {
        return fmt.Errorf("Пользователь с таким именем уже существует")
    }
	
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, user_password) VALUES ($1, $2)`
	_, err = db.DB.Exec(query, username, hashedPassword)
	return err
}

func GetUserByUsername(username string) (User , error) {
	var user User
	query := `SELECT id, username, user_password FROM users WHERE username = $1`
	row := db.DB.QueryRow(query, username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}

func GetJWTKey() (string) {
	err := godotenv.Load("project/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return ""
	}

	JWTKey := os.Getenv("JWT_SECRET")
	return  JWTKey
}