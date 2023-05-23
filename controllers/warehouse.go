package controllers

import (
	// "Capstone/middleware"
	"Capstone/models/payload"
	"Capstone/usecase"
	"net/http"

	"github.com/labstack/echo"
)

func DeleteWarehouse(c echo.Context) error {
	payloadWarehouse := payload.DeleteWarehouseRequest{}

	// adminID, err := middleware.IsAdmin(c)
	// if err != nil {
	// 	return c.JSON(http.StatusNotFound, "Admin not found")
	// }

	c.Bind(&payloadWarehouse)
	err := usecase.DeleteWarehouse(&payloadWarehouse)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "delete complete")
}

func GetAllWarehouse(c echo.Context) error {
	response, err := usecase.GetAllWarehouse()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, response)
}
