package routes

import (
	"Capstone/controllers"
	"Capstone/middleware"
	"Capstone/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	e.Pre(mid.RemoveTrailingSlash())
	e.Use(mid.CORS())

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
	adm.GET("/warehouse", controllers.GetStatusWarehouseController)
	adm.POST("/staff", controllers.CreateStaffController)
	adm.PUT("/staff/:id", controllers.UpdateStaffController)
	adm.GET("/staff", controllers.GetAllStaffController)
	adm.DELETE("/staff/:id", controllers.DeleteStaffController)

	wh := e.Group("/warehouse",  middleware.IsLoggedIn)
	wh.GET("", controllers.GetStatusWarehouseController)				//query params
	wh.POST("/favorite", controllers.AddFavoriteWarehouseController) // second task

}
