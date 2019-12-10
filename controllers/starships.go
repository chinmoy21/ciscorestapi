package controllers

import (
	"github.com/labstack/echo"
	"github.com/chin/ciscorestapi/fetch"
	"net/http"
)

// Call the Fetch package function to read all the entries from starships table
func StarshipsGetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, fetch.AllStarships())
}

// Call the Fetch package function to read a particular entry from starships table
func ParticularStarships(c echo.Context) error {
	idStr := c.Param("id")

	vehicle, error := fetch.FindStarships(idStr)

	if error != nil {
		return c.String(http.StatusNotFound, "Starships NOT FOUND")
	} else {
		return c.JSON(http.StatusOK, vehicle)
	}
}
