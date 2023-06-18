package main // 该文件在main包下

import "fmt" // 需要用到fmt的格式化函数

var ( // 这种声明变量的方法一般用于声明全局变量
	name = "czp"
	age  = 21
)

const (
	zero = iota
	one  = iota // iota 可理解为 const 语句块中的行索引
	two
	tree = iota
	four
	five = iota
)

func main() {
	var a, b int = 1, 2   // 声明变量并初始化   `var` vName** vType = val**
	c := 3                // := 是声明并初始化的简写符号, "初始化声明" 这种格式只能在函数体中出现
	const e float32 = 2.6 // 常量和变量的声明初始化方式基本一致  把`var`换成`const`即可

	fmt.Println(a, b, c) // 被声明的局部变量必须要使用, 否则报错 !
	a, b, _ = b, a, c    // `_`是只写变量
	fmt.Println(a, b, c)
	fmt.Println()

	fmt.Println(zero, one, two, tree, four, five)
	fmt.Println()

	// TODO: select语句
	// TODO: 各种数据类型 (包括但不限于指针)

	// for 有四种形式 除了Java中的普通for和增强迭代for,
	// 还有 for condition {}  , 和 for {} (/无限循环)

	map1 := make(map[int]float32) // new Map<int,float>()
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 读取 key 和 value
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}
	fmt.Println()

	str1, str2 := swap("Google", "Runoob")
	fmt.Println(str1, str2)
	fmt.Println()

	Seq := getSequence()
	fmt.Println(Seq(), Seq(), Seq(), Seq(), Seq())

	Seq = getSequence()
	fmt.Println(Seq(), Seq(), Seq(), Seq(), Seq())

}

func swap(x, y string) (string, string) {
	return y, x
}

func getSequence() func() int { // 闭包: 返回一个拥有局部静态变量的函数
	i := 0              // 初始化
	return func() int { // 每次+1返回
		i += 1
		return i
	}
}
