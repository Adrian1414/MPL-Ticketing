package database

import (
	// "database/sql"

	// "fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	// "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nedpals/supabase-go"
)

var DB *sqlx.DB
var Client *supabase.Client

func InitDB() {
	// Memuat file .env
    // err := godotenv.Load()
    // if err != nil {
    //     log.Fatal("Error loading .env file")
    // }

	//setup supabase
	supabaseUrl := os.Getenv("SUPABASE_URL")
    supabaseKey := os.Getenv("SUPABASE_KEY")

	if supabaseUrl == "" || supabaseKey == "" {
		log.Fatal("SUPABASE_URL atau SUPABASE_KEY tidak ditemukan dalam environment variables")
	}

	Client = supabase.CreateClient(supabaseUrl, supabaseKey)

    // Mengambil DSN dari environment variable
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        log.Fatal("DB_DSN tidak ditemukan dalam .env")
    }

	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	log.Printf("Connected to the database")
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
}
