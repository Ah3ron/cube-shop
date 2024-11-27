package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID          uint            `json:"user_id"`
	Status          string          `json:"status"`
	TotalPrice      float64         `json:"total_price"`
	OrderDate       time.Time       `json:"order_date"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	ShippingDetails ShippingDetails `json:"shipping_details" gorm:"foreignKey:OrderID"`
	OrderItems      []OrderItem     `json:"order_items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
}

type ShippingDetails struct {
	gorm.Model
	OrderID uint   `json:"order_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	City    string `json:"city"`
	Country string `json:"country"`
	ZipCode string `json:"zipCode"`
}
