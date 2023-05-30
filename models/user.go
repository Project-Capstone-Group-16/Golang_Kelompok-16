package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string     `json:"email" form:"email" gorm:"unique"`
	Password    string     `json:"password" form:"password"`
	Fullname    string     `json:"fullname" form:"fullname"`
	BirthDate   *time.Time `json:"birth_date" form:"birth_date"`
	Gender      string     `json:"gender" form:"gender" gorm:"type:enum('MALE', 'FEMALE', '');default:''"`
	PhoneNumber string     `json:"phone_number" form:"phone_number"`
	Address     string     `json:"address" form:"address"`
	Favorite    []Favorite
	Token       string `json:"-" form:"-"`
	Role        string `json:"role" form:"role" gorm:"type:enum('USER', 'ADMIN');default:'USER'"`
}
