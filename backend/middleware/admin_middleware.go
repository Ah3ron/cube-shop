package middleware

import (
	"cube-shop/database"
	"cube-shop/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))

		var dbUser models.User
		if err := database.DB.First(&dbUser, userID).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Пользователь не найден",
			})
		}

		if dbUser.Role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Требуются права администратора",
			})
		}

		return c.Next()
	}
}
