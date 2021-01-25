package main

import "fmt"

//fmt占位符
func main() {
	var n = 100
	//查看类型
	fmt.Printf("%T\n", n)
	//查看变量值
	fmt.Printf("%v\n", n)
	//转换二进制
	fmt.Printf("%b\n", n)
	//十进制
	fmt.Printf("%d\n", n)
	//八进制
	fmt.Printf("%o\n", n)
	//十六进制
	fmt.Printf("%x\n", n)
	var s = "Hello"
	//字符串
	fmt.Printf("%s\n", s)

	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)

}
