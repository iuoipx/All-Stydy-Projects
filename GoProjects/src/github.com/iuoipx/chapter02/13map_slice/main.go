package main

import "fmt"

//map和slice组合

func main() {
	//定义元素类型为map的切片
	var s1 = make([]map[int]string, 1, 10)
	//对map进行初始化
	s1[0] = make(map[int]string, 1)
	s1[0][01] = "zhangsan"
	s1[0][02] = "lisi"
	//s1[1][10] = "dsad"  //err:index out of range [1] with length 1  索引越界
	fmt.Println(s1) //[map[1:zhangsan 2:lisi]]

	//定义值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30}
	fmt.Println(m1) //map[北京:[10 20 30]]
}
