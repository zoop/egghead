package util

import (
	"errors"
	"fmt"
	"log"
	"time"

	postgres "go.elastic.co/apm/module/apmgormv2/driver/postgres"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DatabaseManager is a singleton struct to manage the database connection and migrations
type DatabaseManager struct {
	DB *gorm.DB
}

var dbManagerInstance *DatabaseManager

// GetDatabaseManagerInstance returns a singleton instance of DatabaseManager
func GetDatabaseManager() *DatabaseManager {
	if dbManagerInstance == nil {
		dbManagerInstance = &DatabaseManager{}
	}
	return dbManagerInstance
}

// Migrate migrate the database model to the database tables
func (dm *DatabaseManager) Migrate(models ...interface{}) error {
	err := dm.DB.AutoMigrate(models)
	if err != nil {
		return fmt.Errorf("DB: Failed to migrate the database: %v", err)
	}

	log.Println("DB: GORM migrations completed")
	return nil
}

// InitDB initializes the database connection
func (dm *DatabaseManager) InitDB(connectionString string, models ...interface{}) (*gorm.DB, error) {
	var err error
	dm.DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("DB: Failed to initialize the database: %v", err)
	}

	// Register Elastic APM callbacks for Gorm
	// dm.DB = dm.DB.WithContext(context.TODO())
	// apmgormv2.RegisterCallbacks(dm.DB)

	// Perform GORM migrations
	if err = dm.Migration(models...); err != nil {
		return nil, fmt.Errorf("DB: Failed while migration: %v", err)
	}

	sqlDB, err := dm.DB.DB()

	if err != nil {
		return nil, fmt.Errorf("DB: Failed to fetch the database: %v", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("DB: Database Intialization completed")

	return dm.DB, nil
}

// GetDB returns the GORM DB instance
func (dm *DatabaseManager) GetDB() (*gorm.DB, error) {
	if dm.DB != nil {
		return dm.DB, nil
	}
	return nil, errors.New("DB: Database connection not initialized")
}

// CloseDB closes the database connection
func (dm *DatabaseManager) CloseDB() {
	if dm.DB != nil {
		sqlDB, _ := dm.DB.DB()
		if sqlDB != nil {
			sqlDB.Close()
			log.Println("DB: Closed the connection")
		}
	}
}

// PingDB checks the status of the database connection
func (dm *DatabaseManager) PingDB() error {
	db, err := dm.GetDB()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Printf("DB: Failed to ping the database: %v", err)
		return fmt.Errorf("DB: Failed to ping the database: %v", err)
	}
	log.Println("DB: Database connection is active")
	return nil

}

// Migrate migrates the database model to the database tables
func (dm *DatabaseManager) Migration(models ...interface{}) error {
	err := dm.DB.AutoMigrate(models...)
	if err != nil {
		return fmt.Errorf("DB: Failed to migrate the database: %v", err)
	}

	log.Println("DB: GORM migrations completed")
	return nil
}
