package handlers

import (
	"cube-shop/database"
	"cube-shop/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// GetCart возвращает корзину текущего пользователя
func GetCart(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var cartItems []models.CartItem
	result := database.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось получить содержимое корзины",
		})
	}

	// Подсчет общей суммы
	var total float64
	for _, item := range cartItems {
		total += item.Product.Price * float64(item.Quantity)
	}

	return c.JSON(fiber.Map{
		"items": cartItems,
		"total": total,
	})
}

// AddToCart добавляет товар в корзину
func AddToCart(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var input struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}

	// Проверка наличия товара
	var product models.Product
	if err := database.DB.First(&product, input.ProductID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Товар не найден",
		})
	}

	// Проверка наличия достаточного количества товара
	if product.Stock < input.Quantity {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Недостаточно товара на складе",
			"available": product.Stock,
		})
	}

	// Проверка существующего товара в корзине
	var existingItem models.CartItem
	result := database.DB.Where("user_id = ? AND product_id = ?", userID, input.ProductID).First(&existingItem)
	
	if result.Error == nil {
		// Обновляем количество существующего товара
		newQuantity := existingItem.Quantity + input.Quantity
		if newQuantity > product.Stock {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Превышен доступный остаток товара",
				"available": product.Stock,
			})
		}

		existingItem.Quantity = newQuantity
		database.DB.Save(&existingItem)
		return c.JSON(existingItem)
	}

	// Создаем новый товар в корзине
	cartItem := models.CartItem{
		UserID:    userID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
	}

	if err := database.DB.Create(&cartItem).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось добавить товар в корзину",
		})
    }

    return c.Status(fiber.StatusCreated).JSON(cartItem)
}

// UpdateCartItem обновляет количество товара в корзине
func UpdateCartItem(c *fiber.Ctx) error {
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    userID := uint(claims["user_id"].(float64))
    itemID := c.Params("id")

    var input struct {
        Quantity int `json:"quantity"`
    }

    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Неверный формат данных",
        })
    }

    var cartItem models.CartItem
    if err := database.DB.Preload("Product").Where("id = ? AND user_id = ?", itemID, userID).First(&cartItem).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Товар не найден в корзине",
        })
    }

    // Проверка наличия товара
    if cartItem.Product.Stock < input.Quantity {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Недостаточно товара на складе",
            "available": cartItem.Product.Stock,
        })
    }

    cartItem.Quantity = input.Quantity
    if err := database.DB.Save(&cartItem).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Не удалось обновить количество товара",
        })
    }

    return c.JSON(cartItem)
}

// RemoveFromCart удаляет товар из корзины
func RemoveFromCart(c *fiber.Ctx) error {
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    userID := uint(claims["user_id"].(float64))
    itemID := c.Params("id")

    result := database.DB.Where("id = ? AND user_id = ?", itemID, userID).Delete(&models.CartItem{})
    if result.RowsAffected == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Товар не найден в корзине",
        })
    }

    return c.SendStatus(fiber.StatusNoContent)
}

// ClearCart очищает корзину пользователя
func ClearCart(c *fiber.Ctx) error {
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    userID := uint(claims["user_id"].(float64))

    if err := database.DB.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Не удалось очистить корзину",
        })
    }

    return c.SendStatus(fiber.StatusNoContent)
}
 