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
		DB_Username: os.Getenv("DB_USERNAME"),
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
	err := DB.AutoMigrate(&models.User{}, &models.Favorite{}, &models.Warehouse{}, &models.Staff{}, &models.Locker{}, &models.LockerType{})
	if err != nil {
		panic("failed to migrate database")
	}
}

func Seeders() {
	lockerType := []models.LockerType{
		{Name: "Small", PriceDay: 10000, PriceMonth: 200000, PriceYear: 2000000, Height: 30, Width: 30, Length: 30},
		{Name: "Medium", PriceDay: 20000, PriceMonth: 400000, PriceYear: 4000000, Height: 50, Width: 50, Length: 50},
		{Name: "Large", PriceDay: 30000, PriceMonth: 600000, PriceYear: 6000000, Height: 70, Width: 70, Length: 70},
	}

	for _, v := range lockerType {
		var exist models.LockerType
		err := DB.Where("name = ?", v.Name).First(&exist).Error
		if err != nil {
			DB.Create(&v)
		}
	}
}
