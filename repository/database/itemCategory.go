package database

import (
	"Capstone/config"
	"Capstone/models"
)

func GetItemCategorys() (ItemCategory []models.ItemCategory, err error) {
	if err = config.DB.Find(&ItemCategory).Error; err != nil {
		return nil, err
	}

	return ItemCategory, nil
} // new

func GetItemCategoryById(id uint) (ItemCategory *models.ItemCategory, err error) {
	if err = config.DB.Where("id = ?", id).First(&ItemCategory).Error; err != nil {
		return ItemCategory, err
	}

	return ItemCategory, nil
} // new
