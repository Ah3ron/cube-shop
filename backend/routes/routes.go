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

	// Static file serving
	app.Static("/", "./build")

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
