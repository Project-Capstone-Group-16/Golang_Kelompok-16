package controllers

import (
	"Capstone/middleware"
	"Capstone/models/payload"
	"Capstone/usecase"
	"net/http"

	"github.com/labstack/echo"
)

// Register User
func RegisterUserController(c echo.Context) error {
	payloadUser := payload.CreateUserRequest{}
	c.Bind(&payloadUser)

	if err := c.Validate(&payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error payload create user",
			"error":    "password minimum length has to be 5 character",
		})
	}

	response, err := usecase.CreateUser(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error create user",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success register user",
		Data:    response,
	})
}

// Register Admin
func RegisterAdminController(c echo.Context) error {
	payloadUser := payload.CreateAdminRequest{}
	c.Bind(&payloadUser)

	if err := c.Validate(&payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error payload create admin",
			"error":    "password minimum length has to be 5 character",
		})
	}

	response, err := usecase.CreateAdmin(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error create admin",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success register admin",
		Data:    response,
	})
}

// Logic User
func LoginUserController(c echo.Context) error {
	payloadUser := payload.LoginUserRequest{}

	c.Bind(&payloadUser)

	if err := c.Validate(&payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Field can't be empty")
	}

	response, err := usecase.LoginUser(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Login",
		Data:    response,
	})
}

// Logic Admin
func LoginAdminController(c echo.Context) error {
	payloadUser := payload.LoginAdminRequest{}

	c.Bind(&payloadUser)

	if err := c.Validate(&payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Field can't be empty")
	}

	response, err := usecase.LoginAdmin(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(200, payload.Response{
		Message: "Success Login",
		Data:    response,
	})
}

// Generate OTP
func GenerateOTPController(c echo.Context) error {
	payloadUser := payload.ForgotPasswordRequest{}

	c.Bind(&payloadUser)

	if err := c.Validate(&payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Field can't be empty")
	}

	err := usecase.GenerateOTPEndpoint(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "OTP sent successfully, please check your email for the OTP  token ")
}

// Verify OTP
func VerifyngOtpController(c echo.Context) error {
	payloadUser := payload.VerifyngOtpRequest{}

	c.Bind(&payloadUser)

	if err := c.Validate(&payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, "OTP has to be 4 digit")
	}

	err := usecase.VerifyOTP(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "OTP confirmed")
}

// Update Password User
func UpdatePasswordController(c echo.Context) error {
	payloadUser := payload.UpdatePasswordRequest{}

	userId, err := middleware.IsUser(c)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Unauthorized")
	}

	c.Bind(&payloadUser)

	if err := c.Validate(&payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Field cannot be empty")
	}

	err = usecase.UpdatePassword(userId, &payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Change password success")

}

// User add favorite warehouse
func AddFavoriteWarehouseController(c echo.Context) error {
	payloadUser := payload.CreateFavoriteRequest{}

	userId, err := middleware.IsUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	c.Bind(&payloadUser)

	if err := c.Validate(&payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Field cannot be empty")
	}

	response, err := usecase.CreateFavoriteWarehouse(userId, &payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "Succes Favorite Warehouse",
		Data:    response,
	})
}
