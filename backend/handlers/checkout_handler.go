package handlers

import (
	"cube-shop/database"
	"cube-shop/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Checkout(c *fiber.Ctx) error {
	// Get user ID from JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	// Start a database transaction
	tx := database.DB.Begin()

	// Get all cart items for the user
	var cartItems []models.CartItem
	if err := tx.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not retrieve cart items",
		})
	}

	if len(cartItems) == 0 {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cart is empty",
		})
	}

	// Create order
	order := models.Order{
		UserID:     userID,
		Status:     "pending",
		OrderDate:  time.Now(),
		TotalPrice: 0,
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create order",
		})
	}

	// Create order items and calculate total
	var totalPrice float64 = 0
	for _, cartItem := range cartItems {
		var product models.Product
		if err := tx.First(&product, cartItem.ProductID).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not find product",
			})
		}

		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
			Price:     product.Price,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not create order item",
			})
		}

		totalPrice += product.Price * float64(cartItem.Quantity)
	}

	// Update order with total price
	order.TotalPrice = totalPrice
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update order total",
		})
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
