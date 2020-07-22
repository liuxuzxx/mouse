package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

//
//使用go语言来连接Redis,一个是可以熟悉go语言
//另外一个就是可以尽可能最大的提高Redis的读写性能
//
var ctx = context.Background()

func RedisClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
}
