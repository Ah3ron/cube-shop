package handlers

import (
	"cube-shop/database"
	"cube-shop/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type CheckoutRequest struct {
	Items []struct {
		ID       uint    `json:"id"`
		Quantity int     `json:"quantity"`
		Price    float64 `json:"price"`
		Product  struct {
			ID    uint    `json:"id"`
			Name  string  `json:"name"`
			Price float64 `json:"price"`
		} `json:"product"`
	} `json:"items"`
	Total float64 `json:"total"`
}

func Checkout(c *fiber.Ctx) error {
	// Parse request body
	var req CheckoutRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Get user ID from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	// Start a database transaction
	tx := database.DB.Begin()

	// Create order
	order := models.Order{
		UserID:     userID,
		Status:     "pending",
		OrderDate:  time.Now(),
		TotalPrice: req.Total,
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create order",
		})
	}

	// Create order items
	for _, item := range req.Items {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.Product.ID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not create order items",
			})
		}
	}

	// Clear the user's cart
	if err := tx.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not clear cart",
		})
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not complete checkout",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Checkout successful",
		"order":   order,
	})
}
