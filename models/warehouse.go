package models

import (
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	Name      string `json:"name" form:"name" gorm:"unique;not null"`
	Location  string `json:"location" form:"location"`
	Capacity  uint   `json:"capacity" form:"capacity"`
	Status    string `json:"status" form:"status" gorm:"type:enum('Available','Not Available')"`
	Favorites int   `json:"favorites" form:"favorites"`
	ImageURL  string `json:"image_url" form:"image_url" gorm:"unique;not null"`
	Favorite  *[]Favorite
}
