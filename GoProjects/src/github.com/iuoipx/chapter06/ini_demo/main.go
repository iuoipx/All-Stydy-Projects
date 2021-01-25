package main

import (
	"fmt"
	"reflect"
)

//ini配置文件解析器

//MysqlConfig Mysql配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

//RedisConfig redis配置结构体
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"post"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

func loadIni(fileName string, data interface{}) (err error) {
	//0.参数校验
	//0.1 传进来的data参数必须是指针类型(因为需要在函数中对其赋值)
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	//0.2 传进来的data参数必须是结构体类型指针(因为配置文件中各种键值对需要赋值给结构体的字段)
	//1.读文件得到字节类型数据
	//2.一行一行的读数据
	//2.1 如果是注释就跳过
	//2.2 如果是[开头的表示是节(section)
	//2.3 如果不是[开头就是=分割的键值对
	return
}

func main() {
	var mc MysqlConfig
	err := loadIni("./conf.ini", &mc)
	if err != nil {

	}
}
