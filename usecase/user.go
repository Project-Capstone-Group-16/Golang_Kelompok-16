package usecase

import (
	"Capstone/constants"
	"Capstone/models"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"errors"
	"time"

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
	if !database.IsEmailAvailable(req.Email) {
		return resp, errors.New("email is already registered")
	}

	if req.ConfirmPassword != req.Password {
		return resp, errors.New("Password not match")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	newUser := &models.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: "0" + req.PhoneNumber, // masih bimbang apakah +62 atau 0
		Password:    string(passwordHash),
		Role:        constants.Admin,
	}

	err = database.CreateUser(newUser)
	if err != nil {
		return
	}

	resp = payload.CreateAdminResponse{
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
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
		return errors.New("user not found")
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

// Logic get user by id
func GetUser(id int) (user *models.User, err error) {
	user, err = database.GetuserByID(id)
	if err != nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}

// Logic Update Profile User
func UpdateProfile(user *models.User, req *payload.UpdateProfileUser) (res payload.UpdateProfileUserResponse, err error) {
	birthDate, err := time.Parse("02/01/2006", req.BirthDate)
	if err != nil {
		return
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.BirthDate = &birthDate
	user.Gender = req.Gender
	user.PhoneNumber = "0" + req.PhoneNumber
	user.Address = req.Address
	user.ImageUrl = req.ImageURL

	err = database.UpdateUser(user)
	if err != nil {
		return res, errors.New("Can't update profile user")
	}

	res = payload.UpdateProfileUserResponse{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		BirthDate:   user.BirthDate,
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		ImageURL:    user.ImageUrl,
	}

	return
}

// Logic Get All Users
func GetUsers() (users []models.User, err error) {
	users, err = database.GetUsers()
	if err != nil {
		return nil, errors.New("Error getting users")
	}

	return users, nil
}
