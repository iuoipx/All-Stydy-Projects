package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      = 0
	wg     sync.WaitGroup
	rwlock sync.RWMutex
)

func read() {
	defer wg.Done()
	rwlock.Lock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwlock.Unlock()
}

func write() {
	defer wg.Done()
	rwlock.RLock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	rwlock.RUnlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
