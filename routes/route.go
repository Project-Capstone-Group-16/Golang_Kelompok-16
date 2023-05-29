package routes

import (
	"Capstone/controllers"
	"Capstone/middleware"
	"Capstone/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	e.Pre(mid.RemoveTrailingSlash())
	e.Use(mid.CORSWithConfig(mid.CORSConfig{
		AllowOrigins: []string{"https://inventron-indonesia.netlify.app/"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Static("/images/warehouse", "./images/warehouse")

	// login register user dan admin routes
	e.POST("/register", controllers.RegisterUserController)
	e.POST("/register/admin", controllers.RegisterAdminController)
	e.POST("/login", controllers.LoginUserController)
	e.POST("/login/admin", controllers.LoginAdminController)

	// generete otp routes
	fp := e.Group("/forgot-password")
	fp.POST("/generate", controllers.GenerateOTPController)
	fp.POST("/verify", controllers.VerifyngOtpController)
	fp.PUT("/update", controllers.UpdatePasswordController, middleware.IsLoggedIn)

	// admin routes
	adm := e.Group("/admin", middleware.IsLoggedIn)
	adm.POST("/warehouse", controllers.CreateWarehouseController)
	adm.PUT("/warehouse/:id", controllers.UpdateWarehouseController)
	adm.DELETE("/warehouse/:id", controllers.DeleteWarehouseController)
	adm.GET("/warehouse", controllers.GetAllWarehouseController)

	wh := e.Group("/warehouse")
	wh.GET("", controllers.GetStatusWarehouseController)                                    //query params
	wh.POST("/favorite", controllers.AddFavoriteWarehouseController, middleware.IsLoggedIn) // second task

}
