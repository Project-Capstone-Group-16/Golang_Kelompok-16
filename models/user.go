package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string     `json:"email" form:"email" gorm:"unique"`
	Password    string     `json:"password" form:"password"`
	FirstName   string     `json:"first_name" form:"first_name"`
	LastName    string     `json:"last_name" form:"last_name"`
	BirthDate   *time.Time `json:"birth_date" form:"birth_date"`
	Gender      string     `json:"gender" form:"gender" gorm:"type:enum('PRIA', 'WANITA', '');default:''"`
	PhoneNumber string     `json:"phone_number" form:"phone_number"`
	Address     string     `json:"address" form:"address"`
	Favorite    []Favorite
	Token       string `json:"-" form:"-" gorm:"-"`
	Role        string `json:"role" form:"role" gorm:"type:enum('USER', 'ADMIN');default:'USER'"`
}
