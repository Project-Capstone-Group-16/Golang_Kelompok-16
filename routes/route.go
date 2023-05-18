package routes

import (
	"Capstone/controllers"
	"Capstone/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	e.Pre(mid.RemoveTrailingSlash())

	e.POST("/register", controllers.RegisterUserController)
	e.POST("/login", controllers.LoginUserController)
}
