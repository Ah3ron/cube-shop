package models

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	ID        uint    `json:"id" gorm:"primarykey"`
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
}
