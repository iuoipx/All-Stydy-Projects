package main

import "fmt"

//pointer
//对变量进行取地址(&)操作，可以获得这个变量的指针变量
//指针变量的值是指针地址
//对指针变量进行取值(*)操作，可以获得指针变量指向的原变量的值

func main() {
	//1.&:取地址
	n := 18
	p := &n
	fmt.Println(&n)       //0xc0000120a0
	fmt.Printf("%T\n", p) //*int 表示int类型的指针

	//2.*:根据地址取值
	m := *p
	fmt.Println(m) //18

	//error
	// var a *int //int类型指针(内存地址) 值为nli
	// *a = 100
	// fmt.Println(*a)

	//new函数申请一个内存地址
	var a1 = new(int)
	fmt.Println(a1)  //0xc00005e0c0
	fmt.Println(*a1) //0
	*a1 = 100
	fmt.Println(*a1) //100

	//make和new的区别
	//1.make和new都是用来申请内存的
	//2.new很少用，一般用来给基本数据类型申请内存,string、int,返回的是对应类型的指针(*string,*int)
	//3.make是用来给slice、map、chan申请内存的,make函数返回的是对应的这三个类型本身

}
