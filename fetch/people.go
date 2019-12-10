package fetch

import (
	"database/sql"
	"fmt"
)

type People struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Birth_year string `json:"birth_year"`
	Created string `json:"created"`
	Vehicles  []Vehicle  `json:"vehicles"`
	Starships []Starship `json:"starships"`
	
}

type Peoples []People

// Function to insert an entry into people table
func (people *People) Insert() (sql.Result, error) {
	res, err := GetDatabase().Exec("INSERT INTO people(id, name, gender, birth_year, created) values(?,?,?,?,?)", people.Id, people.Name, people.Gender, people.Birth_year, people.Created)

	if err != nil {
		return res, err
	}

	return res, nil
}

// Function to update an entry in people table
func (people *People) Update() (sql.Result, error) {
	res, err := GetDatabase().Exec("UPDATE people SET name = ?, gender = ? WHERE id = ?", people.Name, people.Gender, people.Id)

	if err != nil {
		return res, err
	}

	return res, nil
}

// Function to delete an entry from people table
func (people *People) Delete() (sql.Result, error) {
	res, err := GetDatabase().Exec("DELETE FROM people WHERE id = ?", people.Id)

	if err != nil {
		return res, err
	}

	return res, nil
}

// Function to fetch all the entries from people table
func PeopleGetAll() Peoples {
	rows, error := GetDatabase().Query("SELECT id, name, gender, birth_year, created FROM people")

	peoples := Peoples{}

	if error != nil {
		panic(error)
	}

	for rows.Next() {
		peoples = append(peoples, CreatePeople(rows))
	}

	return peoples
}

// Function to fetch a particular entry from people table
func ParticularPeople(id string) (People, error) {
	rows, error := GetDatabase().Query("SELECT id, name, gender, birth_year, created FROM people WHERE id = ? LIMIT 1", id)
	defer rows.Close()

	if error != nil {
		return People{}, error
	}

	for rows.Next() {
		return CreatePeople(rows), nil
	}

	return People{}, fmt.Errorf("Could not find a people")
}

// Function to scan the response and create the struct
func CreatePeople(rows *sql.Rows) People {
	people := People{}
	rows.Scan(&people.Id, &people.Name, &people.Gender, &people.Birth_year, &people.Created)

	people.fetchStarships()
	people.fetchVehicles()

	return people
}

// Function to fetch starships associated with the person
func (people *People) fetchStarships() error {
	rows, error := GetDatabase().Query(`
		SELECT id, name, model, crew, passengers, created
		FROM starships
		INNER JOIN people_starships ON people_starships.starships = starships.id
		WHERE people = ?;
	`, people.Id)
	defer rows.Close()

	if error != nil {
		return error
	}

	starships := Starships{}

	for rows.Next() {
		starships = append(starships, CreateStarships(rows))
	}

	people.Starships = starships

	return nil
}

// Function to fetch vehicles associated with the person
func (people *People) fetchVehicles() error {
	rows, error := GetDatabase().Query(`
		SELECT id, name, model, crew, passengers, created
		FROM vehicles
		INNER JOIN people_vehicles on people_vehicles.vehicles = vehicles.id
		WHERE people = ?;
	`, people.Id)
	defer rows.Close()

	if error != nil {
		return error
	}

	vehicles := Vehicles{}

	for rows.Next() {
		vehicles = append(vehicles, CreateVehicles(rows))
	}

	people.Vehicles = vehicles

	return nil
}
