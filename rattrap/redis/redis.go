package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"mouse/rattrap/config"
)

var ctx = context.Background()
var rdb *redis.Client

func redisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.Address,
		Password: config.Conf.Redis.Password,
		DB:       config.Conf.Redis.Db,
	})
}

func Incrby(key string, value int64) int64 {
	number := rdb.IncrBy(ctx, key, value)
	return number.Val()
}

func init() {
	rdb = redisClient()
}
