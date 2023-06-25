package database

import (
	"Capstone/config"
	"Capstone/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetTransactions() (transaction []models.Transaction, err error) {
	if err = config.DB.Preload("User").Preload("Locker.Warehouse").Preload("Locker.LockerType").Preload("ItemCategory").Find(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func GetTransactionsPaymentStatus(status string) (transaction *[]models.Transaction, err error) {
	if err = config.DB.Where("payment_status = ?", status).Find(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func GetTransactionByOrderId(orderId string) (transaction *models.Transaction, err error) {
	if err = config.DB.Preload("User").Preload("Locker").Preload("ItemCategory").Where("order_id = ?", orderId).First(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func GetTransactionByUserId(id uint) (transaction []*models.Transaction, err error) {
	if err = config.DB.Preload("User").Preload("Locker.Warehouse").Preload("Locker.LockerType").Preload("ItemCategory").Where("user_id = ?", id).Find(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
}

func CountTransactionByUserId(id uint) (count int64) {
	count = 0
	transaction := []models.Transaction{}
	if err := config.DB.Model(&transaction).Where("user_id = ?", id).Count(&count).Error; err != nil {
		return 0
	}

	return count
}

func CreateTransaction(tx *gorm.DB, transaction *models.Transaction) error {
	db := config.DB
	if tx != nil {
		db = tx
	}

	if err := db.Preload("User").Preload("Locker").Preload("ItemCategory").Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func UpdateTransaction(tx *gorm.DB, transaction *models.Transaction) error {
	db := config.DB
	if tx != nil {
		db = tx
	}

	if err := db.Clauses(clause.Returning{}).Model(transaction).Where("order_id = ?", transaction.OrderID).Updates(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func UpdateTransactionDone(transaction *models.Transaction) error {
	if err := config.DB.Model(&models.Transaction{}).Where("id = ?", transaction.ID).Updates(&models.Transaction{Status: "Done"}).Error; err != nil {
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

func SumTransactionsAmount() (income int, err error) {
	if err := config.DB.Table("transactions").Select("COALESCE(sum(amount),0)").Where("payment_status = ?", "Paid").Row().Scan(&income); err != nil {
		return income, err
	}

	return income, nil
}

func CountTransactionActiveByUserId(id uint) (count int64) {
	count = 0
	transaction := []models.Transaction{}
	if err := config.DB.Model(&transaction).Where("user_id = ? AND status = ?", id, "On Going").Count(&count).Error; err != nil {
		return 0
	}

	return count
}

func SumTransactionsByUserId(id uint) (expenditure uint, err error) {
	if err := config.DB.Table("transactions").Select("COALESCE(sum(amount),0)").Where("payment_status = ? AND user_id = ?", "Paid", id).Row().Scan(&expenditure); err != nil {
		return expenditure, err
	}

	return expenditure, nil
}
