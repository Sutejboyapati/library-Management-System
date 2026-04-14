package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
<<<<<<< HEAD
	"strings"
	"time"

	"library/backend/apiutil"
=======
	"os"
	"time"

>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
	"library/backend/config"
	"library/backend/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type claims struct {
	UserID int    `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func Register(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := validateCredentials(u.Username, u.Password); err != nil {
		apiutil.WriteError(w, http.StatusBadRequest, err.Error())
=======
		http.Error(w, "Invalid request", http.StatusBadRequest)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	var exists bool
	err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM Users WHERE username = ?)", u.Username).Scan(&exists)
	if err != nil {
		log.Println("DB error:", err)
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if exists {
		apiutil.WriteError(w, http.StatusBadRequest, "Username already taken")
=======
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Username already taken", http.StatusBadRequest)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Server error")
=======
		http.Error(w, "Server error", http.StatusInternalServerError)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	role := u.Role
	if role == "" {
		role = "user"
	}
	_, err = config.DB.Exec("INSERT INTO Users (username, password, role) VALUES (?, ?, ?)", u.Username, string(hash), role)
	if err != nil {
		log.Println("DB insert:", err)
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Database error")
		return
	}
	apiutil.WriteJSON(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
=======
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.LoginCredentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if strings.TrimSpace(creds.Username) == "" || strings.TrimSpace(creds.Password) == "" {
		apiutil.WriteError(w, http.StatusBadRequest, "Username and password are required")
=======
		http.Error(w, "Invalid request", http.StatusBadRequest)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	var id int
	var username, password, role string
	err := config.DB.QueryRow("SELECT id, username, password, role FROM Users WHERE username = ?", creds.Username).
		Scan(&id, &username, &password, &role)
	if err == sql.ErrNoRows {
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusUnauthorized, "Invalid username or password")
=======
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
		return
	}
	if err != nil {
		log.Println("DB error:", err)
<<<<<<< HEAD
		apiutil.WriteError(w, http.StatusInternalServerError, "Server error")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(creds.Password)); err != nil {
		apiutil.WriteError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}
=======
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "change-me-in-production"
	}
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
	c := claims{
		UserID: id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
<<<<<<< HEAD
	tokenStr, err := token.SignedString([]byte(apiutil.JWTSecret()))
	if err != nil {
		apiutil.WriteError(w, http.StatusInternalServerError, "Server error")
		return
	}
	apiutil.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message":  "Login successful",
		"token":    tokenStr,
		"username": username,
		"role":     role,
		"userId":   id,
=======
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Login successful",
		"token":    tokenStr,
		"username": username,
>>>>>>> 2bf3000141f874abc2bfb95f4c31477fae075504
	})
}
