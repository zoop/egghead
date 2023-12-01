package users

import (
	"egghead/app/constants"
	"egghead/app/models"

	"github.com/gofiber/fiber/v2"
)

// RegisterUser registers a new user for a specific product.
func (uc *UserController) RegisterUser(c *fiber.Ctx) error {
	// Read and validate data from the request body
	// productID := c.Params("productID")
	userData := new(models.Users)
	if err := c.BodyParser(userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": constants.InvalidUserData})
	}

	// Call the product service to validate the product
	// product, err := pc.ProductService.GetProductByUID(productID)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": util.GetErrorMessage(err, constants.FailedToRetrieveProduct)})
	// }

	// Call the service to register the user
	exists := uc.UserService.IsValidUserID(userData.UID)
	if exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": constants.UserAlreadyExists})
	}
	// Get the product from the context
	product, ok := c.Locals("product").(*models.Products)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": constants.SomethingWentWrong})
	}

	userData.ProductID = int(product.ID)
	if err := uc.UserService.RegisterUser(product.ID, userData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": constants.FailedToCreateUser})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": constants.UserCreatedSuccessfully, "user": userData})
}

// ArchiveUser archives a user for a specific product.
func (uc *UserController) ArchiveUser(c *fiber.Ctx) error {
	// Read productID and userID from the request params
	// productID := c.Params("productID")
	userID := c.Params("userID")

	// Get the product from the context
	product, ok := c.Locals("product").(*models.Products)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": constants.SomethingWentWrong})
	}

	// Call the service to archive the user
	if err := uc.UserService.ArchiveUser(product.ID, userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": constants.FailedToDeleteUser})
	}

	return c.JSON(fiber.Map{"message": constants.UserDeletedSuccessfully})
}
