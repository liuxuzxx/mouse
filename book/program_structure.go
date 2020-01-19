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
