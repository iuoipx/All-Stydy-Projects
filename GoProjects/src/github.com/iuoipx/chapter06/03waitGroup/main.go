package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//waitGroup

func f() {
	rand.Seed(time.Now().UnixNano()) //随机数种子
	for i := 0; i < 5; i++ {
		r1 := rand.Int()    //int64
		r2 := rand.Intn(10) //0<=x<10
		fmt.Println(r1, r2)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	//如何知道这10个goroutine都结束了
	//time.Sleep
	wg.Wait() //等待wg的计数器减为0
}
