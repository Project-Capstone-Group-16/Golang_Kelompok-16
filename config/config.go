package config

import (
	"Capstone/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() *gorm.DB {
	config := Config{
		DB_Username: "root",
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Name:     "inventron",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	InitMigrate()
	Seeders()

	return DB
}

func InitMigrate() {
	// Migrate the schema
	err := DB.AutoMigrate(&models.User{}, &models.Favorite{}, &models.Warehouse{}, &models.Staff{}, &models.Locker{}, &models.LockerType{}, &models.ItemCategory{}, &models.Transaction{})
	if err != nil {
		panic("failed to migrate database")
	}
}

func Seeders() {
	lockerType := []models.LockerType{
		{Name: "Small", Price: 15000},
		{Name: "Medium", Price: 20000},
		{Name: "Large", Price: 30000},
	}

	itemCategory := []models.ItemCategory{
		{ID: 1, Name: "Tas Ransel", ImageURL: "https://res.cloudinary.com/ddf2m61gv/image/upload/v1686493003/Inventron/Tas%20Ransel.png.png"},
		{ID: 2, Name: "Sepatu", ImageURL: "https://res.cloudinary.com/ddf2m61gv/image/upload/v1686493074/Inventron/Sepatu.png.png"},
		{ID: 3, Name: "Pakaian", ImageURL: "hhttps://res.cloudinary.com/ddf2m61gv/image/upload/v1686493093/Inventron/Pakaian.png.png"},
		{ID: 4, Name: "Kerdus", ImageURL: "https://res.cloudinary.com/ddf2m61gv/image/upload/v1686493120/Inventron/Kerdus.png.png"},
	}

	for _, v := range lockerType {
		var exist models.LockerType
		err := DB.Where("name = ?", v.Name).First(&exist).Error
		if err != nil {
			DB.Create(&v)
		}
	}

	for _, v := range itemCategory {
		var exist models.ItemCategory
		err := DB.Where("name = ?", v.Name).First(&exist).Error
		if err != nil {
			DB.Create(&v)
		}
	}
}
