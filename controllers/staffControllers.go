package controllers

import (
	"Capstone/middleware"
	"Capstone/models/payload"
	"Capstone/usecase"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateStaffController(c echo.Context) error {
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for admin",
		})
	}

	payloadStaff := payload.CreateStaffRequest{}
	c.Bind(&payloadStaff)

	if err := c.Validate(payloadStaff); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error payload create staff",
			"error":   err.Error(),
		})
	}

	response, err := usecase.CreateStaff(&payloadStaff)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "error create staff",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success create staff",
		Data:    response,
	})
}

func UpdateStaffController(c echo.Context) error {
	if _, err := middleware.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"Message": "this route only for admin",
		})
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	staff, err := usecase.GetStaffByID(id)
	if err != nil {
		return errors.New("Staff not found")
	}

	c.Bind(staff)

	response, err := usecase.UpdateStaff(staff)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success update staff",
		Data:    response,
	})
}