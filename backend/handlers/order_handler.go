package handlers

import (
	"cube-shop/database"
	"cube-shop/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func GetAllOrders(c *fiber.Ctx) error {
	var orders []models.Order
	result := database.DB.Preload("OrderItems.Product").Preload("ShippingDetails").Find(&orders)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch orders"})
	}
	return c.JSON(orders)
}

// GetOrdersByUser retrieves all orders for the authenticated user
func GetOrdersByUser(c *fiber.Ctx) error {
	// Extract user ID from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	// Initialize orders slice
	var orders []models.Order

	// Query the database with preloaded relationships
	result := database.DB.
		Preload("ShippingDetails").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders)

	// Handle potential errors
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch orders",
		})
	}

	// Return the orders as JSON
	return c.JSON(orders)
}

// GetOrderByID retrieves a specific order by ID
func GetOrderByID(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	orderID := c.Params("id")

	var order models.Order
	result := database.DB.
		Preload("ShippingDetails").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Where("id = ? AND user_id = ?", orderID, userID).
		First(&order)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	return c.JSON(order)
}

// UpdateOrderStatus updates the status of an order
func UpdateOrderStatus(c *fiber.Ctx) error {
	type UpdateStatusRequest struct {
		Status string `json:"status"`
	}

	var req UpdateStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	orderID := c.Params("id")

	var order models.Order
	result := database.DB.First(&order, orderID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	// Update status
	order.Status = req.Status
	order.UpdatedAt = time.Now()

	if err := database.DB.Save(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update order status",
		})
	}

	return c.JSON(order)
}

// CancelOrder cancels an order
func CancelOrder(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	orderID := c.Params("id")

	tx := database.DB.Begin()

	var order models.Order
	if err := tx.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	if order.Status != "pending" {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Only pending orders can be cancelled",
		})
	}

	// Update order status
	order.Status = "cancelled"
	order.UpdatedAt = time.Now()

	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to cancel order",
		})
	}

	// Restore product stock
	var orderItems []models.OrderItem
	if err := tx.Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch order items",
		})
	}

	for _, item := range orderItems {
		if err := tx.Model(&models.Product{}).
			Where("id = ?", item.ProductID).
			UpdateColumn("stock", gorm.Expr("stock + ?", item.Quantity)).
			Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to restore product stock",
			})
		}
	}

	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to complete order cancellation",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Order cancelled successfully",
		"order":   order,
	})
}
