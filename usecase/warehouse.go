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

	newWarehouse := models.Warehouse{
		Name:     req.Name,
		City:     req.City,
		Province: req.Province,
		Status:   constants.Available,
		ImageURL: req.WarehouseImage,
	}

	err = database.CreateWarehouse(&newWarehouse)
	if err != nil {
		return
	}

	resp = payload.CreateWarehouseResponse{
		Name:     newWarehouse.Name,
		City:     newWarehouse.City,
		Province: newWarehouse.Province,
		Status:   newWarehouse.Status,
		ImageURL: newWarehouse.ImageURL,
	}

	return
}

// Logic Delete Warahouse
func DeleteWarehouse(warehouses *models.Warehouse) error {

	err := database.DeleteWarehouse(warehouses)
	if err != nil {
		return err
	}
	return nil
}

// Logic Get All Warehouse

// logic by status warehouse
func GetWarehouses(warehouse *models.Warehouse) (resp []payload.GetAllWarehouseResponse, err error) {
	warehouses, err := database.GetWarehouses(warehouse)
	if err != nil {
		return resp, err
	}

	var totalFavorite []int
	for _, v := range warehouses {
		warehouse_id := v.ID
		totalCount := database.CountFavoriteByWarehouseId(warehouse_id)
		totalFavorite = append(totalFavorite, int(totalCount))
	}

	resp = []payload.GetAllWarehouseResponse{}
	for i, warehouse := range warehouses {
		resp = append(resp, payload.GetAllWarehouseResponse{
			ID:       warehouse.ID,
			Name:     warehouse.Name,
			City:     warehouse.City,
			Province: warehouse.Province,
			Favorite: totalFavorite[i],
			Status:   warehouse.Status,
			Capacity: warehouse.Capacity,
			ImageURL: warehouse.ImageURL,
		})
	}

	return
}

func GetRecomendedWarehouse(warehouse *models.Warehouse) (resp []payload.GetAllWarehouseResponse, err error) {
	warehouses, err := database.GetRecomendedWarehouses(warehouse)
	if err != nil {
		return resp, err
	}

	var totalFavorite []int
	for _, v := range warehouses {
		warehouse_id := v.ID
		totalCount := database.CountFavoriteByWarehouseId(warehouse_id)
		totalFavorite = append(totalFavorite, int(totalCount))
	}

	resp = []payload.GetAllWarehouseResponse{}
	for i, warehouse := range warehouses {
		resp = append(resp, payload.GetAllWarehouseResponse{
			ID:       warehouse.ID,
			Name:     warehouse.Name,
			City:     warehouse.City,
			Province: warehouse.Province,
			Favorite: totalFavorite[i],
			Status:   warehouse.Status,
			Capacity: warehouse.Capacity,
			ImageURL: warehouse.ImageURL,
		})
	}

	return
}

// logic update warehouse
func UpdateWarehouse(warehouse *models.Warehouse) (resp payload.UpdateWarehouseResponse, err error) {

	// warehouse.Name = req.Name
	// warehouse.City = req.City
	// warehouse.Province = req.Province
	// warehouse.Status = req.Status
	// warehouse.ImageURL = req.WarehouseImage

	err = database.UpdateWarehouse(warehouse)
	if err != nil {
		return resp, errors.New("Can't update warehouse")
	}

	resp = payload.UpdateWarehouseResponse{
		Name:     warehouse.Name,
		City:     warehouse.City,
		Province: warehouse.Province,
		Status:   warehouse.Status,
		ImageURL: warehouse.ImageURL,
	}

	return resp, nil
}

// Get Warehouse By Id
func GetWarehouseByID(id uint64) (warehouse *models.Warehouse, err error) {
	warehouse, err = database.GetWarehouseByID(id)
	if err != nil {
		return warehouse, errors.New("Warehouse not found")
	}

	return warehouse, nil
}
