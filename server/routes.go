package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/metrics", monitor.New())

	// auth := app.Group("/user")
	// userrouter.Router(auth)

	// cities := app.Group("/cities", middleware.JWTHandler(), middleware.RedisHandler())
	// cityrouter.Router(cities)

	// travels := app.Group("/travels", middleware.JWTHandler())
	// travelRouter.Router(travels)
}
