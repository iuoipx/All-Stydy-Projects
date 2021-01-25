package main

import "fmt"

//类型断言
//我想知道空接口接收的值具体是什么

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("error")
	} else {
		fmt.Println("传进来的是一个字符串:", str) //传进来的是一个字符串: 123
	}
}

//类型断言2
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串:", t)
	case int:
		fmt.Println("是一个int:", t)
	case bool:
		fmt.Println("是一个bool:", t)
	}
}

func main() {
	assign("123")
	assign2(19) //是一个int: 19
}
