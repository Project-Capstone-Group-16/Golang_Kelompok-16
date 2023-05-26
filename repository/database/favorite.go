package database

import (
	"Capstone/config"
	"Capstone/models"
)

func CreateFavorite(user *models.Favorite) error {
	if err := config.DB.Preload("Warehouse").Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func CheckFavorite(favorite *models.Favorite) error {
	if err := config.DB.Where("user_id = ? AND warehouse_id = ?", favorite.UserID, favorite.WarehouseID).First(&favorite).Error; err != nil {
		return err
	}

	return nil
}

func GetAllFavorite() (favorite []models.Favorite, err error) {
	if err = config.DB.Preload("Warehouse").Find(&favorite).Error; err != nil {
		return
	}

	return favorite, nil
}

func GetFavoriteWarehouseByID(id int) (favorite *models.Favorite, err error) {
	if err = config.DB.Preload("Warehouse").Where("id = ?", id).First(&favorite).Error; err != nil {
		return
	}

	return favorite, nil
}

func DeleteFavorite(favorite *models.Favorite) error {
	if err := config.DB.Delete(&favorite).Error; err != nil {
		return err
	}

	return nil
}
