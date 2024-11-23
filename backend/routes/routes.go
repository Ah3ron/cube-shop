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
	cart.Post("/checkout", handlers.Checkout)
}

func setupStaticFiles(app *fiber.App) {
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.Dir("build"),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "index.html", // Important for SPA routing
	}))
}
