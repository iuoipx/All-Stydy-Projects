package main

import "fmt"

//Go语言推荐使用驼峰式命名
//var student_name string
var studentName string //推荐使用
//var StudentName string

//声明变量
// var name string
// var age int
// var isOk bool

//批量声明
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	//Go语言中非全局变量声明必须使用，不使用就不能编译
	//在终端输出要打印的内容
	fmt.Print(isOk)
	fmt.Println()
	//%s:占位符，使用name这个变量的值去替换占位符
	fmt.Printf("name:%s\n", name)
	//打印完指定的内容之后会在后面加一个换行符
	fmt.Println(age)

	//声明变量同时赋值
	var s1 = "whb"
	//类型推导(根据值判断该变量是什么类型)
	fmt.Println(s1)
	//同一个作用域({})中不能重复声明同名的变量
	//s1 := "10"
	var s2 = "20"
	fmt.Println(s2)

	//简短变量声明 只能在函数里用
	s3 := "哈撒还"
	fmt.Println(s3)

	//匿名变量是一个特殊的变量: _
}
