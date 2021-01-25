package main

import "fmt"

//结构体是值类型

type person struct {
	name, gender string
}

//Go语言中函数传参永远传的是拷贝
func f(x person) {
	x.gender = "女" //修改的是副本的gender
}

func f2(x *person) {
	// (*x).gender = "女" //根据内存地址找个原变量，修改的是原变量
	x.gender = "女" //语法糖,自动根据指针找对应的变量
}

func main() {
	var p person
	p.name = "iuoip"
	p.gender = "男"
	f(p)
	fmt.Println(p) //{iuoip 男}
	f2(&p)
	fmt.Println(p) //{iuoip 女}

	//1.结构体指针1
	var p2 = new(person)    //new返回一个指针
	fmt.Printf("%T\n", p2)  //*main.person
	fmt.Printf("%p\n", p2)  //0xc000004520  p2保存的值就是一个内存地址
	fmt.Printf("%p\n", &p2) //0xc000006030  求p2的内存地址

	//2.结构体指针2
	//2.1 key value初始化
	var p3 = &person{
		name:   "asd",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3) //main.person{name:"asd", gender:"男"}
	fmt.Printf("%p\n", p3)  //0xc000004540
	fmt.Printf("%p\n", &p3) //0xc000006038

	//2.2使用值列表的形式初始化，值的顺序要和结构体定义时字段的顺序一致
	p4 := person{
		"qwe",
		"女",
	}
	fmt.Printf("%#v\n", p4) //main.person{name:"qwe", gender:"女"}
}
