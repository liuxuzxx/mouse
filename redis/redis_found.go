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
		Addr:     "172.16.16.37:6379",
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

func PipelineList(data *map[string]interface{}) {
	pipeline := rdb.Pipeline()
	for key, value := range *data {
		pipeline.LPush(ctx, key, value)
	}
	_, _ = pipeline.Exec(ctx)
}

func Get(key string) *redis.StringSliceCmd {
	return rdb.LRange(ctx, key, 0, -1)
}

func searchPhoneNumber(key string, number int) int {
	end, _ := rdb.LLen(ctx, key).Result()
	start := int64(0)
	return binarySearch(key, number, start, end)
}

func binarySearch(key string, number int, start, end int64) int {
	for start <= end {
		middle := (end + start) / 2
		binaryResult, _ := rdb.LRange(ctx, key, middle, middle).Result()
		if len(binaryResult) == 0 {
			fmt.Printf("查看错误数据信息:%s,%d,%d  分割信息", key, middle, number)
		}
		phoneNumber, _ := strconv.Atoi(binaryResult[0])
		if phoneNumber > number {
			start = middle + 1
			binarySearch(key, number, start, end)
		} else if phoneNumber < number {
			end = middle - 1
			binarySearch(key, number, start, end)
		} else {
			return phoneNumber
		}
	}
	return -1
}

func init() {
	rdb = redisClient()
}
