package usecase

import (
	"Capstone/models"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"errors"
)

// get favorite by user id
func GetFavoriteUserByID(userId int) (favorite []models.Favorite, err error) {

	favorite, err = database.GetFavoriteUserByID(userId)
	if err != nil {
		return nil, errors.New("Failed to get favorite")
	}

	return favorite, nil
}

// Add Favorite Warehouse
func CreateFavoriteWarehouse(id int, req *payload.CreateFavoriteRequest) (resp any, err error) {
	user, err := database.GetuserByID(id)
	if err != nil {
		return resp, errors.New("User not found")
	}

	warehouse, err := database.GetWarehouseByID(uint64(req.WarehouseID))
	if err != nil {
		return resp, errors.New("Warehouse not found")
	}

	newFavorite := &models.Favorite{
		UserID:      user.ID,
		WarehouseID: req.WarehouseID,
	}

	favorite, err := database.CheckFavorite(newFavorite)
	if err != nil {
		err = database.CreateFavorite(newFavorite)
		if err != nil {
			return resp, errors.New("Can't Create Favorite")
		}
		warehouse.Favorites += 1
		err = database.UpdateWarehouse(nil, warehouse)
		if err != nil {
			return resp, errors.New("Can't Update Warehouse")
		}

		resp = payload.CreateFavoriteResponse{
			WarehouseID: newFavorite.WarehouseID,
			Warehouse: payload.GetAllWarehouseResponse{
				ID:       warehouse.ID,
				Name:     warehouse.Name,
				City:     warehouse.City,
				Address:  warehouse.Address,
				Capacity: warehouse.Capacity,
				Favorite: warehouse.Favorites,
				Status:   warehouse.Status,
				ImageURL: warehouse.ImageURL,
			},
		}
	} else {
		err = database.DeleteFavorite(favorite)
		if err != nil {
			return resp, errors.New("Can't Delete Favorite")
		}

		warehouse.Favorites -= 1
		err = database.UpdateWarehouse(nil, warehouse)
		if err != nil {
			return resp, errors.New("Can't Update Warehouse")
		}

		resp = "Success Delete Favorite"

		return
	}

	return
}
