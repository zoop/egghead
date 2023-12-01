package services

import (
	"egghead/app/constants"
	"egghead/app/models"
	"egghead/app/util"
	"errors"

	"github.com/rs/xid"
	"gorm.io/gorm"
)

// UserService represents the service for user related actions
type UserService struct {
	DB *gorm.DB
}

// NewUserService creates a new instance of UserService.
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

// RegisterUser register a user with some base balance amount
func (s *UserService) RegisterUser(productID uint, user *models.Users) error {
	// Start the transaction
	tx := s.DB.Begin()

	// Create the user data
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update the amount if the user has a default balance
	if user.Balance != 0 {
		transaction := &models.TransactionHistory{
			UID:             xid.New().String(),
			UserID:          int(user.ID),
			Amount:          user.Balance,
			ProductID:       int(productID),
			Reason:          "Amount credited",
			TransactionType: "CREDIT",
			Balance:         user.Balance,
		}

		// Insert a new transaction record in the TransactionHistory table
		if err := tx.Create(transaction).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// ArchiveUser deletes a user and all associated transaction history records
func (s *UserService) ArchiveUser(productID uint, userID string) error {
	// Start the transaction
	tx := s.DB.Begin()

	// Fetch the user data to create a transaction record
	user := &models.Users{}
	if err := tx.Model(&models.Users{}).Where("uid = ?", userID).First(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete the user
	// if err := tx.Model(&models.Users{}).Where("uid = ?", userID).Unscoped().Delete(&models.Users{}).Error; err != nil {
	if err := tx.Unscoped().Delete(&models.Users{}, user.ID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete all associated transaction history records
	if err := tx.Where("user_id = ? and product_id = ?", user.ID, productID).Unscoped().Delete(&models.TransactionHistory{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// ListUsers list all the users
func (s *UserService) ListUsers(search string, page int, limit int) (UserPaginateResult, error) {
	var users []models.Users
	var totalItems int64
	var filter map[string]interface{} = nil
	if search != "" {
		cleanedSearchQuery := util.CleanString(search)
		filter = map[string]interface{}{
			"Name ILIKE": "%" + cleanedSearchQuery + "%",
		}
	}

	models.FindAndCount(s.DB, models.Users{}, filter, page, limit)

	// Calculate page response
	pageResponse, err := util.GetPageResponse(int(totalItems), page, limit)
	if err != nil {
		return UserPaginateResult{}, err
	}

	// Create and return paginated result
	result := UserPaginateResult{
		PaginatedResult: pageResponse,
		Users:           users,
	}

	return result, nil
}

// GetUserByUID retrieves a user by its UID.
func (s *UserService) GetUserByUID(userID string) (*models.Users, error) {
	var user models.Users
	result := s.DB.Where("uid = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	if result == nil {
		return nil, errors.New(constants.UserNotFound)
	}
	return &user, nil
}

// IsValidUserID checks if the given userID is valid.
func (s *UserService) IsValidUserID(userID string) bool {
	var user models.Users
	result := s.DB.Where("uid = ?", userID).First(&user)
	return result.Error == nil
}
