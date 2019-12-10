package fetch

import (
	"database/sql"
	"fmt"
	"log"
)

type Vehicle struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Model string `json:"model"`
	Crew  string `json:"crew"`
	Passengers string `json:"passengers"`
	Created string `json:"created"`

	
}

type Vehicles []Vehicle

// Function to fetch all Vehicles from Database
func AllVehicles() Vehicles {
	rows, error := GetDatabase().Query("SELECT id, name, model, crew, passengers, created FROM vehicles")
	defer rows.Close()

	vehicles := Vehicles{}

	if error != nil {
		log.Print(error)
	}

	for rows.Next() {
		vehicles = append(vehicles, CreateVehicles(rows))
	}

	return vehicles
}

// Function to fetch a particular entry from Database
func FindVehicles(id string) (Vehicle, error) {
	rows, error := GetDatabase().Query("SELECT id, name, model, crew, passengers, created FROM vehicles WHERE id = ? LIMIT 1", id)
	defer rows.Close()
	
	if error != nil {
		log.Print(error)
		return Vehicle{}, error
	}

	for rows.Next() {
		return CreateVehicles(rows), nil
	}

	return Vehicle{}, fmt.Errorf("Vehicle NOT FOUND")
}

// Function to scan the response and create the struct
func CreateVehicles(rows *sql.Rows) Vehicle {
	vehicle := Vehicle{}
	rows.Scan(&vehicle.Id, &vehicle.Name, &vehicle.Model, &vehicle.Crew, &vehicle.Passengers, &vehicle.Created)

	return vehicle
}
