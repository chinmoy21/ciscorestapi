package controllers

import (
	"github.com/labstack/echo"
	"github.com/chin/ciscorestapi/fetch"
	"net/http"
)

// Call the Fetch package function to read all the entries from vehicles table
func VehiclesGetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, fetch.AllVehicles())
}

// Call the Fetch package function to read a particular entry from vehicles table
func ParticularVehicles(c echo.Context) error {
	idStr := c.Param("id")

	vehicle, error := fetch.FindVehicles(idStr)

	if error != nil {
		return c.String(http.StatusNotFound, "Vehicles NOT FOUND")
	} else {
		return c.JSON(http.StatusOK, vehicle)
	}
}
