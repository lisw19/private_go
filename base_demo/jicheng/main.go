package main

import "fmt"

//结构体模拟实现'继承'

type animal struct {
	name string
}

// 给动物实现一个移动方法
func (a animal) move() {
	fmt.Println(a.name, "会动")
}

//狗类
type dog struct {
	foot   uint8
	animal //此时dog就继承了动物类的move属性
}

// 给狗叫一个
func (d dog) wang() {
	fmt.Println(d.name, "wang~~~~~")
}
func main() {
	var d1 = dog{foot: 4, animal: animal{name: "lsw"}}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}
