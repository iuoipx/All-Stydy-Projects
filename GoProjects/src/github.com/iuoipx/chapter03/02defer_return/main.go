package main

import "fmt"

//Go语言中函数的return不是原子操作，在底层时分为两步来执行
//第一步:返回值赋值
//第二布:真正的RET返回
//函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间

func f1() int {
	x := 5
	defer func() {
		x++ //修改的是x不是返回值
	}()
	return x //返回值是int类型 5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回值是int类型的x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //返回值是int类型的y
}

func f4() (x int) {
	defer func(x int) {
		x++ //改变的是函数中的副本
	}(x) //把x当作参数传到匿名函数中
	return 5 //返回值是int类型的x
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
