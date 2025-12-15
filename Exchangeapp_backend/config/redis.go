package config

import (
	"exchangeapp/Exchangeapp/Exchangeapp_backend/global"
	"log"

	"github.com/go-redis/redis"
)

func initRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     AppConfig.Redis.Addr,
		DB:       AppConfig.Redis.Db,
		Password: AppConfig.Redis.Password,
	})

	_, err := RedisClient.Ping().Result()

	if err != nil {
		log.Fatal("Failed to connect to Redis, got error: &v", err)
	}

	global.RedisDB = RedisClient
}
