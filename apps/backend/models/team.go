package models

import (
	"example.com/mpl/database"
)

type Team struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func GetAllTeams() ([]Team, error) {
	// Get all teams from the database
	var teams []Team
	result := database.DB.Find(&teams)

	if result.Error != nil {
		return nil, result.Error
	}
	return teams, nil
}