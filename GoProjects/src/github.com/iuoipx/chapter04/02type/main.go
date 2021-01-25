package main

import "fmt"

//自定义类型和类型别名

//type后面跟的是类型
type myInt int     //自定义类型
type yourInt = int //类型别名

func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n) //main.myInt

	var m yourInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m) //int

	var c rune
	c = '中'
	fmt.Printf("%T\n", c) //int32
}
