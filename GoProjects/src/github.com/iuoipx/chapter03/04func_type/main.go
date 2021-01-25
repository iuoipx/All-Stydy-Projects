package main

import "fmt"

//函数类型

func f1() {
	fmt.Println("Hello")
}

func f2() int {
	return 10
}

func f4(x, y int) int {
	return x + y
}

//函数也可以作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func ff(a, b int) int {
	return a + b
}

//函数还可以作为返回值
//有一个 func()int的参数x，有一个func(int,int)int类型的返回值
func f5(x func() int) func(int, int) int {
	return ff
}

func main() {
	a := f1               //:=变量声明并且赋值
	fmt.Printf("%T\n", a) //func()
	b := f2
	fmt.Printf("%T\n", b) //func() int
	f3(f2)                //10
	f3(b)                 //10

	c := f5(f2)
	fmt.Printf("%T\n", c) //func(int, int) int
}
