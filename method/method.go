package method

import (
	"fmt"
)

//go没有class的概念，所以没有类型方法

type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) Abs() float64 {
	return v.X * v.Y
}

func (v Vertex) Scale(number int) {
	v.X = v.X * float64(number)
	v.Y = v.Y * float64(number)
	fmt.Println("The sum is :", v)
}

type MyFloat float64

//可以使用type关键字来处理类型的问题
func (f MyFloat) Abs(){

}


//其实也很好操作
//其实弄过来弄过去，go还是分值传递和指针传递
func MethodOperation() {
	v := Vertex{1.23, 3.6}
	v.Abs()
	v.Scale(10)
	fmt.Println("v:",v)
}
