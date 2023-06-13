package usecase

import (
	"Capstone/constants"
	"Capstone/models"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"errors"
	"fmt"
)

// logic create new warehouse
func CreateWarehouse(req *payload.CreateWarehouseRequest) (resp payload.CreateWarehouseResponse, err error) {

	newWarehouse := models.Warehouse{
		Name:        req.Name,
		City:        req.City,
		Address:     req.Address,
		Capacity:    40,
		Status:      constants.Available,
		Description: req.Description,
		ImageURL:    req.ImageURL,
	}

	err = database.CreateWarehouse(&newWarehouse)
	if err != nil {
		return resp, errors.New("error create warehouse")
	}

	for i := 1; i <= 15; i++ {
		LockerSmall := models.Locker{
			Name:         fmt.Sprintf("S%d", i),
			WarehouseID:  newWarehouse.ID,
			LockerTypeID: 1,
			LockerNumber: uint(i),
			Availability: constants.Available,
		}
		err = database.CreateLocker(&LockerSmall)
		if err != nil {
			return resp, errors.New("error create locker small")
		}
	}

	for i := 16; i <= 30; i++ {
		LockerMedium := models.Locker{
			Name:         fmt.Sprintf("M%d", i),
			WarehouseID:  newWarehouse.ID,
			LockerTypeID: 2,
			LockerNumber: uint(i),
			Availability: constants.Available,
		}
		err = database.CreateLocker(&LockerMedium)
		if err != nil {
			return resp, errors.New("error create locker medium")
		}
	}

	for i := 31; i <= 40; i++ {
		LockerLarge := models.Locker{
			Name:         fmt.Sprintf("L%d", i),
			WarehouseID:  newWarehouse.ID,
			LockerTypeID: 3,
			LockerNumber: uint(i),
			Availability: constants.Available,
		}
		err = database.CreateLocker(&LockerLarge)
		if err != nil {
			return resp, errors.New("error create locker large")
		}
	}

	resp = payload.CreateWarehouseResponse{
		Name:        newWarehouse.Name,
		City:        newWarehouse.City,
		Address:     newWarehouse.Address,
		Capacity:    newWarehouse.Capacity,
		Description: newWarehouse.Description,
		Status:      newWarehouse.Status,
		ImageURL:    newWarehouse.ImageURL,
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
			ID:          warehouse.ID,
			Name:        warehouse.Name,
			City:        warehouse.City,
			Address:     warehouse.Address,
			Favorite:    totalFavorite[i],
			Status:      warehouse.Status,
			Capacity:    warehouse.Capacity,
			Description: warehouse.Description,
			ImageURL:    warehouse.ImageURL,
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
			ID:          warehouse.ID,
			Name:        warehouse.Name,
			City:        warehouse.City,
			Address:     warehouse.Address,
			Favorite:    totalFavorite[i],
			Status:      warehouse.Status,
			Capacity:    warehouse.Capacity,
			Description: warehouse.Description,
			ImageURL:    warehouse.ImageURL,
		})
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
		Name:        warehouse.Name,
		City:        warehouse.City,
		Address:     warehouse.Address,
		Capacity:    warehouse.Capacity,
		Status:      warehouse.Status,
		Description: warehouse.Description,
		ImageURL:    warehouse.ImageURL,
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
