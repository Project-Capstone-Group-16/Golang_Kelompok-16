package controllers

import (
	"Capstone/models/payload"
	"Capstone/utils"
	"net/http"

	"github.com/labstack/echo"
)

func UploadImage(c echo.Context) error {
	payloadimage := payload.UploadImageCloudinaryRequest{}
	c.Bind(payloadimage)
	ctx := c.Request().Context()
	resp, err := utils.UploadImageCloud(payloadimage.Image, ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To upload Image")
	}

	return c.JSON(http.StatusOK, payload.Response{
		Message: "image upload success",
		Data:    resp,
	})
}
