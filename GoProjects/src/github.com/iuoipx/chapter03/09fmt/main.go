package main

import "fmt"

func main() {
	//Printf("格式化字符串",值)
	//%T:查看类型
	//%d:十进制数
	//%b:二进制数
	//%o:八进制数
	//%x:十六进制数
	//%c:字符
	//%s:字符串
	//%p:指针
	//%v:值
	//%f:浮点数
	//%t:布尔值

	var m1 = make(map[string]int, 1)
	m1["理想"] = 100
	fmt.Printf("%v\n", m1)  //map[理想:100]
	fmt.Printf("%#v\n", m1) //map[string]int{"理想":100}

	num := 90
	fmt.Printf("%d%%\n", num) //90%

	fmt.Printf("%v\n", 100)

	fmt.Printf("%q\n", 65) //'A' 整数->字符

	fmt.Printf("%b\n", 3.141592637247232) //浮点数和复数

	fmt.Printf("%q\n", "理想") //字符串

	//获取用户输入
	var s string
	fmt.Scan(&s)
	fmt.Println("输入的内容是:", s)

	var (
		name  string
		age   int
		class string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
}