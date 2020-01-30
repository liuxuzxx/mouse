package book

import (
	"fmt"
	"math"
)

//method方法，就和Java的成员方法差不多

func (c Circle) area() float64 {
	return math.Pi * float64(c.Radius*c.Radius)
}

//计算面积
func OperationArea() {
	circle := Circle{
		Point: Point{
			10,
			12,
		},
		Radius: 20,
	}

	fmt.Printf("半径:%d 面积是:%f\n", circle.Radius, circle.area())
}
