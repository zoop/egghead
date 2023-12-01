package products

import (
	"egghead/app/constants"
	"egghead/app/models"
	"egghead/app/util"

	"github.com/gofiber/fiber/v2"
)

func (pc *ProductController) CreateProduct(c *fiber.Ctx) error {
	// Read and validate data from the request body
	var productData CreateProductRequest
	if err := c.BodyParser(&productData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": constants.InvalidRequestBody})
	}

	// Check for the uniquness of the product
	slug := productData.Slug
	if slug == "" {
		slug = util.GenerateSlug(productData.Name)
	}

	exists, err := pc.ProductService.IsValidProduct(slug)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": constants.SomethingWentWrong})
	}

	if exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": constants.ProductAlreadyExists})
	}

	// Call the service to create the product
	// Prepare the database model
	productModelData := &models.Products{
		Name:     productData.Name,
		Slug:     slug,
		Metadata: productData.Metadadta,
	}
	if err := pc.ProductService.CreateProduct(productModelData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": constants.FailedToCreateProduct})
	}
	// Prepare the response using the ProductResponse schema
	response := ProductResponse{
		UID:       productModelData.UID,
		Slug:      productModelData.Slug,
		Name:      productModelData.Name,
		Metadadta: productModelData.Metadata,
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": constants.ProductCreatedSuccessfully, "result": response})
}

func (pc *ProductController) ListProducts(c *fiber.Ctx) error {
	// Read query parameters for pagination
	page := util.ConvertStrToInt(c.Query("page_no", "1"))
	limit := util.ConvertStrToInt(c.Query("page_size", "10"))
	search := c.Query("search", "")

	// Call the service to retrieve paginated products
	paginatedResult, err := pc.ProductService.ListProducts(search, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": constants.FailedToRetrieveProducts})
	}

	// Return the paginated list of products in the response
	responseData := fiber.Map{
		"products":   paginatedResult.Products,
		"page":       paginatedResult.Page,
		"totalItems": paginatedResult.TotalItems,
		"totalPages": paginatedResult.TotalPages,
	}
	return c.JSON(responseData)
}

func (pc *ProductController) UpdateProduct(c *fiber.Ctx) error {
	// Read and validate data from the request body
	productID := c.Params("productID")
	var updateProductDetails UpdateProductRequest
	if err := c.BodyParser(&updateProductDetails); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": constants.InvalidRequestBody})
	}

	// Call the service to update the product data
	// @TODO need to add a validation for this if no data is sent
	productData := &models.Products{
		Name:     updateProductDetails.Name,
		Metadata: updateProductDetails.Metadadta,
	}
	if err := pc.ProductService.UpdateProduct(productID, productData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": util.GetErrorMessage(err, constants.FailedToUpdateProduct)})
	}

	// return c.JSON(productData)
	return c.JSON(fiber.Map{"message": constants.ProductUpdatedSuccessfully})
}

func (pc *ProductController) ArchiveProduct(c *fiber.Ctx) error {
	// Read and validate data from the request body
	productID := c.Params("productID")

	// Call the service to get product details

	if err := pc.ProductService.DeleteProductByUID(productID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": constants.FailedToDeleteProduct})
	}

	// Return the product details in the response
	return c.JSON(fiber.Map{"message": constants.ProductDeletedSuccessfully})
}

func (pc *ProductController) GetProductDetails(c *fiber.Ctx) error {
	// Read and validate data from the request body
	productID := c.Params("productID")

	// Call the service to get product details
	product, err := pc.ProductService.GetProductByUID(productID)
	if product == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": constants.ProductNotFound})
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": constants.FailedToRetrieveProduct})
	}

	// Return the product details in the response
	return c.JSON(fiber.Map{"product": product})
}
