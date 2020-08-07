package redis

import (
	"crypto/md5"
	"fmt"
	"hash/fnv"
	"os"
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
	ch := make(chan struct{})
	for _, section := range sections {
		go func(section uint8) {
			//doPhoneNumberMD5Slot(section)
			//doPhoneNumberMD5File(section)
			doFetchList(section)
			ch <- struct{}{}
		}(section)
	}
	for range sections {
		<-ch
	}
}

func doFetchList(section uint8) {
	slotNumber := 1 << 15
	basicNumber := int(section) * 100000000
	hash := fnv.New32()
	for index := 0; index < 100000000; index = index + 1 {
		phoneNumber := basicNumber + index
		md5Value := convertToMD5(int64(phoneNumber))
		_, _ = hash.Write([]byte(md5Value))
		hashValue := hash.Sum32() % uint32(slotNumber)
		key := fmt.Sprintf("%s_%s", strconv.Itoa(int(section)), strconv.FormatUint(uint64(hashValue), 10))
		number := searchPhoneNumber(key, index)
		if number == -1 {
			fmt.Printf("查找失败:%d,%s\n", index, key)
		}
		if index%10000 == 0 {
			fmt.Printf("查找了数据:%d\n", index)
		}
	}
}

func doPhoneNumberMD5Slot(section uint8) {
	slotNumber := 1 << 15
	hash := fnv.New32()
	basicNumber := int(section) * 100000000
	slotStatistics := make(map[string]interface{}, 50000)
	for index := 0; index < 100000000; index = index + 1 {
		phoneNumber := basicNumber + index
		md5Value := convertToMD5(int64(phoneNumber))
		_, _ = hash.Write([]byte(md5Value))
		hashValue := hash.Sum32() % uint32(slotNumber)
		key := fmt.Sprintf("%s_%s", strconv.Itoa(int(section)), strconv.FormatUint(uint64(hashValue), 10))
		slotStatistics[key] = strconv.Itoa(index)
		if index%50000 == 0 {
			PipelineList(&slotStatistics)
			slotStatistics = make(map[string]interface{}, 50000)
			fmt.Printf("已经存储了:%d\n", index)
		}
	}
}

func doPhoneNumberMD5File(section uint8) {
	slotNumber := 1 << 20
	hash := fnv.New32()
	basicNumber := int(section) * 100000000
	slotStatistics := make(map[uint32]int)
	for index := 0; index < 100000000; index = index + 1 {
		phoneNumber := basicNumber + index
		md5Value := convertToMD5(int64(phoneNumber))
		_, _ = hash.Write([]byte(md5Value))
		hashValue := hash.Sum32() % uint32(slotNumber)
		slotStatistics[hashValue] = slotStatistics[hashValue] + 1
	}
	file, err := os.Create(fmt.Sprintf("/home/liuxu/Documents/slot-%d", section))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for index, value := range slotStatistics {
		_, _ = file.WriteString(fmt.Sprintf("slot:%d,count:%d\n", index, value))
	}
}
