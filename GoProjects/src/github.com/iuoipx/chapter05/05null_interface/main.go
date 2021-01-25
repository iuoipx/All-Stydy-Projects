package main

import "fmt"

//空接口
//interface:关键字
//interface{}:空接口类型

//1.作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	//2.作为map的value类型
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "zcx"
	m1["age"] = 30
	m1["merried"] = true
	m1["hobby"] = [...]string{"s", "d", "r"}
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)
}
