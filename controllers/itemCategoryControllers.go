package controllers

import (
	"Capstone/models/payload"
	"Capstone/usecase"

	"github.com/labstack/echo"
)

func GetItemCategorysController(c echo.Context) error {
	response, err := usecase.GetAllItemCategorys()
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, payload.Response{
		Message: "success get all item categories",
		Data:    response,
	})
} //new
