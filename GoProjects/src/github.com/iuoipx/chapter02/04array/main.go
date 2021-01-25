package main

import "fmt"

//数组
//存放元素的容器
//必须指定存放的元素的类型和容量 (长度)
//数组的长度是数组类型的一部分

func main() {
	var a1 [3]bool
	var a2 [4]bool

	fmt.Printf("a1:%T a2:%T\n", a1, a2) //a1:[3]bool a2:[4]bool

	//数组的初始化
	//如果不初始化，默认元素都是零值(布尔值:false,整形和浮点型:0,字符串:"")
	fmt.Println(a1, a2) //[false false false] [false false false false]
	//1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	//2.初始化方式2
	///... 根据初始值自动推断数组的长度是多少
	a3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a3)
	//3.初始化方式3
	//根据索引初始化
	a4 := [5]int{0: 1, 4: 2}
	fmt.Println(a4)

	//数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}
	//1.根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//2.for range遍历
	for index, value := range citys {
		fmt.Println(index, value)
	}
	//多维数组
	//[[1,2],[3,4],[5,6]]
	var a5 [3][2]int
	a5 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a5)

	//多维数组的遍历
	for _, v1 := range a5 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2) //[1 2 3] [100 2 3]
}
