package server

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/citywalker-app/go-api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Setup() *fiber.App {
	config := Config()
	app := fiber.New(config)
	middleware.AddFiberMiddleware(app)

	SetupRoutes(app)

	return app
}

func StartServerWithGracefulShutdown(app *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := app.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Error("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Build Fiber connection URL.
	fiberConnURL := fmt.Sprintf(
		"%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)

	// Run server.
	if err := app.Listen(fiberConnURL); err != nil {
		log.Error("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) error {
	// Build Fiber connection URL.
	fiberConnURL := fmt.Sprintf(
		"%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)

	// Run server.
	return a.Listen(fiberConnURL)
}
