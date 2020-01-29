package book

import "fmt"

//函数，这个是一个重要的地方

//Function Declarations:函数的声明

func compareObj(left AncientArticle, right AncientArticle) bool {
	return left == right
}

//Recursion:递归操作，就是一个很有特性的东西

//Multiple Return Values:多个返回值，这个其实很有意思，go语言的一个新特性

//Errors:异常或者说是错误处理，就和Java中的Exception没什么区别
//对了，这个错误信息不让大写字母，不知道为啥
//error也是当作一个返回值，有点意思，还是和Java的异常不怎么一样
//怀疑Java是不是底层也是这个原理
func ReadFile(path string) error {
	if len(path) == 0 {
		return fmt.Errorf("parameter path can not be empty:%s", path)
	}
	return nil
}

//Functions Values.应该是函数指针的意思

func sumValue(left int64, right int64) int64 {
	return left - right
}

func multiValue(left int64, right int64) int64 {
	result := left * right
	fmt.Printf("计算乘法:%d * %d = %d\n", left, right, result)
	return result
}

func FunctionPoint() {
	var twoParam func(int64, int64) int64 = sumValue
	result := twoParam(32, 8)
	fmt.Printf("结果是:%d\n", result)
	twoParam = multiValue
	result = twoParam(32, 8)
	fmt.Printf("结果是:%d\n", result)
}

//既然与函数指针了，那么返回值是可以返回一个函数的了，那么换句话来说，一切都是对象，这个应用有点广啊
//学习另外一们语言有时候会让我们从另外一个角度看待问题

//defer:延迟执行的关键字，主要是用来执行close关闭资源的一些操作
//查看defer的执行顺序,有点相似finally语句
//defer的执行顺序为：后defer的先执行。
//defer的执行顺序在return之后，但是在返回值返回给调用方之前，所以使用defer可以达到修改返回值的目的

func DeferOperation() string {
	userName := "This is User Name"
	defer overwrite(&userName)
	fmt.Printf("执行正常的操作:%s\n", userName)
	return userName
}

func overwrite(userName *string) {
	*userName = "liuxu"
	fmt.Printf("执行了defer语句操作\n")
}

//Panic,有点相似Java的RuntimeException,Error类似于Exception

func DivideNumber(left int64) {
	if left == 0 {
		panic(fmt.Sprintf("错误信息......"))
	}
}
