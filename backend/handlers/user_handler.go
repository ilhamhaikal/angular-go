package handlers

import (
	"encoding/json"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"backend/models"
	"backend/database"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var storedUser models.User

	json.NewDecoder(r.Body).Decode(&user)

	row := database.DB.QueryRow("SELECT id, username, password, role FROM users WHERE username=$1", user.Username)
	row.Scan(&storedUser.ID, &storedUser.Username, &storedUser.Password, &storedUser.Role)

	if bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)) != nil {
		http.Error(w, "Invalid login credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(storedUser)
}
