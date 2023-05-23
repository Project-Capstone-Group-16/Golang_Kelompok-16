package usecase

import (
	"Capstone/constants"
	"Capstone/models"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Logic Create User
func CreateUser(req *payload.CreateUserRequest) (resp payload.CreateUserResponse, err error) {
	if req.ConfirmPassword != req.Password {
		return resp, errors.New("Password not match")
	}

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

// Logic Create Admin
func CreateAdmin(req *payload.CreateAdminRequest) (resp payload.CreateAdminResponse, err error) {
	if req.ConfirmPassword != req.Password {
		return resp, errors.New("Password not match")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	if !database.IsEmailAvailable(req.Email) {
		return resp, errors.New("email is already registered")
	}

	newUser := &models.User{
		Fullname:    req.Fullname,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    string(passwordHash),
		Role:        constants.Admin,
	}

	err = database.CreateUser(newUser)
	if err != nil {
		return
	}

	resp = payload.CreateAdminResponse{
		Fullname:    newUser.Fullname,
		Email:       newUser.Email,
		PhoneNumber: newUser.PhoneNumber,
		Password:    newUser.Password,
	}

	return
}

// Logic Update Password User
func UpdatePassword(id int, req *payload.UpdatePasswordRequest) error {

	user, err := database.GetuserByID(id)
	if err != nil {
		return err
	}

	if req.ConfirmPassword != req.Password {
		return errors.New("Password not match")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(passwordHash)

	err = database.UpdateUser(user)
	if err != nil {
		return errors.New("Can't update password")
	}
	return nil
}

