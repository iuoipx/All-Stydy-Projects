package main

import "fmt"

//接口
//接口是一种类型，是一种特殊的类型，它规定了变量有哪些方法
//不关心一个变量是什么类型，只关心能调用它的什么方法
//一个变量如果实现了接口中的所有方法,那么这个变量就实现了这个接口,可以称为这个接口类型的变量

type speaker interface {
	speak() //只要实现了speak方法的变量都是speaker类型,方法签名
}

type cat struct {
}

type dog struct {
}

func (c cat) speak() {
	fmt.Println("mmmm~")
}

func (d dog) speak() {
	fmt.Println("wwww~")
}

func da(x speaker) {
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	da(c1)
	da(d1)

	var ss speaker //定义一个接口类型:speaker 的变量:ss
	ss = c1
	fmt.Println(ss)
}
