package userhandler

import (
	userapplication "github.com/citywalker-app/go-api/pkg/user/application"
	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
	"github.com/citywalker-app/go-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sethvargo/go-password/password"
)

func ContinueWithGoogle() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user userdomain.User
		if err := c.BodyParser(&user); err != nil {
			return utils.NewErrorHandler(c, ErrBadRequest, fiber.StatusBadRequest)
		}

		if err := validator.New().Var(user.Email, "required,email"); err != nil {
			return utils.NewErrorHandler(c, ErrBadRequest, fiber.StatusBadRequest)
		}

		if err := userapplication.Get(&user); err != nil {
			pass, err := password.Generate(10, 8, 2, false, false)
			if err != nil {
				return utils.NewErrorHandler(c, err, fiber.StatusInternalServerError)
			}

			user.Password = pass

			user.InitializeUser()

			if err := userapplication.Register(&user); err != nil {
				return utils.NewErrorHandler(c, err, fiber.StatusBadRequest)
			}
		}

		token, err := utils.GenerateJWT(user.Email)

		if err != nil {
			return utils.NewErrorHandler(c, ErrJWTGeneration, fiber.StatusInternalServerError)
		}

		response := Response{JWT: token}

		return utils.NewSuccessHandler(c, response)
	}
}
