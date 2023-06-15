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

	// Routes Public
	e.POST("/register", controllers.RegisterUserController)
	e.POST("/register/admin", controllers.RegisterAdminController)
	e.POST("/login", controllers.LoginUserController)
	e.POST("/login/admin", controllers.LoginAdminController)
	e.POST("/notification", controllers.GetNotificationController)
	e.GET("/lockertypes", controllers.GetLockerTypesController)
	e.GET("/itemcategorys", controllers.GetItemCategorysController)

	// generete otp routes
	fp := e.Group("/forgot-password")
	fp.POST("/generate", controllers.GenerateOTPController)
	fp.POST("/verify", controllers.VerifyngOtpController)
	fp.PUT("/update", controllers.UpdatePasswordController, middleware.IsLoggedIn)

	// admin routes
	adm := e.Group("/admin", middleware.IsLoggedIn)
	adm.GET("/users", controllers.GetUsersController)
	adm.POST("/warehouse", controllers.CreateWarehouseController)
	adm.PUT("/warehouse/:id", controllers.UpdateWarehouseController)
	adm.DELETE("/warehouse/:id", controllers.DeleteWarehouseController)
	adm.GET("/warehouse", controllers.GetWarehousesController)
	adm.POST("/staff", controllers.CreateStaffController)
	adm.PUT("/staff/:id", controllers.UpdateStaffController)
	adm.GET("/staff", controllers.GetAllStaffController)
	adm.DELETE("/staff/:id", controllers.DeleteStaffController)
	adm.GET("/locker", controllers.GetLockersController)
	adm.GET("/locker/small/:id", controllers.GetLockerSmallController)
	adm.GET("/locker/medium/:id", controllers.GetLockerMediumController)
	adm.GET("/locker/large/:id", controllers.GetLockerLargeController)
	adm.GET("/lockertypes", controllers.GetLockerTypesController)
	adm.GET("/transactions", controllers.GetTransactionsController)
	adm.GET("/dashboard", controllers.DashboardAdminController)

	wh := e.Group("/warehouse")
	wh.GET("", controllers.GetWarehousesController)
	wh.GET("/recomended", controllers.GetRecomendedWarehouseController)
	wh.POST("/favorite", controllers.AddFavoriteWarehouseController, middleware.IsLoggedIn)

	us := e.Group("/profile", middleware.IsLoggedIn)
	us.GET("", controllers.GetUserController)
	us.GET("/favorite", controllers.GetFavoriteUserByIDController)
	us.GET("/transaction", controllers.GetTransactionByUserIDController)
	us.PUT("/update", controllers.UpdateProfileController)

	t := e.Group("/transaction", middleware.IsLoggedIn)
	t.POST("", controllers.CreateTransactionController)

	// br := e.Group("/beranda", middleware.IsLoggedIn)
	// br.GET("", controllers.GetBerandaController)

	// ex := e.Group("/exlore", middleware.IsLoggedIn)
	// ex.GET("", controllers.GetExploreController)

	picture := e.Group("/upload")
	picture.POST("/image", controllers.UploadImageController)
}
