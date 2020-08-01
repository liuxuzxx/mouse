package redis

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

//
// 这个文件里面主要是放置存储号码数据信息的东西
//　我们首先是通过各种优化得到最好的方式来存储数据
// 在数据结构当中有那么一句话space trade off time
//　就是所谓的时间和空间的交换工作
//

var phoneNumberSection = map[string][]uint8{
	"mobile":  {139, 138, 137, 136, 135, 134, 147, 150, 151, 152, 157, 158, 159, 172, 178, 182, 183, 184, 187, 188, 198},
	"unicom":  {130, 131, 132, 140, 145, 146, 155, 156, 166, 167, 185, 186, 145, 175, 176},
	"telecom": {133, 149, 153, 177, 173, 180, 181, 189, 191, 199},
}

//
// 号码的信息包含有：运营商，所在城市，状态
// 运营商类型可能改变，但是不会太多，所在城市更不用说了，这个基本不会改变，所以我们暂时不存储了
//
type PhoneNumberDetail struct {
	Section     int
	Status      uint8
	PhoneNumber int
}

func PhoneCache() {
	ch := make(chan struct{})
	for _, value := range phoneNumberSection {
		for _, section := range value {
			go func(section uint8) {
				doPhoneCache(section)
				ch <- struct{}{}
			}(section)
		}
		for range value {
			<-ch
		}
	}
}

func doPhoneCache(section uint8) {
	fmt.Printf("开始存放到redis缓存:%d\n", section)
	data := make([]PhoneNumberDetail, 0)
	start := time.Now().UnixNano()
	for index := 0; index < 99999999; index = index + 1 {
		detail := PhoneNumberDetail{
			Section:     int(section),
			Status:      2,
			PhoneNumber: index,
		}
		data = append(data, detail)
		if index%100000 == 0 {
			PipelineBitSet(&data)
			end := time.Now().UnixNano()
			fmt.Printf("数据量:%d 耗费时间为:%d\n", index, end-start)
			start = end
			data = make([]PhoneNumberDetail, 0)
		}
	}
}

//
// 使用JSON的形式存储到Redis当中去，看下消耗的内存的大小
//
func jsonDetail(detail *PhoneNumberDetail) interface{} {
	marshalCodes, _ := json.Marshal(detail)
	return string(marshalCodes)
}

func stringDetail(detail *PhoneNumberDetail) interface{} {
	return strconv.Itoa(int(detail.Status))
}

func shareCity(detail *PhoneNumberDetail) interface{} {
	return detail.Status
}
