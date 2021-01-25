package main

import "fmt"

//结构体

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	//声明一个person类型的变量p
	var p person
	//通过字段赋值
	p.name = "iuoip"
	p.age = 20
	p.gender = "男"
	p.hobby = []string{"r", "d", "b"}
	fmt.Println(p) //{iuoip 20 男 [r d b]}
	//访问变量p的字段
	fmt.Println(p.name)   //iuoip
	fmt.Printf("%T\n", p) //main.person

	//匿名结构体:多用于临时场景
	var s struct {
		x string
		y int
	}
	s.x = "嘿嘿"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s) //type:struct { x string; y int } value:{嘿嘿 100}
}
