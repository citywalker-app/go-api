package userrouter

import (
	userhandler "github.com/citywalker-app/go-api/pkg/user/infrastructure/handler"
	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	router.Post("/login", userhandler.Login())
	router.Post("/register", userhandler.Register())
	router.Post("/resetPassword", userhandler.ResetPassword())
	router.Post("/confirmCode", userhandler.ConfirmCode())
	router.Post("/continueWithGoogle", userhandler.ContinueWithGoogle())
}
