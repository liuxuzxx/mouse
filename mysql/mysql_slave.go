package mysql

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

//使用go语言模拟mysql的slave，查看具体的协议实现
//主要目的就是观察这个协议是如何运作的，也就是主从
//复制是怎么交互的

const (
	mysqlIP   = "172.16.16.46"
	mysqlPort = 3306
)

func SlaveBinlog() {
	establishConnection()
}

//首先和mysql的master节点建立链接
func establishConnection() {
	conn, err := net.Dial("tcp", strings.Join([]string{mysqlIP, strconv.Itoa(mysqlPort)}, ":"))
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	answerByte := make([]byte, 1000)
	_, readErr := conn.Read(answerByte)
	if readErr != nil {
		fmt.Printf(readErr.Error())
		return
	}
	protocol := GreetingProtocol{}
	protocol.parse(answerByte)
	fmt.Printf("查看对象%v\n", protocol)
}

func certification() {

}
