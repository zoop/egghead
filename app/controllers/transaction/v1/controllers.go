package transaction

import (
	"egghead/app/constants"
	"egghead/app/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (tc *TransactionController) CreditAmount(c *fiber.Ctx) error {
	// Read and validate data from the request body
	// userID := util.ConvertStrToInt(c.Params("userID"))
	userID := c.Params("userID")

	var transactionData models.TransactionHistory
	if err := c.BodyParser(&transactionData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": constants.InvalidCreditRequest})
	}

	// Set the transaction type to CREDIT
	transactionData.TransactionType = "CREDIT"

	// Get the product from the context
	product, ok := c.Locals("product").(*models.Products)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": constants.SomethingWentWrong})
	}
	transactionData.ProductID = int(product.ID)

	if transactionData.Reason == "" {
		// Set a default reason if it's not provided
		transactionData.Reason = fmt.Sprintf("%f INR amount credited", transactionData.Amount)
	}

	// Call the service to credit an amount
	userData, err := tc.TransactionService.US.GetUserByUID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": constants.UserNotFound})
	}

	transactionData.UserID = int(userData.ID)

	if err := tc.TransactionService.CreditAmount(userData.ID, &transactionData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": constants.FailedToCreditAmount})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": constants.TransactionCreditedSuccessfully, "result": transactionData})
}

func (tc *TransactionController) DebitAmount(c *fiber.Ctx) error {
	// Read and validate data from the request body
	// userID := util.ConvertStrToInt(c.Params("userID"))
	userID := c.Params("userID")

	var transactionData models.TransactionHistory
	if err := c.BodyParser(&transactionData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": constants.InvalidDebitRequest})
	}

	// Set the transaction type to DEBIT
	transactionData.TransactionType = "DEBIT"

	// Get the product from the context
	product, ok := c.Locals("product").(*models.Products)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": constants.SomethingWentWrong})
	}
	transactionData.ProductID = int(product.ID)

	if transactionData.Reason == "" {
		// Set a default reason if it's not provided
		transactionData.Reason = fmt.Sprintf("%f INR amount debited", transactionData.Amount)
	}

	// Call the service to debit an amount
	userData, err := tc.TransactionService.US.GetUserByUID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": constants.UserNotFound})
	}
	transactionData.UserID = int(userData.ID)

	if userData.Balance < transactionData.Amount {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": constants.InsufficientFunds})
	}

	if err := tc.TransactionService.DebitAmount(userData.ID, &transactionData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": constants.FailedToDebitAmount})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": constants.TransactionDebitedSuccessfully, "result": transactionData})
}

func (tc *TransactionController) ListTransactions(c *fiber.Ctx) error {
	// Read and validate user ID from the request params
	// userID := util.ConvertStrToInt(c.Params("userID"))
	userID := c.Params("userID")

	// Call the service to get the list of transactions
	userData, err := tc.TransactionService.US.GetUserByUID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": constants.UserNotFound})
	}

	transactions, err := tc.TransactionService.ListTransactions(userData.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": constants.FailedToListTransactions})
	}

	// Return the list of transactions in the response
	return c.JSON(fiber.Map{"result": transactions})
}

func (tc *TransactionController) GetBalance(c *fiber.Ctx) error {
	// Read and validate user ID from the request params
	// userID := util.ConvertStrToInt(c.Params("userID"))
	userID := c.Params("userID")

	// Call the service to get the user's balance
	userData, err := tc.TransactionService.US.GetUserByUID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": constants.UserNotFound})
	}

	// balance, err := tc.TransactionService.GetBalance(userData.ID)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": constants.FailedToRetrieveBalance})
	// }

	// Return the user's balance in the response
	return c.JSON(fiber.Map{"balance": userData.Balance})
}

func (tc *TransactionController) TransactionDetail(c *fiber.Ctx) error {
	// Read and validate user ID and transaction ID from the request params
	// userID := util.ConvertStrToInt(c.Params("userID"))
	userID := c.Params("userID")
	// transactionID := util.ConvertStrToInt(c.Params("transactionID"))
	transactionID := c.Params("transactionID")

	// Call the service to get details of a specific transaction
	userData, err := tc.TransactionService.US.GetUserByUID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": constants.UserNotFound})
	}

	transactionDetail, err := tc.TransactionService.TransactionDetail(userData.ID, transactionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": constants.FailedToRetrieveTransaction})
	}

	// Return the transaction details in the response
	return c.JSON(fiber.Map{"result": transactionDetail})
}
