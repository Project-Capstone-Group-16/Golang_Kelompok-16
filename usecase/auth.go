package usecase

import (
	"Capstone/middleware"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(req *payload.LoginUserRequest) (res payload.LoginUserResponse, err error) {

	user, err := database.GetuserByEmail(req.Email)
	if err != nil {
		return res, echo.NewHTTPError(http.StatusBadRequest, "Email not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return res, errors.New("Wrong Password")
	}

	token, err := middleware.CreateToken(int(user.ID))
	if err != nil {
		return res, echo.NewHTTPError(http.StatusBadRequest, "Failed to generate token")
	}

	user.Token = token

	res = payload.LoginUserResponse{
		Email: user.Email,
		Token: user.Token,
	}

	return
}
