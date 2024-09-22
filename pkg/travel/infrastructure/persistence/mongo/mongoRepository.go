package mongo

import (
	"os"

	"github.com/citywalker-app/go-api/database"
	traveldomain "github.com/citywalker-app/go-api/pkg/travel/domain"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Collection *mongo.Collection
}

func NewMongoRepository() traveldomain.Repository {
	db, ok := database.DB.GetDB().(*mongo.Database)
	if !ok {
		log.Panic("database.DB not of type *mongo.Database")
	}

	collectionName := os.Getenv("MDB_COLLECTION_TRAVELS")
	if collectionName == "" {
		log.Panic("MDB_COLLECTION_TRAVELS not set")
	}

	return &Repository{
		Collection: db.Collection(collectionName),
	}
}
