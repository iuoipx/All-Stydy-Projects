package main

import (
	"fmt"
	"sort"
)

//切片练习

func main() {
	var a = make([]string, 5, 10) //创建切片，len 5，cap 10
	fmt.Println(a)                //[    ]
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a) //[     0 1 2 3 4 5 6 7 8 9]

	var a1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:]) //对切片进行排序
	fmt.Println(a1)
}
