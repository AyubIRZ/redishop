package util

import (
	"github.com/RediSearch/redisearch-go/redisearch"
	"os"
	"sync"
)

// RedisClient is a wrapper around redis client
type RedisSearchClient struct {
	*redisearch.Client
}

var onceRedisSearch sync.Once
var redisSearchClient *RedisSearchClient

//GetRedisClient handles and returns a singleton Redis client
func GetRedisSearchClient() *RedisSearchClient {
	onceRedisSearch.Do(func() {
		client := redisearch.NewClient(os.Getenv("REDIS_SEARCH_HOST")+":"+os.Getenv("REDIS_SEARCH_PORT"), "redishop")
		redisSearchClient = &RedisSearchClient{client}
	})

	return redisSearchClient
}
