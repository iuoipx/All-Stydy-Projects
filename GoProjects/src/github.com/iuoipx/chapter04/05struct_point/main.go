package main

import "fmt"

func main() {
	var a int
	a = 100
	b := &a
	fmt.Printf("type a:%T type b:%T\n", a, b)
	//将a的十六进制内存地址打印出来
	fmt.Printf("%p\n", &a) //0xc0000120a0
	fmt.Printf("%p\n", b)  //0xc0000120a0
	fmt.Printf("%v\n", b)  //0xc0000120a0
	fmt.Printf("%p\n", &b) //0xc000006028
}
