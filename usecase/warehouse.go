package usecase

import (
	"Capstone/models/payload"
	"Capstone/repository/database"
)

func DeleteWarehouse(req *payload.DeleteWarehouseRequest) error {

	warehouse, err := database.GetWarehouseByID(req.WarehouseID)
	if err != nil {
		return err
	}

	err = database.DeleteWarehouse(warehouse)
	if err != nil {
		return err
	}
	return nil
}

func GetAllWarehouse() (resp []payload.GetAllWarehouseResponse, err error) {
	warehouses, err := database.GetAllWarehouses()
	if err != nil {
		return resp, err
	}

	resp = make([]payload.GetAllWarehouseResponse, len(warehouses))
	for i, warehouse := range warehouses {
		resp[i] = payload.GetAllWarehouseResponse{
			Name:     warehouse.Name,
			Location: warehouse.Location,
			Status:   warehouse.Status,
		}
	}
// logic create new warehouse
func CreateWarehouse(req *payload.CreateWarehouseRequest) (resp payload.CreateWarehouseResponse, err error) {

	newWarehouse := &models.Warehouse{
		Name:     req.Name,
		Location: req.Location,
		Status:   constants.Available,
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
		Name:     warehouse.Name,
		Location: warehouse.Location,
		Status:   warehouse.Status,
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
