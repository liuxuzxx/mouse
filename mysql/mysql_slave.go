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
	mysqlIP   = "localhost"
	mysqlPort = 3306
)

var (
	conn net.Conn
	err  error
)

func SlaveBinlog() {
	establishConnection()
}

//首先和mysql的master节点建立链接
func establishConnection() {
	conn, err = net.Dial("tcp", strings.Join([]string{mysqlIP, strconv.Itoa(mysqlPort)}, ":"))
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
	fmt.Printf("协议长度:%d\n", protocol.payloadLength)
	fmt.Printf("序列号:%d\n", protocol.sequenceId)
	fmt.Printf("版本号:%d\n", protocol.protocolVersion)
	fmt.Printf("数据库版本信息:%s\n", protocol.serverVersionInformation)
	fmt.Printf("线程ID:%d\n", protocol.threadId)
	fmt.Printf("挑战随机数(第一部分):%s 长度:%d\n", protocol.scrambleDataPart1, len(protocol.scrambleDataPart1))
	fmt.Printf("能力标记:%d\n", protocol.serverCapability)
	fmt.Printf("字符编码:%d\n", protocol.serverCharset)
	fmt.Printf("插件名字:%s\n", protocol.scramblePluginName)

	certification(append([]byte(protocol.scrambleDataPart1), []byte(protocol.scrambleDataPart2)...))
}

func certification(seed []byte) {
	request := CertificationRequest{
		userName:     "root",
		databaseName: "ancient_article",
		password:     "root",
		pluginName:   "mysql_native_password",
		scrambleData: seed,
	}
	result := request.encode()
	count, writeErr := conn.Write(result)
	if writeErr != nil {
		fmt.Printf("查看错误信息:%v\n", writeErr.Error())
	}
	fmt.Printf("查看写入数量:%d\n", count)

	authResult := make([]byte, 1000)
	readCount, readErr := conn.Read(authResult)
	if readErr != nil {
		fmt.Printf("读取错误:%v\n", readErr.Error())
	}
	fmt.Printf("读取内容长度:%d %s\n", readCount, string(authResult))
}
