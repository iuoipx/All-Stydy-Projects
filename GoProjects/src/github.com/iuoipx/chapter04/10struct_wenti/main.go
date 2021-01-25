package main

import "fmt"

//结构体遇到的问题

//结构体初始化
type person struct {
	name string
	age  int
}

//为什么要有构造函数
func newPerson(name string, age int) person {
	//别人调用时，返回一个person类型的变量
	return person{
		name: name,
		age:  age,
	}
}

// func newPerson(name string, age int) *person {
// 	//别人调用时，返回一个person类型的变量
// 	return &person{
// 		name: name,
// 		age:  age,
// 	}
// }

func main() {
	//声明一个person类型的变量p
	var p person
	p.name = "qwe"
	p.age = 18
	fmt.Println(p)
	var p1 person
	p1.name = "asd"
	p1.age = 22

	//键值对初始化
	p2 := person{
		name: "qwe",
		age:  22,
	}
	fmt.Println(p2)

	//值列表初始化
	var p3 = person{
		"asd",
		30,
	}
	fmt.Println(p3)

}
