package routes

import (
	"egghead/app/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRouter(app *fiber.App, db *gorm.DB) {
	// Create a new router group
	route := app.Group("/private/api", middleware.WithDB(db), middleware.AuthenticationMiddleware)

	// Define routes for all the private endpoints
	SetupUserRouter(route.Group("v1"), db)
	SetupTransactionRouter(route.Group("v1"), db)

}
