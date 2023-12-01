package models

import (
	"time"
)

// TransactionType is a custom type for the transaction type enum
type TransactionType string

// TransactionType enum
const (
	Credit TransactionType = "CREDIT"
	Debit  TransactionType = "DEBIT"
	Refund TransactionType = "REFUND"
)

type TransactionHistory struct {
	BaseModel
	// ID              uint                   `gorm:"primaryKey" json:"id"`
	UID             string          `gorm:"type:varchar(255);not null;uniqueIndex:idx_users_product" json:"uid"`
	TransactionType TransactionType `gorm:"varchar(10);index" json:"transaction_type"`
	ProductID       int             `gorm:"not null;index" json:"-"`
	UserID          int             `gorm:"not null;index" json:"-"`
	Amount          float64         `gorm:"not null" json:"amount"`
	Reason          string          `gorm:"type:varchar(255);not null;index" json:"reason"`
	Timestamp       time.Time       `gorm:"default:current_timestamp" json:"timestamp"`
	Metadata        Metadata        `gorm:"type:jsonb" json:"metadata"`
	Balance         float64         `gorm:"not null" json:"balance"`
	// Product         Products
	// User            Users
}
