package main

import (
	"mouse/article"
	"mouse/command"
)

/**
 还是需要紧跟语言啊，毕竟，语言挺重要，是表达的唯一途径
go run main.go 是运行
go build main.go 是编译成所在系统的可执行文件
go clean 类似于mvn clean，就是清除已经编译好的可执行的文件一类的东西
*/
func main() {
	/*	fmt.Println("This is the first go main")
		fmt.Println("现在是:", time.Now())
		packages.SelfPackage("main")

		fmt.Println("幸运数字:", rand.Intn(125))

		sum := packages.Sum(10, 20)
		fmt.Println("The sum is:", sum)

		fmt.Println(packages.Swap("liu", "xu"))
		x, y := packages.Split(25)
		fmt.Println("查看两个返回值:", x, y)
		packages.Variable()

		flow.DoubleNumber()
		flow.IfStatement("liuxu")
		flow.SwitchStatement()
		flow.DeferStatement()

		name := "liuxuzxx"
		complex.ChangeName(&name)
		fmt.Println("This is name:", name)
		complex.CreateStruct()

		complex.SlicesStatement()
		complex.SliceOperation()
		complex.MapOperation()
		complex.StrategyOperation("real")

		method.MethodOperation()

		method.InterfaceOperation()

		goroutines.StartGoRoutines("No-Die")
		goroutines.LogInformation("Main")

		goroutines.InvokeChanSum()
		goroutines.BufferChannel()*/
	article.LoadChuci()
	command.ShowCommandArguments()
	//command.Lissajous(os.Stdout)
}
