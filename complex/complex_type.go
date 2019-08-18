package complex

import (
	"fmt"
	"time"
)

var name string = "source"
var count = 90

func ChangeName(name *string) {
	*name = "change name"
	showName()
}

func showName() {
	fmt.Println("查看修改之后的name:", name)
}

var (
	IntPointer    *int
	StringPointer *string
	BoolPointer   *bool
)

type User struct {
	UserName string
	Age      int
	Birthday time.Time
	Address  string
}

//A struct is a collection of fields.
//意思就是说:我其实至于属性，没有方法
func CreateStruct() {
	account := User{
		"liuxu", 28, time.Now(), "刘庄村",
	}
	fmt.Println(account)

	users := [12]User{{}, {}, {}}
	tempUser := users[2:5]
	fmt.Print(tempUser)
}

func SlicesStatement() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	s := primes[2:4]
	fmt.Println(s)
	s[0] = 900
	fmt.Println(primes)

	//And this creates the same array as above, then builds a slice that references it:
	//换句话来说，就只执行了首先是弄成一个Array，之后是利用切片
	numbers := []int{5, 6, 7, 8, 9, 99, 87, 5674, 4, 67, 45}
	fmt.Println(numbers)
}

func SliceOperation() {
	var users []User
	users = append(users, User{}, User{})
	fmt.Println(users)
	for index, value := range users {
		fmt.Println("查看一下信息:", index, value)
	}
}

func MapOperation() {
	var resultMap = make(map[int]string)
	resultMap[1] = "liuxu"
	resultMap[2] = "zxx"

	fmt.Println(resultMap)

	request := map[string]User{
		"liuxu":        {},
		"zhongxiaoxia": {},
	}

	request["xiaoxu"] = User{}

	fmt.Println(request)
}

var PersistenceStrategy = map[string]func(string, int) string{
	"liuxu": func(name string, number int) string {
		return name + string(number)
	},
}

//只要有函数指针，咱们就能写策略模式
func StrategyOperation(name string) {

	strategy := map[string]func(string, int) string{
		"real": func(name string, number int) string {
			fmt.Println("执行了策略:", name)
			return name + string(number)
		},
		"empty": func(name string, number int) string {
			fmt.Println("执行了策略:", name)
			return name + string(number)
		},
	}

	strategy[name]("liuxu", 988)
}

func createFunction() func(string, int) (string, int, string) {
	return func(name string, number int) (string, int, string) {
		return "wo", 3, "liuxu"
	}
}

type Clue struct{

}