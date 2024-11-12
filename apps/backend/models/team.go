package models

import (
	// "database/sql"

	"errors"
	"log"

	"example.com/mpl/database"
	// "github.com/jmoiron/sqlx"
)

type Team struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func GetAllTeams() ([]Team, error) {
	if database.DB == nil {
		log.Println("Database connection is nil")
		return nil, errors.New("database connection is not initialized")
	}
	// Get all teams from the database
	var teams []Team
	query := "SELECT id, name FROM teams ORDER BY name"
	// select untuk mendapatkan banyak rows
	// err := sqlx.Select(database.DB, &teams, query)
	err := database.DB.Select(&teams, query)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func GetTeamByID(id int64) (*Team, error) {
	// Get team by id from the database
	var team Team
	query := "SELECT id, name FROM teams WHERE id = $1"

	err := database.DB.Get(&team, query, id)
	if err != nil {
		return nil, err
	}
	return &team, nil
}

// Fungsi untuk mendapatkan ID terkecil yang kosong
func getAvailableIDTeam() (int64, error) {
	var id int64
	query := `
        SELECT COALESCE(
            (SELECT MIN(id) + 1 FROM teams WHERE (id + 1) NOT IN (SELECT id FROM teams)),
            (SELECT COALESCE(MAX(id), 0) + 1 FROM teams)
        ) AS available_id;
    `
	err := database.DB.Get(&id, query)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (t *Team) SaveTeam() error {
	availableID, err := getAvailableIDTeam()
	if err != nil {
		return err
	}
	t.ID = availableID
	// Create new team in the database
	query := "INSERT INTO teams (id, name) VALUES ($1, $2)"
	// Menggunakan QueryRowx untuk mendapatkan ID yang baru
	// err = database.DB.QueryRowx(query,t.ID, t.Name)
	_, err = database.DB.Exec(query, t.ID, t.Name)
	if err != nil {
		return err
	}
	return nil
}

func (t *Team) UpdateTeam() error {
	// Update existing team in the database
	query := "UPDATE teams SET name = $1 WHERE id = $2"
	_, err := database.DB.Exec(query, t.Name, t.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t *Team) DeleteTeam() error {
	// Delete team by id from the database
	query := "DELETE FROM teams WHERE id = $1"
	// exec untuk menghapus satu row
	_, err := database.DB.Exec(query, t.ID)
	if err != nil {
		return err
	}

	return nil
}
