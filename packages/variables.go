package packages

import "fmt"

//The var statement declares a list of variables;
// as in function argument lists, the type is last.
// A var statement can be at package or function level.
// We see both in this example.
var c, java, python, cpp bool

func Variable() {
	var c int
	fmt.Println("How to do define the variable!", c)
}

//Go存在类型推断的功能，所以，可以直接赋值，
// 然后推断类型
var index, count int = 10, 9
var segment, slot = "this", false

func Initializer() {
	count := 90
	fmt.Println("In the function number:", count)

	var sum = JoinNumber(10, 20)
	number := JoinNumber(10, 25)
	fmt.Println(sum, number)
}

func JoinNumber(left, right int) int {
	return left + right
}

//也可以这么声明多个变量，类似于import多个package
var (
	Flag       bool    = false
	PageSize   int     = 90
	PageNumber float32 = 2.3
)

//这种进行类型的转换还是需要通过构造函数
func Conversion() {
	var x, y int = 34, 23
	var f float32 = float32(x)

	fmt.Println(x, y, f)
}

//这么看，var和const其实就是一个修饰词语，本来语言定义一个量
//也就是三要素：类型，名字，修饰 var修饰变量，const修饰常量

const (
	TableName string = "article"
	IsFlag    bool   = false
)

var (
	Left int
	Right int
)

func Table() {
	const InnerName string = "asdfasdf"
}
