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
		return resp, err
	}

	StartDate, err := time.Parse("02/01/2006", req.StartDate)
	if err != nil {
		return
	}

	EndDate, err := time.Parse("02/01/2006", req.EndDate)
	if err != nil {
		return
	}

	locker, err := database.GetLockerByStatus(req.LockerID)
	if err != nil {
		return resp, err
	}

	itemCategory, err := database.GetItemCategoryById(req.ItemCategoryID)
	if err != nil {
		return resp, err
	}

	countDate := EndDate.Sub(StartDate)

	newTransaction := models.Transaction{
		UserID:         user.ID,
		LockerID:       locker.ID,
		ItemCategoryID: req.ItemCategoryID,
		Amount:         uint(countDate.Hours()/24) * locker.LockerType.Price,
		StartDate:      StartDate,
		EndDate:        EndDate,
		PaymentStatus:  "Unpaid",
	}

	paymentURL, err := utils.GetPaymentURL(newTransaction, user)
	if err != nil {
		return newTransaction, err
	}

	if user.FirstName == "" || user.LastName == "" || user.PhoneNumber == "" || user.Address == "" || user.Gender == "" || user.BirthDate == nil {
		return resp, errors.New("Please complete your profile first")
	}

	resp = models.Transaction{
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

	return
}
