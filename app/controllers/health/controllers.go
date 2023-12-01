package health

import "github.com/gofiber/fiber/v2"

// CheckHealth is an endpoint to check the health of the application and its dependencies.
func (hc *HealthController) CheckHealth(c *fiber.Ctx) error {
	// Check the health of the database
	if err := hc.HealthService.PingDB(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database is not healthy"})
	}

	return c.JSON(fiber.Map{"status": "ok"})
}
