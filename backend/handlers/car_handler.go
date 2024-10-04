package handlers

import (
	"encoding/json"
	"net/http"
	"backend/models"
	"backend/database"
)

func CreateCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	json.NewDecoder(r.Body).Decode(&car)
	query := "INSERT INTO cars (name, description, price, imageUrl) VALUES ($1, $2, $3, $4)"
	_, err := database.DB.Exec(query, car.Name, car.Description, car.Price, car.ImageURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(car)
}

func GetCars(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name, description, price, imageUrl FROM cars")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		rows.Scan(&car.ID, &car.Name, &car.Description, &car.Price, &car.ImageURL)
		cars = append(cars, car)
	}
	json.NewEncoder(w).Encode(cars)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	// logic for delete by ID (using r.URL.Query().Get("id"))
}
