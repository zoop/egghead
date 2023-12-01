package middleware

import "github.com/gofiber/fiber/v2"

func JSONMiddleware(c *fiber.Ctx) error {
	c.Accepts("application/json")
	return c.Next()
}
