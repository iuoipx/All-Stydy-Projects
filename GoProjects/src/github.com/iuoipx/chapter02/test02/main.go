package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//判断字符串中汉字的数量
	s1 := "hello沙河world世界"
	//1.依次拿到字符串中的字符
	var count int
	for _, value := range s1 {
		//2.判断当前这个字符是不是汉字
		if unicode.Is(unicode.Han, value) {
			count++
		}
	}
	//3.把汉字出现的次数累加得到最终结果
	fmt.Println(count)

	//how do you do 'do'单词出现次数
	s2 := "how do you do"
	//1.把字符串按照空格切割得到切片
	s3 := strings.Split(s2, " ")
	//2.遍历切片存储到一个map
	m1 := make(map[string]int, 10)
	for _, value := range s3 {
		//1.如果原来map中不存在value这个key,那么出现次数等于1
		if _, ok := m1[value]; !ok {
			m1[value] = 1
			//m1["how"]=1
			//m1["do"]=1
			//m1["you"]=1
		} else {
			//2.如果map中存在value这个key，那么出现次数+1
			m1[value]++
			//mi["do"]++
		}
	}
	//3.累加出现的次数
	for key, value := range m1 {
		fmt.Println(key, value)
	}

	//回文判断
	//字符串从左往右读和从右往左读是一样的，那么就是回文
	//黄山落叶松叶落山黄
	ss := "黄山落叶松叶落山黄"
	//把字符串中的字符拿出来放到一个[]rune中
	r := make([]rune, 0, len(ss))
	for _, value := range ss {
		r = append(r, value)
	}
	for i := 0; i < len(r)/2; i++ {
		//ss[index]==ss[len(ss)-1-i]
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
