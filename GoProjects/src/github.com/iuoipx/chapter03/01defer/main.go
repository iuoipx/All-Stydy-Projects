package main

import "fmt"

//defer多用于函数结束之前释放资源(文件句柄、数据库连接、socket连接)
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿") //defer把后面的语句延迟到函数即将返回时执行
	defer fmt.Println("呵呵") //一个函数中可以有多个defer语句
	defer fmt.Println("哈哈") //多个defer语句按照先进后出(后进先出)的顺序延迟执行
	fmt.Println("end")
}

func main() {
	deferDemo()
}
