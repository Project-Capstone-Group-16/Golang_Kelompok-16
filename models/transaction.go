package models

import (
	"time"
)

type Transaction struct {
	OrderID        uint `json:"order_id" form:"order_id" gorm:"primarykey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	UserID         uint      `json:"user_id" form:"user_id"`
	LockerID       uint      `json:"locker_id" form:"locker_id"`
	ItemCategoryID uint      `json:"item_category_id" form:"item_category_id"`
	Amount         uint      `json:"amount" form:"amount"`
	StartDate      time.Time `json:"start_date" form:"start_date"`
	EndDate        time.Time `json:"end_date" form:"end_date"`
	PaymentStatus  string    `json:"payment_status" form:"payment_status" gorm:"type:enum('Paid','Unpaid')"`
	PaymentUrl     string    `json:"payment_url" form:"payment_url"`
	User           User
	Locker         Locker
	ItemCategory   ItemCategory
}
