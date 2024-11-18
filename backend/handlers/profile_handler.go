package handlers

import (
	"cube-shop/database"
	"cube-shop/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetProfile(c *fiber.Ctx) error {
	// Get user from JWT token using the correct type
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	// Fetch user from database
	var userData models.User
	result := database.DB.First(&userData, userID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Return formatted user data
	return c.JSON(fiber.Map{
		"name":     userData.Name,
		"email":    userData.Email,
		"joinDate": userData.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	})
}
