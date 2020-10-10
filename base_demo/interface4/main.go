package main

import "fmt"

//接口嵌套

type animal interface {
	mover
	eater
}
type mover interface {
	move()
}

type eater interface {
	eat()
}
type cat struct {
	name string
}

func (c cat) move() {
	fmt.Println(c.name, "在move")
}

func (c cat) eat() {
	fmt.Println(c.name, "在吃")
}
func main() {
	var a1 animal
	cat1 := cat{
		"Tom",
	}
	cat1.move()
	cat1.eat()
	a1 = cat1
	a1.eat()

}
