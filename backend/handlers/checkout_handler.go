package handlers

import (
	"cube-shop/database"
	"cube-shop/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type ShippingDetails struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	City    string `json:"city"`
	Country string `json:"country"`
	ZipCode string `json:"zipCode"`
}

type CheckoutRequest struct {
	Items []struct {
		ID        uint       `json:"id"`
		CreatedAt time.Time  `json:"CreatedAt"`
		UpdatedAt time.Time  `json:"UpdatedAt"`
		DeletedAt *time.Time `json:"DeletedAt"`
		UserID    uint       `json:"user_id"`
		ProductID uint       `json:"product_id"`
		Quantity  int        `json:"quantity"`
		Product   struct {
			ID          uint       `json:"ID"`
			CreatedAt   time.Time  `json:"CreatedAt"`
			UpdatedAt   time.Time  `json:"UpdatedAt"`
			DeletedAt   *time.Time `json:"DeletedAt"`
			Name        string     `json:"name"`
			Description string     `json:"description"`
			Price       float64    `json:"price"`
			Stock       int        `json:"stock"`
			ImageURL    string     `json:"image_url"`
			Category    string     `json:"category"`
			Difficulty  string     `json:"difficulty"`
			Brand       string     `json:"brand"`
		} `json:"product"`
	} `json:"items"`
	Total           float64         `json:"total"`
	ShippingDetails ShippingDetails `json:"shippingDetails"`
}

func Checkout(c *fiber.Ctx) error {
	var req CheckoutRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	tx := database.DB.Begin()

	order := models.Order{
		UserID:     userID,
		Status:     "pending",
		OrderDate:  time.Now(),
		TotalPrice: req.Total,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create order",
		})
	}

	for _, item := range req.Items {
		// Verify product stock
		var product models.Product
		if err := tx.First(&product, item.Product.ID).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Product not found",
			})
		}

		if product.Stock < item.Quantity {
			tx.Rollback()
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Insufficient stock for " + product.Name,
			})
		}

		// Update stock
		if err := tx.Model(&product).Update("stock", product.Stock-item.Quantity).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not update product stock",
			})
		}

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

	if err := tx.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not clear cart",
		})
	}

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
