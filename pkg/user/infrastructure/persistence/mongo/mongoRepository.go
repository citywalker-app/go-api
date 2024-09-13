package mongo

import (
	"os"

	"github.com/citywalker-app/go-api/database"
	userdomain "github.com/citywalker-app/go-api/pkg/user/domain"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Collection *mongo.Collection
}

func NewMongoRepository() userdomain.Repository {
	db, ok := database.DB.GetDB().(*mongo.Database)
	if !ok {
		log.Error("database.DB not of type *mongo.Database")
	}

	collectionName := os.Getenv("MDB_COLLECTION_USERS")
	if collectionName == "" {
		log.Error("MDB_COLLECTION_USERS not set, check .env file")
	}
	return &Repository{
		Collection: db.Collection(collectionName),
	}
}
