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
	"github.com/gofiber/fiber/v2/middleware/limiter"
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

	// Rate limiting
	app.Use(limiter.New(limiter.Config{
		Max:        1488,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
	}))

	// CORS with optimized settings
	app.Use(cors.New())

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
			if c.Method() != "GET" && c.Method() != "HEAD" {
				return true
			}
			path := c.Path()
			isStaticAsset := strings.HasSuffix(path, ".css") ||
				strings.HasSuffix(path, ".js") ||
				strings.HasSuffix(path, ".jpg") ||
				strings.HasSuffix(path, ".png") ||
				strings.HasSuffix(path, ".ico") ||
				strings.HasSuffix(path, ".woff2") ||
				strings.HasSuffix(path, ".webp")
			return !isStaticAsset || c.Query("noCache") == "true"
		},
		Expiration:           30 * time.Minute,
		CacheControl:         true,
		StoreResponseHeaders: true,
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
