package book

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

//
// @Date 2020-01-15 10:10-56
// @Author 刘旭
// 基本的数据类型
//

//integer:整数类型
//纯int占据的字节和平台有关系，也就是说还是有点像c语言，看平台，但是其他集中类型没啥子问题
//T(value) 这种可以让value的类型转换为T类型

func IntegerType() {
	var (
		color      int8  = 120
		unicode    int16 = 1280
		number     int32 = 100000
		loginCount int64 = 1200000000000000
		normal     int   = 90
	)

	fmt.Printf("%d %d %d %d %d\n", unsafe.Sizeof(color), unsafe.Sizeof(unicode), unsafe.Sizeof(number), unsafe.Sizeof(loginCount), unsafe.Sizeof(normal))
}

//float point numbers

func CompareFloat() {
	var money float32 = 16777217 //1<<24 因为32位，24位标识整数部分，剩下8位标识小数部分，所以增加1就是越界了
	fmt.Printf("%f %f\n", money, money+1)
}

//complex numbers
func ComplexNumbers() {
	var (
		x complex64 = 1 + 2i
		y complex64 = 3 + 5i
	)
	fmt.Println(x * y)
}

//boolean,没啥子好说的，就是true和false这两个类型

//string类型

func StringFound() {
	var (
		welcome string = "welcome to china!"
		china   string = "Hello 你好中国!"
		comment string = `这是一段注释性质的内容,
                         主要是想验证一下这个所谓的raw string内容
所展现出来的结果性质的东西!`
	)
	fmt.Printf("%d %d\n", len(welcome), utf8.RuneCountInString(china))
	fmt.Printf("展示最原始的内容信息:%s\n", comment)

	//这种方式也是可以遍历utf8类型的字符串
	for _, r := range china {
		fmt.Printf("%c ", r)
	}

	fmt.Println()
	//这种方式可以遍历一个utf8的字符串
	for i := 0; i < len(china); {
		r, size := utf8.DecodeRuneInString(china[i:])
		fmt.Printf("%d\t%c ", i, r)
		i += size
	}
	fmt.Println()
}

type Money float64
type Price float64

var (
	price = Price(10.23)
)

//这种type的形式目的就是为了给变量一个别名，但是好像还有严格的类型验证的校验权限
func SumMoney(count int) Money {
	return Money(Price(count) * price)
}
