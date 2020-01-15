package book

//
// @Date 2020-01-15 10:10-56
// @Author 刘旭
// 基本的数据类型
//

type Money float64
type Price float64

var (
	price = Price(10.23)
)

//这种type的形式目的就是为了给变量一个别名，但是好像还有严格的类型验证的校验权限
func SumMoney(count int) Money {
	return Money(Price(count) * price)
}
