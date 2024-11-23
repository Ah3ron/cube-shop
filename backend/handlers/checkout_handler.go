package handlers

import (
	"cube-shop/database"
	"cube-shop/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type CheckoutRequest struct {
	Items []struct {
		Quantity int            `json:"quantity"`
		Product  models.Product `json:"product"`
	} `json:"items"`
	Total           float64 `json:"total"`
	ShippingDetails struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Address string `json:"address"`
		City    string `json:"city"`
		Country string `json:"country"`
		ZipCode string `json:"zipCode"`
	} `json:"shippingDetails"`
}

func Checkout(c *fiber.Ctx) error {
	var req CheckoutRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate request
	if len(req.Items) == 0 || req.Total <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order details",
		})
	}

	// Extract user ID from JWT
	userID := uint(c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64))

	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

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
			"error": "Failed to create order",
		})
	}

	// Create shipping details
	shippingDetails := models.ShippingDetails{
		OrderID: order.ID,
		Name:    req.ShippingDetails.Name,
		Email:   req.ShippingDetails.Email,
		Address: req.ShippingDetails.Address,
		City:    req.ShippingDetails.City,
		Country: req.ShippingDetails.Country,
		ZipCode: req.ShippingDetails.ZipCode,
	}

	if err := tx.Create(&shippingDetails).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save shipping details",
		})
	}

	// Process order items
	for _, item := range req.Items {
		// Create order item
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.Product.ID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create order item",
			})
		}

		// Update product stock
		if err := tx.Model(&models.Product{}).
			Where("id = ?", item.Product.ID).
			Where("stock >= ?", item.Quantity).
			UpdateColumn("stock", gorm.Expr("stock - ?", item.Quantity)).
			Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update product stock",
			})
		}
	}

	// Clear user's cart
	if err := tx.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to clear cart",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to complete checkout",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Checkout successful",
		"order":   order,
	})
}
