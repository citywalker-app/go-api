package server

import (
	"os"
	"strconv"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func Config() fiber.Config {
	readTimeoutSecondsCount, err := strconv.Atoi(os.Getenv("READ_TIMEOUT_SECONDS"))
	if err != nil {
		readTimeoutSecondsCount = 5
	}

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		Prefork:     os.Getenv("PREFORK") == "true",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}
}
