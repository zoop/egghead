package transaction

import (
	"egghead/app/services"

	"gorm.io/gorm"
)

// TransactionController represents the controller for transaction-related actions.
type TransactionController struct {
	TransactionService *services.TransactionService
}

// NewTransactionController creates a new instance of TransactionController.
func NewTransactionController(db *gorm.DB) *TransactionController {
	us := services.NewUserService(db)
	return &TransactionController{
		TransactionService: services.NewTransactionService(db, us),
	}
}
