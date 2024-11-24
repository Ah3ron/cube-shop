package middleware

import (
	"cube-shop/database"
	"cube-shop/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RequireRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get user claims from JWT token
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))

		// Fetch user data from database
		var userData models.User
		result := database.DB.First(&userData, userID)
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		// Check if user has required role
		if userData.Role != role {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   "Access denied",
				"message": "Insufficient permissions",
			})
		}

		return c.Next()
	}
}
