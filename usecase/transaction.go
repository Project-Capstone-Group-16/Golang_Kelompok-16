package usecase

import (
	"Capstone/models"
	"Capstone/models/payload"
	"Capstone/repository/database"
	"Capstone/utils"
	"errors"
	"time"
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

	warehouse, err := database.GetWarehouseByID(uint64(req.WarehouseID))
	if err != nil {
		return resp, errors.New("Warehouse not found")
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

	orderId := "TRX-" + time.Now().Format("20060102150405")

	newTransaction := models.Transaction{
		OrderID:        orderId,
		UserID:         user.ID,
		LockerID:       locker.ID,
		ItemCategoryID: itemCategory.ID,
		Amount:         uint(countDate.Hours()/24) * lockerType.Price,
		StartDate:      StartDate,
		EndDate:        EndDate,
		PaymentStatus:  "Unpaid",
	}

	paymentURL, err := utils.GetPaymentURL(&newTransaction, user)
	if err != nil {
		return newTransaction, err
	}

	if user.FirstName == "" || user.LastName == "" || user.PhoneNumber == "" || user.Address == "" || user.Gender == "" || user.BirthDate == nil {
		return resp, errors.New("Please complete your profile first")
	}

	resp = models.Transaction{
		OrderID:        newTransaction.OrderID,
		UserID:         newTransaction.UserID,
		User:           *user,
		LockerID:       newTransaction.LockerID,
		Locker:         *locker,
		ItemCategoryID: newTransaction.ItemCategoryID,
		ItemCategory:   *itemCategory,
		Amount:         newTransaction.Amount,
		StartDate:      newTransaction.StartDate,
		EndDate:        newTransaction.EndDate,
		PaymentStatus:  newTransaction.PaymentStatus,
		PaymentUrl:     paymentURL,
	}

	err = database.CreateTransaction(&resp)
	if err != nil {
		return resp, err
	}

	warehouse.Capacity -= 1

	err = database.UpdateWarehouse(warehouse)
	if err != nil {
		return resp, errors.New("Failed to update warehouse capacity")
	}

	locker.Availability = "Not Available"

	err = database.UpdateLockerStatus(locker)
	if err != nil {
		return resp, errors.New("Failed to update locker status")
	}

	return
}

func GetTransactionsByUserId(id int) (resp []*models.Transaction, err error) {
	resp, err = database.GetTransactionByUserId(uint(id))
	if err != nil {
		return resp, err
	}

	return
}
