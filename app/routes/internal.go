package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InternalRoute(app *fiber.App, db *gorm.DB) {
	// Create routes group.
	route := app.Group("/internal/api")

	// Define routes for all the internal endpoints
	SetupProductRoutes(route.Group("/v1"), db)

}
