package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/citywalker-app/go-api/utils"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func JWTProtected() func(*fiber.Ctx) error {
	config := jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}

	return jwtware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	return utils.NewErrorHandler(c, err, fiber.StatusUnauthorized)
}

func JWTHandler() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		bearToken := c.Get("Authorization")
		token := strings.TrimPrefix(bearToken, "Bearer ")

		if token == bearToken {
			return utils.NewErrorHandler(c, ErrInvalidTokenFormat, fiber.StatusUnauthorized)
		}

		claims := &jwt.StandardClaims{}
		parsedToken, err := jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !parsedToken.Valid {
			return utils.NewErrorHandler(c, ErrInvalidToken, fiber.StatusUnauthorized)
		}

		if claims.ExpiresAt < time.Now().Unix() {
			return utils.NewErrorHandler(c, ErrExpiredToken, fiber.StatusUnauthorized)
		}

		c.Locals("email", claims.Issuer)
		return c.Next()
	}
}
