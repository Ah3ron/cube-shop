package main

import (
	"cube-shop/database"
	"cube-shop/routes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	// Initialize app with optimized config
	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		ServerHeader:            "Fiber",
		BodyLimit:               10 * 1024 * 1024,
		ReadTimeout:             5 * time.Second,
		WriteTimeout:            10 * time.Second,
		IdleTimeout:             120 * time.Second,
		DisableStartupMessage:   true,
		ReduceMemoryUsage:       true,
		JSONEncoder:             json.Marshal,
		JSONDecoder:             json.Unmarshal,
	})

	// CORS with secure settings
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080, https://cube-shop.up.railway.app/",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           int((24 * time.Hour).Seconds()),
	}))

	// Compression with enhanced settings
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
		Next: func(c *fiber.Ctx) bool {
			return !strings.Contains(c.Get("Accept-Encoding"), "gzip")
		},
	}))

	// Enhanced caching strategy
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			if c.Method() != "GET" {
				return true
			}
			path := c.Path()
			return strings.HasPrefix(path, "/api")
		},
		Expiration:           24 * time.Hour,
		CacheControl:         true,
		StoreResponseHeaders: true,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Path()
		},
	}))

	// Add ETag support
	app.Use(etag.New())

	// Performance monitoring
	app.Get("/metrics", monitor.New(monitor.Config{
		Title: "Metrics",
	}))

	// Optimize database connection pool
	database.Connect()

	// Setup routes
	routes.SetupRoutes(app)

	// Graceful shutdown
	if err := app.Listen(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
