package main

import "fmt"

//copy

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1 //赋值
	var a3 = make([]int, 3, 3)
	copy(a3, a1)            //copy
	fmt.Println(a1, a2, a3) //[1 3 5] [1 3 5] [1 3 5]
	a1[0] = 100
	fmt.Println(a1, a2, a3) //[100 3 5] [100 3 5] [1 3 5]

	//从切片中输出元素
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)      //[100 5]
	fmt.Println(cap(a1)) //3

	s1 := [...]int{1, 3, 5}           //数组
	s2 := s1[:]                       //切片
	fmt.Println(s2, len(s2), cap(s2)) //[1 3 5] 3 3
	//1.切片不保存具体的值
	//2.切片对应一个底层数组
	//3.底层数组都是占用一块连续的内存
	fmt.Printf("%p\n", s2)            //0xc00005c1a0
	s2 = append(s2[:1], s2[2:]...)    //修改了底层数组
	fmt.Printf("%p\n", s2)            //0xc00005c1a0
	fmt.Println(s2, len(s2), cap(s2)) //[1 5] 2 3
	//
	fmt.Println(s1) //[1 5 5]
}
