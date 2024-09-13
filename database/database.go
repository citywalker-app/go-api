package database

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
)

type Database interface {
	Connect()
	GetDB() interface{}
}

var DB = func() Database {
	switch os.Getenv("DATABASE") {
	case "MONGODB":
		var db MongoDB
		db.Connect()
		return &db
	default:
		log.Error("Unsupported database: modify .env file")
		return nil
	}
}()
