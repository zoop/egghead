package middleware

import (
	"egghead/app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthenticationMiddleware(c *fiber.Ctx) error {
	productID := c.Get("x-product-id")

	if productID == "" {
		// If the x-product-id header is missing, return 401 Unauthorized
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "x-product-id header is missing",
		})
	}

	// Validate the productID against the database
	if product, _ := services.NewProductService(c.Locals("db").(*gorm.DB)).GetProductByUID(productID); product != nil {
		c.Locals("product", product)
		return c.Next()
	}

	// If the product ID is not valid, return 401 Unauthorized
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   "Unauthorized",
		"message": "Invalid x-product-id",
	})

}

// func isValidProductID(productID string, db *gorm.DB) bool {
// 	var product models.Products
// 	result := db.Where("uid = ?", productID).First(&product)
// 	return result.Error == nil
// }
