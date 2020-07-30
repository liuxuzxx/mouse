package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
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

//
// Pipeline的速度果然很快，基本速度可到25w/s的写入量,网速基本上顶格到8-9MB的速度了
//查看结果：[]redis.Cmder存储5w个数据，耗费的时间是:201851939
//查看结果：[]redis.Cmder存储5w个数据，耗费的时间是:192541394
//查看结果：[]redis.Cmder存储5w个数据，耗费的时间是:209681272
//查看结果：[]redis.Cmder存储5w个数据，耗费的时间是:197705737
//查看结果：[]redis.Cmder存储5w个数据，耗费的时间是:195451363
//查看结果：[]redis.Cmder存储5w个数据，耗费的时间是:198647618
//查看结果：[]redis.Cmder存储5w个数据，耗费的时间是:214687613
//查看结果：[]redis.Cmder存储5w个数据，耗费的时间是:203231755
//查看结果：[]redis.Cmder存储5w个数据，耗费的时间是:198738337
//查看结果：[]redis.Cmder存储5w个数据，耗费的时间是:203646099
//查看结果：[]redis.Cmder存储5w个数据，耗费的时间是:193644458
// 这个时间是纳秒，纳秒-->微秒-->毫秒 1000进制
//
func PipelineSet(data *map[string]interface{}) {
	pipeline := rdb.Pipeline()
	for key, value := range *data {
		pipeline.Set(ctx, key, value, 0)
	}
	result, err := pipeline.Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("查看结果：%T", result)
}

func PipelineBitSet(data *[]PhoneNumberDetail) {
	pipeline := rdb.Pipeline()
	step := 5
	for _, value := range *data {
		key := value.Section
		offset := value.PhoneNumber
		status := value.Status
		pipeline.SetBit(ctx, strconv.Itoa(key), int64(offset*step), 1)

		for index := 0; index < 4; index = index + 1 {
			bitValue := status >> index & 1
			pipeline.SetBit(ctx, strconv.Itoa(key), int64(offset*step+1+index), int(bitValue))
		}
	}
	_, _ = pipeline.Exec(ctx)
}

func init() {
	rdb = redisClient()
}
