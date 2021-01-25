package main

import "fmt"

//运算符

func main() {
	var (
		a = 5
		b = 2
	)
	//算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句,不能放在=的右边赋值  ==> a=a+1
	b--

	//关系运算符
	fmt.Println(a == b) //Go语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b) //不等于
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a > b)
	fmt.Println(a < b)

	//逻辑运算符
	//如果年龄大于18并且年龄小于60
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("苦逼上班的!")
	} else {
		fmt.Println("不用上班")
	}

	//如果年龄小于18或者年龄大于60
	if age < 18 || age > 60 {
		fmt.Println("不用上班")
	} else {
		fmt.Println("苦逼上班的!")
	}

	//not取反
	isMarried := false
	fmt.Println(isMarried)
	fmt.Println(!isMarried)

	//位运算:针对的是二进制数
	//5的二进制:101
	//2的二进制: 10

	//&:按位与(两位均为1才为1)
	fmt.Println(5 & 2) //000  =>0
	//|:按位或(两位有一个为1就为1)
	fmt.Println(5 | 2) //111  =>7
	//^:按位异或(两位不一样为1)
	fmt.Println(5 ^ 2) //111  =>7
	//<<:将二进制位左移指定位数
	fmt.Println(5 << 1) //1010 =>10
	//>>:将二进制位右移指定的位数
	fmt.Println(5 >> 1) //10  =>2

	// m := int8(1)  //只能存8位
	// fmt.Println(m << 10)

	//赋值运算符,用来给变量赋值
	var x int8
	x = 10
	x += 2 //x=x+2
	x -= 2 //x=x-2
	x *= 2 //x=x*2
	x /= 2 //x=x/2
	//x %= 2  //x=x%2
	x <<= 2 //x=x<<2
	x &= 2  //x=x&2
	x |= 3  //x-x|3
	x ^= 4  //x=x^4
	x >>= 2 //x=x>>2
	fmt.Println(x)
}
