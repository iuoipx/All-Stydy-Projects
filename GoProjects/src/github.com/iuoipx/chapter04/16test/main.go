package main

import (
	"fmt"
	"os"
)

var smr studentMgr

//菜单函数
func showMenu() {
	//1.打印菜单
	fmt.Println("---------------------")
	fmt.Println("欢迎光临学生管理系统!")
	fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.修改学生
		4.删除学生
		5.退出
	`)
}

func main() {
	smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		//2.等待用户选择
		fmt.Print("请输入编号:")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项!\n", choice)
		//3.执行对应的函数
		switch choice {
		case 1:
			smr.showAllStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			os.Exit(1) //退出
		default:
			fmt.Println("错误-请重新选择")
		}
	}
}
