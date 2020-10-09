package main

import "fmt"

type mover interface {
	move()
}

type eater interface {
	eat()
}

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// 使用值接收者 实现接口方法
//func (c cat) move() {
//	fmt.Println("走猫步")
//}
//func (c cat) eat(string2 string) {
//	fmt.Println("猫吃鱼")
//}

// 使用值接收者 实现接口方法
func (c *cat) move() {
	fmt.Println("走猫步")
}
func (c *cat) eat(string2 string) {
	fmt.Println("猫吃鱼")
}
func main() {
	fmt.Println("interface")
	var a1 animal
	c1 := cat{"tom", 4}  // cat
	c2 := &cat{"假老链", 4} //*cat
	a1 = &c1             //实现animal接口的是cat的指针
	fmt.Println(c1.name)
	a1 = c2
	a1.move()

}
