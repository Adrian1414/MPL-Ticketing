package main

import (
	// "fmt"
	"log"
	"os"

	"example.com/mpl/database"
	"example.com/mpl/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	// defer database.DB.Close()
	
	server := gin.Default()

	routes.RegisterRoute(server)

	// Ambil port dari environment variable PORT (default 3000 jika tidak ada)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Port default untuk pengembangan lokal
	}

	// Jalankan server pada port yang benar
	if err := server.Run(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}

	// log.Println("Server is running on http://localhost:3000")
	// fmt.Println("Server is running on http://localhost:3000")
}
