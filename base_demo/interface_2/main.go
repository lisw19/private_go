package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	color string
	feet  int8
}

func (c chicken) eat(bug string) {
	fmt.Println(c.color, "的鸡在吃", bug)
}

func (c chicken) move() {
	fmt.Println("鸡在动")
}

func (c cat) move() {
	fmt.Println("走猫步")
}
func (c cat) eat(string2 string) {
	fmt.Println("猫吃鱼")
}
func main() {
	fmt.Println("interface")
	var a1 animal                  // 定义一个接口类型变量
	bc := cat{name: "淘气", feet: 4} //定义一个cat类型变量
	a1 = bc
	a1.eat("sdfs")
	a1.move()

	l := chicken{color: "yellow", feet: 2}
	a1 = l
	a1.eat("蚯蚓")
}
