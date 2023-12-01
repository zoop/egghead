package main

import (
	"egghead/app/config"
	"egghead/app/middleware"
	"egghead/app/models"
	"egghead/app/routes"
	"egghead/app/util"

	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Load configuration
	appConfig := config.LoadConfig()

	log.Println("APP: Application config configured")

	// Initialize database
	dbManager := util.GetDatabaseManager()

	// Construct the basic connection string
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		appConfig.DatabaseHost, appConfig.DatabasePort, appConfig.DatabaseUser, appConfig.DatabasePassword, appConfig.DatabaseName,
	)

	// Add SSL parameters to the connection string if SSL is configured
	if appConfig.DatabaseSSLMode != "disable" {
		connectionString += fmt.Sprintf(" sslmode=%s", appConfig.DatabaseSSLMode)
		// If SSL root certificate is provided, add it to the connection string
		if appConfig.DatabaseServerCA != "" {
			connectionString += fmt.Sprintf(" sslrootcert=%s", appConfig.DatabaseServerCA)
		}
		// If SSL client certificate and key are provided, add them to the connection string
		if appConfig.DatabaseClientCA != "" && appConfig.DatabaseClientKey != "" {
			connectionString += fmt.Sprintf(" sslcert=%s sslkey=%s", appConfig.DatabaseClientCA, appConfig.DatabaseClientKey)
		}
	} else {
		// Log a message when SSL is not configured
		log.Println("APP: Could not find the values for SSL of database")
	}

	db, err := dbManager.InitDB(connectionString, models.Products{}, models.Users{}, models.TransactionHistory{})
	if err != nil {
		log.Fatalf("APP: Failed to initialize database: %v", err)
	}

	// Set up the middleware
	middleware.SetupMiddleware(app, db)

	// Set up the performance monitor
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Egghead Performance Monitor"}))

	// Set up routes
	routes.SetupRoutes(app, db)

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Use a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Increment the WaitGroup
	wg.Add(1)

	// Start a goroutine for graceful shutdown
	go func() {
		<-c
		log.Println("APP: Received interrupt. Starting graceful shutdown")

		// Close all open resources
		dbManager.CloseDB()
		// defer wg.Done()

		// Wait for existing requests to finish before shutting down the server
		// wg.Wait()

		// Shutdown the Fiber app
		if err := app.Shutdown(); err != nil {
			log.Printf("APP: Error during shutdown: %v\n", err)
		}

		log.Println("APP: Server stopped")
		wg.Done()
	}()

	// Start the server
	go func() {
		defer wg.Done()
		log.Printf("APP: Server listening on :%s", appConfig.ServerPort)
		if err := app.Listen(":" + appConfig.ServerPort); err != nil {
			log.Fatalf("APP: Error while starting server: %v", err)
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
