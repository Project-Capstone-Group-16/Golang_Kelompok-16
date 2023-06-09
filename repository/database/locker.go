package database

import (
	"Capstone/config"
	"Capstone/models"

	"gorm.io/gorm"
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

func GetLockerByStatus(idWarehouse, idLockerType uint) (locker *models.Locker, err error) {
	err = config.DB.Preload("Warehouse").Preload("LockerType").Where("warehouse_id = ? AND locker_type_id = ? AND Availability = ?", idWarehouse, idLockerType, "Available").First(&locker).Error
	if err != nil {
		return nil, err
	}

	return locker, nil
}

func GetLockerById(lockerId uint) (locker *models.Locker, err error) {
	err = config.DB.Preload("Warehouse").Preload("LockerType").Where("id = ?", lockerId).First(&locker).Error
	if err != nil {
		return nil, err
	}

	return locker, nil
} // new

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

func UpdateLockerStatus(tx *gorm.DB, locker *models.Locker) error {
	db := config.DB
	if tx != nil {
		db = tx
	}

	err := db.Updates(&locker).Error
	if err != nil {
		return err
	}

	return nil
}

func CountAllLockers() (countedLockers int64, err error) {
	lockers := []models.Locker{}
	if err := config.DB.Model(&lockers).Count(&countedLockers).Error; err != nil {
		return 0, err
	}
	return
}

func CountUsedLockers() (countUsed int64, err error) {
	lockers := []models.Locker{}
	if err := config.DB.Model(&lockers).Where("Availability = ?", "Not Available").Count(&countUsed).Error; err != nil {
		return 0, err
	}
	return
}
