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

// func GetAllWarehouse() (resp payload.GetAllWarehouseResponse, err error) {
// 	warehouse, err := database.GetAllWarehouses()
// 	if err != nil {
// 		return resp, err
// 	}

//		resp.Warehouses = warehouses
//		return resp,nil
//	}

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

	return resp, nil
}
