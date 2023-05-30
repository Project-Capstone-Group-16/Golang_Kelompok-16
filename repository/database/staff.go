package database

import (
	"Capstone/config"
	"Capstone/models"
)

func CreateStaff(staff *models.Staff) error {
	if err := config.DB.Create(staff).Error; err != nil {
		return err
	}
	return nil
}

func UpdateStaff(staff *models.Staff) error {
	if err := config.DB.Updates(&staff).Error; err != nil {
		return err
	}
	return nil
}

func GetStaffByID(id uint64) (staff *models.Staff, err error) {
	if err = config.DB.Where("id = ?", id).First(&staff).Error; err != nil {
		return nil, err
	}

	return staff, nil
}
