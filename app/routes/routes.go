package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Route for all the health endpoints
	HealthCheckRoute(app, db)
	SwaggerRoute(app)

	// Route for all the internal endpoints for inter-service comms
	InternalRoute(app, db)

	// Route for all the private endpoints accessible within the cluster
	PrivateRouter(app, db)

	NotFoundRoute(app)

	log.Println("APP: Routes setup is completed")
}
