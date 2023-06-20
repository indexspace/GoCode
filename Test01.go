package main

import "fmt"

type Circle struct { // 定义名为Circle的结构体
	radius float64 // 含有一个属性
}

// 该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 { // 方法名用`(c className)`来指定对象
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

var nums1 = [...]int{1, 2, 3, 4, 5} // 数组: [size]type{el1,el2,el3}
var slice1 = []int{1, 2, 3}         // 未指定大小的数组为切片
var m = make(map[string]int, 10)    // 创建map, 可以指定大小

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
	fmt.Println()

	for _, val := range nums1 {
		fmt.Println(val)
	}
	fmt.Println()

	slice1 = append(slice1, 4) // (slice, element) -> slice
	for _, val := range slice1 {
		fmt.Println(val)
	}
	fmt.Println()

	kvs := map[string]string{"a": "apple", "b": "banana"} // map[keyType]valType
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	fmt.Println()

	m["a"] = 1
	m["b"] = 2
	for i := range m {
		fmt.Println(i)
	}
	fmt.Println()

	// 接口类型可以转化其他类型变量的格式: `interface_var.(target_type)`
	var i interface{} = "Hello, World" // 定义接口类型
	str, ok := i.(string)              // 转化为字符串, ok是一个表示是否转化成功的bool值
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
	fmt.Println()

}
