package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	db *mongo.Database
}

func (m *MongoDB) Connect() {
	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user, pass := os.Getenv("MDB_USERNAME"), os.Getenv("MDB_PASSWORD")
	database, dbHost := os.Getenv("MDB_DATABASE"), os.Getenv("DB_HOST")
	uri := fmt.Sprintf("mongodb://%s:%s@%s", user, pass, dbHost)

	// Connecting to DB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic("Error connecting to MongoDB: ", err)
	}

	m.db = client.Database(database)
}

func (m *MongoDB) ConnectTest() {
	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := "mongodb://test:test@localhost:27017"

	// Connecting to DB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic("Error connecting to MongoDB: ", err)
	}

	m.db = client.Database("citywalker")
}

func (m *MongoDB) GetDB() interface{} {
	return m.db
}
