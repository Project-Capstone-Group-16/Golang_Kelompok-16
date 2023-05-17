package main

import (
	"Capstone/config"

	"github.com/labstack/echo"
)

func main() {
	config.InitDB()
	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))
}
