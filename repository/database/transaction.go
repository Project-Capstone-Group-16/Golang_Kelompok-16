package database

import (
	"Capstone/config"
	"Capstone/models"

	"gorm.io/gorm/clause"
)

func GetTransactions() (transaction []models.Transaction, err error) {
	if err = config.DB.Preload("User").Preload("Locker").Preload("ItemCategory").Find(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
} // new

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
} // new

func GetTransactionByUserId(id uint) (transaction []*models.Transaction, err error) {
	if err = config.DB.Preload("User").Preload("Locker.Warehouse").Preload("Locker.LockerType").Preload("ItemCategory").Where("user_id = ?", id).Find(&transaction).Error; err != nil {
		return transaction, err
	}

	return transaction, nil
} // new

func CountTransactionByUserId(id uint) (count int64) {
	count = 0
	transaction := []models.Transaction{}
	if err := config.DB.Model(&transaction).Where("user_id = ?", id).Count(&count).Error; err != nil {
		return 0
	}

	return count
}

func CreateTransaction(transaction *models.Transaction) error {
	if err := config.DB.Preload("User").Preload("Locker").Preload("ItemCategory").Create(&transaction).Error; err != nil {
		return err
	}

	return nil
} // new

func UpdateTransaction(transaction *models.Transaction) error {
	if err := config.DB.Clauses(clause.Returning{}).Model(transaction).Where("order_id = ?", transaction.OrderID).Updates(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func UpdateTransactionDone(transaction *models.Transaction) error {
	if err := config.DB.Clauses(clause.Returning{}).Exec("UPDATE transactions SET status = 'Done' WHERE end_date < NOW() AND status = 'On Going'").Error; err != nil {
		return err
	}

	return nil
} // new

func DeleteTransaction(transaction *models.Transaction) error {
	if err := config.DB.Delete(&transaction).Error; err != nil {
		return err
	}

	return nil
} // new

func SumTransactionsAmount() (income int, err error) {
	if err := config.DB.Table("transactions").Select("sum(amount)").Where("payment_status = ?", "Paid").Row().Scan(&income); err != nil {
		return income, err
	}

	return income, nil
}
