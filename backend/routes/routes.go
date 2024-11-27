package routes

import (
	"cube-shop/handlers"
	"cube-shop/middleware"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	setupAuthRoutes(api)
	setupPublicProductRoutes(api)

	protected := api.Group("/", middleware.Protected())
	setupUserRoutes(protected)
	setupProductRoutes(protected)
	setupCartRoutes(protected)
	setupOrderRoutes(protected)

	setupStaticFiles(app)
}

func setupAuthRoutes(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)
	auth.Post("/refresh", handlers.RefreshToken)
}

func setupPublicProductRoutes(api fiber.Router) {
	products := api.Group("/products")
	products.Get("/", handlers.GetProducts)
	products.Get("/:id", handlers.GetProduct)
	products.Get("/categories", handlers.GetCategories)
	products.Get("/search", handlers.SearchProducts)
}

func setupUserRoutes(protected fiber.Router) {
	users := protected.Group("/users")
	users.Get("/me", handlers.GetProfile)
	users.Put("/me", handlers.UpdateProfile)
	users.Put("/me/password", handlers.UpdatePassword)
	users.Delete("/me", handlers.DeleteAccount)
}

func setupProductRoutes(protected fiber.Router) {
	products := protected.Group("/products")
	products.Post("/", middleware.AdminOnly(), handlers.CreateProduct)
	products.Put("/:id", middleware.AdminOnly(), handlers.UpdateProduct)
	products.Delete("/:id", middleware.AdminOnly(), handlers.DeleteProduct)
}

func setupCartRoutes(protected fiber.Router) {
	cart := protected.Group("/cart")
	cart.Get("/", handlers.GetCart)
	cart.Post("/items", handlers.AddToCart)
	cart.Put("/items/:id", handlers.UpdateCartItem)
	cart.Delete("/items/:id", handlers.RemoveFromCart)
	cart.Delete("/", handlers.ClearCart)
}

func setupOrderRoutes(protected fiber.Router) {
	orders := protected.Group("/orders")
	orders.Post("/checkout", handlers.Checkout)
	orders.Get("/", handlers.GetOrders)
	orders.Get("/:id", handlers.GetOrderByID)
	orders.Put("/:id/cancel", handlers.CancelOrder)

	orders.Put("/:id/status", middleware.AdminOnly(), handlers.UpdateOrderStatus)
	orders.Get("/all", middleware.AdminOnly(), handlers.GetAllOrders)
}

func setupStaticFiles(app *fiber.App) {
	app.Use("/", filesystem.New(filesystem.Config{
		Root:         http.Dir("build"),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "index.html",
	}))
}
