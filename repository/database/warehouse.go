package database

import (
	"Capstone/config"
	"Capstone/models"

	"gorm.io/gorm/clause"
)

// get all warehouse query database
func GetWarehouses(warehouseParam *models.Warehouse) (warehouse []models.Warehouse, err error) {
	db := config.DB

	if warehouseParam.Status != "" {
		db = db.Where("status = ?", warehouseParam.Status)
	}

	if warehouseParam.Location != "" {
		db = db.Where("location = ?", warehouseParam.Location)
	}

	if err := db.Order("capacity desc").Find(&warehouse).Error; err != nil {
		return nil, err
	}

	return warehouse, nil
}

func GetRecomendedWarehouses(warehouseParam *models.Warehouse) (warehouse []models.Warehouse, err error) {
	db := config.DB

	if warehouseParam.Status != "" {
		db = db.Where("status = ?", warehouseParam.Status)
	}

	if warehouseParam.Location != "" {
		db = db.Where("location = ?", warehouseParam.Location)
	}

	if err := db.Order("favorites desc").Find(&warehouse).Error; err != nil {
		return nil, err
	}

	return warehouse, nil
}

// delete warehouse query database
func DeleteWarehouse(warehouse *models.Warehouse) error {
	if err := config.DB.Delete(warehouse).Error; err != nil {
		return err
	}

	return nil
}

// create warehouse query database
func CreateWarehouse(warehouse *models.Warehouse) error {
	if err := config.DB.Clauses(clause.Returning{}).Create(warehouse).Error; err != nil {
		return err
	}

	return nil
}

// update warehouse query database
func UpdateWarehouse(warehouse *models.Warehouse) error {
	if err := config.DB.Model(&warehouse).Save(&warehouse).Error; err != nil {
		return err
	}

	return nil
}

// get warehouse by id query database
func GetWarehouseByID(id uint64) (warehouse *models.Warehouse, err error) {
	if err = config.DB.Where("id = ?", id).First(&warehouse).Error; err != nil {
		return nil, err
	}

	return warehouse, nil
}
