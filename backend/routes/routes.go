package routes

import (
	"backend/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	
	// Gunakan router untuk semua rute API
	router.HandleFunc("/api/login", handlers.LoginUser).Methods("POST")
	router.HandleFunc("/api/cars", handlers.GetCars).Methods("GET")
	router.HandleFunc("/api/cars", handlers.CreateCar).Methods("POST")
	router.HandleFunc("/api/cars/{id}", handlers.DeleteCar).Methods("DELETE")

	// Tambahkan rute lain ke dalam router
	router.HandleFunc("/admin/add-car", handlers.AddCar)
	router.HandleFunc("/login", handlers.LoginUser)

	// Kembalikan router untuk digunakan di main.go
	return router
}
