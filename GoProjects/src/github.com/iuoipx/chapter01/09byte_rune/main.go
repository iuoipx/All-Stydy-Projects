package main

import "fmt"

//byte和rune类型

//Go语言为了处理非ASCII码类型的字符，定义了新的rune类型

func main() {
	s := "iuoipx"
	//len()获取byte字节的数量
	n := len(s)
	fmt.Println(n) //6

	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(s[i])        //105 117 111 105 112 120
	// 	fmt.Printf("%c\n", s[i]) //%c:打印字符
	// }

	// for _, c := range s {
	// 	//fmt.Println(c)
	// 	fmt.Printf("%c\n", c)
	// }

	//字符串修改
	s2 := "白萝卜"             // =>'白' '萝' '卜'
	s3 := []rune(s2)        //把字符串强制转换成一个rune切片  [白,萝,卜]
	s3[0] = '红'             //双引号表示字符串，单引号表示字符
	fmt.Println(string(s3)) //把rune切片强制转换为字符串  红萝卜

	//字符串中,英文字符 byte ,中文字符 rune
	c1 := "红"
	c2 := '红'                           //rune(int32)
	fmt.Printf("c1:%T c2:%T\n", c1, c2) //c1:string c2:int32
	c3 := "H"
	c4 := 'H'
	fmt.Printf("c3:%T c4:%T\n", c3, c4) //c3:string c4:int32

	//类型转换
	n1 := 10 //int
	var f float64
	f = float64(n1)
	fmt.Println(f)        //10
	fmt.Printf("%T\n", f) //float64
}
