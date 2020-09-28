package main

import (
	"fmt"
	"os"
)

// 声明一个学生管理对象
var smr studentM

func showMenu() {
	fmt.Println("---------------------weCome student sms")
	fmt.Println(`
1.查看所有学生
2.添加学生
3.修改学生
4.删除学生
5.退出`)
}
func main() {
	smr = studentM{
		allStudent: make(map[int64]student, 100),
	}
	for true {
		showMenu()
		//等待用户输入
		fmt.Print("\n\n请输入操作:")
		var change int
		fmt.Scanln(&change)
		switch change {
		case 1:
			smr.showS()
		case 2:
			smr.addS()
		case 3:
			smr.changeS()
		case 4:
			smr.delS()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("输入操作错误")
		}
		fmt.Print("执行操作", change)
	}
}
