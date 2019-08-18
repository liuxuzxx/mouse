package packages

//有没有发现go的函数定义很特别
//直接就是我们正常语言的掉转头
func Sum(left, right int) int {
	return left + right
}

//可以这么声明参数，每个参数都指定类型
func Add(left int, right int) int {
	return left + right
}

//定义了一个函数返回两个结果
//我估计，内部肯定是使用Map来存储，然后给扔出去
func Swap(left, right string) (string, string) {
	return right, left
}

//可以定义返回的变量
//Naked return statements
// should be used only in short functions,
// as with the example shown here. They can harm readability in longer functions
func Split(sum int) (x, y int) {
	x = sum + 10
	y = sum * 2
	return
}
