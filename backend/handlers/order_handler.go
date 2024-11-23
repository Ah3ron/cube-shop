package handlers

import (
	"cube-shop/database"
	"cube-shop/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetOrders(c *fiber.Ctx) error {
	var orders []models.Order
	result := database.DB.Preload("OrderItem").Find(&orders)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch orders",
		})
	}
	return c.JSON(orders)
}

func GetOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.Order
	result := database.DB.Preload("OrderItem").First(&order, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	}
	return c.JSON(order)
}

func CreateOrder(c *fiber.Ctx) error {
	// Parse the order request
	type OrderRequest struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	var orderReq OrderRequest
	if err := c.BodyParser(&orderReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Get product to calculate total
	var product models.Product
	if err := database.DB.First(&product, orderReq.ProductID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	// Create the order
	order := models.Order{
		UserID:     c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(uint),
		Status:     "pending",
		OrderDate:  time.Now(),
		TotalPrice: float64(orderReq.Quantity) * product.Price,
	}

	// Start transaction
	tx := database.DB.Begin()

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create order",
		})
	}

	// Create order item
	orderItem := models.OrderItem{
		OrderID:   order.ID,
		ProductID: orderReq.ProductID,
		Quantity:  orderReq.Quantity,
		Price:     product.Price,
	}

	if err := tx.Create(&orderItem).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create order item",
		})
	}

	tx.Commit()
	return c.Status(fiber.StatusCreated).JSON(order)
}
