package controllers

import (
	"Capstone/middleware"
	"Capstone/models/payload"
	"Capstone/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateWarehouseController(c echo.Context) error {
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(401, "Unauthorized")
	}

	payloadWarehouse := payload.CreateWarehouseRequest{}
	c.Bind(&payloadWarehouse)

	if err := c.Validate(payloadWarehouse); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error payload create Warehouse",
			"error":    "field cannot be empty",
		})
	}

	response, err := usecase.CreateWarehouse(&payloadWarehouse)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error create warehouse",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success create warehouse",
		Data:    response,
	})
}

func UpdateWarehouseController(c echo.Context) error {
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(401, "Unauthorized")
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	warehouse, err := usecase.GetWarehouseByID(id)
	if err != nil {
		return err
	}

	c.Bind(warehouse)

	response, err := usecase.UpdateWarehouse(warehouse)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success update warehouse",
		Data:    response,
	})
}

func DeleteWarehouse(c echo.Context) error {
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(401, "Unauthorized")
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	warehouse, err := usecase.GetWarehouseByID(id)
	if err != nil {
		return err
	}

	err = usecase.DeleteWarehouse(warehouse)
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
