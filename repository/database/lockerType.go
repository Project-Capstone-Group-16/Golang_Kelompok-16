package database

import (
	"Capstone/config"
	"Capstone/models"
)

func GetLockerTypes() (lockerType []models.LockerType, err error) {
	err = config.DB.Find(&lockerType).Error
	if err != nil {
		return nil, err
	}
	return lockerType, nil
}

func GetLockerTypeById(id uint64) (lockerType *models.LockerType, err error) {
	err = config.DB.Where("id = ?", id).First(&lockerType).Error
	if err != nil {
		return nil, err
	}

	return lockerType, nil
}
