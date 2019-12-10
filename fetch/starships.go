package fetch

import (
	"database/sql"
	"fmt"
)

type Starship struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Model string `json:"model"`
	Crew  string `json:"crew"`
	Passengers string `json:"passengers"`
	Created string `json:"created"`

}

type Starships []Starship

// Function to fetch all Starships from Database
func AllStarships() Starships {
	rows, error := GetDatabase().Query("SELECT id, name, model, crew, passengers, created FROM starships")
	defer rows.Close()

	starships := Starships{}

	if error != nil {
		return starships
	}

	for rows.Next() {
		starships = append(starships, CreateStarships(rows))
	}

	return starships
}

// Function to fetch a particular entry from database
func FindStarships(id string) (Starship, error) {
	rows, error := GetDatabase().Query("SELECT id, name, model, crew, passengers, created FROM starships WHERE id = ? LIMIT 1", id)
	defer rows.Close()

	if error != nil {
		return Starship{}, error
	}

	for rows.Next() {
		return CreateStarships(rows), nil
	}

	return Starship{}, fmt.Errorf("Starships NOT FOUND")
}

// Function to scan the response and create the struct
func CreateStarships(rows *sql.Rows) Starship {
	starship := Starship{}
	rows.Scan(&starship.Id, &starship.Name, &starship.Model, &starship.Crew, &starship.Passengers, &starship.Created)

	return starship
}
