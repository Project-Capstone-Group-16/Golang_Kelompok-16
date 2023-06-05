package models

import (
	"time"

	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	WarehouseID uint       `json:"warehouse_id"`
	FullName    string     `json:"full_name" form:"full_name"`
	Occupation  string     `json:"occupation" form:"occupation" gorm:"type:enum('PIC', 'FRONT DESK', '');default:''"`
	BirthDate   *time.Time `json:"birth_date" form:"birth_date"`
	PhoneNumber string     `json:"phone_number" form:"phone_number"`
	Address     string     `json:"address" form:"address"`
	Warehouse   *Warehouse
}
