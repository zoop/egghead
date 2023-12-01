package middleware

import (
	"egghead/app/constants"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	message string
	trace   interface{}
}

func ErrorHandleMiddleware(c *fiber.Ctx, err interface{}, debug bool) {
	// Default unhandled error code
	defaultCode := fiber.StatusInternalServerError
	errorResponse := fiber.Map{"message": constants.InternalServerError}

	if debug {
		errorResponse = fiber.Map{
			"message": constants.InternalServerError,
			"trace":   err,
		}
	}

	c.Status(defaultCode).JSON(errorResponse)
}

// Handle the unhandler exception
func RecoverErrorHandler(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			ErrorHandleMiddleware(c, r, true)
		}
	}()

	// Next middleware or route handler
	return c.Next()
}
