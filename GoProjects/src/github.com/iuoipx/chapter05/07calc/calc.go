package calc

import "fmt"

//包中的标识符(变量名\函数名\结构体\接口等)如果首字母是小写的，
//表示私有(只能在当前这个包中使用)

func init() {
	fmt.Println("improt调用时自动执行")
}

func Add(x, y int) int {
	return x + y
}
