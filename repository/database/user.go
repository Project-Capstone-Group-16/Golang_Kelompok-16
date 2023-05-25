package database

import (
	"Capstone/config"
	"Capstone/models"
	"net/http"

	"github.com/labstack/echo"
)

func CreateUser(user *models.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func IsEmailAvailable(email string) bool {
	var count int64
	user := models.User{}
	if err := config.DB.Model(&user).Where("email = ?", email).Count(&count).Error; err != nil {
		echo.NewHTTPError(http.StatusNotFound, err.Error())
		return false
	}
	return count == 0
}

func GetuserByEmail(email string) (user models.User, err error) {
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetUsers() (users []models.User, err error) {
	if err = config.DB.Model(&models.User{}).Find(&users).Error; err != nil {
		return
	}
	return
}

// update user query database
func UpdateUser(user *models.User) error {
	if err := config.DB.Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// delete user query database
func DeleteUser(user *models.User) error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

// login user query database
func LoginUser(user *models.User) error {
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}

// get user by id query database
func GetuserByID(id int) (user *models.User, err error) {
	if err := config.DB.Preload("Favorite").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
