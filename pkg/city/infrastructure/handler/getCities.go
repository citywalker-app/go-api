package cityhandler

import (
	cityapplication "github.com/citywalker-app/go-api/pkg/city/application"
	"github.com/citywalker-app/go-api/utils"
	"github.com/gofiber/fiber/v2"
)

func GetCities() fiber.Handler {
	return func(c *fiber.Ctx) error {
		lng := c.Params("lng")

		cities, err := cityapplication.GetCities(&lng)
		if err != nil {
			return utils.NewErrorHandler(c, err, fiber.StatusNotFound)
		}

		utils.SetCache(c.Context(), c.Path(), cities)

		response := Response{Cities: *cities}

		return utils.NewSuccessHandler(c, response)
	}
}
