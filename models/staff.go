package models

import (
	"time"

	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	WarehouseID  uint       `json:"warehouse_id"`
	FullName     string     `json:"full_name" form:"full_name"`
	Occupation   string     `json:"occupation" form:"occupation"`
	WorkingHours string     `json:"working_hours" form:"working_hours"`
	BirthDate    *time.Time `json:"birth_date" form:"birth_date"`
	PhoneNumber  string     `json:"phone_number" form:"phone_number"`
	Warehouse    *Warehouse
}
