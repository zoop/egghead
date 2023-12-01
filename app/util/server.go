package util

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShutdown(app *fiber.App) {
	// Create a channel for idle connecions
	idleConnsClosed := make(chan struct{})

	// new goroutine (concurrent execution unit) to handle graceful shutdown.
	go func() {
		sigint := make(chan os.Signal, 1)   // creates a buffered channel with name sigint
		signal.Notify(sigint, os.Interrupt) // registers the channel sigint to receive notifications for the interrupt signal
		<-sigint                            // blocks until a signal is received on the sigint channel. When a signal is received (e.g., user presses Ctrl+C), the subsequent code is executed

		// @TODO: Close the database connection is exists

		// Received an interrupt signal, shutdown
		if err := app.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed) // signaling that the idle connections have been closed
	}()

	// Build Fiber connection URL.
	fiberConnURL, _ := ConnectionURLBuilder("fiber")

	// Run server.
	if err := app.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
