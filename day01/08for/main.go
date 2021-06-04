package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	var i = 2
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	s := "hello中国"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
	var c1 rune = '1'
	var c2 rune = '中'
	fmt.Println(c1, c2)
	fmt.Printf("%T %T\n", c1, c2)
}
