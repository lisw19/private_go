package main

import (
	"fmt"
	"os"
)

/*
学生信息管理
*/

type student struct {
	name string
	id   int32
}

//定义一个student类型map
var students map[int32]*student

func showStudent() {
	//遍历s
	for k, v := range students {
		fmt.Printf("name: %s, id: %d\n", v.name, k)
	}

}

// 学生类型构造函数
func newStudent(name string, id int32) *student{
	return &student{name: name,
		id: id}
}
func addStudent() {
	//根据参数构造学生
	var (
		id int32
		name string
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)
	fmt.Println("")
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Printf("添加完成name：%s， id：%d\n", name, id)
	newStu := newStudent(name, id)
	students[id] = newStu
}



func delStudent() {
	//根据学号删除数据
	var (
		id int32
		name string
	)
	fmt.Print("请输入学号(del)：")
	fmt.Scanln(&id)
	fmt.Println("")
	fmt.Print("请输入姓名(del)：")
	fmt.Scanln(&name)
	fmt.Printf("添加完成删除name：%s， id：%d\n", name, id)
	delete(students, id)
}
func main() {
	// 初始化 开辟内存
	students = make(map[int32]*student, 30)
	for {
		fmt.Println("欢迎使用students manages")
		// 打印功能菜单
		fmt.Println(`
 					  1.查看所有学生
					  2.add student
					  3.del student
					  4.退出`)
		// 等待用户选择操作
		fmt.Print("选择操作：")
		var input int
		fmt.Scanln(&input)
		fmt.Printf("已选择动作%d！\n", input)
		// 执行对应功能
		switch input {
		case 1:
			showStudent()
		case 2:
			addStudent()
		case 3:
			delStudent()
		case 4:
			os.Exit(1)
		default:
			println("quit")

		}
	}
}
