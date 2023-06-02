package controllers

import (
	"Capstone/models/payload"
	"Capstone/usecase"
	"net/http"

	"github.com/labstack/echo"
)

func UploadImage(c echo.Context) error {
	PayloadImage := payload.UploadImage{}
	c.Bind(&PayloadImage)
	ctx := c.Request().Context()
	resp, err := usecase.UploadImageCloud(PayloadImage.Image, ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To upload Image")
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "success register user",
		Data:    resp,
	})
}
