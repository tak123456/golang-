package main

import "fmt"

func main() {
	//定义数组，长度和类型
	var a1 [3]int32
	var a2 [4]int32
	fmt.Printf("%T %T\n", a1, a2)
	fmt.Println(a1, a2)
	//数组的初始化,根据索引初始化
	a3 := [3]int32{1, 2, 3}
	a4 := []int32{1, 2, 3, 4, 5}
	a5 := [5]int32{0: 1, 4: 2}
	//根据索引访问数组
	citys := []string{"北京", "上海", "深圳"}
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
	//普通遍历
	for i := 0; i < len(citys); i++ {
		fmt.Print(citys[i] + " ")
	}
	fmt.Println()
	//range遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}
	//二维数组
	arr := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(arr)
	//先定义后使用
	var arr1 [3][2]int
	arr1 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(arr1)
	//索引遍历
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	//range遍历
	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}
	//数组是值类型
	arr2 := [3]int32{1, 2, 3}
	arr3 := arr2
	arr3[0] = 100
	fmt.Println(arr2, arr3)
}
