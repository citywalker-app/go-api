package userhandler

import (
	userapplication "github.com/citywalker-app/go-api/pkg/user/application"
	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
	"github.com/citywalker-app/go-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ConfirmCode() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(userdomain.User)
		if err := c.BodyParser(&user); err != nil {
			return utils.NewErrorHandler(c, ErrBadRequest, fiber.StatusBadRequest)
		}

		if err := validator.New().Var(user.Email, "required,email"); err != nil {
			return utils.NewErrorHandler(c, ErrBadRequest, fiber.StatusBadRequest)
		}

		if user.FullName == "" {
			if err := userapplication.Get(user); err != nil {
				return utils.NewErrorHandler(c, ErrUserNotFound, fiber.StatusNotFound)
			}
		}

		code, err := utils.GenerateRandomCode()
		if err != nil {
			return utils.NewErrorHandler(c, err, fiber.StatusInternalServerError)
		}

		utils.SetCache(c.Context(), user.Email, code)

		html := `<!DOCTYPE html>
						<html>
							<head>
								<title>City Walker code</title>
							</head>
							<body>
								<p>Your City Walker code is <strong>` + code + `</strong></p>
							</body>
						</html>`

		if err := utils.SendEmail(user.Email, "City Walker code", html); err != nil {
			return utils.NewErrorHandler(c, ErrEmail, fiber.StatusInternalServerError)
		}

		return utils.NewSuccessHandler(c, map[string]interface{}{"confirmCode": code})
	}
}
