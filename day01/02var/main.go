package main

import "fmt"

//批量声明
var (
	name string
	age  int
	sex  bool
)

/*非全局变量声明后必须使用，全局变量没有这个限制
*:=不能使用在函数外边
*函数外的每个语句必须都以关键字（var,const,func等）
 */

func main() {
	name = "田安康"
	age = 22
	// sex = true
	//先定义变量，再赋值
	var score int
	score = 3
	//声明变量的同时赋值
	var s1 string = "s1s1"
	//类型推导根据值推导出类型
	var s2 = "s2s2"
	//简短变量声明,只能在函数中使用
	s3 := "s3s3"
	//匿名变量,多用于占位符，表示忽略值
	x, _ := foo()
	_, y := foo()
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(s3)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("var")
	fmt.Println(sex)
	fmt.Println(score)
	// fmt.Printf("name:%s\n", name)
	// fmt.Println(age)
}
func foo() (int, string) {
	return 10, "哈哈"
}
