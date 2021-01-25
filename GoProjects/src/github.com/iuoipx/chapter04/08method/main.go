package main

import "fmt"

//方法
//标识符:变量名 函数名 类型名 方法名
//Go语言中如果标识符首字母是大写的，就表示对外部可见(暴露的，公有的)

type dog struct {
	name string
	age  int
}

//构造函数
func newDog(name string, age int) dog {
	return dog{
		name: name,
		age:  age,
	}
}

//方法时作用与指定类型的函数
//接收者表示的是调用方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:wangawng", d.name)
}

//使用值接收者:传拷贝进去
//值类型传参拷贝，不影响原值
func (d dog) guonian() {
	d.age++
}

//指针接收者:传内存地址进去
func (d *dog) zhenguonian() {
	d.age++
}

func (d *dog) dream() {
	fmt.Println("no work")
}

func main() {
	d1 := newDog("xza", 18)
	d1.wang()
	fmt.Println(d1.age) //18
	d1.guonian()
	fmt.Println(d1.age) //18
	d1.zhenguonian()    //把内存地址传进去
	fmt.Println(d1.age) //19
	d1.dream()
}
