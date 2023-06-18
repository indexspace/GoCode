package main

import "fmt"

type Circle struct { // 定义名为Circle的结构体
	radius float64 // 含有一个属性
}

var nums = [5]int{1, 2, 3, 4, 5}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
	fmt.Println()

	for i, val := range nums {
		fmt.Println(i, val)
	}
}

// 该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 { // 方法名用`(c className)`来指定对象
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

// TODO: 数组
