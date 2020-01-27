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
func ReadFile(path string) error {
	if len(path) == 0 {
		return fmt.Errorf("parameter path can not be empty:%s", path)
	}
	return nil
}
