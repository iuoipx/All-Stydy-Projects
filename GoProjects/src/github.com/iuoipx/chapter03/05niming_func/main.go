package main

import "fmt"

//匿名函数

func main() {
	//函数内部没有办法声明带名字函数
	//匿名函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	//如果只调用一次的函数，还可以简写成立即执行函数
	func(x, y int) {
		fmt.Println(x + y)
		fmt.Println("hello")
	}(100, 200) //加()调用函数
}
