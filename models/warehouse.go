package models

import (
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"unique;not null"`
	Location string `json:"location" form:"location"`
	Status   string `json:"status" form:"status"`
	Favorite []Favorite
}
