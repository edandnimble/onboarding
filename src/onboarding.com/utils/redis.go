package utils

import (
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
		redisSingletonClient = &redisClient{
			client: redis.NewUniversalClient(
				&redis.UniversalOptions{
					Addrs: []string{":6379"},
				}),
		}
	})

	return redisSingletonClient
}

func (c *redisClient) IncreaseGuess(id uint32) (int64, error) {
	num := strconv.FormatUint(uint64(id), 10)
	return c.client.Incr(num).Result()
}
