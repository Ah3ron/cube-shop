package handlers

import (
	"cube-shop/database"
	"cube-shop/models"

	"github.com/gofiber/fiber/v2"
)

func GetAdminStats(c *fiber.Ctx) error {
	var stats struct {
		TotalOrders   int64   `json:"totalOrders"`
		TotalProducts int64   `json:"totalProducts"`
		TotalUsers    int64   `json:"totalUsers"`
		TotalRevenue  float64 `json:"totalRevenue"`
	}

	// Get total orders
	database.DB.Model(&models.Order{}).Count(&stats.TotalOrders)

	// Get total products
	database.DB.Model(&models.Product{}).Count(&stats.TotalProducts)

	// Get total users
	database.DB.Model(&models.User{}).Count(&stats.TotalUsers)

	// Calculate total revenue
	database.DB.Model(&models.Order{}).
		Where("status != ?", "cancelled").
		Select("COALESCE(SUM(total_amount), 0)").
		Scan(&stats.TotalRevenue)

	return c.JSON(fiber.Map{
		"stats": stats,
	})
}
