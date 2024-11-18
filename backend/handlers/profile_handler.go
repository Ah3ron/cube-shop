package handlers

import (
	"cube-shop/database"
	"cube-shop/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetProfile(c *fiber.Ctx) error {
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

	return c.JSON(fiber.Map{
		"name":     userData.Name,
		"email":    userData.Email,
		"joinDate": userData.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	})
}
