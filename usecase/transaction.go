package usecase

import (
	"Capstone/config"
	"Capstone/models"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"Capstone/utils"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func CreateTransaction(id int, req *payload.CreateTransactionRequest) (resp models.Transaction, err error) {
	user, err := database.GetuserByID(id)
	if err != nil {
		return resp, errors.New("user not found")
	}

	StartDate, err := time.Parse("02/01/2006", req.StartDate)
	if err != nil {
		return resp, errors.New("Failed to parse start date")
	}

	EndDate, err := time.Parse("02/01/2006", req.EndDate)
	if err != nil {
		return resp, errors.New("Failed to parse start date")
	}

	if StartDate.Before(time.Now().AddDate(0, 0, -1)) {
		return resp, errors.New("Start date must be after today")
	}

	if StartDate == EndDate {
		return resp, errors.New("Start date cannot be the same as end date")
	}

	if StartDate.After(EndDate) {
		return resp, errors.New("Start date must be before end date")
	}

	warehouse, err := database.GetWarehouseByID(uint64(req.WarehouseID))
	if err != nil {
		return resp, errors.New("Warehouse not found")
	}

	if warehouse.Capacity == 0 {
		return resp, errors.New("Warehouse is full")
	}

	lockerType, err := database.GetLockerTypeById(uint64(req.LockerTypeID))
	if err != nil {
		return resp, errors.New("Locker type not found")
	}

	locker, err := database.GetLockerByStatus(warehouse.ID, lockerType.ID)
	if err != nil {
		return resp, errors.New("Locker not found")
	}

	itemCategory, err := database.GetItemCategoryById(req.ItemCategoryID)
	if err != nil {
		return resp, errors.New("Item category not found")
	}

	countDate := EndDate.Sub(StartDate)

	uuid := uuid.New()

	if user.FirstName == "" || user.LastName == "" || user.PhoneNumber == "" || user.Address == "" || user.Gender == "" || user.BirthDate == nil {
		return resp, errors.New("Please complete your profile first")
	}

	newTransaction := models.Transaction{
		OrderID:        "TRX-" + uuid.String(),
		UserID:         user.ID,
		User:           *user,
		LockerID:       locker.ID,
		Locker:         *locker,
		ItemCategoryID: itemCategory.ID,
		ItemCategory:   *itemCategory,
		Amount:         uint(countDate.Hours()/24) * lockerType.Price,
		StartDate:      StartDate,
		EndDate:        EndDate,
		Status:         "Waiting for Payment",
		PaymentStatus:  "Unpaid",
	}

	tx := config.DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	err = database.CreateTransaction(tx, &newTransaction)
	if err != nil {
		return resp, err
	}

	responseMidtrans, err := utils.GetPaymentURL(&newTransaction, user)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentUrl = responseMidtrans.RedirectURL

	err = database.UpdateTransaction(tx, &newTransaction)
	if err != nil {
		fmt.Println("Failed to update transaction")
		return
	}

	err = tx.Commit().Error
	if err != nil {
		errors.New("Failed to commit transaction")
		return
	}

	return newTransaction, nil
}

func GetTransactionsByUserId(id int) (resp []*models.Transaction, err error) {
	resp, err = database.GetTransactionByUserId(uint(id))
	if err != nil {
		return resp, err
	}

	return
}

func GetAllTransactions() ([]models.Transaction, error) {
	transactions, err := database.GetTransactions()
	if err != nil {
		return nil, err
	}
	return transactions, nil
} //new

func ProcessPayemnt(req *payload.TransactionNotificationInput) error {
	transaction, err := database.GetTransactionByOrderId(req.OrderID)
	if err != nil {
		fmt.Println("Failed to get transactions with unpaid payment status")
		return err
	}

	locker, err := database.GetLockerById(transaction.LockerID)
	if err != nil {
		return errors.New("Locker not found")
	}

	warehouse, err := database.GetWarehouseByID(uint64(locker.WarehouseID))
	if err != nil {
		return errors.New("Warehouse not found")
	}

	transaction.PaymentMethod = req.PaymentType

	if req.TransactionStatus == "settlement" || req.TransactionStatus == "capture" {
		transaction.PaymentStatus = "Paid"
		transaction.Status = "On Going"

		tx := config.DB.Begin()
		defer func() {
			if err != nil {
				tx.Rollback()
			}
		}()

		date, _ := time.Parse("2006-01-02 15:04:05", req.TransactionTime)

		transaction.PaymentDate = &date
		err = database.UpdateTransaction(tx, transaction)
		if err != nil {
			fmt.Println("Failed to update transaction")
			return err
		}

		warehouse.Capacity = warehouse.Capacity - 1
		err = database.UpdateWarehouse(tx, warehouse)
		if err != nil {
			return errors.New("Failed to update warehouse capacity")
		}

		if warehouse.Capacity == 0 {
			warehouse.Status = "Not Available"
			err = database.UpdateWarehouse(tx, warehouse)
			if err != nil {
				return errors.New("Failed to update warehouse status")
			}
		}

		locker.Availability = "Not Available"
		err = database.UpdateLockerStatus(tx, locker)
		if err != nil {
			return errors.New("Failed to update locker status")
		}

		err = tx.Commit().Error
		if err != nil {
			return errors.New("Failed to commit transaction")
		}
	} else if req.TransactionStatus != "pending" {
		transaction.PaymentStatus = "Canceled"
		transaction.Status = "Canceled" // new
		err = database.UpdateTransaction(nil, transaction)
		if err != nil {
			fmt.Println("Failed to update transaction")
			return err
		}
	}

	return nil

}

// func UpdateStatusDone() {
// 	transaction, err := database.GetTransactions()
// 	if err != nil {
// 		fmt.Println("Failed to get transactions")
// 		return
// 	}

// 	for _, v := range transaction {
// 		if v.Status == "On Going" && v.EndDate.After(time.Now()) {
// 			v.Status = "Done"
// 			err = database.UpdateTransactionDone(&v)
// 			if err != nil {
// 				fmt.Println("Failed to update transaction")
// 				return
// 			}

// 			locker, err := database.GetLockerById(v.LockerID)
// 			if err != nil {
// 				fmt.Println("Failed to get locker")
// 				return
// 			}

// 			locker.Availability = "Available"
// 			err = database.UpdateLockerStatus(locker)
// 			if err != nil {
// 				fmt.Println("Failed to update locker status")
// 				return
// 			}

// 			warehouse, err := database.GetWarehouseByID(uint64(locker.WarehouseID))
// 			if err != nil {
// 				fmt.Println("Failed to get warehouse")
// 				return
// 			}

// 			warehouse.Capacity += 1
// 			err = database.UpdateWarehouse(warehouse)
// 			if err != nil {
// 				fmt.Println("Failed to update warehouse capacity")
// 				return
// 			}
// 		}
// 	}
// } // new
