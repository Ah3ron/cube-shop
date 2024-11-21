package main

import (
	"cube-shop/database"
	"cube-shop/routes"
	"log"
	"strings"
	"time"

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
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			path := c.Path()
			isStaticAsset := strings.HasSuffix(path, ".css") || strings.HasSuffix(path, ".js")
			return !isStaticAsset || c.Query("noCache") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	// Connect to database
	database.Connect()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
