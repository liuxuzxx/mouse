package redis

import (
	"context"
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
// 纳秒数:80014132424
// 总量是:19950001 差5w到2000w
// 速度是: 249330.9668632495/s 25w的速度
//
func PipelineSet(data *map[string]interface{}) {
	pipeline := rdb.Pipeline()
	for key, value := range *data {
		pipeline.Set(ctx, key, value, 0)
	}
	_, err := pipeline.Exec(ctx)
	if err != nil {
		panic(err)
	}
}

func PipelineBitSet(data *[]PhoneNumberDetail) {
	pipeline := rdb.Pipeline()
	step := 5
	for _, value := range *data {
		key := value.Section
		offset := value.PhoneNumber
		status := value.Status
		pipeline.BitField(ctx, strconv.Itoa(key), "set", "u5", int64(offset*step), 1<<4+status)
	}
	_, _ = pipeline.Exec(ctx)
}

func init() {
	rdb = redisClient()
}
