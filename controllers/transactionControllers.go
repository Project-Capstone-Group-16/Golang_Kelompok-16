package controllers

import (
	"Capstone/middleware"
	"Capstone/models/payload"
	"Capstone/usecase"
	"net/http"

	"github.com/labstack/echo"
)

func CreateTransactionController(c echo.Context) error {
	palyloadTransaction := payload.CreateTransactionRequest{}

	userId, err := middleware.IsUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for user",
		})
	}

	c.Bind(&palyloadTransaction)

	if err := c.Validate(&palyloadTransaction); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload create transaction",
			"error":   err.Error(),
		})
	}

	response, err := usecase.CreateTransaction(userId, &palyloadTransaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success create transaction",
		Data:    response,
	})
}
