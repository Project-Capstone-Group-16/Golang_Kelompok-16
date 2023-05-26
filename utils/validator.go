package utils

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			// Jika terjadi kesalahan validasi internal
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}

		validationErrors := err.(validator.ValidationErrors)
		errorMessage := ""
		for _, fieldError := range validationErrors {
			// Ubah pesan validasi default menjadi pesan yang lebih informatif
			switch fieldError.Tag() {
			case "required":
				errorMessage += fmt.Sprintf(`Field %s is required`+"\n", fieldError.Field())
			case "email":
				errorMessage += fmt.Sprintf("Field %s must be a valid email address"+"\n", fieldError.Field())
			// Tambahkan penanganan pesan validasi lainnya sesuai kebutuhan
			case "min":
				errorMessage += fmt.Sprintf("Field %s must be 6 character"+"\n", fieldError.Field())
			default:
				errorMessage += fmt.Sprintf("Field %s is invalid"+"\n", fieldError.Field())
			}
		}
		return echo.NewHTTPError(http.StatusBadRequest, errorMessage)
	}

	return nil
}
