package main

import "fmt"

//map

func main() {
	var m1 map[string]int         //nil还没有初始化(没有在内存中开辟空间)
	fmt.Println(m1 == nil)        //true
	m1 = make(map[string]int, 10) //10时键值对的数量
	m1["理想"] = 18
	m1["jiwuming"] = 35
	fmt.Println(m1)
	fmt.Println(m1["理想"])
	fmt.Println(m1["sad"]) //如果不存在这个key拿到对应该类型零值

	//约定成俗用ok接受返回的布尔值
	//如果存在sad key返回ok为true，value为值
	v, ok := m1["sad"]
	if !ok {
		fmt.Println("没有该key")
	} else {
		fmt.Println(v)
	}

	//map的遍历
	for k, v := range m1 {
		fmt.Println(k, v) //理想18jiwuming35
	}
	for k := range m1 {
		fmt.Println(k) //理想 jiwuming
	}
	for _, v := range m1 {
		fmt.Println(v) //18 35
	}

	//删除
	delete(m1, "jiwuming")
	fmt.Println(m1)   //map[理想:18]
	delete(m1, "sad") //删除不存在的key时不做操作
}
