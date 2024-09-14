package userhandler

import (
	userapplication "github.com/citywalker-app/go-api/pkg/user/application"
	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
	"github.com/citywalker-app/go-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user userdomain.User
		if err := c.BodyParser(&user); err != nil {
			return utils.NewErrorHandler(c, ErrBadRequest, fiber.StatusBadRequest)
		}

		if err := validator.New().StructPartial(&user, "Email", "Password"); err != nil {
			return utils.NewErrorHandler(c, ErrBadRequest, fiber.StatusBadRequest)
		}

		err := userapplication.Register(&user)
		if err != nil {
			return utils.NewErrorHandler(c, err, fiber.StatusBadRequest)
		}

		token, err := utils.GenerateJWT(user.Email)
		if err != nil {
			return utils.NewErrorHandler(c, ErrJWTGeneration, fiber.StatusInternalServerError)
		}

		response := Response{JWT: token, User: user}

		return utils.NewSuccessHandler(c, response)
	}
}
