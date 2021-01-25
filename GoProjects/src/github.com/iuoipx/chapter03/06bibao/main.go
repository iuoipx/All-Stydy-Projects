package main

import "fmt"

//闭包
//闭包是一个函数，这个函数包含了外部作用域的一个变量

//底层原理:
//1.函数可以作为返回值
//2.函数内部查找变量的顺序,现在自己内部找，找不到往外层找
//闭包=函数+外部变量的引用

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

//要求 f1(f2)
func f1(f func()) {
	fmt.Println("--f1--")
	f()
}

func f2(x, y int) {
	fmt.Println("--f2--")
	fmt.Println(x + y)
}

func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		//fmt.Println(x + y)
		f(x, y)
	}
	return tmp
}

func main() {
	ret := adder(100)
	ret2 := ret(200)
	fmt.Println(ret2)

	ret3 := f3(f2,100, 200)
	f1(ret3)
}
