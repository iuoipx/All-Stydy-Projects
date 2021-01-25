package main

import "fmt"

//切片slice

func main() {
	//1.切片的定义
	var s1 []int //定义一个存放int类型的切片
	var s2 []string
	fmt.Println(s1, s2)    //[] []
	fmt.Println(s1 == nil) //true
	fmt.Println(s2 == nil) //true
	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"wqe", "asd", "zxc"}
	fmt.Println(s1, s2)    //[1 2 3] [wqe asd zxc]
	fmt.Println(s1 == nil) //false
	fmt.Println(s2 == nil) //false
	//长度和容量
	fmt.Printf("len(s1):%d,cap(s1):%d\n", len(s1), cap(s1)) //len(s1):3,cap(s1):3
	fmt.Printf("len(s2):%d,cap(s2):%d\n", len(s2), cap(s2)) //len(s1):3,cap(s1):3

	//2.由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] //基于一个数组切割，左包含右不包含,[0,4}
	fmt.Println(s3)
	s4 := a1[:4]            //[0:4]        相当于从0切到四 [0,4}
	s5 := a1[3:]            //[3:len(a1)]  相当于从3切到最后 [3,len(a1)]
	s6 := a1[:]             //[0:len(a1)]  [0,len(a1)]
	fmt.Println(s4, s5, s6) //[1 3 5 7] [7 9 11 13] [1 3 5 7 9 11 13]
	//切片的容量是指底层数组的容量
	fmt.Printf("len(s4):%d,cap(s4):%d\n", len(s4), cap(s4)) //len(s1):4,cap(s1):7
	//底层数组从切片的第一个元素到最后的元素数量
	fmt.Printf("len(s5):%d,cap(s5):%d\n", len(s5), cap(s5)) //len(s1):4,cap(s1):7
	//切片再切片
	s7 := s5[3:]
	fmt.Printf("len(s7):%d,cap(s7):%d\n", len(s7), cap(s7)) //len(s7):1,cap(s7):1
	//切片是引用类型，都指向了底层数组
	fmt.Println("s5:", s5) //s5: [7 9 11 13]
	a1[6] = 1300           //修改了底层数组
	fmt.Println("s5:", s5) //s5: [7 9 11 1300]

}
