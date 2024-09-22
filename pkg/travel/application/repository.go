package travelapplication

import (
	"os"

	traveldomain "github.com/citywalker-app/go-api/pkg/travel/domain"
	"github.com/citywalker-app/go-api/pkg/travel/infrastructure/persistence/mongo"
)

var Repo = func() traveldomain.Repository {
	switch os.Getenv("DATABASE") {
	case "MONGODB":
		return mongo.NewMongoRepository()
	default:
		panic("Unsupported database: modify .env file")
	}
}()
