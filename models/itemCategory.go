package models

type ItemCategory struct {
	ID       uint   `json:"id" form:"id" gorm:"primary_key"`
	Name     string `json:"name" form:"name"`
	ImageURL string `json:"image_url" form:"image_url"`
} // new
