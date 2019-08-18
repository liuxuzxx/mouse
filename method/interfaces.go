package method

import (
	"fmt"
	"time"
)

//接口就是java的接口思想啊
type ToString interface {
	ConvertToString() string
}

type EmptyInterface interface {
}

type User struct {
	UserName string
	Password string
	Age      int
	Birthday time.Time
}

func (u User) ConvertToString() string {
	return u.UserName + ":" + u.Password + ":" + string(u.Age)
}

func InterfaceOperation() {
	var inter ToString
	user := User{UserName: "liuxu", Password: "root", Age: 28, Birthday: time.Now()}
	inter = user
	information := inter.ConvertToString()
	fmt.Println("打印的信息内容:", information)

	var empty EmptyInterface
	empty = user

	fmt.Println("打印空的interface:", empty)

}
