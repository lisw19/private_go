package main

import "fmt"
//自定义类型和类型别名 

type my_int int


type persion struct{
	name string
	age int
	gender bool
}
func main()  {
	var n my_int
	n = 100
	fmt.Println(n)
	fmt.Printf("%T", n) //main.my_int
	var lsw  persion
	lsw.name = "li"
	lsw.age = 120
	lsw.gender = true
	fmt.Println(lsw)
}