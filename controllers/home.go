package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

// Home function to display readme file
func Home(c echo.Context) error {
	return c.File("README.md")
	return c.String(http.StatusOK, "Hello Cisco")
}
