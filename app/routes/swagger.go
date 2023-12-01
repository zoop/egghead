package routes

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(app *fiber.App) {
	// Create routes group.
	route := app.Group("/swagger")

	// Routes for GET method:
	// route.Get("*", swagger.HandlerDefault) // get one user by ID
	route.Get(
		"/",
		func(c *fiber.Ctx) error { // Read the JSON file from the file system
			data, err := os.ReadFile("swagger.json")
			if err != nil {
				log.Printf("Error reading JSON file: %v", err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to read JSON file",
				})
			}

			// Set the response content type to JSON.
			c.Set("Content-Type", "application/json")

			// Send the JSON data as the response.
			return c.Send(data)
		})
}
