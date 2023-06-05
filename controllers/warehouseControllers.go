package controllers

import (
	"Capstone/constants"
	"Capstone/middleware"
	"Capstone/models"
	"Capstone/models/payload"
	"Capstone/usecase"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

// Create Warehouse
func CreateWarehouseController(c echo.Context) error {
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for admin",
		})
	}

	fileHeader, err := c.FormFile("warehouse_image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload create warehouse",
			"error":   "Warehouse image can't be empty",
		})
	}

	// file, _ := fileHeader.Open()

	payloadWarehouse := payload.CreateWarehouseRequest{}

	c.Bind(&payloadWarehouse)

	payloadWarehouse.WarehouseImage = fileHeader.Filename

	if err := c.Validate(payloadWarehouse); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload create warehouse",
			"error":   err.Error(),
		})
	}

	response, err := usecase.CreateWarehouse(fileHeader, &payloadWarehouse)
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

	fileHeader, _ := c.FormFile("warehouse_image")

	file, _ := fileHeader.Open()

	c.Bind(&warehouse)

	warehouse.ImageURL = fileHeader.Filename

	response, err := usecase.UpdateWarehouse(file, warehouse)
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
	rmPath := strings.TrimLeft(warehouse.ImageURL, constants.Base_Url+"/")
	os.Remove(rmPath)

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
		Location: c.QueryParam("location"),
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
		Location: c.QueryParam("location"),
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
