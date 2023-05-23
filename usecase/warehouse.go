package usecase

import (
	"Capstone/constants"
	"Capstone/models"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"errors"
)

// logic create new warehouse
func CreateWarehouse(req *payload.CreateWarehouseRequest) (resp payload.CreateWarehouseResponse, err error) {

	newWarehouse := &models.Warehouse{
		Name:     req.Name,
		Location: req.Location,
		Status:   constants.NotAvailable,
	}

	err = database.CreateWarehouse(newWarehouse)
	if err != nil {
		return
	}

	resp = payload.CreateWarehouseResponse{
		Name:     newWarehouse.Name,
		Location: newWarehouse.Location,
		Status:   newWarehouse.Status,
	}

	return
}

// logic update warehouse
func UpdateWarehouse(warehouse *models.Warehouse) (resp payload.UpdateWarehouseResponse, err error) {

	err = database.UpdateWarehouse(warehouse)

	if err != nil {
		return resp, errors.New("Can't update warehouse")
	}
	resp = payload.UpdateWarehouseResponse{
		Name: warehouse.Name,
		Location: warehouse.Location,
		Status: warehouse.Status,	
	}
	
	return resp, nil
}

func GetWarehouseByID(id uint64) (warehouse *models.Warehouse, err error) {
	warehouse, err = database.GetWarehouseByID(id)
	if err != nil {
		return warehouse, errors.New("Warehouse not found")
	}
	return warehouse, nil
}