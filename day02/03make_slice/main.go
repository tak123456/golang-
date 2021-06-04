package main

import "fmt"

func main() {
	s1 := make([]int, 5, 10) //长度为5容量10的int[]切片
	fmt.Printf("%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
