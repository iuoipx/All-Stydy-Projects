package main

import (
	"fmt"

	//包名路径是从GOPATH/src下开始写起,路径分隔符用/
	//想被别的包调用的标识符都要首字母大写
	//单行导入和多行导入
	//导入包的时候可以指定别名
	//导入包不想使用包内部的标识符，需要使用匿名导入
	//每个包导入的时候会自动执行一个名为init()的函数，它没有参数和返回值，也不能手动调用
	//多个包中都定义了init()函数,则它们的执行顺序main->A->B->C
	calc "githup.com/iuoipx/chapter05/07calc"
)

func init() {
	fmt.Println("自动执行!")
}

func main() {
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}
