package command

import (
	"fmt"
	"os"
	"strings"
)

/**
接收命令行执行传递过来的参数
中国 我知道 liushan zhongxiaoxia liuxu C:\Users\1040\AppData\Local\Temp\go-build589176354\b001\exe\main.exe
这个参数读取和shell差不多，都是运行的命令算作第一个参数，index是0

s := ""
var s string
var s = ""
var s string = ""

记不住go的变量声明规则啊，有点小复杂啊，但是肯定是有逻辑在里面的
包级别只允许出现以下几个关键字
const type func import var
*/
func ShowCommandArguments() {
	var s, sep string
	fmt.Println("sep信息:", sep, "{}")
	for i := len(os.Args) - 1; i >= 0; i-- {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	for index, arg := range os.Args {
		fmt.Println(index, arg)
	}

	fmt.Println(strings.Join(os.Args, ","))
}
