package main

import (
	"fmt"
	"strings"
)

func main() {
	//转移字符\
	path := "D:\\GO\\src"
	fmt.Println(path)
	//多行字符输出,反引号，可以原样输出，不适用转义字符
	names := `
	小明
	    小   红
	小李
	`
	fmt.Println(names)
	path1 := `"D:\GO\src1"`
	fmt.Println(path1)
	//字符串长度
	fmt.Println(len(path))
	//字符串切割
	ret := strings.Split(path, "\\")
	fmt.Println(ret)
	//字符串合并，+或者fmt.Sprintf()
	name := "小黑"
	home := "中华人民共和国"
	nameHome := fmt.Sprintf("%s%s\n", name, home)
	fmt.Println(nameHome)
	//判断是否包含
	fmt.Println(strings.Contains(name, "黑"))
	//前缀
	fmt.Println(strings.HasPrefix(home, "中华"))
	//后缀
	fmt.Println(strings.HasSuffix(home, "共和国"))
	s1 := "abcdabcd"
	//子串出现的位置
	fmt.Println(strings.Index(s1, "ab"))
	fmt.Println(strings.LastIndex(s1, "c"))
	//拼接
	fmt.Println(strings.Join(ret, "-"))

}
