package mongo

import (
	"os"

	"github.com/citywalker-app/go-api/database"
	citydomain "github.com/citywalker-app/go-api/pkg/city/domain"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Collection *mongo.Collection
}

func NewMongoRepository() citydomain.Repository {
	db, ok := database.DB.GetDB().(*mongo.Database)
	if !ok {
		log.Error("database.DB not of type *mongo.Database")
	}

	collectionName := os.Getenv("MDB_COLLECTION_CITIES")
	if collectionName == "" {
		log.Error("MDB_COLLECTION_CITIES not set")
	}
	return &Repository{
		Collection: db.Collection(collectionName),
	}
}
