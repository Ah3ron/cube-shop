package routes

import (
	"cube-shop/handlers"
	"cube-shop/middleware"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func SetupRoutes(app *fiber.App) {
	// API routes group
	api := app.Group("/api")

	// Public routes
	setupPublicRoutes(api)

	// Protected routes
	setupProtectedRoutes(api)

	// Static files serving
	setupStaticFiles(app)
}

func setupPublicRoutes(api fiber.Router) {
	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// Public product routes
	api.Get("/products", handlers.GetProducts)
	api.Get("/products/:id", handlers.GetProduct)
}

func setupProtectedRoutes(api fiber.Router) {
	// Protected routes group
	protected := api.Group("/", middleware.Protected())

	// Checkout route
	protected.Post("/checkout", handlers.Checkout)

	// User routes
	protected.Get("/profile", handlers.GetProfile)

	// Product management routes
	products := protected.Group("/products")
	products.Post("/", handlers.CreateProduct)
	products.Put("/:id", handlers.UpdateProduct)
	products.Delete("/:id", handlers.DeleteProduct)

	// Cart routes
	cart := protected.Group("/cart")
	cart.Get("/", handlers.GetCart)
	cart.Post("/", handlers.AddToCart)
	cart.Patch("/:id", handlers.UpdateCartItem)
	cart.Delete("/:id", handlers.RemoveFromCart)

	// Order routes
	orders := protected.Group("/orders")
	orders.Get("/", handlers.GetOrdersByUser)
	orders.Get("/:id", handlers.GetOrderByID)
	orders.Put("/:id/status", handlers.UpdateOrderStatus)
	orders.Put("/:id/cancel", handlers.CancelOrder)

	admin := protected.Group("/admin", middleware.RequireRole("admin"))

	// Order Management
	admin.Get("/orders", handlers.GetAllOrders)
	admin.Get("/orders/:id", handlers.GetOrderByID)
	admin.Put("/orders/:id/status", handlers.UpdateOrderStatus)

	// Product Management
	admin.Get("/products", handlers.GetProducts)
	admin.Post("/products", handlers.CreateProduct)
	admin.Put("/products/:id", handlers.UpdateProduct)
	admin.Delete("/products/:id", handlers.DeleteProduct)

	// User Management

	admin.Get("/users", handlers.GetAllUsers)
	admin.Get("/users/:id", handlers.GetUserByID)
	admin.Put("/users/:id/role", handlers.UpdateUserRole)
}

func setupStaticFiles(app *fiber.App) {
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.Dir("build"),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "index.html", // Important for SPA routing
	}))
}
