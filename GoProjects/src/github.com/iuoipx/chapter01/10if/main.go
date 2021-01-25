package main

import "fmt"

//if条件判断
func main() {
	// age := 19
	// if age > 18 {
	// 	fmt.Println("true>18")
	// } else {
	// 	fmt.Println("flase<18")
	// }

	//作用域
	//age变量此时只在if条件判断语句中生效
	if age := 19; age > 18 {
		fmt.Println("true>18")
	} else {
		fmt.Println("flase<18")
	}

	//fmt.Println(age) //underfind
}
