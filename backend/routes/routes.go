package routes

import (
	"cube-shop/handlers"
	"cube-shop/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Public Routes
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	api.Get("/products", handlers.GetProducts)
	api.Get("/products/:id", handlers.GetProduct)

	// Protected Routes
	protected := api.Group("/", middleware.Protected())

	// User Routes
	protected.Get("/profile", handlers.GetProfile)

	// Cart Routes
	cart := protected.Group("/cart")
	cart.Get("/", handlers.GetCart)
	cart.Post("/", handlers.AddToCart)
	cart.Patch("/:id", handlers.UpdateCartItem)
	cart.Delete("/:id", handlers.RemoveFromCart)

	// Product Routes
	products := protected.Group("/products")
	products.Post("/", handlers.CreateProduct)
	products.Put("/:id", handlers.UpdateProduct)
	products.Delete("/:id", handlers.DeleteProduct)

	// Order Routes
	orders := protected.Group("/orders")
	orders.Get("/", handlers.GetOrdersByUser)
	orders.Get("/:id", handlers.GetOrderByID)
	orders.Put("/:id/status", handlers.UpdateOrderStatus)
	orders.Put("/:id/cancel", handlers.CancelOrder)

	// Checkout Route
	protected.Post("/checkout", handlers.Checkout)

	// Admin Routes
	admin := protected.Group("/admin", middleware.RequireRole("admin"))

	// Admin Order Management
	adminOrders := admin.Group("/orders")
	adminOrders.Get("/", handlers.GetAllOrders)
	adminOrders.Get("/:id", handlers.GetOrderByID)
	adminOrders.Put("/:id/status", handlers.UpdateOrderStatus)

	// Admin Product Management
	adminProducts := admin.Group("/products")
	adminProducts.Get("/", handlers.GetProducts)
	adminProducts.Post("/", handlers.CreateProduct)
	adminProducts.Put("/:id", handlers.UpdateProduct)
	adminProducts.Delete("/:id", handlers.DeleteProduct)

	// Admin User Management
	adminUsers := admin.Group("/users")
	adminUsers.Get("/", handlers.GetAllUsers)

	// Admin Stats Route
	admin.Get("/stats", handlers.GetAdminStats)
}
