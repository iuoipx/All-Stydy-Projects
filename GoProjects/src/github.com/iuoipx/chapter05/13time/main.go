package main

import (
	"fmt"
	"time"
)

//时间

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())

	//时间戳
	fmt.Println(now.Unix())
	//纳秒时间戳
	fmt.Println(now.UnixNano())

	//time.Unix()
	ret := time.Unix(1578451254, 0)
	fmt.Println(ret)

	//时间间隔
	fmt.Println(time.Hour) //1h0m0s
	//now+1小时
	fmt.Println(now.Add(24 * time.Hour)) //2020-01-09 10:44:06.3650852 +0800 CST m=+86400.007995001
	//两个时间间隔
	fmt.Println(now.Sub(now.Add(24 * time.Hour))) //-24h0m0s
	//定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t) //1s执行一次
	// }

	//格式化时间 把语言中时间对象转换成字符串类型的时间
	//2020-01-08
	fmt.Println(now.Format("2006-01-02")) //2020-01-08
	//2019/01/08 10:53:48
	fmt.Println(now.Format("2006/01/02 15:04:05")) //2020/01/08 10:53:48
	//2019/01/08 10:55:40 AM
	fmt.Println(now.Format("2006/01/02 03:04:05 PM")) //2020/01/08 10:55:40 AM

	//按照对应格式解析字符串类型时间
	timeObj, err := time.Parse("2006-01-02", "1999-01-08")
	if err != nil {
		fmt.Printf("parse time failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj) //1999-01-08 00:00:00 +0000 UTC
	fmt.Println(timeObj.Unix())

	// sleep
	n := 5
	fmt.Println("开始sleep")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println(n, "秒过去了")
}
