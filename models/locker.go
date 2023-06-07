package models

import "gorm.io/gorm"

type Locker struct {
	gorm.Model
	WarehouseID  uint   `json:"warehouse_id" form:"warehouse_id"`
	LockerTypeID uint   `json:"locker_type_id" form:"locker_type_id"`
	LockerNumber uint   `json:"locker_number" form:"locker_number"`
	Availability string `json:"availability" form:"availability" gorm:"type:enum('Available','Not Available')"`
	Warehouse    Warehouse
	LockerType   LockerType
}

type LockerType struct {
	gorm.Model
	Name       string `json:"name" form:"name" gorm:"type:enum('Small','Medium','Large')"`
	PriceDay   uint   `json:"price_day" form:"price_day"`
	PriceMonth uint   `json:"price_month" form:"price_month"`
	PriceYear  uint   `json:"price_year" form:"price_year"`
	Height     uint   `json:"height" form:"height"`
	Width      uint   `json:"width" form:"width"`
	Length     uint   `json:"length" form:"length"`
}
