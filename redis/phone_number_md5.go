package redis

import (
	"crypto/md5"
	"fmt"
	"hash/fnv"
	"strconv"
)

//
// 这次实验的是存储号码和他的MD5映射关系
//

func convertToMD5(phoneNumber int64) string {
	sum := md5.Sum([]byte(strconv.FormatInt(phoneNumber, 10)))
	return fmt.Sprintf("%x", sum)
}

func convertToMD5Bytes(phoneNumber int64) [md5.Size]byte {
	return md5.Sum([]byte(strconv.FormatInt(phoneNumber, 10)))
}

//
// 存储手机号码---MD5的映射关系到Redis当中
// 查看 10000000(1000w)数据的内存耗费量 934286456字节，大概是900MB的样子
// 单个占据的是76个字节，主要是还有其他的数据占用
// 46亿数据就是400GB的内存，有点大，并且不好维护
//

func PhoneNumberMD5() {
	hash := fnv.New32()
	section := 13900000000
	data := make(map[string]interface{})
	for index := 0; index < 100000000; index = index + 1 {
		phoneNumber := section + index
		md5Value := convertToMD5(int64(phoneNumber))
		_, _ = hash.Write([]byte(md5Value))
		hashValue := hash.Sum32()
		data[strconv.FormatInt(int64(hashValue), 10)] = strconv.Itoa(phoneNumber)
		if index%10000 == 0 {
			PipelineSet(&data)
			fmt.Printf("存储了:%d\n", index)
			data = make(map[string]interface{})
		}
	}
}
