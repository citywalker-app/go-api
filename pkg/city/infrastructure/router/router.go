package cityrouter

import (
	cityhandler "github.com/citywalker-app/go-api/pkg/city/infrastructure/handler"
	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Get("/all/:lng", cityhandler.GetCities())
	router.Get("/:city", cityhandler.GetCity())
}
