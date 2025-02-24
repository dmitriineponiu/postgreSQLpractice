package authorization

import (
	// "bytes"
	"encoding/json"
	// "fmt"
	// "io"
	"net/http"
	"pstgSQL/project/models/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


var jwtKey = []byte(user.GetJWTKey())

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"user_password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var cred Credentials

	// // Читаем тело запроса для отладки
    // bodyBytes, _ := io.ReadAll(r.Body)
    // fmt.Println("Raw Request Body:", string(bodyBytes))
	
	// // Перечитываем тело запроса (после ReadAll его нужно восстановить)
    // r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	
	// fmt.Println("Username:", cred.Username)
    // fmt.Println("Password:", cred.Password)

	err = user.RegisterUser(cred.Username, cred.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully!"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var cred Credentials
	json.NewDecoder(r.Body).Decode(&cred)

	user, err := user.GetUserByUsername(cred.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

 