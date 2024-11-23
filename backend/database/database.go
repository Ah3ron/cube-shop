package database

import (
	"cube-shop/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db_url := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	DB = db
	db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
		&models.ShippingDetails{},
	)

	var count int64
	DB.Model(&models.Product{}).Count(&count)
	if count == 0 {
		SeedData()
	}
}
