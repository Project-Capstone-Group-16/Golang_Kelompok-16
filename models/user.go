package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string    `json:"username" form:"username" gorm:"unique;not null"`
	Email       string    `json:"email" form:"email" gorm:"unique;not null"`
	Passwrod    string    `json:"password" form:"password" gorm:"unique;not null"`
	Fullname    string    `json:"fullname" form:"fullname" gorm:"unique;not null"`
	Gender      string    `json:"gender" form:"gender" gorm:"type:enum('Male', 'Female')"`
	Token       string    `json:"-" form:"-"`
	BirthDate   time.Time `json:"birth_date" form:"birth_date"`
	PhoneNumber int       `json:"phone_number" form:"phone_number" gorm:"unique;not null"`
	Addres      string    `json:"address" form:"address" gorm:"unique;not null"`
	Role        string    `json:"role" form:"role" gorm:"type:enum('USER', 'ADMIN');default:'USER'"`
}
