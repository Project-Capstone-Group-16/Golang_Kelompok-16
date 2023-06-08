package controllers

import (
	"Capstone/models/payload"
	"Capstone/utils"
	"net/http"

	"github.com/labstack/echo"
)

func UploadImage(c echo.Context) error {

	file, err := c.FormFile("image_url")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To upload Image")
	}

	resp, err := utils.UploadImageCloud(file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To upload Image")
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "image upload success",
		Data:    resp,
	})
}