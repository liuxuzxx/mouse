package redis

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

//
//手机号码的基本信息
//

var phoneNumberSection = map[string][]uint8{
	"mobile":  {139, 138, 137, 136, 135, 134, 147, 150, 151, 152, 157, 158, 159, 172, 178, 182, 183, 184, 187, 188, 198},
	"unicom":  {130, 131, 132, 140, 145, 146, 155, 156, 166, 167, 185, 186, 145, 175, 176},
	"telecom": {133, 149, 153, 177, 173, 180, 181, 189, 191, 199},
}

type PhoneNumberDetail struct {
	City    string `json:"city"`
	Status  int8   `json:"status"`
	Carrier int8   `json:"carrier"`
}

func MobilePhoneNumbers() {
	for index, value := range phoneNumberSection["mobile"] {
		fmt.Printf("序号是:%d 数值是:%d\n", index, value)
		basicValue := int64(value) * 100000000
		startTime := time.Now().Nanosecond()
		for start := 0; start <= 99999999; start = start + 1 {
			phoneNumber := basicValue + int64(start)
			phone := strconv.FormatInt(phoneNumber, 10)
			detail := &PhoneNumberDetail{
				City:    "北京市",
				Status:  11,
				Carrier: 3,
			}
			jsonBytes, _ := json.Marshal(detail)
			Set(phone, string(jsonBytes))
			if start%100000 == 0 {
				endTime := time.Now().Nanosecond()
				fmt.Printf("号码段:%d 10万条数据耗费的时间为:%d 纳秒\n", value, endTime-startTime)
				startTime = endTime
			}
			if start > 1000000 {
				fmt.Printf("已经查过了100万数据，暂时停止输送了!")
				return
			}
		}
	}
}
