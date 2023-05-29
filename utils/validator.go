package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
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
			// Customize  pesan error menjadi pesan yang lebih informatif
			switch fieldError.Tag() {
			case "required":
				errorMessage += fmt.Sprintf(`Field %s is required`+"\n", fieldError.Field())
			case "email":
				errorMessage += fmt.Sprintf("Field %s must be a valid email address"+"\n", fieldError.Field())
			case "min":
				errorMessage += fmt.Sprintf("Field %s minimum 6 character"+"\n", fieldError.Field())
			case "number":
				errorMessage += fmt.Sprintf("Field %s must be number"+"\n", fieldError.Field())
			case "len":
				errorMessage += fmt.Sprintf("Field %s must be 11 length"+"\n", fieldError.Field())
			default:
				errorMessage += fmt.Sprintf("Field %s is invalid"+"\n", fieldError.Field())
			}
		}
		return errors.New(errorMessage)
	}

	return nil
}
