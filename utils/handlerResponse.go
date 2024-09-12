package utils

import "github.com/gofiber/fiber/v2"

func NewErrorHandler(c *fiber.Ctx, err error, status int) error {
	return c.Status(status).JSON(fiber.Map{
		"error":   true,
		"message": err.Error(),
	})
}

func NewSuccessHandler(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(data)
}
