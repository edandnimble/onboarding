package utils

import (
	"os"
	"strconv"
	"sync"

	"github.com/go-redis/redis"
)

type redisClient struct {
	client redis.UniversalClient
}

var redisSingletonClient *redisClient
var redisOnce sync.Once

func GetRedisClient() *redisClient {
	redisOnce.Do(func() {
		redisPort := os.Getenv("REDIS_PORT")
		redisSingletonClient = &redisClient{
			client: redis.NewUniversalClient(
				&redis.UniversalOptions{
					Addrs: []string{":" + redisPort},
				}),
		}
	})

	return redisSingletonClient
}

func (c *redisClient) IncreaseGuess(id uint32) (int64, error) {
	num := strconv.FormatUint(uint64(id), 10)
	return c.client.Incr(num).Result()
}
