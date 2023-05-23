package models

import "gorm.io/gorm"

type Warehouse struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"unique;not null"`
	Location string `json:"location" form:"location" gorm:"not null"`
	Status 	 string `json:"status" form:"location" gorm:"not null"`
}
