package main

import "fmt"

//结构体占有一块连续的内存空间

type x struct {
	a int8
	b int8
	c int8
}

func main() {
	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n", &(m.a)) //0xc0000120a0
	fmt.Printf("%p\n", &(m.b)) //0xc0000120a1
	fmt.Printf("%p\n", &(m.c)) //0xc0000120a2
}
