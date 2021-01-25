package main

import (
	"time"

	"github.com/iuoipx/chapter05/mylogger"
)

//测试我们写的日志
func main() {
	log := mylogger.NewFileLogger("Info", "./", "asdd.log", 10*1024*1024)
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second)
	}
}
