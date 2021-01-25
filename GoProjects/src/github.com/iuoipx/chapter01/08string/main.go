package main

import (
	"fmt"
	"strings"
)

//字符串
func main() {
	path := "\tD:\\Projects\\GoProjects\\src\\githup.com\\iuoipx\\chapter01\""
	fmt.Println(path) //        D:\Projects\GoProjects\src\githup.com\iuoipx\chapter01"

	//多行字符串
	s2 := `
		世情薄
		  人情恶
		雨送黄昏花易落
	`
	fmt.Println(s2)

	s3 := `
		D:\Projects\GoProjects\src\githup.com\iuoipx\chapter01
	`

	//字符串相关操作

	fmt.Println(len(s3)) //长度

	//字符串转换
	name := "理想"
	world := "dddddd"
	ss := name + world //拼接字符串
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world) //拼接字符串
	//fmt.Printf("%s%s", name, world)
	fmt.Println(ss1)

	//分隔
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	//包含
	fmt.Println(strings.Contains(ss, "理性")) //fasle
	fmt.Println(strings.Contains(ss, "理想")) //true

	//前缀
	fmt.Println(strings.HasPrefix(ss, "理想")) //true
	//后缀
	fmt.Println(strings.HasSuffix(ss, "理想")) //false

	//字串出现位置
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))     //2
	fmt.Println(strings.LastIndex(s4, "b")) //5

	//拼接
	fmt.Println(strings.Join(ret, "+"))
}
