package usecase

import (
	"Capstone/models"
	"Capstone/repository/database"
)

func GetAllItemCategorys() ([]models.ItemCategory, error) {
	itemCategorys, err := database.GetItemCategorys()
	if err != nil {
		return nil, err
	}
	return itemCategorys, nil
} //new
