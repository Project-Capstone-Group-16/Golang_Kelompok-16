package middleware

import (
	"Capstone/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningMethod: "HS256",
	SigningKey:    []byte(constants.SECRET_JWT),
})

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	byteSecret := []byte(constants.SECRET_JWT)
	return token.SignedString(byteSecret)
}

func IsUser(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["userId"].(float64))

	return userId, nil
}
func IsAdmin(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return 0, echo.NewHTTPError(401, "Unauthorized")
	}
	claims := user.Claims.(jwt.MapClaims)
	userId := int(claims["user_id"].(float64))

	return userId, nil
}
