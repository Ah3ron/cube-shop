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

	// Serve static files from the build directory
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.Dir("build"),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "index.html", // Important for SPA routing
	}))

}
