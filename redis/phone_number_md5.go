package redis

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

//
// 这次实验的是存储号码和他的MD5映射关系
//

func ConvertToMD5(){
	phoneNumber := 18270719298
	fmt.Printf("手机号码是:%d,MD5是:%x\n",phoneNumber,md5.Sum([]byte(strconv.Itoa(phoneNumber))))
}