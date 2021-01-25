package main

import "fmt"

//for循环

func main() {
	//基本格式
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// //变种1
	// var i = 5
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	//变种2
	// var i = 5
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//无限循环
	// for{
	// 	fmt.Println("1231")
	// }

	//for range循环
	s := "Hello沙河"
	for i, v := range s { //i索引
		//fmt.Println(i, v)
		fmt.Printf("%d %c\n", i, v)
		// 0 H
		// 1 e
		// 2 l
		// 3 l
		// 4 o
		// 5 沙
		// 8 河
	}
}
