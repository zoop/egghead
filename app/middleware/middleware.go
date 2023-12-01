package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.elastic.co/apm/module/apmfiber/v2"
	"gorm.io/gorm"
)

func SetupMiddleware(app *fiber.App, db *gorm.DB) error {
	// Custom Middleware
	// app.Use(
	// 	func(c *fiber.Ctx) error {
	// 		c.Locals("db", db)
	// 		return c.Next()
	// 	},
	// )

	// Middleware for JSON conversion
	app.Use(JSONMiddleware)

	// Middleware for CORS
	app.Use(cors.New())

	// Middleware for logging
	// app.Use(LoggingMiddleware)
	app.Use(logger.New())

	// Sample authentication middleware (replace it with your actual auth logic)
	// app.Use(AuthenticationMiddleware, WithDB(db))

	// Error Handler
	app.Use(RecoverErrorHandler)

	// Elastic APM middleware to track metrics
	app.Use(apmfiber.Middleware())

	log.Println("APP: Middleware setup is completed")

	return nil
}

func WithDB(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	}
}
