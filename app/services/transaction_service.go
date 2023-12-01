package services

import (
	"egghead/app/constants"
	"egghead/app/models"
	"errors"

	"github.com/rs/xid"
	"gorm.io/gorm"
)

// TransactionService represents the sercide for handling transaction.
type TransactionService struct {
	DB *gorm.DB
	US *UserService
}

// NewTransactionService creates a new TransactionService instance.
func NewTransactionService(db *gorm.DB, us *UserService) *TransactionService {
	return &TransactionService{DB: db, US: us}
}

// CreditAmount credits an amount into the user account
func (s *TransactionService) CreditAmount(userId uint, transaction *models.TransactionHistory) error {
	tx := s.DB.Begin()

	// Get the latest balance of the user
	var userBalance float64
	if err := tx.Model(&models.Users{}).Where("id = ?", userId).Pluck("balance", &userBalance).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update the user's balance in the Users table
	if err := tx.Model(&models.Users{}).Where("id = ?", userId).Update("balance", gorm.Expr("balance + ?", transaction.Amount)).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Insert a new transaction record in the TransactionHistory table
	transaction.Balance = userBalance + transaction.Amount
	transaction.UID = xid.New().String()
	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// DebitAmount debits an amoutn from the user account
func (s *TransactionService) DebitAmount(userId uint, transaction *models.TransactionHistory) error {
	tx := s.DB.Begin()
	// Retrieve the user's current balance
	var user models.Users
	if err := tx.Where("id = ?", userId).First(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Check if the user's balance is sufficient
	if user.Balance < transaction.Amount {
		tx.Rollback()
		return errors.New(constants.InsufficientFunds)
	}

	// Update the user's balance in the Users table
	if err := tx.Model(&models.Users{}).Where("id = ?", userId).Update("balance", gorm.Expr("balance - ?", transaction.Amount)).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Insert a new transaction record in the TransactionHistory table
	transaction.Balance = user.Balance - transaction.Amount
	transaction.UID = xid.New().String()
	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
func (s *TransactionService) TransactionDetail(userID uint, transactionID string) (*models.TransactionHistory, error) {
	var transaction models.TransactionHistory

	// Assuming you have a proper relationship between Users and Transactions
	err := s.DB.Where("user_id = ? AND uid = ?", userID, transactionID).First(&transaction).Error
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

// GetBalance retrieves the current balance for a user
func (s *TransactionService) GetBalance(userID uint) (float64, error) {
	var user models.Users

	err := s.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return 0.0, err
	}

	return user.Balance, nil
}

// ListTransactions retrieves a list of transactions for a user
func (s *TransactionService) ListTransactions(userID uint) ([]models.TransactionHistory, error) {
	var transactions []models.TransactionHistory

	err := s.DB.Where("user_id = ?", userID).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
