//自定义自己的包名
package packages

import (
	"fmt"
	"math/rand"
)

func SelfPackage(name string) {
	fmt.Println("使用自己的包:", name)
	privateFunction()
}

//开头字母小写的是private 形式权限
//只能在本包内使用
func privateFunction(){
	fmt.Println("This is a private function:",rand.Intn(100))
}
