package main

import (
	"fmt"
	"math"
)

func main() {
	//go中浮点数默认为64位
	f1 := 1.234
	fmt.Println(math.MaxFloat32)
	fmt.Printf("%T\n", f1)
	//定义32位的浮点数
	f2 := float32(1.234)
	fmt.Printf("%T\n", f2)
	//复数,实数和虚数都是32位
	var c1 complex64 = 1 + 2i
	fmt.Println(c1)
	//都是64位
	var c2 complex128 = 1 + 2i
	fmt.Println(c2)
	fmt.Printf("%T %T\n", c1, c2)
}
