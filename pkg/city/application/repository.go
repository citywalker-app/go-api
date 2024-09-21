package cityapplication

import (
	"os"

	citydomain "github.com/citywalker-app/go-api/pkg/city/domain"
	"github.com/citywalker-app/go-api/pkg/city/infrastructure/persistence/mongo"
	"github.com/gofiber/fiber/v2/log"
)

var repo = func() citydomain.Repository {
	switch os.Getenv("DATABASE") {
	case "MONGODB":
		return mongo.NewMongoRepository()
	default:
		log.Panic("Unsupported database: modify .env file")
		return nil
	}
}()
