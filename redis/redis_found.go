package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

//
//使用go语言来连接Redis,一个是可以熟悉go语言
//另外一个就是可以尽可能最大的提高Redis的读写性能
//
var ctx = context.Background()
var rdb *redis.Client

func redisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func Set(key string, value interface{}) {
	_ = rdb.Set(ctx, key, value, 0).Err()
}

func init() {
	rdb = redisClient()
}
