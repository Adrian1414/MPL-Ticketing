package database

import (
	// "database/sql"

	// "fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	// Memuat file .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Mengambil DSN dari environment variable
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        log.Fatal("DB_DSN tidak ditemukan dalam .env")
    }

	DB, err = sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
}
