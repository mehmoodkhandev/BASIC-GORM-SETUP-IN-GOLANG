package config

import (
	"fmt"

	"github.com/MehmoodNadeemKhan1/CRUD_API/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config Config) *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUsername, config.DBPassword, config.DBName, config.DBPort)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)
	fmt.Println("Database connected successfully")
	return db

}
