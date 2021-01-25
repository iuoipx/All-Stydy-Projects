package main

import "fmt"

func main() {
	array := [...]int{1, 3, 5, 7, 8}
	var result int
	for _, value := range array {
		result += value
	}
	fmt.Println(result) //24

	//定义两个for循环，外层的从第一个开始遍历
	//内层的for循环从外层后面的那个开始找
	//它们两个数的和为8
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
