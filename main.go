package main

import (
	"fmt"
	"math/rand"
	"mouse/complex"
	"mouse/flow"
	"mouse/method"
	"mouse/packages"
	"time"
)

func main() {
	fmt.Println("This is the first go main")
	fmt.Println("现在是:", time.Now())
	packages.SelfPackage("main")

	fmt.Println("幸运数字:", rand.Intn(125))

	sum := packages.Sum(10, 20)
	fmt.Println("The sum is:", sum)

	fmt.Println(packages.Swap("liu","xu"))
	x, y := packages.Split(25)
	fmt.Println("查看两个返回值:",x,y)
	packages.Variable()

	flow.DoubleNumber()
	flow.IfStatement("liuxu")
	flow.SwitchStatement()
	flow.DeferStatement()

	name := "liuxuzxx"
	complex.ChangeName(&name)
	fmt.Println("This is name:",name)
	complex.CreateStruct()

	complex.SlicesStatement()
	complex.SliceOperation()
	complex.MapOperation()
	complex.StrategyOperation("real")

	method.MethodOperation()
}
