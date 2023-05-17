package main

import (
	"Capstone/config"
	"Capstone/middleware"
	"Capstone/routes"

	"github.com/labstack/echo"
)

func main() {
	db := config.InitDB()
	e := echo.New()

	routes.Routes(e, db)
	middleware.Logmiddleware(e)

	e.Logger.Fatal(e.Start(":8080"))
}
