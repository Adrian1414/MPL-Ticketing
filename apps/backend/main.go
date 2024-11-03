package main

import (
	// "fmt"
	// "log"

	"example.com/mpl/database"
	"example.com/mpl/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	defer database.DB.Close()
	
	server := gin.Default()

	routes.RegisterRoute(server)

	server.Run(":3000")

	// log.Println("Server is running on http://localhost:3000")
	// fmt.Println("Server is running on http://localhost:3000")
}
