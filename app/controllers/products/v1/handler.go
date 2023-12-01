package products

import (
	"egghead/app/services"

	"gorm.io/gorm"
)

// ProductController represents the controller for product-related actions.
type ProductController struct {
	ProductService *services.ProductService
}

// NewProductController creates a new instance of ProductController.
func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		ProductService: services.NewProductService(db),
	}
}
