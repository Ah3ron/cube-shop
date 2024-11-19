package main

import (
	"cube-shop/database"
	"cube-shop/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize app
	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		DisableStartupMessage:   false,
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(cache.New())

	// Connect to database
	database.Connect()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
