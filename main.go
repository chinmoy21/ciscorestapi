package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/chin/ciscorestapi/controllers"
)

const port = ":8000"

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// README
	e.GET("/", controllers.Home)

	// PEOPLE
	e.GET("/people/", controllers.PeopleGetAll)
	e.GET("/people/:id", controllers.ParticularPeople)
	e.POST("/people/", controllers.PeopleCreate)
	e.PUT("/people/:id", controllers.PeopleUpdate)
	e.DELETE("/people/:id", controllers.PeopleDelete)

	// VEHICLES
	e.GET("/vehicles/", controllers.VehiclesGetAll)
	e.GET("/vehicles/:id", controllers.ParticularVehicles)
	
	// STARSHIPS
	e.GET("/starships/", controllers.StarshipsGetAll)
	e.GET("/starships/:id", controllers.ParticularStarships)
	
	// PORT
	e.Logger.Fatal(e.Start(":1200"))
}
