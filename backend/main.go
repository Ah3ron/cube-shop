package main

import (
	"cube-shop/database"
	"cube-shop/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize app
	app := fiber.New(fiber.Config{StrictRouting: true})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Connect to database
	database.Connect()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
