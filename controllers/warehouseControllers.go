package controllers

import (
	"Capstone/constants"
	"Capstone/middleware"
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
		return c.JSON(401, "Unauthorized")
	}

	file, err := c.FormFile("warehouse_image")
	if err != nil {
		return err
	}

	payloadWarehouse := payload.CreateWarehouseRequest{}
	c.Bind(&payloadWarehouse)
	payloadWarehouse.WarehouseImage = file.Filename

	if err := c.Validate(payloadWarehouse); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error payload create Warehouse",
			"error":    "field cannot be empty",
		})
	}

	response, err := usecase.CreateWarehouse(file, &payloadWarehouse)
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
		return c.JSON(401, "Unauthorized")
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	warehouse, err := usecase.GetWarehouseByID(id)
	if err != nil {
		return errors.New("Warehouse not found")
	}

	file, _ := c.FormFile("warehouse_image")

	c.Bind(warehouse)

	if file != nil {
		warehouseImage, _ := usecase.UploadImage(file, warehouse.Name)
		path := fmt.Sprintf("%s/%s", constants.Base_Url, warehouseImage)
		if warehouseImage != "" {
			warehouse.ImageURL = path
		}
	}

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

// Get all warehouse
func GetAllWarehouseController(c echo.Context) error {
	response, err := usecase.GetAllWarehouse()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "Succes get all warehouse",
		Data:    response,
	})
}

// get all warehouse by status
func GetStatusWarehouseController(c echo.Context) error {
	status := c.QueryParam("status")

	response, err := usecase.GetAllByStatusWarehouse(status)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "Succes get all warehouse by status",
		Data:    response,
	})
}
