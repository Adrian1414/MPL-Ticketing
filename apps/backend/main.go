package main

import (
	// "fmt"
	// "log"

	"example.com/mpl/database"
	"example.com/mpl/migration"
	"example.com/mpl/routes"
	"example.com/mpl/seeder"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	migration.MigrateTeamsTable()
	seeder.SeedData()

	server := gin.Default()

	routes.RegisterRoute(server)

	server.Run(":3000")

	// log.Println("Server is running on http://localhost:3000")
	// fmt.Println("Server is running on http://localhost:3000")
}
