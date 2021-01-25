package main

import "fmt"

//函数

//函数是一段代码的封装
//把一段逻辑抽象出来封装到一个函数中，每次用到它时直接使用函数名调用
//使用函数能够让代码结构更清晰、更简洁
//函数的定义
func sum(x int, y int) (ret int) {
	return x + y
}

//没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

//没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

//没有参数但有返回值
func f3() int {
	return 3
}

//返回值可以命名也可以不命名
//命名的返回值就相当与在函数中声明一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return //使用命名返回值return后面可以省略
}

//多个返回值
func f5() (int, string) {
	return 1, "sad"
}

//参数类型简写
//当参数中连续多个参数的类型一致时，我们可以把非最后一个参数的类型省略
func f6(x, y int, a, b, c string) int {
	return x + y
}

//可变长参数
//可变长参数必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x, y) //y的类型是切片 []int
}

//Go语言中没有默认参数概念

func main() {
	r := sum(1, 2)
	fmt.Println(r)

	_, n := f5()
	fmt.Println(n) //sad

	f7("下雨了")                     //下雨了 []
	f7("下雨了", 12, 213, 34, 56, 5) //下雨了 [12 213 34 56 5]
}
