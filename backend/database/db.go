package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

var DB *sql.DB

// Fungsi Connect yang digunakan untuk koneksi ke PostgreSQL
func Connect() {
	// Mengambil detail koneksi dari variabel lingkungan
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Format koneksi string PostgreSQL
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Membuka koneksi ke database
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi: %v", err)
	}

	// Cek koneksi
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	log.Println("Koneksi ke database berhasil!")
}
