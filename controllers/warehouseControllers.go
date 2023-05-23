package controllers

import (
	"Capstone/constants"
	"Capstone/middleware"
	"Capstone/models/payload"
	"Capstone/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

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

func UpdateWarehouseController(c echo.Context) error {
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(401, "Unauthorized")
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	warehouse, err := usecase.GetWarehouseByID(id)
	if err != nil {
		return err
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
