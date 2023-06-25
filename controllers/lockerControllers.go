package controllers

import (
	"Capstone/middleware"
	"Capstone/models/payload"
	"Capstone/usecase"
	"strconv"

	"github.com/labstack/echo"
)

func GetLockersController(c echo.Context) error {
	_, err := middleware.IsAdmin(c)
	if err != nil {
		return c.JSON(400, "this route for admin only")
	}

	response, err := usecase.GetLockers()
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, payload.Response{
		Message: "success get lockers",
		Data:    response,
	})
}

func GetLockerTypesController(c echo.Context) error {
	response, err := usecase.GetAllLockerTypes()
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, payload.Response{
		Message: "success get all locker types",
		Data:    response,
	})
} //new

func GetLockerSmallController(c echo.Context) error {
	_, err := middleware.IsAdmin(c)
	if err != nil {
		return c.JSON(400, "this route for admin only")
	}

	WarehouseId, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	response, err := usecase.GetLockerSmall(uint(WarehouseId))
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, payload.Response{
		Message: "success get lockers",
		Data:    response,
	})
}

func GetLockerMediumController(c echo.Context) error {
	_, err := middleware.IsAdmin(c)
	if err != nil {
		return c.JSON(400, "this route for admin only")
	}

	WarehouseId, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	response, err := usecase.GetLockerMedium(uint(WarehouseId))
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, payload.Response{
		Message: "success get lockers",
		Data:    response,
	})
}

func GetLockerLargeController(c echo.Context) error {
	_, err := middleware.IsAdmin(c)
	if err != nil {
		return c.JSON(400, "this route for admin only")
	}

	WarehouseId, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	response, err := usecase.GetLockerLarge(uint(WarehouseId))
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, payload.Response{
		Message: "success get lockers",
		Data:    response,
	})
}
