package models

import (
	"example.com/mpl/database"
	"github.com/jmoiron/sqlx"
)

type Team struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func GetAllTeams() ([]Team, error) {
	// Get all teams from the database
	var teams []Team
	query := "SELECT id, name FROM teams"
	err := sqlx.Select(database.DB, &teams, query)
	if err != nil {
		return nil, err
	}
	return teams, nil
}
