package main

import "fmt"

//接口是一种特殊的数据类型， 规定了变量要有哪些方法
type dog struct {
}

type cat struct {
}

type person struct {
}

//定一个能speak的类型
type can_speaker interface {
	speak() //方法签名：只要实现了speak方法的对象 就是speaker类型
}

func (d dog) speak() {
	fmt.Println("wang wang wang ～")
}

func (c cat) speak() {
	fmt.Println("miao miao miao ～")
}

func (s person) speak() {
	fmt.Println("a a a ～")
}
func da(x can_speaker) {
	x.speak()
}

// 一个变量如果实现了接口中规定的所有方法，那么这个变量就实现了这个接口，就可以称为是这个接口类型的变量。
func main() {
	fmt.Println("ss")
	c1 := cat{}
	d1 := dog{}
	p1 := person{}
	da(c1)
	da(d1)
	da(p1)
	// 定一个接口类型的变量
	var ss can_speaker
	ss = c1
	ss.speak()
	fmt.Println(ss)
}
