package routes

import (
	"egghead/app/controllers/health"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// HealthCheckRoute sets up the routes for the HealthController.
func HealthCheckRoute(app *fiber.App, db *gorm.DB) {
	// Create routes group.
	route := app.Group("/health")

	healthController := health.NewHealthController(db)

	// Routes for GET method:
	route.Get("/liveness", healthController.CheckHealth)
	route.Get("/readiness", healthController.CheckHealth)
	route.Get("/startup", healthController.CheckHealth)
}
