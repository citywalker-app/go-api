package travelapplication

import (
	"os"

	traveldomain "github.com/citywalker-app/go-api/pkg/travel/domain"
	"github.com/citywalker-app/go-api/pkg/travel/infrastructure/persistence/mongo"
	"github.com/gofiber/fiber/v2/log"
)

var Repo = func() traveldomain.Repository {
	switch os.Getenv("DATABASE") {
	case "MONGODB":
		return mongo.NewMongoRepository()
	default:
		log.Panic("Unsupported database: modify .env file")
		return nil
	}
}()
