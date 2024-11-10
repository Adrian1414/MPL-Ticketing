package models

import "example.com/mpl/database"

type Roster struct {
	ID     int64  `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Role   string `json:"role" db:"role"`
	TeamID int64  `json:"team_id" db:"team_id"`
	Photo  string `json:"photo" db:"photo"`
}

func GetAllRosters() ([]Roster, error) {
	// Get all teams from the database
	var rosters []Roster
	query := "SELECT id, name, role, team_id, photo FROM rosters"
	// select untuk mendapatkan banyak rows
	// err := sqlx.Select(database.DB, &teams, query)
	err := database.DB.Select(&rosters, query)
	if err != nil {
		return nil, err
	}
	return rosters, nil
}

func getAvailableIDRoster() (int64, error) {
	var id int64
	query := `
        SELECT COALESCE(
            (SELECT MIN(id) + 1 FROM rosters WHERE (id + 1) NOT IN (SELECT id FROM rosters)),
            (SELECT COALESCE(MAX(id), 0) + 1 FROM rosters)
        ) AS available_id;
    `
	err := database.DB.Get(&id, query)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Roster) SaveRoster() error {
	availableID, err := getAvailableIDRoster()
	if err != nil {
		return err
	}
	r.ID = availableID
	// Create new team in the database
	query := "INSERT INTO rosters (id, name, role, team_id, photo) VALUES ($1, $2, $3, $4, $5)"
	// Menggunakan QueryRowx untuk mendapatkan ID yang baru
	// err = database.DB.QueryRowx(query,t.ID, t.Name)
	_, err = database.DB.Exec(query, r.ID, r.Name, r.Role, r.TeamID, r.Photo)
	if err != nil {
		return err
	}
	return nil
}
