package services

import (
	"egghead/app/models"
	"egghead/app/util"
	"errors"

	"github.com/rs/xid"
	"gorm.io/gorm"
)

// ProductService represents the service for handling products.
type ProductService struct {
	DB *gorm.DB
}

// NewProductService creates a new ProductService instance.
func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{DB: db}
}

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(product *models.Products) error {
	// product.UID = uuid.New().String()
	if product.Slug == "" {
		product.Slug = util.GenerateSlug(product.Name)
	}
	product.UID = xid.New().String()
	return s.DB.Create(product).Error
}

// ProductDetail retrieves a product by its ID
func (s *ProductService) ProductDetail(uid string) (*models.Products, error) {
	var product models.Products
	err := s.DB.First(&product, uid).Error
	return &product, err
}

// UpdateProdict updates a product by its ID
func (s *ProductService) UpdateProduct(uid string, updatedProduct *models.Products) error {
	// Set the ID of the updated product
	updatedProduct.UID = uid

	// Update the product in the database
	return s.DB.Model(&models.Products{}).Where("uid = ?", uid).Updates(updatedProduct).Error
}

// DeleteProduct deletes a product by its ID
func (s *ProductService) DeleteProduct(uid uint) error {
	return s.DB.Delete(&models.Products{}, uid).Error
}

// DeleteProductByUID deletes a product by UID
func (s *ProductService) DeleteProductByUID(uid string) error {
	// product, err := s.GetProductByUID(uid)
	// if product == nil {
	// 	return errors.New("Product is already deleted")
	// }

	// if err != nil {
	// 	return err
	// }

	// return s.DeleteProduct(product.ID)

	return s.DB.Where("uid = ?", uid).Unscoped().Delete(&models.Products{}).Error
	// return s.DB.Unscoped().Delete(&models.Products{UID: uid}).Error
}

// ListProducts retrieves a list of all products.
func (s *ProductService) ListProducts(search string, page int, limit int) (ProductPaginateResult, error) {
	var products []models.Products
	var totalItems int64
	// var filter map[string]interface{} = nil
	// if search != "" {
	// 	cleanedSearchQuery := util.CleanString(search)
	// 	filter = map[string]interface{}{
	// 		"Name ILIKE": "%" + cleanedSearchQuery + "%",
	// 	}
	// }

	// // Calculate offset for pagination
	offset := (page - 1) * limit

	// // Perform paginated query
	if err := s.DB.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return ProductPaginateResult{}, err
	}

	// // Count total number of items for pagination
	if err := s.DB.Model(&models.Products{}).Count(&totalItems).Error; err != nil {
		return ProductPaginateResult{}, err
	}

	// models.FindAndCount(s.DB, models.Products{}, filter, page, limit)

	// Calculate page response
	pageResponse, err := util.GetPageResponse(int(totalItems), page, limit)
	if err != nil {
		return ProductPaginateResult{}, err
	}

	// Create and return paginated result
	result := ProductPaginateResult{
		PaginatedResult: pageResponse,
		Products:        products,
	}

	return result, nil
}

// GetProductByID retrieves a product by its ID.
func (s *ProductService) GetProductByID(productID uint) (*models.Products, error) {
	var product models.Products
	result := s.DB.First(&product, productID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

// GetProductByUID retrieves a product by its ID.
func (s *ProductService) GetProductByUID(productID string) (*models.Products, error) {
	var product models.Products
	result := s.DB.Where("uid = ?", productID).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

// GetProductBySlug retrieves a product by its slug.
func (s *ProductService) GetProductBySlug(productSlug string) (*models.Products, error) {
	var product models.Products
	err := s.DB.Where("uid = ?", productSlug).First(&product).Error
	if err == nil {
		// If a product with the same slug already exists, return an error
		return &product, errors.New("product with the same slug already exists")
	}
	return &product, nil
}

// GetProductBySlug retrieves a product by its slug.
func (s *ProductService) GetProductByName(productName string) (*models.Products, error) {
	var product models.Products
	err := s.DB.Where("name = ?", productName).First(&product).Error
	if err == nil {
		// If a product with the same slug already exists, return an error
		return &product, errors.New("product with the same slug already exists")
	}
	return &product, nil
}

// IsValidProduct checks if a product with the given ID exists in the database.
func (s *ProductService) IsValidProduct(productSlug string) (bool, error) {
	var product models.Products
	err := s.DB.Where("slug = ?", productSlug).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Product with the given ID does not exist
			return false, nil
		}
		// Other database error occurred
		return false, err
	}

	// Product with the given ID exists
	return true, nil
}

// IsValidProductID checks if the given product ID is valid.
func (s *ProductService) IsValidProductID(productID string) bool {
	var product models.Products
	result := s.DB.Where("uid = ?", productID).First(&product)
	return result.Error == nil
}
