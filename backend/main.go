package main

import (
	"log"
	"net/http"
	"os"

	"backend/database"
	"backend/routes"
)

func main() {
	// Ambil variabel lingkungan
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Koneksi ke database
	database.Connect()

	// Register routes
	router := routes.RegisterRoutes()

	// Mulai server
	log.Printf("Server running on port %s", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Gagal memulai server: %v", err)
	}
}
