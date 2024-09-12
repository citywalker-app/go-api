package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func AddFiberMiddleware(app *fiber.App) {
	app.Use(
		// Add CORS to each route
		cors.New(),
		logger.New(),
		healthcheck.New(),
		helmet.New(),
		limiter.New(),
		favicon.New(),
	)
}
