package server

import (
	"github.com/citywalker-app/go-api/middleware"
	cityrouter "github.com/citywalker-app/go-api/pkg/city/infrastructure/router"
	travelrouter "github.com/citywalker-app/go-api/pkg/travel/infrastructure/router"
	userrouter "github.com/citywalker-app/go-api/pkg/user/infrastructure/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/metrics", monitor.New())

	auth := app.Group("/user")
	userrouter.Router(auth)

	cities := app.Group("/cities", middleware.JWTHandler(), middleware.RedisHandler())
	cityrouter.Router(cities)

	travels := app.Group("/travels", middleware.JWTHandler())
	travelrouter.Router(travels)
}
