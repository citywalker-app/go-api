package travelrouter

import (
	travelhandler "github.com/citywalker-app/go-api/pkg/travel/infrastructure/handler"
	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Post("/create", travelhandler.Create())
}
