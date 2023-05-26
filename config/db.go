package config

import (
	"fmt"
	"go-trial-class/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func DBConnect() {
	conn := "host=localhost user=postgres password=root dbname=trial_golang port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database", err.Error())
	} else {
		fmt.Println("DB Connected")
		DB = db
	}

	DB.AutoMigrate(&entity.Product{}, &entity.Order{})
}
