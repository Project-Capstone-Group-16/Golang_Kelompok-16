package database

import (
	"Capstone/config"
	"Capstone/models"
)

func CreateLocker(locker *models.Locker) (err error) {
	err = config.DB.Create(&locker).Error
	return
}
