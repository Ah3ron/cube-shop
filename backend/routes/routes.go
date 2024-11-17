package routes

import (
	"cube-shop/handlers"
	"cube-shop/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Static files
	app.Static("/", "./build/")

	// Public routes
	auth := app.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// Protected routes
	api := app.Group("/api", middleware.Protected())

	// Products
	products := api.Group("/products")
	products.Get("/", handlers.GetProducts)
	products.Get("/:id", handlers.GetProduct)
	products.Post("/", handlers.CreateProduct)
	products.Put("/:id", handlers.UpdateProduct)
	products.Delete("/:id", handlers.DeleteProduct)

	// Orders
	orders := api.Group("/orders")
	orders.Get("/", handlers.GetOrders)
	orders.Get("/:id", handlers.GetOrder)
	orders.Post("/", handlers.CreateOrder)
	orders.Put("/:id", handlers.UpdateOrder)
}
