package envloader

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func init() {
	log.Info("Loading environment variables...")

	// Load environment variables
	if err := godotenv.Load(os.Getenv("ENV_FILE")); err != nil {
		log.Warn("No env gotten, maybe it's production. If not, check the file path.")
	} else {
		log.Info("Environment variables loaded.")
	}
}
