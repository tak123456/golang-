package main

import "fmt"

func main() {
	age := 20
	if age > 21 {
		fmt.Println("大于21岁")
	} else {
		fmt.Println("小于21岁233333")
	}
	//多条件
	if age >= 22 {
		fmt.Println("毕业了")
	} else if age > 21 {
		fmt.Println("大四了")
	} else {
		fmt.Println("大三")
	}
}
