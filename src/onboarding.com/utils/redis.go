package utils

import (
	"fmt"
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
		ip, port, err := GetServiceDNS("redis")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		redisSingletonClient = &redisClient{
			client: redis.NewUniversalClient(
				&redis.UniversalOptions{
					Addrs: []string{ip + ":" + port},
				}),
		}
	})

	return redisSingletonClient
}

func (c *redisClient) IncreaseGuess(id uint32) (int64, error) {
	num := strconv.FormatUint(uint64(id), 10)
	return c.client.Incr(num).Result()
}
