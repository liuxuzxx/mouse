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
var rdb *redis.Client

func redisClient() *redis.Client{
	rdb = redis.NewClient(&redis.Options{
		Addr:     "172.16.16.37:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Set(ctx, "liuxu", "zhongxiaoxia", 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("初始化Redis的客户端对象成功")
	return rdb
}

func Set(key,value string){
	_ = rdb.Set(ctx, key, value, 0).Err()
}

func init(){
	rdb = redisClient()
}