package book

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

//
// @Date 2020-01-15 15:49-44
// @Author 刘旭
// Go的并发编程,go有两种并发编程的模型，一种是goroutine,另外一种是共享内存
//The main function then returns. When this happens, all goroutines are abruptly terminated
//and the program exits
//这句话耐人寻味啊：当main退出，都退出了
//

func Goroutine() {
	startGoroutine()
}

func startGoroutine() {
	go spinner(100 * time.Millisecond)
	const n = 45
	result := fib(n)
	fmt.Printf("The fib result is:%d\n", result)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "range" {
			fmt.Printf("\r%c\n", r)
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

//使用go语言制作一个简单的服务器

func Server() {
	listen, err := net.Listen("tcp", "localhost:16380")
	if err != nil {
		log.Fatal(err)
	}

	for {
		connection, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	for {
		_, err := io.WriteString(connection, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
	}
}
