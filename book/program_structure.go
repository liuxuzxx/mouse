package book

//<The Go Programming Language>这本书的实验工作

//Names:主要是说GO语言的变量命名规范，正常的操作

//Declaration:声明东西
import (
	"fmt"
	"os"
)

const boiling = 0.98

func CalMoney(count int) {
	var f = boiling
	var sum = (f - 32) * 5 / 9
	fmt.Printf("boiling point:%g\n", sum)
}

//Variables:变量，程序的灵魂
var userName string
var age int

var (
	root  string
	count int
)

//Short variables
//GO语言的short variables name,这种方式主要是为了处理local
//变量的，所以不用写类型，毕竟你开始的时候就已经初始化了
//还有一个就是不可以重复定义声明，但是，好像只要有一个不一样就可以
//这个规则很奇怪啊，按照道理来说应该是都不一样才行啊.
func animal(count int) string {
	animalNumber := 10
	fmt.Print(animalNumber)
	open, err := os.Open("lx")
	if err != nil {
		fmt.Print(open)
	}
	var clues, defect string
	clues, defect = "liuxu", "wangzhang"

	fmt.Print(clues, defect)
	return "liuxu"
}
