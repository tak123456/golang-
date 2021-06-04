package main

import "fmt"

func main() {
	var age = 16
	fmt.Printf("%d\n", age)
	//十进制换成二进制
	fmt.Printf("%b\n", age)
	//十进制换成八进制
	fmt.Printf("%o\n", age)
	//十进制换成16进制
	fmt.Printf("%x\n", age)
	//八进制
	a := 010
	fmt.Println(a)
	//16进制
	b := 0x0010
	fmt.Println(b)
	//查看变量类型
	fmt.Printf("%T\n", b)
	//声明int8,int16类型变量,否则为int类型默认
	c, d := int8(9), int16(10)
	fmt.Println(c, d)
	fmt.Printf("%T %T", c, d)
}
