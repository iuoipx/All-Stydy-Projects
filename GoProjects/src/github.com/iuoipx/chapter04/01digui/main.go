package main

import "fmt"

//递归:函数自己调用自己
//递归适合处理那种问题相同、问题的规模越来越小的场景
//递归要有一个明确的退出条件

//计算n的阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

//上台阶:n个台阶，一次可以走一步，也可以走两步，有多少种走法
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func main() {
	ret := f(7)
	fmt.Println(ret)
	ret1 := taijie(4)
	fmt.Println(ret1)
}
