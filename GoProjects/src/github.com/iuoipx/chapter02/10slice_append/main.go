package main

import "fmt"

func main() {
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15}
	s1 := a1[:]

	//输出索引为1的3
	s1 = append(a1[:1], a1[2:]...)
	fmt.Println(s1) //[1 5 7 9 11 13 15]
	fmt.Println(a1) //[1 5 7 9 11 13 15 15]
}
