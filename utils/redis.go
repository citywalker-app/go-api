package utils

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2/log"
)

var RedisDB = func() *redis.Client {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Error("Error converting REDIS_DB to int: %v", err)
	}

	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       db, // use default DB
	})
}()

func SetCache(ctx context.Context, key string, value interface{}) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return RedisDB.Set(ctx, key, jsonData, 60*time.Minute).Err()
}

func GetCache(ctx context.Context, key string, value interface{}) error {
	data, err := RedisDB.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, value)
}
