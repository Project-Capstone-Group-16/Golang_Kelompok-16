package usecase

import (
	"Capstone/models"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(req *payload.CreateUserRequest) (resp payload.CreateUserResponse, err error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	if !database.IsEmailAvailable(req.Email) {
		return resp, errors.New("email is already registered")
	}

	newUser := &models.User{
		Email:    req.Email,
		Password: string(passwordHash),
	}

	err = database.CreateUser(newUser)
	if err != nil {
		return
	}

	resp = payload.CreateUserResponse{
		Email:    newUser.Email,
		Password: newUser.Password,
	}

	return
}
