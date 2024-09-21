package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/citywalker-app/go-api/utils"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func RedisHandler() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		path := c.Path()
		cachedResponse, err := utils.RedisDB.Get(context.Background(), path).Bytes()

		if errors.Is(err, redis.Nil) {
			return c.Next()
		} else if err != nil {
			return utils.NewErrorHandler(c, err, fiber.StatusInternalServerError)
		}

		var result interface{}
		if err := json.Unmarshal(cachedResponse, &result); err != nil {
			return utils.NewErrorHandler(c, err, fiber.StatusInternalServerError)
		}

		return utils.NewSuccessHandler(c, map[string]interface{}{getResponseKey(path): result})
	}
}

func getResponseKey(path string) string {
	switch {
	case strings.HasPrefix(path, "/cities/all"):
		return "cities"
	case strings.HasPrefix(path, "/cities/"):
		return "city"
	default:
		return "noMatter"
	}
}
