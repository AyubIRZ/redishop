package util

import (
	"github.com/go-redis/redis"
	"log"
	"os"
	"sync"
)

// RedisClient is a wrapper around redis client
type RedisClient struct {
	*redis.Client
}

var onceRedis sync.Once
var redisClient *RedisClient

//GetRedisClient handles and returns a singleton Redis client
func GetRedisClient() *RedisClient {
	onceRedis.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		redisClient = &RedisClient{client}
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to redis %v", err)
	}

	return redisClient
}
