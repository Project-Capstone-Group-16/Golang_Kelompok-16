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

func GetWarehouseByID(id uint) (warehouse *models.Warehouse, err error) {
	warehouse.ID = id
	if err := config.DB.Where("id = ?", id).First(&warehouse).Error; err != nil {
		return nil, err
	}

	return warehouse, nil
}
func DeleteWarehouse(warehouse *models.Warehouse)error {
	if err := config.DB.Delete(warehouse).Error; err != nil {
		return err
	}
	return nil
}