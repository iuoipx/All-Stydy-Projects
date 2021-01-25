package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int //声明一个传递整形的通道
//<-
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b)     //<nil>
	b = make(chan int) //不带缓冲区通道的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道中取到了", x)
	}()
	b <- 10 //没有缓冲区
	fmt.Println("10发送到了通道b中..")
	wg.Wait()
}

func bufChannel() {
	fmt.Println(b)        //<nil>
	b = make(chan int, 1) //不带缓冲区通道的初始化
	b <- 10               //没有缓冲区
	fmt.Println("10发送到了通道b中..")
	close(b) //关闭通道
}

func main() {
	bufChannel()
	b = make(chan int, 16) //带缓冲区通道的初始化
	fmt.Println(b)         //0xc000044060
}
