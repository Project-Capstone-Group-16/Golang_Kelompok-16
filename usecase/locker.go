package usecase

import (
	"Capstone/models"
	"Capstone/repository/database"
)

func GetLockers() (locker []models.Locker, err error) {
	locker, err = database.GetLockers()
	if err != nil {
		return nil, err
	}
	return locker, nil
}

func GetLockerSmall(warehouseId uint) (locker []models.Locker, err error) {
	locker, err = database.GetLockerSmall(warehouseId)
	if err != nil {
		return nil, err
	}
	return locker, nil
}

func GetLockerMedium(warehouseId uint) (locker []models.Locker, err error) {
	locker, err = database.GetLockerMedium(warehouseId)
	if err != nil {
		return nil, err
	}
	return locker, nil
}

func GetLockerLarge(warehouseId uint) (locker []models.Locker, err error) {
	locker, err = database.GetLockerLarge(warehouseId)
	if err != nil {
		return nil, err
	}
	return locker, nil
}

