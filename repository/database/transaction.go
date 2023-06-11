package database

import (
	"Capstone/config"
	"Capstone/models"
)

func GetTransactions() (transaction []models.Transaction, err error) {
	if err = config.DB.Preload("User").Preload("Locker").Preload("ItemCategory").Find(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func GetTransactionById(id uint64) (transaction *models.Transaction, err error) {
	if err = config.DB.Preload("User").Preload("Locker").Preload("ItemCategory").Where("id = ?", id).First(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func CreateTransaction(transaction *models.Transaction) error {
	if err := config.DB.Preload("User").Preload("Locker").Preload("ItemCategory").Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func UpdateTransaction(transaction *models.Transaction) error {
	if err := config.DB.Updates(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTransaction(transaction *models.Transaction) error {
	if err := config.DB.Delete(&transaction).Error; err != nil {
		return err
	}

	return nil
}
