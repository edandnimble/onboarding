package utils

import (
	"context"
	"strconv"
	"sync"

	"github.com/go-redis/redis"
)

var redisCtx = context.Background()

type redisClient struct {
	client     redis.UniversalClient
	ctx        *context.Context
	numKey     string
	guesserKey string
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
			numKey:     "NUM_",
			guesserKey: "GUESS_",
		}
	})

	return redisSingletonClient
}

func (c *redisClient) IncreaseGuess(id string) (err error) {
	return c.client.Incr(id).Err()
}

func (c *redisClient) AddNumber(num uint32) (err error) {
	err = c.client.Set(c.numKey+strconv.FormatUint(uint64(num), 10), true, 0).Err()
	return
}

func (c *redisClient) RemoveNumber(num uint32) (err error) {
	err = c.client.Del(c.numKey + strconv.FormatUint(uint64(num), 10)).Err()
	return
}

func (c *redisClient) IsNumberExist(num uint32) (exists bool, err error) {
	val, err := c.client.Exists(c.numKey + strconv.FormatUint(uint64(num), 10)).Result()
	if val == 0 {
		return false, err
	}

	return true, err
}
