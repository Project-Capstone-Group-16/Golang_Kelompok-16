package database

import (
	"Capstone/config"
	"Capstone/models"
)

func GetAllWarehouses() (warehouse []models.Warehouse, err error) {
	if err := config.DB.Find(&warehouse).Error; err != nil {
		return nil, err
	}
	return warehouse, nil
}

func GetWarehouseByid(id uint) (warehouse *models.Warehouse, err error) {
	warehouse.ID = id
	if err := config.DB.Where("id = ?", id).First(&warehouse).Error; err != nil {
		return nil, err
	}

	return warehouse, nil
}

func DeleteWarehouse(warehouse *models.Warehouse) error {
	if err := config.DB.Delete(warehouse).Error; err != nil {
		return err
	}
	return nil
}

// create warehouse query database
func CreateWarehouse(warehouse *models.Warehouse) error {
	if err := config.DB.Create(warehouse).Error; err != nil {
		return err
	}
	return nil
}

// update warehouse query database
func UpdateWarehouse(warehouse *models.Warehouse) error {
	if err := config.DB.Model(&warehouse).Updates(&warehouse).Error; err != nil {
		return err
	}
	return nil
}

// get all warehouse query database
func GetWarehouses() (warehouses []models.Warehouse, err error) {
	if err = config.DB.Model(&models.Warehouse{}).Find(&warehouses).Error; err != nil {
		return
	}
	return
}

func GetWarehouseByID(id uint64) (warehouse *models.Warehouse, err error) {
	if err = config.DB.Where("id = ?", id).First(&warehouse).Error; err != nil {
		return nil, err
	}
	return warehouse, nil
}
