package handlers

import (
	"encoding/json"
	"net/http"

	"backend/database"
	"backend/models"
)

func AddCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	json.NewDecoder(r.Body).Decode(&car)

	_, err := database.DB.Exec("INSERT INTO cars (name, description, price) VALUES ($1, $2, $3)", car.Name, car.Description, car.Price)
	if err != nil {
		http.Error(w, "Error adding car", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Car added successfully")
}
