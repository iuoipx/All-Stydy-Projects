package main

import (
	"fmt"
	"strconv"
	"sync"
)

//map不支持并发安全
var m = make(map[string]int)
var m2 = sync.Map{}

var lock sync.Mutex

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			// key := strconv.Itoa(n)
			// lock.Lock()
			// set(key, n)
			// lock.Unlock()
			// fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			// wg.Done()

			key := strconv.Itoa(n)
			m2.Store(key, n)         //使用sync.Map内置的Store方法设置键值对
			value, _ := m2.Load(key) //使用sync.Map内置的Load方法根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
