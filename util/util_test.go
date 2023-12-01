package tests

// import (
// 	"egghead/app/config"
// 	"egghead/app/util"
// 	"log"
// 	"os"
// 	"testing"
// )

// var testDBConnectionString = "user=youruser dbname=testdb sslmode=disable"
// var dbManager *util.DatabaseManager

// func TestMain(m *testing.M) {
// 	// Set up test environment
// 	setup()

// 	// Run tests
// 	exitCode := m.Run()

// 	// Tear down test environment
// 	teardown()

// 	// Exit with the status code from tests
// 	os.Exit(exitCode)
// }

// func setup() {
// 	// Load test configurations
// 	// You might use a test-specific configuration file for this purpose
// 	// For simplicity, using a direct connection string here
// 	config := config.LoadConfig()

// 	// create the database connection string

// 	// Initialise the database manager
// 	dbManager := util.GetDatabaseManager()

// 	// Perform migrations
// 	err := dbManager.Migrate()
// 	if err != nil {
// 		log.Printf("Error migrating database: %v", err)
// 	}

// 	// Initialize the database
// 	_, err = dbManager.InitDB(config.DatabaseHost, []interface{}{})
// 	if err != nil {
// 		log.Printf("Error initializing database: %v", err)
// 	}

// }

// func teardown() {
// 	// Close the database connection
// 	dbManager := util.GetDatabaseManager()
// 	dbManager.CloseDB()
// }

// func TestDatabaseManager_Migrate(t *testing.T) {
// 	// Mock models for testing
// 	type User struct {
// 		ID   uint
// 		Name string
// 	}

// 	// Perform migration
// 	err := dbManager.Migrate(&User{})
// 	if err != nil {
// 		t.Errorf("Error migrating database: %v", err)
// 	}

// 	// Check if the migration was successful by querying the database schema
// 	var userTableExists bool
// 	dbManager.DB.Migrator().HasTable(&User{})
// 	if userTableExists {
// 		t.Log("User table migration successful")
// 	} else {
// 		t.Error("User table migration failed")
// 	}
// }

// func TestDatabaseManager_InitDB(t *testing.T) {
// 	// Set up a test database connection string
// 	// (This will use the connection string from TestMain)
// 	connectionString := testDBConnectionString

// 	// Initialize the database
// 	_, err := dbManager.InitDB(connectionString, []interface{}{})
// 	if err != nil {
// 		t.Errorf("Error initializing database: %v", err)
// 	}

// 	// TODO: Perform additional tests based on your application requirements
// }

// func TestDatabaseManager_PingDB(t *testing.T) {
// 	// Perform a ping to check the database connection
// 	err := dbManager.PingDB()
// 	if err != nil {
// 		t.Errorf("Error pinging database: %v", err)
// 	}

// 	// TODO: Check other conditions based on your application requirements
// }
