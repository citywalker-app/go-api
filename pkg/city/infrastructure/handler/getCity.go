package cityhandler

import (
	cityapplication "github.com/citywalker-app/go-api/pkg/city/application"
	"github.com/citywalker-app/go-api/utils"
	"github.com/gofiber/fiber/v2"
)

func GetCity() fiber.Handler {
	return func(c *fiber.Ctx) error {
		city := c.Params("city")

		cityFounded, err := cityapplication.GetCity(&city)
		if err != nil {
			return utils.NewErrorHandler(c, err, fiber.StatusNotFound)
		}

		utils.SetCache(c.Context(), c.Path(), cityFounded)

		response := Response{City: *cityFounded}

		return utils.NewSuccessHandler(c, response)
	}
}
