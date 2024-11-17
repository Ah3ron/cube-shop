package routes

import (
	"cube-shop/handlers"
	"cube-shop/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// API routes group
	api := app.Group("/api")

	// Public routes
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// Protected routes
	protected := api.Group("/", middleware.Protected())

	// Products
	products := protected.Group("/products")
	products.Get("/", handlers.GetProducts)
	products.Get("/:id", handlers.GetProduct)
	products.Post("/", handlers.CreateProduct)
	products.Put("/:id", handlers.UpdateProduct)
	products.Delete("/:id", handlers.DeleteProduct)

	// Orders
	orders := protected.Group("/orders")
	orders.Get("/", handlers.GetOrders)
	orders.Get("/:id", handlers.GetOrder)
	orders.Post("/", handlers.CreateOrder)
	orders.Put("/:id", handlers.UpdateOrder)

	// Static files - should be last to not interfere with API routes
	app.Static("/", "./build")

	// SPA fallback - handle client-side routing
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("./build/index.html")
	})
}
