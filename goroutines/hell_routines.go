package goroutines

import (
	"fmt"
	"time"
)

//
//go-routines和java的thread还不一样，当你的main结束，他就跟着结束了，真是奇怪啊
//
func LogInformation(routineName string) {
	for index := 0; index < 10; index++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Go-Routines:", routineName)
	}
}

func StartGoRoutines(routineName string) {
	fmt.Println("Start a go routine:", routineName)
	go LogInformation("Go-Liuxu")
}
