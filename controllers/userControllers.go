package controllers

import (
	"Capstone/models/payload"
	"Capstone/usecase"
	"net/http"

	"github.com/labstack/echo"
)

func RegisterUserController(c echo.Context) error {
	payloadUser := payload.CreateUserRequest{}
	c.Bind(&payloadUser)

	if err := c.Validate(payloadUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error payload create user",
			"error":    "password minimum length has to be 5 character",
		})
	}

	response, err := usecase.CreateUser(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": "error create user",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success register user",
		Data: response,
	})
}
