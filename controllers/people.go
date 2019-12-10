package controllers

import (
	"github.com/labstack/echo"
	"github.com/chin/ciscorestapi/fetch"
	"net/http"
)

// Call the Fetch package function to read all the entries from people table
func PeopleGetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, fetch.PeopleGetAll())
}

// Call the Fetch package function to read a particular entry from people table
func ParticularPeople(c echo.Context) error {
	idStr := c.Param("id")
	people, error := fetch.ParticularPeople(idStr)

	if error != nil {
		return c.String(http.StatusNotFound, "People NOT FOUND")
	} else {
		return c.JSON(http.StatusOK, people)
	}
}

// Call the Fetch package function to create a new entry in people table
func PeopleCreate(c echo.Context) error {
	people := &fetch.People{}

	if err := c.Bind(people); err != nil {
		return err
	}

	if _, err := people.Insert(); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, people)
}

// Call the Fetch package function to update an entry in people table
func PeopleUpdate(c echo.Context) error {
	people := &fetch.People{}
	people.Id = c.Param("id")

	if err := c.Bind(people); err != nil {
		return err
	}

	if _, err := people.Update(); err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, people)
}

// Call the Fetch package function to delete an entry from people table
func PeopleDelete(c echo.Context) error {
	idStr := c.Param("id")
	people, error := fetch.ParticularPeople(idStr)

	if error != nil {
		return c.String(http.StatusNotFound, "People NOT FOUND")
	}

	if _, err := people.Delete(); err != nil {
		return err
	}

	return c.String(http.StatusNoContent, "People DELETED")
}
