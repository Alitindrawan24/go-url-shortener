package store

import (
	"fmt"
	"strconv"
	"time"
	"url-shortener/helper"

	"github.com/go-redis/redis"
)

type StorageService struct {
	redisClient *redis.Client
}

var storageService = &StorageService{}

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	database, _ := strconv.Atoi(helper.GetEnvVariableKey("REDIS_DATABASE", "0"))

	redisClient := redis.NewClient(&redis.Options{
		Addr:     helper.GetEnvVariableKey("REDIS_ADDRESS", "localhost:6379"),
		Password: helper.GetEnvVariableKey("REDIS_PASSWORD", ""),
		DB:       database,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init redis connection: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message - {%s}", pong)
	storageService.redisClient = redisClient
	return storageService
}

func SaveUrlMapping(shortUrl string, originalUrl string, UserID string) {
	err := storageService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed to saving key url | Error: %v", err))
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storageService.redisClient.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to retrieve initial url | Error: %v", err))
	}

	return result
}
