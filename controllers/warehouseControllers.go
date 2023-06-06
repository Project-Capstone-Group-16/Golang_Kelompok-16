package controllers

import (
	"Capstone/middleware"
	"Capstone/models"
	"Capstone/models/payload"
	"Capstone/usecase"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Create Warehouse
func CreateWarehouseController(c echo.Context) error {
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for admin",
		})
	}

	payloadWarehouse := payload.CreateWarehouseRequest{}

	c.Bind(&payloadWarehouse)

	if err := c.Validate(payloadWarehouse); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload create warehouse",
			"error":   err.Error(),
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

// Update Warehouse
func UpdateWarehouseController(c echo.Context) error {
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for admin",
		})
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	warehouse, err := usecase.GetWarehouseByID(id)
	if err != nil {
		return errors.New("Warehouse not found")
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

// Delete Warehouse
func DeleteWarehouseController(c echo.Context) error {
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for admin",
		})
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	warehouse, err := usecase.GetWarehouseByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = usecase.DeleteWarehouse(warehouse)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "delete complete")
}

// Get all warehouse
// get all warehouse by status
func GetWarehousesController(c echo.Context) error {
	warehouseParams := models.Warehouse{
		Status:   c.QueryParam("status"),
		City:     c.QueryParam("city"),
		Province: c.QueryParam("province"),
	}

	response, err := usecase.GetWarehouses(&warehouseParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: fmt.Sprintf("Succes get all warehouse by status %s", warehouseParams.Status),
		Data:    response,
	})
}

func GetRecomendedWarehouseController(c echo.Context) error {
	warehouseParams := models.Warehouse{
		Status:   c.QueryParam("status"),
		City:     c.QueryParam("city"),
		Province: c.QueryParam("province"),
	}

	response, err := usecase.GetRecomendedWarehouse(&warehouseParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "Succes get all recomended warehouse",
		Data:    response,
	})
}
