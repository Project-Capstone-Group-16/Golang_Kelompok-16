package models

import "gorm.io/gorm"

type Locker struct {
	gorm.Model
	WarehouseID  uint   `json:"warehouse_id" form:"warehouse_id"`
	LockerTypeID uint   `json:"locker_type_id" form:"locker_type_id"`
	Name         string `json:"name" form:"name"`
	LockerNumber uint   `json:"locker_number" form:"locker_number"`
	Availability string `json:"availability" form:"availability" gorm:"type:enum('Available','Not Available')"`
	Warehouse    Warehouse
	LockerType   LockerType
}

type LockerType struct {
	gorm.Model
	Name  string `json:"name" form:"name" gorm:"type:enum('Small','Medium','Large')"`
	Price uint   `json:"price" form:"price"`
}
