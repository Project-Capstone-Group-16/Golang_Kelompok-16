package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string    `json:"email" form:"email" gorm:"unique;not null"`
	Password    string    `json:"password" form:"password"`
	Fullname    string    `json:"fullname" form:"fullname" gorm:"not null"`
	Gender      string    `json:"gender" form:"gender" gorm:"type:enum('MALE','FEMALE')"`
	BirthDate   time.Time `json:"birth_date" form:"birth_date"`
	PhoneNumber string    `json:"phone_number" form:"phone_number" gorm:"not null"`
	Address     string    `json:"address" form:"address" gorm:"not null"`
	Favorite    []Favorite
	Token       string `json:"-" form:"-"`
	Role        string `json:"role" form:"role" gorm:"type:enum('USER', 'ADMIN');default:'USER'"`
}
