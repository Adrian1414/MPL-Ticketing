package migration

import (
	"example.com/mpl/database"
	"example.com/mpl/models"
)

func MigrateTeamsTable() {
	err := database.DB.AutoMigrate(&models.Team{})
	if err != nil {
		panic("Could not migrate teams table")
	}
}
