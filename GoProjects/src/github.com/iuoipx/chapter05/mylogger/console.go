package mylogger

import (
	"fmt"
	"time"
)

//LogLevel 日志级别
type LogLevel uint16

//ConsoleLogger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

//NewLog 构造函数
func NewLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

//Debug debug级别
func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

//Trace trace级别
func (c ConsoleLogger) Trace(format string, a ...interface{}) {
	c.log(TRACE, format, a...)
}

//Info info级别
func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

//Warning warning级别
func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

//Error error级别
func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

//Fatal fatal级别
func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
