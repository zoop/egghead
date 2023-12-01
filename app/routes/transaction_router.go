package routes

import (
	"egghead/app/controllers/transaction/v1"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupTransactionRouter(route fiber.Router, db *gorm.DB) {
	//Initialize the user controller
	transactionController := transaction.NewTransactionController(db)

	// Define rotues for the user controller
	route.Get("/user/:userID/balance", transactionController.GetBalance)
	route.Post("/user/:userID/credit", transactionController.CreditAmount)
	route.Post("/user/:userID/debit", transactionController.DebitAmount)
	route.Get("/user/:userID/transactions", transactionController.ListTransactions)
	route.Get("/user/:userID/transactions/:transactionID", transactionController.TransactionDetail)
}
