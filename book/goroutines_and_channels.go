package book

import (
	"fmt"
	"time"
)

//
// @Date 2020-01-15 15:49-44
// @Author 刘旭
// Go的并发编程,go有两种并发编程的模型，一种是goroutine,另外一种是共享内存
//

func Goroutine() {
	startGoroutine()
}

func startGoroutine() {
	go spinner(100 * time.Millisecond)
	const n = 45
	result := fib(n)
	fmt.Printf("The fib result is:%d", result)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "range" {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
