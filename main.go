package main

import (
	"Capstone/config"
	"Capstone/middleware"
	"Capstone/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	err := godotenv.Load("deploy.env")
	if err != nil {
		log.Fatal("failed to load .env")
	}

	db := config.InitDB()
	e := echo.New()

	routes.Routes(e, db)
	middleware.Logmiddleware(e)

	go func() {
		e.Logger.Fatal(e.Start(":8080"))
	}()

	select {}
}
