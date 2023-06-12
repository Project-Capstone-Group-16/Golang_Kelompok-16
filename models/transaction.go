package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	OrderID        string `json:"order_id" form:"order_id"`
	UserID         uint   `json:"user_id" form:"user_id"`
	User           User
	LockerID       uint `json:"locker_id" form:"locker_id"`
	Locker         Locker
	ItemCategoryID uint `json:"item_category_id" form:"item_category_id"`
	ItemCategory   ItemCategory
	Amount         uint      `json:"amount" form:"amount"`
	StartDate      time.Time `json:"start_date" form:"start_date"`
	EndDate        time.Time `json:"end_date" form:"end_date"`
	Status         string    `json:"status" form:"status"`
	PaymentMethod  string    `json:"payment_method" form:"payment_method"`
	PaymentDate    time.Time `json:"payment_date" form:"payment_date"`
	PaymentStatus  string    `json:"payment_status" form:"payment_status" gorm:"type:enum('Paid','Unpaid','Canceled')"`
	PaymentUrl     string    `json:"payment_url" form:"payment_url"`
}
