package main

import "fmt"

func main() {
	//int 和stirng类型的切片
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil, s2 == nil)
	fmt.Println(len(s1), cap(s1), len(s2), cap(s2))
	//初始化
	s1 = []int{0, 1, 2, 3, 4, 5, 6, 7}
	s2 = []string{"北京", "上海"}
	/*由数组得到切片
	*切片是数组的封装
	*切片指向了一个底层数组，
	*切片的长度就是他元素的个数，
	*切片的容量就是他底层数组从第一个元素到最后一个元素
	 */
	s3 := s1[0:2]
	s4 := s1[1:5]
	s5 := s1[:]
	//切片再切片
	s6 := s4[0:2]
	fmt.Println(s1, s2)
	fmt.Println(s3)
	fmt.Println(len(s1), cap(s1), len(s2), cap(s2))
	fmt.Print(s4, s5)
	fmt.Println(len(s4), cap(s4), len(s5), cap(s5), len(s3), cap(s3))
	//切片是引用类型，都指向了一个底层数组
	//可以对底层数组修改，可以对一个切片进行修改
	fmt.Print(s1, s3, s4, s5, s6)
	fmt.Println()
	s1[1] = 10000
	s4[2] = 200
	fmt.Print(s1, s3, s4, s5, s6)
	fmt.Println()
	fmt.Println(len(s6), cap(s6))
}
