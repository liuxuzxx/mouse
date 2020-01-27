package book

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
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
	summer := Month[11:]
	summer = append(summer, "Year")
	fmt.Println(Month)
}

func CompareSlice() {
	summer := Month[4:7]
	partSummer := Month[4:]
	for _, outValue := range summer {
		for _, inValue := range partSummer {
			if outValue == inValue {
				fmt.Printf("找到了一样的数据:%s %s\n", outValue, inValue)
			}
		}
	}
}

func Reverse() {
	doReverse(Month[4:7])
	doReverse(Month[:])
	fmt.Println(Month)
}

//注意[]string的意思是传递一个切片，不是一个数组，虽然切片的底层还是数组
func doReverse(source []string) {
	for i, j := 0, len(source)-1; i < j; i, j = i+1, j-1 {
		source[i], source[j] = source[j], source[i]
	}
}

//怪不得Go语言这本书籍一直在说，数组和Slice不一样，Slice是可变化的，数组是固定死的
func AppendElement() {
	var words []rune
	for _, v := range "Hello,世界你好，我很好！" {
		words = append(words, v)
		fmt.Printf("Capacity:%d Size:%d\n", cap(words), len(words))
	}
	fmt.Println(words)
}

//
//总体来说Array和Slice确实是大不一样啊，Array具有固定性质，类型都给固定死了，但是Slice是可以随便变化的
//Array根据存储的元素类型和元素个数为综合类型，但是Slice只是以存储的元素作为类型
//

//进入map类型的研究，基本上所有的语言都有这些个特性数据结构，所以不用那么仔细的查看学习

func InitialMap() {
	students := map[int64]string{
		1: "Linux",
		2: "Suse",
		3: "Feroa",
	}

	students[9] = "Rondo"
	delete(students, 2)
	//But a map element is not a variable,and
	//One reason that we can’t take the address of a map element is that growing a map might cause
	//rehashing of existing elements into new storage locations, thus potentially invalidating the
	//address.

	for key, value := range students {
		fmt.Printf("Key:%d Value:%s\n", key, value)
	}

	var languages map[string]Language
	languages = make(map[string]Language)
	languages["Java"] = Language{
		name:     "Java",
		age:      24,
		birthday: time.Time{},
	}
}

type Language struct {
	name     string
	age      int8
	birthday time.Time
}

//我感觉<The Go Language>这本书关于map不能进行&操作是一个错误的没有逻辑闭环的描述操作
//首先是：你说他不是变量，那么不是变量如何可以赋值，你说他会重新rehash,那么Slice也会重新分配内存啊
//Slice进行扩充的时候为什么指针地址没有变化
//这个应该是获取的是Slice类型的地址，并不是里面元素的地址
//There is a major difference between slices and maps: Slices are backed by a backing array and maps are not.
//
//If a map grows or shrinks a potential pointer to a map element may become a dangling pointer pointing into nowhere (uninitialised memory). The problem here is not "confusion of the user" but that it would break a major design element of Go: No dangling pointers.
//
//If a slice runs out of capacity a new, larger backing array is created and the old backing array is copied into the new; and the old backing array remains existing. Thus any pointers obtained from the "ungrown" slice pointing into the old backing array are still valid pointers to valid memory.
//
//If you have a slice still pointing to the old backing array (e.g. because you made a copy of the slice before growing the slice beyond its capacity) you still access the old backing array. This has less to do with pointers of slice elements, but slices being views into arrays and the arrays being copied during slice growth.
//
//Note that there is no "reducing the backing array of a slice" during slice shrinkage.
//只能说上面的解释稍微合理一些
//没有想到一点疑问竟然可以挖掘到了Go语言的底层

func AddressSliceAndMap() {
	languages := make([]Language, 2, 2)
	languages[0] = Language{
		name:     "Java",
		age:      1,
		birthday: time.Time{},
	}

	languages[1] = Language{
		name:     "Dart",
		age:      3,
		birthday: time.Time{},
	}

	first := reflect.ValueOf(&languages[1]).Pointer()
	fmt.Println("地址是:", first)
	languages = append(languages, Language{
		name:     "Flutter",
		age:      5,
		birthday: time.Time{},
	})

	second := reflect.ValueOf(&languages[1]).Pointer()
	fmt.Println("地址是:", second)
}

//Structs，结构体，一个迈向面向对象变成的桥梁
//构造一个Node节点的结构体，数据和前面的节点指针和后面的节点指针
//书上所谓复杂的例子其实就是Java的面向对象的思想的封装思维
type Node struct {
	data string
	pre  *Node
	next *Node
}

type Point struct {
	X int64
	Y int64
}

type Circle struct {
	Point
	Radius int64
}

type Wheel struct {
	Circle
	Spokes int64
}

//JSON的数据形式，就是一种序列化的操作
type AncientArticle struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	Context    string `json:"context"`
	CreateTime time.Time
	Commit     string `json:"comment"`
}

func ConvertToJSON() {
	var article = AncientArticle{
		Title:      "诵读少妇之人枢轴",
		Author:     "李白",
		Context:    "成曲阜三嗪，凤眼望五金．浴巾离别意，同时还有人",
		CreateTime: time.Time{},
		Commit:     "作者的离别制作",
	}

	data, err := json.MarshalIndent(article, "", "	")
	if err != nil {
		fmt.Println("出现一个错误！")
	}
	fmt.Printf("%s\n", data)
}

//Text And Template这种就是类似于之前的structs的那种jstl的形式处理标签
//具体的例子我就不实验了，没有必要在某些小细节上一直的扣啊扣的
