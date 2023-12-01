package routes

import (
	"egghead/app/controllers/products/v1"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupProductRoutes(route fiber.Router, db *gorm.DB) {
	// Initialize the product controller
	productController := products.NewProductController(db)

	// Define routes for the product controller
	route.Get("/products", productController.ListProducts)
	route.Post("/products", productController.CreateProduct)
	route.Get("/product/:productID", productController.GetProductDetails)
	route.Put("/product/:productID", productController.UpdateProduct)
	route.Delete("/product/:productID", productController.ArchiveProduct)

}
