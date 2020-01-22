package book

import (
	"crypto/sha256"
	"fmt"
)

//复杂类型，就是简单类型的组合关系

//Array 数组，这个很常见
//Array有点特殊啊，并不是类似于我们在Java或者是c语言当中看到的那种，就是一个引用，其实go定位数组类型是根据数组存储的类型和数组长度
//来唯一定位类型的

func ShowLanguage() {
	var language = [...]string{"Java", "Go", "Flutter", "Dart"}
	for index, value := range language {
		fmt.Printf("顺序:%d 值:%s\n", index, value)
	}
	fmt.Printf("Capacity:%d Size:%d \n", cap(language), len(language))
}

type Week int8

func ShowWeek() {
	const (
		Monday Week = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	//数组这个地方可以根据序号进行排位，这个地方可不是map这玩意啊
	//可以根据前面的序号指定数组当中的某些数值
	week := [...]string{Monday: "周一", Sunday: "周日", Tuesday: "周二", Wednesday: "周三", Thursday: "周四", Friday: "周五", Saturday: "周六"}
	for i, v := range week {
		fmt.Printf("序号:%d 星期:%s\n", i, v)
	}

	//Array这么搞的目的可能是为了，我们有的时候想定义一个数组只是想让数组当中的某个位置初始化．
}

func PartInitial() {
	filter := [...]int{99: -1, 123: 90}
	for i, v := range filter {
		fmt.Printf("下标:%d 数值:%d\n", i, v)
	}
}

//验证<The Go Language>这本书的一个例子
func Sha256() {
	//看书的时候对这个使用方法不是很理解，后面理解为:就是[]byte就是一个type,然后肯定就有构造方法了，那么传入一个参数
	sum256 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n", sum256)
}

//Slice芬片，这个分片暂且不知道作用在什么地方

var (
	Month = [...]string{
		1:  "January",
		2:  "February",
		3:  "Match",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
)

func SliceMonth() {
	summer := Month[4:7]
	fmt.Printf("Capacity:%d Size:%d\n", cap(summer), len(summer))
	for i, v := range summer {
		fmt.Printf("Index:%d Value:%s\n", i, v)
	}
}
