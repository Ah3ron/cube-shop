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

	// Save shipping details
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
			"error": "Could not save shipping details",
		})
	}

	// Rest of the existing checkout logic remains the same
	// ... (process items, update stock, clear cart)

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
