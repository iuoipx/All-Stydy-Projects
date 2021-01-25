package main

import "fmt"

//流程控制之跳出for循环

func main() {
	//当i=5时跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break //跳出for循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	//当i=5时,跳过此次for循环(不执行for循环内部的打印语句),继续下次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue //跳出此次循环，继续下一次循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	//跳出多层for循环
	flag := false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break //跳出内层for循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			break //跳出外层循环
		}
	}

	//goto+label实现跳出多层for循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto XX //跳到指定标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}

XX: //label标签
	fmt.Println("over")
}
