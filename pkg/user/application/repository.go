package userapplication

import (
	"os"

	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
	"github.com/citywalker-app/go-api/pkg/user/infrastructure/persistence/mongo"
	"github.com/gofiber/fiber/v2/log"
)

var repo = func() userdomain.Repository {
	switch os.Getenv("DATABASE") {
	case "MONGODB":
		return mongo.NewMongoRepository()
	default:
		log.Error("Unsupported database: modify .env file")
		return nil
	}
}()
