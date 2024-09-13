// nolint:gci
package main

import (
	"log"
	"os"

	_ "github.com/citywalker-app/go-api/envLoader"
	_ "github.com/joho/godotenv/autoload"

	"github.com/citywalker-app/go-api/server"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	app := server.Setup()

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		server.StartServer(app)
	} else {
		server.StartServerWithGracefulShutdown(app)
	}
}
