package usecase

import (
	"Capstone/middleware"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"errors"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(req *payload.LoginUserRequest) (res payload.LoginUserResponse, err error) {

	user, err := database.GetuserByEmail(req.Email)
	if err != nil {
		echo.NewHTTPError(400, "Email not registered")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return res, errors.New("Wrong Password")
	}

	token, err := middleware.CreateToken(int(user.ID))
	if err != nil {
		echo.NewHTTPError(400, "Failed to generate token")
		return
	}

	user.Token = token

	res = payload.LoginUserResponse{
		Email: user.Email,
		Token: user.Token,
	}

	return
}
