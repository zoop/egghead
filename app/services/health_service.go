package services

import (
	"fmt"

	"gorm.io/gorm"
)

// HealthService represents the service for handling health-related actions.
type HealthService struct {
	DB *gorm.DB
}

// NewHealthService creates a new HealthService instance.
func NewHealthService(db *gorm.DB) *HealthService {
	return &HealthService{DB: db}
}

// PingDB checks the status of the database connection.
func (hs *HealthService) PingDB() error {
	db, err := hs.DB.DB()
	if err != nil {
		return fmt.Errorf("Failed to get database instance: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Failed to ping the database: %v", err)
	}

	return nil
}
