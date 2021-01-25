package main

import "fmt"

//append()为切片追加元素

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	//s1=[北京 上海 深圳] len(s1)=3 cap(s1)=3
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	//s1[3] = "广州" //错误写法，编译错误:索引越界  index out of range[3] with length 3

	//调用append函数必须用原来的切片变量接受返回值
	//append追加元素，原来的底层数组放不下的时候，Go底层就会把底层数组换一个
	//必须用变量接受append返回值
	s1 = append(s1, "广州")
	//s1=[北京 上海 深圳 广州] len(s1)=4 cap(s1)=6
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "成都")
	//s1=[北京 上海 深圳 广州 杭州 成都] len(s1)=6 cap(s1)=6
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 := []string{"武汉", "西安", "苏州"}
	s1 = append(s1, s2...) //...表示拆开
	//s1=[北京 上海 深圳 广州 杭州 成都 武汉 西安 苏州] len(s1)=9 cap(s1)=12
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
