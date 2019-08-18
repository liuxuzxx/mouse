package flow

import (
	"fmt"
	"runtime"
)

// := 这种赋值方式一般是用来处理局部变量的东西
func EqualDifferenceSeries(tail int) int {
	var sum = 0
	var index int = 0
	for index = 0; index < tail; index++ {
		sum += index
	}

	return sum
}

func DoubleNumber() {
	sum := 1
	for ; sum < 500; {
		sum += sum
	}

	fmt.Println("获取大于一个数字的最小的2进制数字", sum)
}

func IfStatement(name string) {
	if len(name) > 5 {
		fmt.Println("这个名字长度超过5")
	} else if len(name) < 4 {
		fmt.Println("This is length is 4")
	}
	//the if statement can start with a short statement to execute before the condition.
	//就是说if语句可以在condition条件比对之前进行一个初始化的赋值操作
	//Go好像对于一个表达式的赋值的返回值不是一个东西
	if length := len(name); length < 5 {

	}

	flag := len(name) > 5

	fmt.Println("看啊看呢和则个是不是bool:", flag)
}

//switch其实就是多if-else版本而已，基本语法和if-else一样子
//switch的case不用使用break关键字，go可能是默认了
func SwitchStatement() {
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("您的系统是:", os)
	case "darwin":
		fmt.Println("OS X")
	default:
		fmt.Printf("%s\n",os)
	}
}
//A defer statement defers the execution of a function until the surrounding function returns.
//
//The deferred call's arguments are evaluated immediately,
// but the function call is not executed until the surrounding function returns.

//总结一下defer关键字的特点：当执行到defer这句话的时候，首先是执行这个函数的参数
//之后等到包含这句defer的函数执行完之后，开始执行defer函数
//defer函数执行起来就像一个stack的操作
func DeferStatement() string{
	name := "liuxu"
	defer fmt.Println(addHead())
	fmt.Println("This name is :",name)
	defer fmt.Println("第二次的defer函数")
	return name
}

func addHead() string{
	fmt.Println("体验一下defer的参数推断")
	return "addHead"
}
