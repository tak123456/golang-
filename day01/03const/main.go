package main

import "fmt"

//批量声明常量时，如果某一行没有赋值，则默认值和上一行相同
const (
	summer = "夏天"
	fall   = "秋天"
	winnter
)

//iota:在const关键字出现时重置为0，const 中每增加一行常量声明,iota计数加一
const (
	a = iota
	a1
	a2
	a3
)
const (
	b  = iota
	b1 = 100
	b2
	b3
)

//插队
const (
	c  = iota
	c1 = 100
	c2 = iota
	c3
)

//多个变量在一行
const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	const pi = 3.14
	fmt.Println(pi)
	fmt.Println(summer + " " + fall + " " + winnter)
	fmt.Println(a, a1, a2, a3)
	fmt.Println(b, b1, b2, b3)
	fmt.Println(c, c1, c2, c3)
	fmt.Println(d1, d2, d3, d4)
	fmt.Println(KB, MB, GB, TB, PB)

}
