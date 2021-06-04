package main

import "fmt"

//rune:int32 ,表示单个字符,byte :uint8
func main() {
	s1 := "白萝卜"
	s2 := []rune(s1)
	s2[0] = '红'
	fmt.Println(s1)
	fmt.Println(string(s2))
	c := byte('A')
	c1 := 'A'
	c2 := "A"
	fmt.Printf("%d %d %s\n", c, c1, c2)
	fmt.Printf("%T %T\n", c, c2)
}
