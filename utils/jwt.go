package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(email string) (*string, error) {
	hourCount, ok := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOUR_COUNT"))
	if ok != nil {
		log.Error("Error parsing JWT_EXPIRE_HOUR")
		hourCount = 1
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    email,
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(hourCount)).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	return &token, nil
}
