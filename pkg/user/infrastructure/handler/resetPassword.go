package userhandler

import (
	userapplication "github.com/citywalker-app/go-api/pkg/user/application"
	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
	"github.com/citywalker-app/go-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ResetPassword() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user userdomain.User
		if err := c.BodyParser(&user); err != nil {
			return utils.NewErrorHandler(c, ErrBadRequest, fiber.StatusBadRequest)
		}

		if err := validator.New().StructPartial(&user, "Email", "Password"); err != nil {
			return utils.NewErrorHandler(c, ErrBadRequest, fiber.StatusBadRequest)
		}

		if err := userapplication.ResetPassword(&user); err != nil {
			return utils.NewErrorHandler(c, err, fiber.StatusBadRequest)
		}

		return utils.NewSuccessHandler(c, map[string]interface{}{})
	}
}
