package main

import (
	"fmt"
	"runtime"
)

//runtime.Caller()

func main() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime.Caller() failed,err")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName) //main.main
	fmt.Println(file)     //d:/Projects/GoProjects/src/githup.com/iuoipx/chapter05/15runtime_demo/main.go
	fmt.Println(line)     //11
}
