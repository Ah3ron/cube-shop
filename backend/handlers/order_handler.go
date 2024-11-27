package handlers

import (
	"cube-shop/database"
	"cube-shop/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func Checkout(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var input struct {
		ShippingDetails models.ShippingDetails `json:"shipping_details"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}

	// Получаем корзину пользователя
	var cartItems []models.CartItem
	if err := database.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось получить корзину",
		})
	}

	if len(cartItems) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Корзина пуста",
		})
	}

	// Начинаем транзакцию
	tx := database.DB.Begin()

	// Создаем заказ
	var totalPrice float64
	for _, item := range cartItems {
		totalPrice += item.Product.Price * float64(item.Quantity)
	}

	order := models.Order{
		UserID:          userID,
		Status:          "pending",
		TotalPrice:      totalPrice,
		OrderDate:       time.Now(),
		ShippingDetails: input.ShippingDetails,
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось создать заказ",
		})
	}

	// Создаем элементы заказа и обновляем склад
	for _, item := range cartItems {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Не удалось создать элемент заказа",
			})
		}

		// Обновляем количество товара на складе
		if err := tx.Model(&models.Product{}).
			Where("id = ? AND stock >= ?", item.ProductID, item.Quantity).
			UpdateColumn("stock", gorm.Expr("stock - ?", item.Quantity)).
			Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Недостаточно товара на складе",
			})
		}
	}

	// Очищаем корзину
	if err := tx.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось очистить корзину",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось завершить заказ",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

func GetOrders(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var orders []models.Order
	result := database.DB.
		Preload("ShippingDetails").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&orders)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось получить заказы",
		})
	}

	return c.JSON(orders)
}

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
			"error": "Заказ не найден",
		})
	}

	return c.JSON(order)
}

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
			"error": "Заказ не найден",
		})
	}

	if order.Status != "pending" {
		tx.Rollback()
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Можно отменить только заказы в статусе 'pending'",
		})
	}

	// Возвращаем товары на склад
	var orderItems []models.OrderItem
	if err := tx.Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось получить элементы заказа",
		})
	}

	for _, item := range orderItems {
		if err := tx.Model(&models.Product{}).
			Where("id = ?", item.ProductID).
			UpdateColumn("stock", gorm.Expr("stock + ?", item.Quantity)).
			Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Не удалось обновить склад",
			})
		}
	}

	order.Status = "cancelled"
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось обновить статус заказа",
		})
	}

	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось отменить заказ",
		})
	}

	return c.JSON(order)
}

func UpdateOrderStatus(c *fiber.Ctx) error {
	orderID := c.Params("id")
	
	var input struct {
		Status string `json:"status"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}

	var order models.Order
	if err := database.DB.First(&order, orderID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Заказ не найден",
		})
	}

	allowedStatuses := map[string]bool{
		"pending":    true,
		"processing": true,
		"shipped":    true,
		"delivered":  true,
		"cancelled":  true,
	}

	if !allowedStatuses[input.Status] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Недопустимый статус заказа",
		})
	}

	order.Status = input.Status
	if err := database.DB.Save(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось обновить статус заказа",
		})
	}

	return c.JSON(order)
}

func GetAllOrders(c *fiber.Ctx) error {
	var orders []models.Order
	result := database.DB.
		Preload("ShippingDetails").
		Preload("OrderItems").
		Preload("OrderItems.Product").
		Order("created_at desc").
		Find(&orders)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось получить заказы",
		})
	}

	return c.JSON(orders)
} 