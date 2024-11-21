package handlers

import (
	"cube-shop/database"
	"cube-shop/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetCart(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var cartItems []models.CartItem
	result := database.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch cart items",
		})
	}
	return c.JSON(cartItems)
}

func AddToCart(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	cartItem := new(models.CartItem)
	if err := c.BodyParser(cartItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	cartItem.UserID = userID

	// Check if product exists
	var product models.Product
	if err := database.DB.First(&product, cartItem.ProductID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	// Check if item already exists in cart
	var existingItem models.CartItem
	result := database.DB.Where("user_id = ? AND product_id = ?", userID, cartItem.ProductID).First(&existingItem)
	if result.Error == nil {
		// Update quantity
		existingItem.Quantity += cartItem.Quantity
		database.DB.Save(&existingItem)
		return c.JSON(existingItem)
	}

	result = database.DB.Create(&cartItem)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not add item to cart",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(cartItem)
}

func UpdateCartItem(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	id := c.Params("id")
	cartItem := new(models.CartItem)

	if err := c.BodyParser(cartItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	result := database.DB.Model(&models.CartItem{}).Where("id = ? AND user_id = ?", id, userID).Updates(cartItem)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update cart item",
		})
	}
	return c.JSON(cartItem)
}

func RemoveFromCart(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	id := c.Params("id")
	result := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.CartItem{})
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not remove item from cart",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
