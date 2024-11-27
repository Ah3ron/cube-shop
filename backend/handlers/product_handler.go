package handlers

import (
	"cube-shop/database"
	"cube-shop/models"
	"github.com/gofiber/fiber/v2"
)

// GetProducts возвращает список всех продуктов с возможностью фильтрации
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	query := database.DB.Model(&models.Product{})

	// Фильтрация по категории
	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}

	// Фильтрация по сложности
	if difficulty := c.Query("difficulty"); difficulty != "" {
		query = query.Where("difficulty = ?", difficulty)
	}

	// Фильтрация по бренду
	if brand := c.Query("brand"); brand != "" {
		query = query.Where("brand = ?", brand)
	}

	// Сортировка
	if sort := c.Query("sort"); sort != "" {
		switch sort {
		case "price_asc":
			query = query.Order("price asc")
		case "price_desc":
			query = query.Order("price desc")
		case "name_asc":
			query = query.Order("name asc")
		case "name_desc":
			query = query.Order("name desc")
		}
	}

	result := query.Find(&products)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось получить список продуктов",
		})
	}

	return c.JSON(products)
}

// GetProduct возвращает информацию о конкретном продукте
func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	result := database.DB.First(&product, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Продукт не найден",
		})
	}

	return c.JSON(product)
}

// GetCategories возвращает список всех категорий
func GetCategories(c *fiber.Ctx) error {
	var categories []string
	result := database.DB.Model(&models.Product{}).Distinct().Pluck("category", &categories)
	
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось получить категории",
		})
	}

	return c.JSON(categories)
}

// SearchProducts выполняет поиск по продуктам
func SearchProducts(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Поисковый запрос не может быть пустым",
		})
	}

	var products []models.Product
	searchQuery := "%" + query + "%"
	
	result := database.DB.Where(
		"name ILIKE ? OR description ILIKE ? OR category ILIKE ? OR brand ILIKE ?",
		searchQuery, searchQuery, searchQuery, searchQuery,
	).Find(&products)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка при выполнении поиска",
		})
	}

	return c.JSON(products)
}

// CreateProduct создает новый продукт (только для админов)
func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}

	// Валидация
	if product.Name == "" || product.Price <= 0 || product.Stock < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверные данные продукта",
		})
	}

	// Устанавливаем значение по умолчанию для ImageURL, если оно не указано
	if product.ImageURL == "" {
		product.ImageURL = "https://placehold.co/400/141414/FFFFFF/png"
	}

	result := database.DB.Create(&product)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось создать продукт",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

// UpdateProduct обновляет существующий продукт (только для админов)
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}

	result := database.DB.Model(&models.Product{}).Where("id = ?", id).Updates(product)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось обновить продукт",
		})
	}

	// Получаем обновленный продукт
	var updatedProduct models.Product
	database.DB.First(&updatedProduct, id)

	return c.JSON(updatedProduct)
}

// DeleteProduct удаляет продукт (только для админов)
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	result := database.DB.Delete(&models.Product{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось удалить продукт",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
} 