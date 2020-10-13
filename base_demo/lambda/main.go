package main

import "fmt"

func f1(f func()) {
	fmt.Println("start f1")
	f()
}

func f2(x, y int) int {
	fmt.Println("start f2")
	fmt.Println(x + y)
	return x + y
}

//定义一个函数对f2进行一个包装
func f3(f func(int, int)) func() {
	tmp := func() {
		fmt.Println("hello f3")
		f(1, 2)
	}
	return tmp
}

func main() {
	fmt.Println("start 闭包")
	var x1 int
	x1 = f2(1, 2)
	fmt.Println(x1)
}
