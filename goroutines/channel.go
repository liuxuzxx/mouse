package goroutines

import "fmt"

//
//Channels：管道，类似于Java的NIO的Channel这个东西
//垃圾的代码，真的是不能修改了，完全是很垃圾的代码
//有种：大鹏战士很天低的感觉
//

func chanSum(numbers []int, channel chan int) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	channel <- sum
}

//
//不是很理解的就是，这个为什么在goroutines中调用channel这个东西，不是非常的理解啊
//难道是为了其他的事情
//
func InvokeChanSum() {
	channel := make(chan int)

	numbers := []int{1, 2, 3, 4, 5}

	go chanSum(numbers, channel)

	sum := <-channel
	fmt.Println("产看获取到的sum数字:", sum)
}

func BufferChannel() {
	channel := make(chan int, 2)
	channel <- 90
	channel <- 100
	fmt.Println(<-channel, <-channel)
}
