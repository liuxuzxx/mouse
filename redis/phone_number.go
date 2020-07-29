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
	City   string
	Status uint8
}

func PhoneCache() {
	for _, value := range phoneNumberSection["mobile"] {
		phoneValue := int64(value) * 100000000
		data := make(map[string]interface{}, 0)
		startTime := time.Now().UnixNano()
		for index := 0; index < 20000000; index = index + 1 {
			resultPhone := phoneValue + int64(index)
			detail := &PhoneNumberDetail{
				City:   "203",
				Status: 2,
			}

			//detailJson := jsonDetail(detail)
			//detailJson := stringDetail(detail)
			detailJson := shareCity(detail)
			data[strconv.FormatInt(resultPhone, 10)] = detailJson
			if index%50000 == 0 {
				PipelineSet(&data)
				data = make(map[string]interface{}, 0)
			}
		}
		endTime := time.Now().UnixNano()
		fmt.Printf("耗费的时间为:%d\n", endTime-startTime)
		return
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
	return detail.City + "-" + strconv.Itoa(int(detail.Status))
}

func shareCity(detail *PhoneNumberDetail) interface{} {
	return detail.Status
}
