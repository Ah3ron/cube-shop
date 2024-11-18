package routes

import (
	"cube-shop/handlers"
	"cube-shop/middleware"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// API routes group
	api := app.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// Unprotected routes
	api.Get("/products", handlers.GetProducts)

	// Protected API routes
	protected := api.Group("/", middleware.Protected())
	protected.Get("/profile", handlers.GetProfile)
	products := protected.Group("/products")
	products.Post("/", handlers.CreateProduct)
	products.Get("/:id", handlers.GetProduct)
	products.Put("/:id", handlers.UpdateProduct)
	products.Delete("/:id", handlers.DeleteProduct)

	// Static file serving with MIME types
	app.Static("/", "./build", fiber.Static{
		ByteRange: true,
		Browse:    false,
		MaxAge:    3600,
		ModifyResponse: func(c *fiber.Ctx) error {
			ext := filepath.Ext(c.Path())
			switch ext {
			case ".js":
				c.Type("application/javascript; charset=utf-8")
			case ".mjs":
				c.Type("application/javascript; charset=utf-8")
			case ".css":
				c.Type("text/css; charset=utf-8")
			}
			return nil
		},
	})

	// Specific _app directory handling
	app.Static("/_app", "./build/_app", fiber.Static{
		ByteRange: true,
		Browse:    false,
		MaxAge:    3600,
		ModifyResponse: func(c *fiber.Ctx) error {
			ext := filepath.Ext(c.Path())
			switch ext {
			case ".js":
				c.Type("application/javascript; charset=utf-8")
			case ".mjs":
				c.Type("application/javascript; charset=utf-8")
			case ".css":
				c.Type("text/css; charset=utf-8")
			}
			return nil
		},
	})

	// SPA fallback route
	app.Get("/*", func(c *fiber.Ctx) error {
		path := c.Path()
		fullPath := filepath.Join("./build", path)

		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			return c.SendFile("./build/index.html")
		}

		return c.Next()
	})
}
