package travelhandler

import (
	travelapplication "github.com/citywalker-app/go-api/pkg/travel/application"
	traveldomain "github.com/citywalker-app/go-api/pkg/travel/domain"
	userapplication "github.com/citywalker-app/go-api/pkg/user/application"
	"github.com/citywalker-app/go-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		email, ok := c.Locals("email").(string)
		if !ok {
			return utils.NewErrorHandler(c, ErrBadRequest, fiber.StatusInternalServerError)
		}
		var travel traveldomain.Travel

		if err := c.BodyParser(&travel); err != nil {
			return utils.NewErrorHandler(c, ErrBadRequest, fiber.StatusBadRequest)
		}

		if err := validator.New().Struct(&travel); err != nil {
			return utils.NewErrorHandler(c, ErrBadRequest, fiber.StatusBadRequest)
		}

		travelCreated, err := travelapplication.Create(&travel)
		if err != nil {
			return utils.NewErrorHandler(c, err, fiber.StatusNotFound)
		}

		err = userapplication.AddTravel(&travelCreated.ID, &email)
		if err != nil {
			return utils.NewErrorHandler(c, err, fiber.StatusNotFound)
		}

		response := Response{Travel: *travelCreated}

		return utils.NewSuccessHandler(c, response)
	}
}
