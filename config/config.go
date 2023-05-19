package config

import (
	"Capstone/models"
	"fmt"

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
		DB_Password: "subaruputra13",
		DB_Port:     "3306",
		DB_Host:     "database-mariadb.covxu3uwiatv.ap-southeast-2.rds.amazonaws.com",
		DB_Name:     "inventron",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
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

	return DB
}

func InitMigrate() {
	// Migrate the schema
	err := DB.AutoMigrate(&models.User{})

	if err != nil {
		panic("Failed to migrate database")
	}
}
