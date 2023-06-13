package database

import (
	"Capstone/config"
	"Capstone/models"
)

func CreateLocker(locker *models.Locker) (err error) {
	err = config.DB.Create(&locker).Error
	if err != nil {
		return err
	}

	return nil
}

func GetLockers() (locker []models.Locker, err error) {
	err = config.DB.Preload("Warehouse").Preload("LockerType").Find(&locker).Error
	if err != nil {
		return nil, err
	}
	return locker, nil
}

func GetLockerSmall(warehouseId uint) (locker []models.Locker, err error) {
	err = config.DB.Preload("Warehouse").Preload("LockerType").Where("locker_type_id = ? AND warehouse_id = ?", 1, warehouseId).Find(&locker).Error
	if err != nil {
		return nil, err
	}
	return locker, nil
}

func GetLockerMedium(warehouseId uint) (locker []models.Locker, err error) {
	err = config.DB.Preload("Warehouse").Preload("LockerType").Where("locker_type_id = ? AND warehouse_id = ?", 2, warehouseId).Find(&locker).Error
	if err != nil {
		return nil, err
	}
	return locker, nil
}

func GetLockerLarge(warehouseId uint) (locker []models.Locker, err error) {
	err = config.DB.Preload("Warehouse").Preload("LockerType").Where("locker_type_id = ? AND warehouse_id = ?", 3, warehouseId).Find(&locker).Error
	if err != nil {
		return nil, err
	}
	return locker, nil
}
