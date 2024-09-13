package userhandler

import (
	userapplication "github.com/citywalker-app/go-api/pkg/user/application"
	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
	"github.com/citywalker-app/go-api/utils"
	"github.com/gofiber/fiber/v2"
)

func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user userdomain.User

		if err := c.BodyParser(&user); err != nil {
			return utils.NewErrorHandler(c, ErrBadRequest, fiber.StatusBadRequest)
		}

		if err := userapplication.Login(&user); err != nil {
			return utils.NewErrorHandler(c, err, fiber.StatusUnauthorized)
		}

		token, err := utils.GenerateJWT(user.Email)
		if err != nil {
			return utils.NewErrorHandler(c, ErrJWTGeneration, fiber.StatusInternalServerError)
		}

		return utils.NewSuccessHandler(c, map[string]interface{}{"jwt": token, "user": user})
	}
}
