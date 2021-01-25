package main

import (
	"fmt"
	"strconv"
)

//strconv

func main() {
	//从字符串解析出对应的数据
	str := "10000"
	// ret1 := int64(str)
	r, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parseint failed,err:", err)
		return
	}
	fmt.Printf("%#v %T\n", r, r) //10000 int64

	//字符串转换int
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt) //10000 int

	i := 68
	ret := string(i) //'D'
	fmt.Println(ret)
	ret1 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n", ret1) //"68"

	//int转string
	ret3 := strconv.Itoa(i)
	fmt.Printf("%#v %T\n", ret3, ret3) //"68" string

	//从字符串解析除布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue) //true bool
	//从字符串解析出浮点数
	floatStr := "1.234"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue) //1.234 float64

}
