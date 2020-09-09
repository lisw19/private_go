package main

import "fmt"



func f1(f func())  {
	fmt.Println("start f1")
	f()
}





//定义一个函数对f2进行一个包装
func f3(f func(int, int)) func() {
	tmp := func ()  {
		fmt.Println("hello f3")
		f(1, 2)
	}
	return tmp
}



func main()  {
	fmt.Println("start 闭包")
	var x1 func(int,int)= func f2(x, y int)  {
		fmt.Println("start f2")
		fmt.Println(x + y)
	}
	fmt.Println(x)
}